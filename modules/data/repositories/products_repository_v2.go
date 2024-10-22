package repositories

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/yehpattana/api-yehpattana/configs"
	commonfolderpath "github.com/yehpattana/api-yehpattana/modules/commons/common_folder_path"
	commonhelpers "github.com/yehpattana/api-yehpattana/modules/commons/common_helpers"
	commonimages "github.com/yehpattana/api-yehpattana/modules/commons/common_images"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	producthelpers "github.com/yehpattana/api-yehpattana/modules/products/product_helpers"
	productsdto "github.com/yehpattana/api-yehpattana/modules/products/products_dto"
	productsresponses "github.com/yehpattana/api-yehpattana/modules/products/products_responses"
	"gorm.io/gorm"
)

type ProductsRepositoryV2Interface interface {
	// Admin side
	CheckIsProductExistByMasterCode(masterCode string) (*productsresponses.CheckIsProductExistByMasterCodeResponse, error)
	CreateProduct(req *productsdto.CreateProductRequest) (*productsresponses.CreatedProductAdminResponse, error)
	CreateStock(req *productsdto.CreateStockRequest) (*productsresponses.CreatedStockAdminResponse, error)
	UpdateStock(req *productsdto.UpdateStockRequest) (*productsresponses.UpdateStockResponse, error)
	DecreaseStock(req *productsdto.DecreaseStockRequest) (*productsresponses.DecreaseStockResponse, error)
	UploadCoverImage(req *productsdto.UploadCoverImageRequest) (*productsresponses.UploadedCoverImageAdminResponse, error)
	UploadSizeChart(req *productsdto.UploadSizeChartRequest) (*productsresponses.UploadedSizeChartAdminResponse, error)
	GetCoverImageByMasterCode(masterCode string) (*productsresponses.GetCoverImageByMasterCodeResponse, error)
	GetSizeChartByMasterCode(masterCode string) (*productsresponses.GetSizeChartByMasterCodeResponse, error)
	UpdateCoverImageByMasterCode(req *productsdto.UpdateCoverImageRequest) (*productsresponses.UpdatedCoverImageAdminResponse, error)
	UpdateSizeChartByMasterCode(req *productsdto.UpdateSizeChartRequest) (*productsresponses.UpdatedSizeChartAdminResponse, error)
	UpdateMainProductDetailByMasterCode(req *productsdto.UpdateMainProductDetailByMasterCodeRequest) (*productsresponses.UpdatedProductAdminResponse, error)
	UpdateMasterCode(req *productsdto.UpdateMasterCodeRequest) (*productsresponses.UpdateMasterCodeResponse, error)
	UpdateProductVaraint(req *productsdto.UpdateProductVariantRequest) (*productsresponses.UpdateProductVaraintResponse, error)
	GetAllProductsAdmin(role string, companyId string) (*productsresponses.GetAllProductsAdminResponse, error)
	GetProductDetailAdmin(masterCode string) (*productsresponses.GetProductDetailAdminResponse, error)
	DeleteAllProductInMasterCode(masterCode string) (*productsresponses.DeleteAllProductInMasterCodeResponse, error)
	DeleteProductVariant(productId string) (*productsresponses.DeleteProductVariantResponse, error)

	// Common side
	GetAllProductByCompany(companyId string) (*productsresponses.GetAllProductByCompanyResponse, error)

	// Web view side
	GetAllProducts() (*productsresponses.GetAllProductsResponse, error)
	GetProductDetail(productId string) (*productsresponses.GetProductDetailResponse, error) // TODO
}

func ProductsRepositoryV2Impl(db *gorm.DB, cfg *configs.ConfigInterface) ProductsRepositoryV2Interface {
	return &productsRepositoryV2Impl{
		DB:  db,
		cfg: cfg,
	}
}

type productsRepositoryV2Impl struct {
	*gorm.DB
	cfg *configs.ConfigInterface
}

func (p *productsRepositoryV2Impl) GetAllProductByCompany(companyId string) (*productsresponses.GetAllProductByCompanyResponse, error) {
	var products []productsresponses.ProductEssentialDetailAdminResponse

	result := p.DB.
		Table("Products").
		Where("created_by_company = ? AND use_as_primary_data = ?", companyId, 1).
		Order("created_at DESC").
		Find(&products)

	if result.Error != nil {
		return &productsresponses.GetAllProductByCompanyResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get all product by company id",
				ErrorMessage: result.Error.Error(),
			},
			Products: nil,
		}, result.Error
	}
	// Sort stocks based on size order
	sort.Slice(products, func(i, j int) bool {
		return producthelpers.ProductOrder[strings.ToLower(products[i].ProductGroup)] < producthelpers.ProductOrder[strings.ToLower(products[j].ProductGroup)]
	})
	for i := range products {
		details, err := p.GetProductDetail(products[i].MasterCode)
		if err != nil {
			fmt.Println("Error fetching variants:", err)
			continue
		}
		minPrice := float64(0)
		minRrpPrice := float64(0)
		currency := ""
		var colorCodes []string
		for _, variant := range details.ProductVaraints {
			for _, stock := range variant.Stock {
				if stock.Price > 0 && (minPrice == 0 || stock.Price < minPrice) {
					minPrice = stock.Price
					currency = stock.Currency
				}
				if stock.RrpPrice > 0 && (minRrpPrice == 0 || stock.RrpPrice < minRrpPrice) {
					minRrpPrice = stock.RrpPrice
					currency = stock.Currency
				}
			}
			colorCodes = append(colorCodes, variant.ColorCode)
		}
		products[i].Variants = details.ProductVaraints
		products[i].Colors = colorCodes
		products[i].Price = minPrice
		products[i].RrpPrice = minRrpPrice
		products[i].Currency = currency
	}

	return &productsresponses.GetAllProductByCompanyResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to get all product by company id!",
		},
		Products: products,
	}, nil
}

func (p *productsRepositoryV2Impl) DeleteProductVariant(productId string) (*productsresponses.DeleteProductVariantResponse, error) {
	// Check product variant exist by product id
	isProductExist, err := checkProductVaraintExistByProductId(productId, p.DB)
	if err != nil {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product variant exist by product id",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	if !isProductExist {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product variant not found",
				ErrorMessage: "Product variant not found by product id: " + productId,
			},
		}, nil
	}

	// delete stock by product id
	result := p.DB.Table("Stock").Where("product_id = ?", productId).Delete(&entities.Stock{})
	if result.Error != nil {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to delete stock by product id",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// fetch product variant by product id
	productVariant, err := getProductVariantDetailByProductId(productId, p.DB)
	if err != nil {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get product variant by product id",
				ErrorMessage: err.Error(),
			},
		}, err

	}

	// if use_as_primary_data is true then update first product in list of master_code to true
	if productVariant.UseAsPrimaryData {
		// fetch all product varaints by master code
		products, err := getAllProductsByMasterCode(p.DB, productVariant.MasterCode)
		if err != nil {
			return &productsresponses.DeleteProductVariantResponse{
				Status: commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to get all product by master code",
					ErrorMessage: err.Error(),
				},
			}, err
		}

		// update first product in list of master_code to true
		for _, product := range products {
			if product.Id != productId {
				result := p.DB.Table("Products").
					Where("id = ?", product.Id).
					Select("use_as_primary_data", "updated_at")

				updates := map[string]interface{}{
					"use_as_primary_data": true,
					"updated_at":          commonhelpers.GetCurrentTimeISO(),
				}

				result = result.Updates(updates)
				if result.Error != nil {
					return &productsresponses.DeleteProductVariantResponse{
						Status: commonresponse.CommonResponse{
							Status:       false,
							Message:      "REPO - Failed to update use_as_primary_data to db",
							ErrorMessage: result.Error.Error(),
						},
					}, result.Error
				}
				break
			}
		}
	}

	// delete product variant by product id
	result = p.DB.Table("Products").Where("id = ?", productId).Delete(&entities.Product{})
	if result.Error != nil {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to delete product variant by product id",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	return &productsresponses.DeleteProductVariantResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to delete product variant!",
		},
	}, nil
}

func (p *productsRepositoryV2Impl) DeleteAllProductInMasterCode(masterCode string) (*productsresponses.DeleteAllProductInMasterCodeResponse, error) {
	// check product exist by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.DeleteAllProductInMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.DeleteAllProductInMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
		}, nil
	}

	// fetch all product varaints by master code
	products, err := getAllProductsByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.DeleteAllProductInMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get all product by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	// fetch stock by each product variant and delete all stock
	for _, product := range products {
		// delete stock by product id
		result := p.DB.Table("Stock").Where("product_id = ?", product.Id).Delete(&entities.Stock{})
		if result.Error != nil {
			return &productsresponses.DeleteAllProductInMasterCodeResponse{
				Status: commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to delete stock by product id",
					ErrorMessage: result.Error.Error(),
				},
			}, result.Error
		}
	}

	result := p.DB.Table("Products").Where("master_code = ?", masterCode).Delete(&entities.Product{})
	if result.Error != nil {
		return &productsresponses.DeleteAllProductInMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to delete all product in master code",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}
	log := &entities.Log{Description: fmt.Sprintf("delete product %s", masterCode), CreatedAt: commonhelpers.GetCurrentTimeISO(), UpdatedBy: ""}
	saveLog(p.DB, log)
	return &productsresponses.DeleteAllProductInMasterCodeResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to delete all product in master code!",
		},
	}, nil

}

func getProductVariantDetailByProductId(productId string, db *gorm.DB) (entities.Product, error) {
	productVariant := entities.Product{}
	db.Table("Products").Where("id = ?", productId).First(&productVariant)
	return productVariant, nil
}

func (p *productsRepositoryV2Impl) UpdateProductVaraint(req *productsdto.UpdateProductVariantRequest) (*productsresponses.UpdateProductVaraintResponse, error) {
	// check product exist by product id
	// TODO recheck why found product by product id is bug
	print("eiei req product id: ", req.ProductId)
	isProductExist, err := checkProductVaraintExistByProductId(req.ProductId, p.DB)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product variant exist by product id",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	if !isProductExist {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by product id: " + req.ProductId,
			},
			Data: nil,
		}, nil
	}

	// fetch product variant by product id
	productVariant, err := getProductVariantDetailByProductId(req.ProductId, p.DB)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get product variant by product id",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	// upload front image
	currentFrontImage := productVariant.FrontImage
	frontImageFolder := commonfolderpath.ProductFrontFolderPath
	frontImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.FrontImage, currentFrontImage, frontImageFolder)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload front image",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	// upload back image
	currentBackImage := productVariant.BackImage
	backImageFolder := commonfolderpath.ProductBackFolderPath
	backImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.BackImage, currentBackImage, backImageFolder)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload back image",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	isColorCodeIsEmtpy := req.ColorCode == ""
	if isColorCodeIsEmtpy {
		req.ColorCode = productVariant.ColorCode
	}

	isProductCodeIsEmtpy := req.ProductCode == ""
	if isProductCodeIsEmtpy {
		req.ProductCode = productVariant.ProductCode
	}

	// if use_as_primary_data is true then update all product with same master code to false
	if req.UseAsPrimaryData {
		result := p.DB.Table("Products").
			Where("master_code = ?", productVariant.MasterCode).
			Select("use_as_primary_data", "updated_at")

		updates := map[string]interface{}{
			"use_as_primary_data": false,
			"updated_at":          commonhelpers.GetCurrentTimeISO(),
		}

		result = result.Updates(updates)
		if result.Error != nil {
			return &productsresponses.UpdateProductVaraintResponse{
				Status: &commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to update use_as_primary_data to db",
					ErrorMessage: result.Error.Error(),
				},
				Data: nil,
			}, result.Error
		}
	}

	// update product variant to db
	result := p.DB.Table("Products").
		Where("id = ?", req.ProductId).
		Select("product_code", "color_code", "front_image", "back_image", "price", "use_as_primary_data", "updated_at")

	updates := map[string]interface{}{
		"product_code":        req.ProductCode,
		"color_code":          req.ColorCode,
		"front_image":         frontImageUrl,
		"back_image":          backImageUrl,
		"price":               req.Price,
		"use_as_primary_data": req.UseAsPrimaryData,
		"updated_at":          commonhelpers.GetCurrentTimeISO(),
	}

	result = result.Updates(updates)
	if result.Error != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update product variant to db",
				ErrorMessage: result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	// fetch updated product variant from db
	updatedProductVariant, err := getProductVariantDetailByProductId(req.ProductId, p.DB)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get product variant by product id",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	return &productsresponses.UpdateProductVaraintResponse{
		Status: &commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update product variant!",
		},
		Data: &productsresponses.ProductVaraintAfterUpdateResponse{
			ProductId:        updatedProductVariant.Id,
			MasterCode:       updatedProductVariant.MasterCode,
			ProductCode:      updatedProductVariant.ProductCode,
			ColorCode:        updatedProductVariant.ColorCode,
			FrontImage:       updatedProductVariant.FrontImage,
			BackImage:        updatedProductVariant.BackImage,
			Price:            updatedProductVariant.Price,
			UseAsPrimaryData: updatedProductVariant.UseAsPrimaryData,
		},
	}, nil
}

func (p *productsRepositoryV2Impl) UpdateMasterCode(req *productsdto.UpdateMasterCodeRequest) (*productsresponses.UpdateMasterCodeResponse, error) {
	// check product exist by master code
	isOldMasterCodeEmpty := req.OldMasterCode == ""
	if isOldMasterCodeEmpty {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Old master code is empty",
				ErrorMessage: "Old master code is empty",
			},
		}, errors.New("Old master code is empty")
	}

	isNewMasterCodeEmpty := req.NewMasterCode == ""
	if isNewMasterCodeEmpty {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "New master code is empty",
				ErrorMessage: "New master code is empty",
			},
		}, errors.New("New master code is empty")

	}

	isUpdatedByEmpty := req.UpdatedBy == ""
	if isUpdatedByEmpty {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Updated by is empty",
				ErrorMessage: "Updated by is empty",
			},
		}, errors.New("Updated by is empty")
	}

	checkProductExist, err := p.CheckIsProductExistByMasterCode(req.OldMasterCode)
	if err != nil {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + req.OldMasterCode,
			},
		}, nil
	}

	checkIsNewMasterCodeExist, err := p.CheckIsProductExistByMasterCode(req.NewMasterCode)
	if err != nil {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check new master code exist",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isNewMasterCodeExist := checkIsNewMasterCodeExist.IsProductExist

	if isNewMasterCodeExist {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "New master code already exist",
				ErrorMessage: "New master code already exist: " + req.NewMasterCode,
			},
		}, nil
	}

	// update all product with old master code to new master code
	result := p.DB.Table("Products").
		Where("master_code = ?", req.OldMasterCode).
		Select("master_code", "updated_at")

	updates := map[string]interface{}{
		"master_code": req.NewMasterCode,
		"updated_at":  commonhelpers.GetCurrentTimeISO(),
		"updated_by":  req.UpdatedBy,
	}

	result = result.Updates(updates)
	if result.Error != nil {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update master code to db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// fetch all updated product from db by new master code and filter only the same master code
	updatedProducts, err := getAllProductsByMasterCode(p.DB, req.NewMasterCode)
	if err != nil {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get product by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return &productsresponses.UpdateMasterCodeResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update master code!",
		},
		Data: updatedProducts,
	}, nil
}

func (p *productsRepositoryV2Impl) UpdateMainProductDetailByMasterCode(req *productsdto.UpdateMainProductDetailByMasterCodeRequest) (*productsresponses.UpdatedProductAdminResponse, error) {
	masterCode := req.MasterCode

	// check product exist by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.UpdatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.UpdatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
		}, nil
	}

	// update product to db
	result := p.DB.Table("Products").
		Where("master_code = ? AND use_as_primary_data = ?", req.MasterCode, true).
		Select("name", "product_group", "season", "gender", "product_class", "collection", "category", "brand", "is_club", "club_name", "remark", "launch_date", "end_of_life", "size_chart", "pack_size", "current_supplier", "description", "fabric_content", "fabric_type", "weight", "updated_at")

	updates := map[string]interface{}{
		"name":               req.Name,
		"product_group":      req.ProductGroup,
		"season":             req.Season,
		"gender":             req.Gender,
		"product_class":      req.ProductClass,
		"collection":         req.Collection,
		"category":           req.Category,
		"brand":              req.Brand,
		"is_club":            req.IsClub,
		"club_name":          req.ClubName,
		"remark":             req.Remark,
		"launch_date":        req.LaunchDate,
		"end_of_life":        req.EndOfLife,
		"size_chart":         req.SizeChart,
		"pack_size":          req.PackSize,
		"current_supplier":   req.CurrentSupplier,
		"description":        req.Description,
		"fabric_content":     req.FabricContent,
		"fabric_type":        req.FabricType,
		"weight":             req.Weight,
		"created_by_company": req.CreatedByCompany,
		"updated_at":         commonhelpers.GetCurrentTimeISO(),
	}

	// TODO ต้องยิงอัพเดต varaint ที่เหลือด้วย

	result = result.Updates(updates)
	if result.Error != nil {
		return &productsresponses.UpdatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update product to db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// fetch updated product from db
	updatedProduct, err := getMainProductDetailPrimaryDataByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.UpdatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to get main product detail by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}
	log := &entities.Log{Description: fmt.Sprintf("update product %s", updatedProduct.Name), CreatedAt: commonhelpers.GetCurrentTimeISO(), UpdatedBy: req.EditedBy}
	saveLog(p.DB, log)
	return &productsresponses.UpdatedProductAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update product!",
		},
		Data: productsresponses.ProductAdminResponse{
			Id:               updatedProduct.Id,
			Name:             updatedProduct.Name,
			ProductCode:      updatedProduct.ProductCode,
			MasterCode:       updatedProduct.MasterCode,
			ColorCode:        updatedProduct.ColorCode,
			ProductStatus:    updatedProduct.ProductStatus,
			CoverImage:       updatedProduct.CoverImage,
			FrontImage:       updatedProduct.FrontImage,
			BackImage:        updatedProduct.BackImage,
			Price:            updatedProduct.Price,
			UseAsPrimaryData: updatedProduct.UseAsPrimaryData,
			ProductGroup:     updatedProduct.ProductGroup,
			Season:           updatedProduct.Season,
			Gender:           updatedProduct.Gender,
			ProductClass:     updatedProduct.ProductClass,
			Collection:       updatedProduct.Collection,
			Category:         updatedProduct.Category,
			Brand:            updatedProduct.Brand,
			IsClub:           updatedProduct.IsClub,
			ClubName:         updatedProduct.ClubName,
			Remark:           updatedProduct.Remark,
			LaunchDate:       updatedProduct.LaunchDate,
			EndOfLife:        updatedProduct.EndOfLife,
			SizeChart:        updatedProduct.SizeChart,
			PackSize:         updatedProduct.PackSize,
			CurrentSupplier:  updatedProduct.CurrentSupplier,
			Description:      updatedProduct.Description,
			FabricContent:    updatedProduct.FabricContent,
			FabricType:       updatedProduct.FabricType,
			Weight:           updatedProduct.Weight,
			CreatedByCompany: updatedProduct.CreatedByCompany,
			CreatedBy:        updatedProduct.CreatedBy,
			CreatedAt:        updatedProduct.CreatedAt,
			UpdatedAt:        updatedProduct.UpdatedAt,
		},
	}, nil

}
func (p *productsRepositoryV2Impl) CheckIsProductExistByMasterCode(masterCode string) (*productsresponses.CheckIsProductExistByMasterCodeResponse, error) {
	productData := p.DB.Table("Products").Where("master_code = ?", masterCode).First(&entities.Product{})

	if productData.RowsAffected == 0 {
		return &productsresponses.CheckIsProductExistByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:  false,
				Message: "Product not found: " + masterCode,
			},
			IsProductExist: false,
		}, nil
	}

	return &productsresponses.CheckIsProductExistByMasterCodeResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Product found",
		},
		IsProductExist: true,
	}, nil
}

func (p *productsRepositoryV2Impl) CreateProduct(req *productsdto.CreateProductRequest) (*productsresponses.CreatedProductAdminResponse, error) {
	// validation checker
	isHaveProductName := req.Name != ""
	isHaveProductCode := req.ProductCode != ""
	isHaveMasterCode := req.MasterCode != ""
	isHaveCreatedBy := req.CreatedBy != ""
	isHaveCreatedByCompany := req.CreatedByCompany != ""
	isHaveRequiredField := !isHaveProductName || !isHaveProductCode || !isHaveMasterCode || !isHaveCreatedBy || !isHaveCreatedByCompany

	isCreateByCompanyIsUUID := commonhelpers.CheckIsValidUUID(req.CreatedByCompany)
	if !isCreateByCompanyIsUUID {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Invalid Created By Company",
				ErrorMessage: "Created By Company must be UUID (company id)",
			},
			Data: productsresponses.ProductAdminResponse{},
		}, nil
	}
	productStatusValue := req.ProductStatus

	isValidProductStatus := producthelpers.CheckIsValidStatus(productStatusValue)

	if !isValidProductStatus {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Invalid Product Status Value",
				ErrorMessage: "Product Status must be only: 'available', 'hidden', 'out_of_stock'",
			},
			Data: productsresponses.ProductAdminResponse{},
		}, nil
	}

	genderValue := req.Gender
	validGenders := []string{"male", "female", "unisex", "kids"}
	isValidGender := false
	for _, validGender := range validGenders {
		if genderValue == validGender {
			isValidGender = true
			break
		}
	}

	coverImageValue := req.CoverImage
	if coverImageValue == "" {
		coverImageValue = commonimages.DefaultCoverImage
	}

	if isHaveRequiredField {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Name, Product Code, Master Code, Created By Company or Created By is empty",
			},
			Data: productsresponses.ProductAdminResponse{},
		}, nil
	}
	if !isValidGender {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Invalid Gender Value",
				ErrorMessage: "Gender must be only: 'male', 'female', 'unisex', 'kids'",
			},
			Data: productsresponses.ProductAdminResponse{},
		}, nil
	}
	// upload front image
	frontImageFolder := commonfolderpath.ProductFrontFolderPath
	frontImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.FrontImage, commonimages.DefaultProductImageFrontSide, frontImageFolder)
	if err != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload front image",
				ErrorMessage: err.Error(),
			},
			Data: productsresponses.ProductAdminResponse{},
		}, err
	}

	// upload back image
	backImageFolder := commonfolderpath.ProductBackFolderPath
	backImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.BackImage, commonimages.DefaultProductImageBackSide, backImageFolder)
	if err != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload back image",
				ErrorMessage: err.Error(),
			},
			Data: productsresponses.ProductAdminResponse{},
		}, err
	}
	// upload back image
	sizeChartImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.SizeChart, commonimages.DefaultProductImageBackSide, backImageFolder)
	if err != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload size chart image",
				ErrorMessage: err.Error(),
			},
			Data: productsresponses.ProductAdminResponse{},
		}, err
	}

	// find product by master code if not exist then assign use_as_primary_data to true
	checkProductExist, err := p.CheckIsProductExistByMasterCode(req.MasterCode)
	if err != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
			Data: productsresponses.ProductAdminResponse{},
		}, err

	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		req.UseAsPrimaryData = true
	}

	// Create product object
	product := &entities.Product{
		Id:               commonhelpers.GenerateUUID(),
		Name:             req.Name,
		ProductCode:      req.ProductCode,
		MasterCode:       req.MasterCode,
		ColorCode:        req.ColorCode,
		ProductStatus:    req.ProductStatus,
		UseAsPrimaryData: req.UseAsPrimaryData,
		CoverImage:       coverImageValue,
		FrontImage:       frontImageUrl,
		BackImage:        backImageUrl,
		Price:            req.Price,
		ProductGroup:     req.ProductGroup,
		Season:           req.Season,
		Gender:           req.Gender,
		ProductClass:     req.ProductClass,
		Collection:       req.Collection,
		Category:         req.Category,
		Brand:            req.Brand,
		IsClub:           req.IsClub,
		ClubName:         req.ClubName,
		Remark:           req.Remark,
		LaunchDate:       req.LaunchDate,
		EndOfLife:        req.EndOfLife,
		SizeChart:        sizeChartImageUrl,
		PackSize:         req.PackSize,
		CurrentSupplier:  req.CurrentSupplier,
		Description:      req.Description,
		FabricContent:    req.FabricContent,
		FabricType:       req.FabricType,
		Weight:           req.Weight,
		CreatedByCompany: req.CreatedByCompany,
		CreatedBy:        req.CreatedBy,
		CreatedAt:        commonhelpers.GetCurrentTimeISO(),
		UpdatedAt:        commonhelpers.GetCurrentTimeISO(),
	}

	// Save product to db
	result := p.DB.Table("Products").Create(product)
	if result.Error != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to save product to db",
				ErrorMessage: result.Error.Error(),
			},
			Data: productsresponses.ProductAdminResponse{},
		}, result.Error
	}

	log := &entities.Log{Description: fmt.Sprintf("create product %s", req.Name), CreatedAt: commonhelpers.GetCurrentTimeISO(), UpdatedBy: req.CreatedBy}
	saveLog(p.DB, log)
	return &productsresponses.CreatedProductAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to create product!",
		},
		Data: productsresponses.ProductAdminResponse{
			Id:               product.Id,
			Name:             product.Name,
			ProductCode:      product.ProductCode,
			MasterCode:       product.MasterCode,
			ColorCode:        product.ColorCode,
			ProductStatus:    product.ProductStatus,
			CoverImage:       product.CoverImage,
			FrontImage:       product.FrontImage,
			BackImage:        product.BackImage,
			Price:            product.Price,
			UseAsPrimaryData: product.UseAsPrimaryData,
			ProductGroup:     product.ProductGroup,
			Season:           product.Season,
			Gender:           product.Gender,
			ProductClass:     product.ProductClass,
			Collection:       product.Collection,
			Category:         product.Category,
			Brand:            product.Brand,
			IsClub:           product.IsClub,
			ClubName:         product.ClubName,
			Remark:           product.Remark,
			LaunchDate:       product.LaunchDate,
			EndOfLife:        product.EndOfLife,
			SizeChart:        product.SizeChart,
			PackSize:         product.PackSize,
			CurrentSupplier:  product.CurrentSupplier,
			Description:      product.Description,
			FabricContent:    product.FabricContent,
			FabricType:       product.FabricType,
			Weight:           product.Weight,
			CreatedByCompany: product.CreatedByCompany,
			CreatedBy:        product.CreatedBy,
			CreatedAt:        product.CreatedAt,
			UpdatedAt:        product.UpdatedAt,
		},
	}, nil

}

func (p *productsRepositoryV2Impl) UploadCoverImage(req *productsdto.UploadCoverImageRequest) (*productsresponses.UploadedCoverImageAdminResponse, error) {
	coverImageFolder := commonfolderpath.ProductCoverFolderPath
	coverImageUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.CoverImage, commonimages.DefaultCoverImage, coverImageFolder)
	if err != nil {
		return &productsresponses.UploadedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload cover image",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return &productsresponses.UploadedCoverImageAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to upload cover image!",
		},

		Data: productsresponses.ProductImageResponse{
			Image: coverImageUrl,
		},
	}, nil
}

func checkIsOutOfStock(stock int) bool {
	return stock == 0
}

func assignVaraintStatus(isOutOfStock bool) string {
	if isOutOfStock {
		return "out_of_stock"
	}
	return "available"
}

func (p *productsRepositoryV2Impl) UploadSizeChart(req *productsdto.UploadSizeChartRequest) (*productsresponses.UploadedSizeChartAdminResponse, error) {
	sizeChartFolder := commonfolderpath.ProductSizeChartFolderPath
	sizeChartUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.SizeChart, commonimages.DefaultChartSizeImage, sizeChartFolder)
	if err != nil {
		return &productsresponses.UploadedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to upload size chart",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return &productsresponses.UploadedSizeChartAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to upload size chart!",
		},
		Data: productsresponses.ProductImageResponse{
			Image: sizeChartUrl,
		},
	}, nil

}

func (p *productsRepositoryV2Impl) GetCoverImageByMasterCode(masterCode string) (*productsresponses.GetCoverImageByMasterCodeResponse, error) {
	// Check required field
	isHaveMasterCode := masterCode != ""

	if !isHaveMasterCode {
		return &productsresponses.GetCoverImageByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
		}, nil
	}

	// check product exist by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetCoverImageByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.GetCoverImageByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
		}, nil
	}

	// fetch cover image from db
	coverImage, err := fetchImageFromDbByMasterCode(p.DB, masterCode, "Products", "cover_image")
	if err != nil {
		return &productsresponses.GetCoverImageByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch cover image from db",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return &productsresponses.GetCoverImageByMasterCodeResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch cover image!",
		},
		Data: productsresponses.ProductImageResponse{
			MasterCode: masterCode,
			Image:      coverImage,
		},
	}, nil
}

func (p *productsRepositoryV2Impl) GetSizeChartByMasterCode(masterCode string) (*productsresponses.GetSizeChartByMasterCodeResponse, error) {

	// Check required field
	isHaveMasterCode := masterCode != ""

	if !isHaveMasterCode {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
		}, nil
	}

	// check product exist by master code

	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
		}, nil
	}

	// fetch size chart from db
	sizeChart, err := fetchImageFromDbByMasterCode(p.DB, masterCode, "Products", "size_chart")
	if err != nil {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch size chart from db",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isHaveSizeChartImage := sizeChart != ""
	if !isHaveSizeChartImage {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       true,
				Message:      "Founded master code and primary data size chart is empty string",
				ErrorMessage: "Size chart image is not yet upload, please upload that or still keep this for no size chart image",
			},
			Data: productsresponses.ProductImageResponse{
				MasterCode: masterCode,
				Image:      sizeChart,
			},
		}, nil
	}

	return &productsresponses.GetSizeChartByMasterCodeResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch size chart!",
		},
		Data: productsresponses.ProductImageResponse{
			MasterCode: masterCode,
			Image:      sizeChart,
		},
	}, nil
}

func (p *productsRepositoryV2Impl) UpdateCoverImageByMasterCode(req *productsdto.UpdateCoverImageRequest) (*productsresponses.UpdatedCoverImageAdminResponse, error) {
	isHaveMasterCode := req.MasterCode != ""
	isHaveCoverImage := req.CoverImage != ""

	if !isHaveMasterCode {
		return &productsresponses.UpdatedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
		}, nil
	}

	if !isHaveCoverImage {
		req.CoverImage = commonimages.DefaultCoverImage
	}

	// check product exist by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(req.MasterCode)
	if err != nil {
		return &productsresponses.UpdatedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.UpdatedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + req.MasterCode,
			},
		}, nil
	}

	// update cover image to db
	result := p.DB.Table("Products").
		Where("master_code = ? AND use_as_primary_data = ?", req.MasterCode, true).
		Select("cover_image")

	updates := map[string]interface{}{
		"cover_image": req.CoverImage,
		"updated_at":  commonhelpers.GetCurrentTimeISO(),
	}

	result = result.Updates(updates)

	if result.Error != nil {
		return &productsresponses.UpdatedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update cover image to db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	return &productsresponses.UpdatedCoverImageAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update cover image!",
		},
		Data: productsresponses.ProductImageResponse{
			MasterCode: req.MasterCode,
			Image:      req.CoverImage,
		},
	}, nil

}

func (p *productsRepositoryV2Impl) UpdateSizeChartByMasterCode(req *productsdto.UpdateSizeChartRequest) (*productsresponses.UpdatedSizeChartAdminResponse, error) {
	isHaveMasterCode := req.MasterCode != ""

	if !isHaveMasterCode {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
		}, nil
	}

	// check product exist by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(req.MasterCode)
	if err != nil {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	isProductExist := checkProductExist.IsProductExist

	if !isProductExist {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + req.MasterCode,
			},
		}, nil
	}
	sizeChartFolder := commonfolderpath.ProductSizeChartFolderPath
	sizeChartUrl, err := commonhelpers.UploadImageOrUseDefaultImage(req.SizeChart, commonimages.DefaultChartSizeImage, sizeChartFolder)
	if err != nil {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update size chart to db",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	// update size chart to db
	result := p.DB.Table("Products").
		Where("master_code = ? AND use_as_primary_data = ?", req.MasterCode, true).
		Select("size_chart")

	update := map[string]interface{}{
		"size_chart": sizeChartUrl,
		"updated_at": commonhelpers.GetCurrentTimeISO(),
	}

	result = result.Updates(update)

	if result.Error != nil {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update size chart to db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	return &productsresponses.UpdatedSizeChartAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update size chart!",
		},
		Data: productsresponses.ProductImageResponse{
			MasterCode: "",
			Image:      sizeChartUrl,
		},
	}, nil
}

func checkSizeInStockExistByProductId(productId string, size string, db *gorm.DB) (bool, error) {
	stock := &entities.Stock{}
	// Query to check if the product and size exist in the stock table.
	result := db.Table("Stock").Where("product_id = ? AND size = ?", productId, size).First(stock)

	// Check for errors in the query execution.
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Return false if no record is found.
			return false, nil
		}
		// Return the error if something else went wrong.
		return false, result.Error
	}

	// Return true if the size is found.
	return true, nil
}

func (p *productsRepositoryV2Impl) CreateStock(req *productsdto.CreateStockRequest) (*productsresponses.CreatedStockAdminResponse, error) {
	// check product exist by product code
	checkProductExist, err := checkProductVaraintExistByProductId(req.ProductId, p.DB)
	if err != nil {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by product code",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	isProductExist := checkProductExist

	if !isProductExist {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by product code: " + req.ProductId,
			},
			Data: nil,
		}, nil
	}

	// TODO check is have this size in product variant
	// Find product_id in Stock table
	isSizeAlreadyExist, err := checkSizeInStockExistByProductId(req.ProductId, req.Size, p.DB)
	if err != nil {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check size exist by product id",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	if isSizeAlreadyExist {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Size already exist",
				ErrorMessage: "This size already exist in stock, Please update the stock instead of create new one",
			},
			Data: nil,
		}, nil
	}

	// validation checker
	isHaveProductCode := req.ProductId != ""
	isHaveSizeRemark := req.SizeRemark != ""
	isSizeOther := req.Size == "Other"
	isSizeRemarkRequired := isSizeOther && !isHaveSizeRemark
	sizeValue := req.Size
	isValidSize := producthelpers.CheckIsValidSize(sizeValue)

	if !isValidSize {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Invalid Size Value",
				ErrorMessage: "Size must be only: 'XXS', 'XS', 'S', 'M', 'L', 'XL', 'XXL', 'XXXL', 'XXXXL', 'XXXXXL', 'Free Size', 'Other'",
			},
			Data: nil,
		}, nil
	}

	if !isHaveProductCode {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Product Code is empty",
			},
			Data: nil,
		}, nil
	}

	if isSizeRemarkRequired {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "You selected size Other, This size is required Size Remark field",
			},
			Data: nil,
		}, nil
	}

	// Create stock object
	stock := &entities.Stock{
		Id:          commonhelpers.GenerateUUID(),
		ProductId:   req.ProductId,
		Size:        req.Size,
		SizeRemark:  req.SizeRemark,
		Quantity:    req.Quantity,
		PreQuantity: req.PreQuantity,
		Price:       req.Price,
		UsdPrice:    req.UsdPrice,
		Currency:    req.Currency,
		ItemStatus:  assignVaraintStatus(checkIsOutOfStock(req.Quantity)),
		CreatedAt:   commonhelpers.GetCurrentTimeISO(),
		UpdatedAt:   commonhelpers.GetCurrentTimeISO(),
	}

	// Save stock to db
	result := p.DB.Table("Stock").Create(stock)
	if result.Error != nil {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to save stock to db",
				ErrorMessage: result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	return &productsresponses.CreatedStockAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to create stock!",
		},
		Data: &productsresponses.StockAdminResponse{
			Id:          stock.Id,
			ProductId:   stock.ProductId,
			Size:        stock.Size,
			SizeRemark:  stock.SizeRemark,
			Quantity:    stock.Quantity,
			PreQuantity: stock.PreQuantity,
			Price:       stock.Price,
			ItemStatus:  stock.ItemStatus,
			CreatedAt:   stock.CreatedAt,
			UpdatedAt:   stock.UpdatedAt,
		},
	}, nil
}

func (p *productsRepositoryV2Impl) UpdateStock(req *productsdto.UpdateStockRequest) (*productsresponses.UpdateStockResponse, error) {
	// Check if StockId is provided
	if req.StockId == "" {
		return &productsresponses.UpdateStockResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Stock Id is empty",
			},
		}, nil
	}

	// Find the stock by stock id
	stock := &entities.Stock{}
	result := p.DB.Table("Stock").Where("id = ?", req.StockId).First(stock)
	if result.Error != nil {
		return &productsresponses.UpdateStockResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to find stock by stock id",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// Prepare the updates
	updates := map[string]interface{}{
		"quantity":     req.Quantity,
		"pre_quantity": req.PreQuantity,
		"price":        req.Price,
		"rrp_price":    req.RrpPrice,
		"usd_price":    req.UsdPrice,
		"currency":     req.Currency,
		"size_remark":  req.SizeRemark,
		"item_status":  assignVaraintStatus(checkIsOutOfStock(req.Quantity)),
		"updated_at":   commonhelpers.GetCurrentTimeISO(),
	}

	// Update the stock in the database
	result = p.DB.Table("Stock").Where("id = ?", req.StockId).Updates(updates)
	if result.Error != nil {
		return &productsresponses.UpdateStockResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update stock to db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// Fetch the updated stock details to ensure response has the latest data
	updatedStock := &entities.Stock{}
	result = p.DB.Table("Stock").Where("id = ?", req.StockId).First(updatedStock)
	if result.Error != nil {
		return &productsresponses.UpdateStockResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch updated stock from db",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// Return the response with updated data
	return &productsresponses.UpdateStockResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to update stock!",
		},
		Data: productsresponses.StockAdminResponse{
			Id:          updatedStock.Id,
			ProductId:   updatedStock.ProductId,
			Size:        updatedStock.Size,
			SizeRemark:  updatedStock.SizeRemark,
			Quantity:    updatedStock.Quantity,
			PreQuantity: updatedStock.PreQuantity,
			Price:       updatedStock.Price,
			UsdPrice:    updatedStock.UsdPrice,
			Currency:    updatedStock.Currency,
			ItemStatus:  updatedStock.ItemStatus,
			CreatedAt:   updatedStock.CreatedAt,
			UpdatedAt:   updatedStock.UpdatedAt,
		},
	}, nil
}

func (p *productsRepositoryV2Impl) DecreaseStock(req *productsdto.DecreaseStockRequest) (*productsresponses.DecreaseStockResponse, error) {
	isValidStockId, err := checkIsValidStockId(req.StockId, p.DB)
	if err != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check valid stock id",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	if !isValidStockId {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "Stock not found",
				ErrorMessage: "Stock not found by stock id: " + req.StockId,
			},
			Data: nil,
		}, nil
	}

	// find stock by stock id and assign to stock object
	stock := &entities.Stock{}
	result := p.DB.Table("Stock").Where("id = ?", req.StockId).First(stock)
	if result.Error != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to find stock by stock id",
				ErrorMessage: result.Error.Error(),
			},
			Data: nil,
		}, result.Error
	}

	// check quantity is valid
	isValidQuantity := req.DecreaseQuantity <= stock.Quantity
	if !isValidQuantity {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Invalid Quantity Value",
				ErrorMessage: "Quantity must be less than or equal to current stock quantity",
			},
			Data:            nil,
			CurrentQuantity: stock.Quantity,
		}, nil
	}

	stockIsNotEnough := req.DecreaseQuantity > stock.Quantity
	if stockIsNotEnough {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Stock is not enough",
				ErrorMessage: "Stock is not enough to decrease",
			},
			Data:            nil,
			CurrentQuantity: stock.Quantity,
		}, nil
	}

	// update new stock data to db
	updates := map[string]interface{}{
		"quantity":    stock.Quantity - req.DecreaseQuantity,
		"item_status": assignVaraintStatus(checkIsOutOfStock(stock.Quantity - req.DecreaseQuantity)),
		"updated_at":  commonhelpers.GetCurrentTimeISO(),
	}

	result = p.DB.Table("Stock").Where("id = ?", req.StockId).Updates(updates)

	if result.Error != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to update stock to db",
				ErrorMessage: result.Error.Error(),
			},
			Data:            nil,
			CurrentQuantity: stock.Quantity,
		}, result.Error
	}

	// fetch updated stock data and update to response
	updatedStock := &entities.Stock{}
	result = p.DB.Table("Stock").Where("id = ?", req.StockId).First(updatedStock)
	if result.Error != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch updated stock data",
				ErrorMessage: result.Error.Error(),
			},
			Data:            nil,
			CurrentQuantity: stock.Quantity,
		}, result.Error
	}

	// find product code in Stocks table by product code
	productId := updatedStock.ProductId
	product := &entities.Product{}

	result = p.DB.Table("Products").Where("id = ?", productId).First(product)
	if result.Error != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to find product by product code",
				ErrorMessage: result.Error.Error(),
			},
			Data:            nil,
			CurrentQuantity: stock.Quantity,
		}, result.Error
	}

	// update product status to out_of_stock if quantity is 0
	if updatedStock.Quantity == 0 {
		updates := map[string]interface{}{
			"item_status": "out_of_stock",
			"updated_at":  commonhelpers.GetCurrentTimeISO(),
		}

		result = p.DB.Table("Products").Where("id = ?", productId).Updates(updates)
		if result.Error != nil {
			return &productsresponses.DecreaseStockResponse{
				Status: &commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to update product status to out_of_stock",
					ErrorMessage: result.Error.Error(),
				},
				Data:            nil,
				CurrentQuantity: stock.Quantity,
			}, result.Error
		}
	}

	return &productsresponses.DecreaseStockResponse{
		Status: &commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to decrease stock!",
		},
		Data: &productsresponses.StockAdminResponse{
			Id:          req.StockId,
			ProductId:   updatedStock.ProductId,
			Size:        updatedStock.Size,
			SizeRemark:  updatedStock.SizeRemark,
			Quantity:    updatedStock.Quantity,
			PreQuantity: updatedStock.PreQuantity,
			CreatedAt:   updatedStock.CreatedAt,
			UpdatedAt:   updatedStock.UpdatedAt,
		},
		CurrentQuantity: updatedStock.Quantity,
	}, nil
}

func (p *productsRepositoryV2Impl) GetAllProductsAdmin(role string, companyId string) (*productsresponses.GetAllProductsAdminResponse, error) {

	var products []productsresponses.ProductEssentialDetailAdminResponse
	var result *gorm.DB
	if role == "SuperAdmin" {
		result = p.DB.Table("(SELECT MAX(id) as id, master_code FROM Products WHERE use_as_primary_data = ? GROUP BY master_code) AS unique_products", 1).
			Joins("INNER JOIN Products ON Products.id = unique_products.id").
			Select("Products.id, Products.name, Products.product_code, Products.master_code, Products.color_code, Products.product_status, Products.cover_image, Products.front_image, Products.back_image, Products.price, Products.product_group, Products.season, Products.gender, Products.product_class, Products.collection, Products.category, Products.brand, Products.is_club, Products.club_name, Products.remark, Products.launch_date, Products.end_of_life, Products.size_chart, Products.pack_size, Products.current_supplier, Products.description, Products.fabric_content, Products.fabric_type, Products.weight, Products.created_by_company, Products.created_by, Products.created_at, Products.updated_at, Products.use_as_primary_data").
			Order("Products.created_at DESC").
			Find(&products)
	} else {
		// filter if same master code then not appear in the list
		result = p.DB.Table("(SELECT MAX(id) as id, master_code FROM Products WHERE use_as_primary_data = ? AND created_by_company = ? GROUP BY master_code) AS unique_products", 1, companyId).
			Joins("INNER JOIN Products ON Products.id = unique_products.id").
			Select("Products.id, Products.name, Products.product_code, Products.master_code, Products.color_code, Products.product_status, Products.cover_image, Products.front_image, Products.back_image, Products.price, Products.product_group, Products.season, Products.gender, Products.product_class, Products.collection, Products.category, Products.brand, Products.is_club, Products.club_name, Products.remark, Products.launch_date, Products.end_of_life, Products.size_chart, Products.pack_size, Products.current_supplier, Products.description, Products.fabric_content, Products.fabric_type, Products.weight, Products.created_by_company, Products.created_by, Products.created_at, Products.updated_at, Products.use_as_primary_data").
			Order("Products.created_at DESC").
			Find(&products)
	}

	if result.Error != nil {
		return &productsresponses.GetAllProductsAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch all products",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}
	for i := range products {
		details, err := p.GetProductDetail(products[i].MasterCode)
		if err != nil {
			fmt.Println("Error fetching variants:", err)
			continue
		}
		minPrice := float64(0)
		currency := ""

		for _, variant := range details.ProductVaraints {
			for _, stock := range variant.Stock {
				if stock.Price > 0 && (minPrice == 0 || stock.Price < minPrice) {
					minPrice = stock.Price
					currency = stock.Currency
				}
			}
		}
		products[i].Price = minPrice
		products[i].Currency = currency
	}
	return &productsresponses.GetAllProductsAdminResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch all products!",
		},
		Products: products,
	}, nil
}

func (p *productsRepositoryV2Impl) GetAllProducts() (*productsresponses.GetAllProductsResponse, error) {
	var products []productsresponses.ProductEssentialDetailAdminResponse

	// Step 1: Fetch unique products based on the master code.
	result := p.DB.Table("(SELECT MAX(id) as id, master_code FROM Products WHERE use_as_primary_data = ? GROUP BY master_code) AS unique_products", 1).
		Joins("INNER JOIN Products ON Products.id = unique_products.id").
		Select("Products.id, Products.name, Products.product_code, Products.master_code, Products.color_code, Products.product_status, Products.cover_image, Products.front_image, Products.back_image, Products.price, Products.product_group, Products.season, Products.gender, Products.product_class, Products.collection, Products.category, Products.brand, Products.is_club, Products.club_name, Products.remark, Products.launch_date, Products.end_of_life, Products.size_chart, Products.pack_size, Products.current_supplier, Products.description, Products.fabric_content, Products.fabric_type, Products.weight, Products.created_by_company, Products.created_by, Products.created_at, Products.updated_at, Products.use_as_primary_data").
		Order("Products.created_at DESC").
		Find(&products)

	if result.Error != nil {
		return &productsresponses.GetAllProductsResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch all products",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// Step 2: Collect all product variants for each master code.
	var masterCodes []string
	for _, product := range products {
		masterCodes = append(masterCodes, product.MasterCode)
	}

	var productVariants []entities.Product
	if len(masterCodes) > 0 {
		result = p.DB.Table("Products").
			Where("master_code IN (?)", masterCodes).
			Find(&productVariants)

		if result.Error != nil {
			return &productsresponses.GetAllProductsResponse{
				Status: commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to fetch product variants",
					ErrorMessage: result.Error.Error(),
				},
			}, result.Error
		}
	}

	// Step 3: Fetch stock information for each product.
	var stockItems []entities.Stock
	result = p.DB.Table("Stock").
		Where("product_id IN (?)", getProductIds(productVariants)).
		Find(&stockItems)

	if result.Error != nil {
		return &productsresponses.GetAllProductsResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch stock items",
				ErrorMessage: result.Error.Error(),
			},
		}, result.Error
	}

	// Step 4: Map data to GetProductDataField and collect color codes and sizes.
	productDataFields := make([]productsresponses.ProductEssentialDetailResponse, 0)
	for _, product := range products {
		colorCodes := producthelpers.CollectColorCodes(product.MasterCode, productVariants)
		sizeRange := producthelpers.CollectAndSortSizes(product.Id, stockItems)

		productDataField := productsresponses.ProductEssentialDetailResponse{
			Id:               product.Id,
			Name:             product.Name,
			MasterCode:       product.MasterCode,
			ProductStatus:    product.ProductStatus,
			CoverImage:       product.CoverImage,
			Colors:           colorCodes,
			SizeRange:        sizeRange,
			CreatedByCompany: product.CreatedByCompany,
			Collection:       product.Collection,
			Category:         product.Category,
		}

		productDataFields = append(productDataFields, productDataField)
	}

	return &productsresponses.GetAllProductsResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch all products!",
		},
		Products: productDataFields,
	}, nil
}

func getProductIds(products []entities.Product) []string {
	var ids []string
	for _, product := range products {
		ids = append(ids, product.Id)
	}
	return ids
}

// FetchProducts fetches all products from the database
func (p *productsRepositoryV2Impl) fetchProducts() ([]entities.Product, error) {
	var products []entities.Product
	result := p.DB.Table("Products").Find(&products)
	return products, result.Error
}

func (p *productsRepositoryV2Impl) fetchProductVariants() ([]entities.Stock, error) {
	var productVariants []entities.Stock
	result := p.DB.Table("Stock").Find(&productVariants)
	return productVariants, result.Error
}

func (p *productsRepositoryV2Impl) GetProductDetailAdmin(masterCode string) (*productsresponses.GetProductDetailAdminResponse, error) {
	// Check required field
	if masterCode == "" {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, nil
	}

	// Check if the product exists by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	if !checkProductExist.IsProductExist {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, nil
	}

	// Fetch main product detail from db by master code
	productDetail, err := getMainProductDetailPrimaryDataByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch product detail from db",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	// Fetch all product variants by master code
	products, err := getAllProductsByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch all products by master code",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	// Prepare response for product variants
	productVariantsField := make([]*productsresponses.ProductVaraintAdminResponse, 0)
	for _, product := range products {
		// Fetch stocks specific to the current product variant
		var stockEntities []entities.Stock
		var color []entities.Colors
		result := p.DB.Table("Stock").Where("product_id = ?", product.Id).Find(&stockEntities)
		if result.Error != nil {
			return &productsresponses.GetProductDetailAdminResponse{
				Status: &commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to fetch stock by product code",
					ErrorMessage: result.Error.Error(),
				},
				MainProductData: nil,
				ProductVaraints: nil,
			}, result.Error
		}

		resultColor := p.DB.Table("Colors").Where("code = ?", product.ColorCode).First(&color)
		if resultColor.Error != nil {
			return &productsresponses.GetProductDetailAdminResponse{
				Status: &commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to fetch color by code",
					ErrorMessage: resultColor.Error.Error(),
				},
				MainProductData: nil,
				ProductVaraints: nil,
			}, resultColor.Error
		}

		// Map stock data to response format
		stocks := make([]*productsresponses.StockAdminResponse, len(stockEntities))
		for i, stock := range stockEntities {
			stocks[i] = &productsresponses.StockAdminResponse{
				Id:          stock.Id,
				ProductId:   stock.ProductId,
				Size:        stock.Size,
				SizeRemark:  stock.SizeRemark,
				Quantity:    stock.Quantity,
				PreQuantity: stock.PreQuantity,
				Price:       stock.Price,
				RrpPrice:    stock.RrpPrice,
				UsdPrice:    stock.UsdPrice,
				Currency:    stock.Currency,
				ItemStatus:  stock.ItemStatus,
				CreatedAt:   stock.CreatedAt,
				UpdatedAt:   stock.UpdatedAt,
			}
		}

		// Sort stocks based on size order
		sort.Slice(stocks, func(i, j int) bool {
			return producthelpers.SizeOrder[stocks[i].Size] < producthelpers.SizeOrder[stocks[j].Size]
		})

		// Create the product variant response field
		productVariantField := &productsresponses.ProductVaraintAdminResponse{
			ProductId:        product.Id,
			ProductCode:      product.ProductCode,
			ColorCode:        color[0].Name,
			FrontImage:       product.FrontImage,
			BackImage:        product.BackImage,
			Price:            float64(product.Price),
			UseAsPrimaryData: product.UseAsPrimaryData,
			Stock:            stocks,
		}

		productVariantsField = append(productVariantsField, productVariantField)
	}

	// Create and return the response
	return &productsresponses.GetProductDetailAdminResponse{
		Status: &commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch product detail!",
		},
		MainProductData: &productsresponses.MainProductDetailResponse{
			Id:               productDetail.Id,
			Name:             productDetail.Name,
			CoverImage:       productDetail.CoverImage,
			MasterCode:       productDetail.MasterCode,
			ProductStatus:    productDetail.ProductStatus,
			ProductGroup:     productDetail.ProductGroup,
			Season:           productDetail.Season,
			Gender:           productDetail.Gender,
			ProductClass:     productDetail.ProductClass,
			Collection:       productDetail.Collection,
			Category:         productDetail.Category,
			Brand:            productDetail.Brand,
			IsClub:           productDetail.IsClub,
			ClubName:         productDetail.ClubName,
			Remark:           productDetail.Remark,
			LaunchDate:       productDetail.LaunchDate,
			EndOfLife:        productDetail.EndOfLife,
			SizeChart:        productDetail.SizeChart,
			PackSize:         productDetail.PackSize,
			CurrentSupplier:  productDetail.CurrentSupplier,
			Description:      productDetail.Description,
			FabricContent:    productDetail.FabricContent,
			FabricType:       productDetail.FabricType,
			Weight:           productDetail.Weight,
			CreatedByCompany: productDetail.CreatedByCompany,
			CreatedBy:        productDetail.CreatedBy,
			EditedBy:         productDetail.EditedBy,
			CreatedAt:        productDetail.CreatedAt,
			UpdatedAt:        productDetail.UpdatedAt,
		},
		ProductVaraints: productVariantsField,
	}, nil
}

func (p *productsRepositoryV2Impl) GetProductDetail(masterCode string) (*productsresponses.GetProductDetailResponse, error) {
	// Check required field
	if masterCode == "" {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Required field is empty",
				ErrorMessage: "Master Code is empty",
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, nil
	}

	// Check if the product exists by master code
	checkProductExist, err := p.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to check product exist by master code",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	if !checkProductExist.IsProductExist {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "Product not found",
				ErrorMessage: "Product not found by master code: " + masterCode,
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, nil
	}

	// Fetch main product detail from db by master code
	productDetail, err := getMainProductDetailPrimaryDataByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch product detail from db",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	// Fetch all product variants by master code
	products, err := getAllProductsByMasterCode(p.DB, masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "REPO - Failed to fetch all products by master code",
				ErrorMessage: err.Error(),
			},
			MainProductData: nil,
			ProductVaraints: nil,
		}, err
	}

	// Prepare response for product variants
	productVariantsField := make([]*productsresponses.ProductVaraintWebViewResponse, 0)
	for _, product := range products {
		// Fetch stocks specific to the current product variant
		var stockEntities []entities.Stock
		result := p.DB.Table("Stock").Where("product_id = ?", product.Id).Find(&stockEntities)
		if result.Error != nil {
			return &productsresponses.GetProductDetailResponse{
				Status: commonresponse.CommonResponse{
					Status:       false,
					Message:      "REPO - Failed to fetch stock by product code",
					ErrorMessage: result.Error.Error(),
				},
				MainProductData: nil,
				ProductVaraints: nil,
			}, result.Error
		}

		// Map stock data to response format
		stocks := make([]*productsresponses.StockWebviewResponse, len(stockEntities))
		for i, stock := range stockEntities {
			stocks[i] = &productsresponses.StockWebviewResponse{
				Id:          stock.Id,
				ProductId:   stock.ProductId,
				Size:        stock.Size,
				SizeRemark:  stock.SizeRemark,
				Price:       stock.Price,
				RrpPrice:    stock.RrpPrice,
				Currency:    stock.Currency,
				Quantity:    stock.Quantity,
				PreQuantity: stock.PreQuantity,
				ItemStatus:  stock.ItemStatus,
			}
		}

		// Sort stocks based on size order
		sort.Slice(stocks, func(i, j int) bool {
			return producthelpers.SizeOrder[stocks[i].Size] < producthelpers.SizeOrder[stocks[j].Size]
		})

		// Create the product variant response field
		productVariantField := &productsresponses.ProductVaraintWebViewResponse{
			ProductId:   product.Id,
			ProductCode: product.ProductCode,
			ColorCode:   product.ColorCode,
			FrontImage:  product.FrontImage,
			BackImage:   product.BackImage,
			Stock:       stocks,
		}

		productVariantsField = append(productVariantsField, productVariantField)
	}

	// Create and return the response
	return &productsresponses.GetProductDetailResponse{
		Status: commonresponse.CommonResponse{
			Status:  true,
			Message: "Success to fetch product detail!",
		},
		MainProductData: &productsresponses.MainProductDetailResponse{
			Id:               productDetail.Id,
			Name:             productDetail.Name,
			CoverImage:       productDetail.CoverImage,
			MasterCode:       productDetail.MasterCode,
			ProductStatus:    productDetail.ProductStatus,
			ProductGroup:     productDetail.ProductGroup,
			Season:           productDetail.Season,
			Gender:           productDetail.Gender,
			ProductClass:     productDetail.ProductClass,
			Collection:       productDetail.Collection,
			Category:         productDetail.Category,
			Brand:            productDetail.Brand,
			IsClub:           productDetail.IsClub,
			ClubName:         productDetail.ClubName,
			Remark:           productDetail.Remark,
			LaunchDate:       productDetail.LaunchDate,
			EndOfLife:        productDetail.EndOfLife,
			SizeChart:        productDetail.SizeChart,
			PackSize:         productDetail.PackSize,
			CurrentSupplier:  productDetail.CurrentSupplier,
			Description:      productDetail.Description,
			FabricContent:    productDetail.FabricContent,
			FabricType:       productDetail.FabricType,
			Weight:           productDetail.Weight,
			CreatedByCompany: productDetail.CreatedByCompany,
			CreatedBy:        productDetail.CreatedBy,
			EditedBy:         productDetail.EditedBy,
			CreatedAt:        productDetail.CreatedAt,
			UpdatedAt:        productDetail.UpdatedAt,
		},
		ProductVaraints: productVariantsField,
	}, nil

}

func fetchImageFromDbByMasterCode(db *gorm.DB, masterCode string, table string, selectedField string) (string, error) {
	// Define a struct to hold the result
	var image = ""

	// Fetch image from db by master code and use_as_primary_data = true
	result := db.Table(table).
		Where("master_code = ? AND use_as_primary_data = ?", masterCode, true).
		Select(selectedField).
		Scan(&image)

	// Return the fetched image and any error encountered
	return image, result.Error
}

func checkProductVaraintExistByProductId(productId string, db *gorm.DB) (bool, error) {
	productData := db.Table("Products").Where("id = ?", productId).First(&entities.Product{})

	print("eiei productId: ", productId)
	if productData.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func checkIsValidStockId(stockId string, db *gorm.DB) (bool, error) {
	stockData := db.Table("Stock").Where("id = ?", stockId).First(&entities.Stock{})

	if stockData.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func updateAllOldMasterCodeToNewMasterCode(db *gorm.DB, oldMasterCode string, newMasterCode string, updatedBy string) error {
	// update product master code
	result := db.Table("Products").
		Where("master_code = ?", oldMasterCode).
		Updates(map[string]interface{}{
			"master_code": newMasterCode,
			"updated_at":  commonhelpers.GetCurrentTimeISO(),
			"updated_by":  updatedBy,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getMainProductDetailPrimaryDataByMasterCode(db *gorm.DB, masterCode string) (*entities.Product, error) {
	product := &entities.Product{}
	result := db.Table("Products").
		Where("master_code = ? AND use_as_primary_data = ?", masterCode, true).
		First(product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func getUniqueProductByMasterCode(db *gorm.DB) ([]*entities.Product, error) {
	products := []*entities.Product{}
	result := db.Table("(SELECT MAX(id) as id, master_code FROM Products WHERE use_as_primary_data = ? GROUP BY master_code) AS unique_products", 1).
		Joins("INNER JOIN Products ON Products.id = unique_products.id").
		Select("Products.id, Products.name, Products.product_code, Products.master_code, Products.color_code, Products.product_status, Products.cover_image, Products.front_image, Products.back_image, Products.price, Products.product_group, Products.season, Products.gender, Products.product_class, Products.collection, Products.category, Products.brand, Products.is_club, Products.club_name, Products.remark, Products.launch_date, Products.end_of_life Products.size_chart, Products.pack_size, Products.current_supplier, Products.description, Products.fabric_content, Products.fabric_type, Products.weight, Products.created_by_company, Products.created_by, Products.created_at, Products.updated_at, Products.use_as_primary_data").
		Order("Products.created_at DESC").
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func getAllProductsByMasterCode(db *gorm.DB, masterCode string) ([]entities.Product, error) {
	products := []entities.Product{}
	result := db.Table("Products").
		Where("master_code = ?", masterCode).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func getStockByProductId(db *gorm.DB, productId string) ([]entities.Stock, error) {
	stocks := []entities.Stock{}
	result := db.Table("Stock").
		Where("product_id = ?", productId).
		Find(&stocks)

	if result.Error != nil {
		return nil, result.Error
	}

	return stocks, nil
}

func saveLog(db *gorm.DB, log *entities.Log) {
	db.Table("Log").Create(&log)
}
