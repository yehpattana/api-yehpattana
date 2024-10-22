package productscontrollers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/auth/pkg"
	commonhelpers "github.com/yehpattana/api-yehpattana/modules/commons/common_helpers"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	productsdto "github.com/yehpattana/api-yehpattana/modules/products/products_dto"
	productsservices "github.com/yehpattana/api-yehpattana/modules/products/products_services"
)

type productControllerErrorCode string

const (
	cantGetAllProducts          productControllerErrorCode = "PRODUCT_ERR_001"
	cantGetProductDetail        productControllerErrorCode = "PRODUCT_ERR_002"
	cantChangeProductStatus     productControllerErrorCode = "PRODUCT_ERR_003"
	cantCreateProduct           productControllerErrorCode = "PRODUCT_ERR_004"
	cantUpdateProduct           productControllerErrorCode = "PRODUCT_ERR_005"
	cantUploadCoverImage        productControllerErrorCode = "PRODUCT_ERR_006"
	cantUploadSizeChart         productControllerErrorCode = "PRODUCT_ERR_007"
	cantEditCoverImage          productControllerErrorCode = "PRODUCT_ERR_008"
	cantEditSizeChart           productControllerErrorCode = "PRODUCT_ERR_009"
	cantDeleteProduct           productControllerErrorCode = "PRODUCT_ERR_010"
	cantUpdateMainProductDetail productControllerErrorCode = "PRODUCT_ERR_011"
	cantUpdateStock             productControllerErrorCode = "PRODUCT_ERR_012"
	cantDecreaseStock           productControllerErrorCode = "PRODUCT_ERR_013"
	cantCreateStock             productControllerErrorCode = "PRODUCT_ERR_014"
	cantUpdateMasterCode        productControllerErrorCode = "PRODUCT_ERR_015"
	cantUpdateProductVariant    productControllerErrorCode = "PRODUCT_ERR_016"
)

type ProductControllerV2Interface interface {
	CheckIsProductExistByMasterCode(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	CreateStock(c *fiber.Ctx) error
	UpdateStock(c *fiber.Ctx) error
	DecreaseStock(c *fiber.Ctx) error
	GetCoverImageByMasterCode(c *fiber.Ctx) error
	GetSizeChartByMasterCode(c *fiber.Ctx) error
	UploadCoverImage(c *fiber.Ctx) error
	UploadSizeChart(c *fiber.Ctx) error
	UpdateCoverImage(c *fiber.Ctx) error
	UpdateSizeChart(c *fiber.Ctx) error
	UpdateMainProductDetailByMasterCode(c *fiber.Ctx) error
	UpdateMasterCode(c *fiber.Ctx) error
	UpdateProductVariant(c *fiber.Ctx) error
	GetAllAdminProducts(c *fiber.Ctx) error
	GetProductDetailAdmin(c *fiber.Ctx) error
	DeleteAllProductInMasterCode(c *fiber.Ctx) error
	DeleteProductVariant(c *fiber.Ctx) error

	// Common Side
	GetAllProductsByCompany(c *fiber.Ctx) error

	// Web View
	GetAllProducts(c *fiber.Ctx) error
	GetProductDetail(c *fiber.Ctx) error
}

func ProductControllerV2Impl(
	config configs.ConfigInterface,
	productService productsservices.ProductServiceV2Interface,
) ProductControllerV2Interface {
	return &productControllerV2Impl{
		config:         config,
		productService: productService,
	}
}

type productControllerV2Impl struct {
	config         configs.ConfigInterface
	productService productsservices.ProductServiceV2Interface
}

// Get All Products By Company godoc
// @Summary ดึงข้อมูล product ทั้งหมดโดยใช้ company id
// @Description Get all products by company id
// @Router /v1/products/admin/get-all-products-by-conmpany-id/{companyID} [get]
// @Param companyID path string true "Company id"
// @Produce json
// @Tags Products-Get
// @Security ApiKeyAuth
func (p *productControllerV2Impl) GetAllProductsByCompany(c *fiber.Ctx) error {
	companyID := c.Params("companyID")
	response, err := p.productService.GetAllProductByCompany(companyID)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetAllProducts), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// DeleteProductVariant godoc
// @Summary ลบ product variant
// @Description ลบ product variant และ stock ของ product variant นั้นทั้งหมด
// @Router /v1/products/admin/delete-product-variant/{productId} [delete]
// @Param productId path string true "Product id"
// @Produce json
// @Tags Products-Delete
// @Security ApiKeyAuth
func (p *productControllerV2Impl) DeleteProductVariant(c *fiber.Ctx) error {
	productId := c.Params("productId")
	response, err := p.productService.DeleteProductVaraint(productId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantDeleteProduct), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// CheckIsProductExistByMasterCode godoc
// @Summary ลบสินค้าทั้งหมดใน master code
// @Description ลบสินค้าทั้งหมดใน master code
// @Router /v1/products/admin/delete-all-product-in-master-code/{masterCode} [delete]
// @Param masterCode path string true "Master code"
// @Produce json
// @Tags Products-Delete
// @Security ApiKeyAuth
func (p *productControllerV2Impl) DeleteAllProductInMasterCode(c *fiber.Ctx) error {
	masterCode := c.Params("masterCode")
	formattedMasterCode := commonhelpers.ReplacePercent20WithSpace(masterCode)

	response, err := p.productService.DeleteAllProductInMasterCode(formattedMasterCode)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantDeleteProduct), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetProductDetail godoc
// @Summary ดึงข้อมูล product detail และ varaints ทั้งหมด
// @Description Get product detail
// @Router /v1/products/{masterCode} [get]
// @Param masterCode path string true "Master code"
// @Produce json
// @Tags Webview-Products
func (p *productControllerV2Impl) GetProductDetail(c *fiber.Ctx) error {
	masterCode := c.Params("masterCode")
	formatParams, err := url.QueryUnescape(masterCode)
	if err != nil {
		fmt.Println("Error decoding string:", err)
	}
	response, err := p.productService.GetProductDetail(formatParams)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetProductDetail), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetAllProducts godoc
// @Summary ดึงข้อมูล product ทั้งหมด
// @Description Get all products
// @Router /v1/products [get]
// @Produce json
// @Tags Webview-Products
func (p *productControllerV2Impl) GetAllProducts(c *fiber.Ctx) error {
	response, err := p.productService.GetAllProducts()
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetAllProducts), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateProductVariant godoc
// @Summary อัพเดท product variant
// @Description Update product variant
// @Router /v1/products/admin/update-product-variant [patch]
// @Param product_id formData string true "Product id"
// @Param product_code formData string false "Product code (ถ้าไม่ส่งมาจะใช้ code เดิม)"
// @Param color_code formData string false "Color code (ถ้าไม่ส่งมาจะใช้สีเดิิม)"
// @Param front_image formData file false "Front image"
// @Param back_image formData file false "Back image"
// @Param price formData string true "Price"
// @Param use_as_primary_data formData boolean true "Use as primary data (ถ้าเลือกเป็น true varaint ที่เหลือจะปรับเป็น false)"
// @Produce json
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateProductVariant(c *fiber.Ctx) error {
	// Parse form values
	productId := c.FormValue("product_id")
	productCode := c.FormValue("product_code")
	colorCode := c.FormValue("color_code")

	// Optional fields
	priceStr := c.FormValue("price")
	useAsPrimaryData := c.FormValue("use_as_primary_data") == "true"

	// Validate required field
	if productId == "" {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUpdateProductVariant), "Product ID is required").Res()
	}

	// Prepare request object
	req := &productsdto.UpdateProductVariantRequest{
		ProductId:        productId,
		UseAsPrimaryData: useAsPrimaryData,
		ProductCode:      productCode,
		ColorCode:        colorCode,
	}

	// Parse optional price
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUpdateProductVariant), "Invalid price format").Res()
		}
		req.Price = price
	}

	// Handle optional front_image
	if frontImage, err := c.FormFile("front_image"); err == nil {
		req.FrontImage = frontImage
	} else if err.Error() != "there is no uploaded file associated with the given key" {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUpdateProductVariant), err.Error()).Res()
	}

	// Handle optional back_image
	if backImage, err := c.FormFile("back_image"); err == nil {
		req.BackImage = backImage
	} else if err.Error() != "there is no uploaded file associated with the given key" {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUpdateProductVariant), err.Error()).Res()
	}

	// Call the service to update the product variant
	response, err := p.productService.UpdateProductVaraint(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantUpdateProductVariant), err.Error()).Res()
	}

	// Return success response
	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetProductDetailAdmin godoc
// @Summary ดึงข้อมูล product detail และ varaints ทั้งหมด โดยใช้ master code
// @Description ไซส์หลังบ้านเรียง auto sort จากเล็กไปใหญ่ให้แล้วจ้า
// @Router /v1/products/admin/{masterCode} [get]
// @Param masterCode path string true "Master code"
// @Produce json
// @Tags Products-Get
// @Security ApiKeyAuth
func (p *productControllerV2Impl) GetProductDetailAdmin(c *fiber.Ctx) error {
	masterCode := c.Params("masterCode")
	formattedMasterCode := commonhelpers.ReplacePercent20WithSpace(masterCode)

	response, err := p.productService.GetProductDetailAdmin(formattedMasterCode)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetProductDetail), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetAllAdminProducts godoc
// @Summary ดึงข้อมูล product ทั้งหมด (admin)
// @Description นางจะ filter เอาข้อมูล master code ที่เซ็ท use as primary data = true มาเท่านั้น ซึ่งก็คือจะมีแค่อันเดียวคือ สินค้าอันแรกที่แอดไหปใน master code นั้น ถ้ามีหลายอัน ก็เอาอันแรกมาใช้แค่อันเดียวจ้า
// @Router /v1/products/all/admin [get]
// @Produce json
// @Tags Products-Get
// @Security ApiKeyAuth
func (p *productControllerV2Impl) GetAllAdminProducts(c *fiber.Ctx) error {
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	if token == "" {
		return commonresponse.NewResponse(c).Error(fiber.StatusUnauthorized, "Token is required", "Missing Authorization token").Res()
	}
	claims, _ := pkg.ParseToken(p.config.Jwt(), token)
	role := ""
	companyId := ""
	if claims != nil {
		companyId = claims.Claims.CompanyName
		role = claims.Claims.Role
	}
	response, err := p.productService.GetAllProductsAdmin(role, companyId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetAllProducts), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateMasterCode godoc
// @Summary อัพเดท master code
// @Description Update master code
// @Router /v1/products/admin/update-new-master-code [patch]
// @Accept json
// @Produce json
// @Param request body productsdto.UpdateMasterCodeRequest true "Request body"
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateMasterCode(c *fiber.Ctx) error {
	req := new(productsdto.UpdateMasterCodeRequest)

	isUpdatedByIsBlank := req.UpdatedBy == ""
	if isUpdatedByIsBlank {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid request format", "updated_by is required").Res()
	}

	isOldMasterCodeIsBlank := req.OldMasterCode == ""
	if isOldMasterCodeIsBlank {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid request format", "old_master_code is required").Res()
	}

	isNewMasterCodeIsBlank := req.NewMasterCode == ""
	if isNewMasterCodeIsBlank {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid request format", "new_master_code is required").Res()
	}

	// Parse JSON request body
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid request format", err.Error()).Res()
	}

	// Call service to update master code
	response, err := p.productService.UpdateMasterCode(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "Failed to update master code", err.Error()).Res()
	}

	// Return success response
	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateMainProductDetailByMasterCode godoc
// @Summary อัพเดท main product detail โดยใช้ master code
// @Description (อัพเดตข้อมูล master code หลักที่ set เป็น use as primary data)
// @Router /v1/products/admin/update-main-product-detail [patch]
// @Param master_code formData string true "Master code (โค้ดที่ต้องการเปลี่ยน main product detail ไม่ใช่่โค้ดใหม่ที่ต้องการเปลี่ยน)"
// @Param name formData string true "Name"
// @Param cover_image formData string false "Cover image (ส่งเป็น img url มาจ้า)"
// @Param product_status formData string true "Product status" Enums(available, hidden, out_of_stock)
// @Param product_group formData string false "Product group"
// @Param season formData string false "Season"
// @Param gender formData string true "gender" Enums(male, female, unisex, kids)
// @Param product_class formData string false "Product class"
// @Param collection formData string false "Collection"
// @Param category formData string false "Category"
// @Param brand formData string false "Brand"
// @Param is_club formData boolean false "Is club"
// @Param club_name formData string false "Club name"
// @Param remark formData string false "Remark"
// @Param launch_date formData string false "Launch date"
// @Param size_chart formData string false "Size chart (ส่งเป็น img url มาจ้า)"
// @Param pack_size formData string false "Pack size"
// @Param current_supplier formData string false "Current supplier"
// @Param description formData string false "Description"
// @Param fabric_content formData string false "Fabric content"
// @Param fabric_type formData string false "Fabric type"
// @Param weight formData string true "Weight"
// @Param created_by_company formData string true "Created by company"
// @Param edited_by formData string true "Edited by (ส่งเป็น id ของ user เด้อ)"
// @Produce json
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateMainProductDetailByMasterCode(c *fiber.Ctx) error {
	req := new(productsdto.UpdateMainProductDetailByMasterCodeRequest)

	masterCode := c.FormValue("master_code")
	req.MasterCode = masterCode
	req.Name = c.FormValue("name")
	req.CoverImage = c.FormValue("cover_image")
	req.ProductStatus = c.FormValue("product_status")
	req.ProductGroup = c.FormValue("product_group")
	req.Season = c.FormValue("season")
	req.Gender = c.FormValue("gender")
	req.ProductClass = c.FormValue("product_class")
	req.Collection = c.FormValue("collection")
	req.Category = c.FormValue("category")
	req.Brand = c.FormValue("brand")
	req.IsClub = c.FormValue("is_club") == "true"
	req.ClubName = c.FormValue("club_name")
	req.Remark = c.FormValue("remark")
	req.LaunchDate = c.FormValue("launch_date")
	req.EndOfLife = c.FormValue("end_of_life")
	req.SizeChart = c.FormValue("size_chart")
	req.PackSize = c.FormValue("pack_size")
	req.CurrentSupplier = c.FormValue("current_supplier")
	req.Description = c.FormValue("description")
	req.FabricContent = c.FormValue("fabric_content")
	req.FabricType = c.FormValue("fabric_type")
	weightStr := c.FormValue("weight")
	if weightStr == "" {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUpdateMainProductDetail), "weight is required").Res()
	}
	req.Weight, _ = strconv.ParseFloat(weightStr, 64)
	req.CreatedByCompany = c.FormValue("created_by_company")
	req.EditedBy = c.FormValue("edited_by")

	response, err := p.productService.UpdateMainProductDetailByMasterCode(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantUpdateMainProductDetail), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()

}

// DecreaseStock godoc
// @Summary ลดจำนวนสินค้าใน stock
// @Description Decrease stock
// @Router /v1/products/decrease-stock [patch]
// @Accept json
// @Produce json
// @Param request body productsdto.DecreaseStockRequest true "Request body"
// @Tags Products-Stock
// @Security ApiKeyAuth
func (p *productControllerV2Impl) DecreaseStock(c *fiber.Ctx) error {
	req := new(productsdto.DecreaseStockRequest)

	// Parse JSON request body
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid JSON format", err.Error()).Res()
	}

	// Validate required fields
	if req.DecreaseQuantity <= 0 {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid decrease quantity", "decrease_quantity must be greater than zero").Res()
	}

	response, err := p.productService.DecreaseStock(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantDecreaseStock), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateStock godoc
// @Summary อัพเดทจำนวนสินค้าใน stock
// @Description *ถ้าจะลดจำนวนสินค้าใน stock เคสที่ user ซื้อสินค้าให้ใช้ decrease-stock แทนเด้อ
// @Router /v1/products/update-stock [patch]
// @Accept json
// @Produce json
// @Param request body productsdto.UpdateStockRequest true "Request body"
// @Tags Products-Stock
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateStock(c *fiber.Ctx) error {
	req := new(productsdto.UpdateStockRequest)

	// Parse JSON request body
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid JSON format", err.Error()).Res()
	}

	if req.StockId == "" {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "StockId is required", "").Res()
	}

	// Call the service to update stock
	response, err := p.productService.UpdateStock(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "Cannot update stock", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// CreateStock godoc
// @Summary สร้าง stock
// @Description Create stock
// @Router /v1/products/create-stock [post]
// @Param product_id formData string true "Product id"
// @Param size formData string true "Size" enums(XXS, XS, S, M, L, XL, XXL, XXXL, XXXXL, XXXXXL, Free Size, Other)
// @Param price formData number true "Price"
// @Param size_remark formData string false "Size remark (ถ้าเลือกไซส์เป็น other บังคับว่าต้องใส่ size remark ด้วย)"
// @Param quantity formData string false "Quantity"
// @Produce json
// @Tags Products-Stock
// @Security ApiKeyAuth
func (p *productControllerV2Impl) CreateStock(c *fiber.Ctx) error {
	req := new(productsdto.CreateStockRequest)

	req.ProductId = c.FormValue("product_id")
	req.Size = c.FormValue("size")
	req.SizeRemark = c.FormValue("size_remark")
	req.Price, _ = strconv.ParseFloat(c.FormValue("price"), 64)
	req.Price, _ = strconv.ParseFloat(c.FormValue("rrp_price"), 64)

	quantityStr := c.FormValue("quantity")
	if quantityStr == "" {
		quantityStr = "0"
	}
	req.Quantity, _ = strconv.Atoi(quantityStr)
	preQuantityStr := c.FormValue("pre_quantity")
	if preQuantityStr == "" {
		preQuantityStr = "0"
	}
	req.PreQuantity, _ = strconv.Atoi(preQuantityStr)

	response, err := p.productService.CreateStock(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateStock), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateCoverImage godoc
// @Summary อัพเดท cover image โดยใช้ master code
// @Description Update cover image by master code
// @Router /v1/products/admin/update-cover-image [patch]
// @Accept json
// @Produce json
// @Param request body productsdto.UpdateCoverImageRequest true "Request body"
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateCoverImage(c *fiber.Ctx) error {
	req := new(productsdto.UpdateCoverImageRequest)

	// Parse JSON request body
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Invalid JSON format", err.Error()).Res()
	}

	// Validate required fields
	if req.MasterCode == "" || req.CoverImage == "" {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, "Master code and cover image URL are required", "").Res()
	}

	response, err := p.productService.UpdateCoverImageByMasterCode(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateStock), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UpdateSizeChart godoc
// @Summary อัพเดท size chart โดยใช้ master code
// @Description Update size chart by master code
// @Router /v1/products/admin/update-size-chart [patch]
// @Accept json
// @Produce json
// @Param request body productsdto.UpdateSizeChartRequest true "Request body"
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UpdateSizeChart(c *fiber.Ctx) error {
	req := new(productsdto.UpdateSizeChartRequest)
	sizeChart, err := c.FormFile("size_chart")
	if err != nil {
		fmt.Println(err)
	}
	masterCode := c.FormValue("master_code")
	if masterCode == "" {
		masterCode = ""
	}
	req.SizeChart = sizeChart
	req.MasterCode = masterCode

	response, err := p.productService.UpdateSizeChartByMasterCode(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateStock), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UploadSizeChart godoc
// @Summary อัพโหลด size chart
// @Description Upload size chart
// @Router /v1/products/admin/upload-size-chart [post]
// @Param size_chart formData file true "Size chart image file"
// @Produce json
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UploadSizeChart(c *fiber.Ctx) error {
	req := new(productsdto.UploadSizeChartRequest)

	sizeChart, err := c.FormFile("size_chart")
	if err != nil {
		sizeChart = nil
	}

	req.SizeChart = sizeChart

	response, err := p.productService.UploadSizeChart(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateProduct), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// UploadCoverImage godoc
// @Summary อัพโหลด cover image
// @Description Upload cover image
// @Router /v1/products/admin/upload-cover-image [post]
// @Param cover_image formData file true "Cover image file"
// @Produce json
// @Tags Products-Update
// @Security ApiKeyAuth
func (p *productControllerV2Impl) UploadCoverImage(c *fiber.Ctx) error {
	req := new(productsdto.UploadCoverImageRequest)

	coverImage, err := c.FormFile("cover_image")
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantUploadCoverImage), err.Error()).Res()
	}

	req.CoverImage = coverImage

	response, err := p.productService.UploadCoverImage(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateProduct), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetSizeChartByMasterCode godoc
// @Summary ดึงข้อมูล size chart โดยใช้ master code
// @Description Get size chart by master code
// @Router /v1/products/admin/get-size-chart/{masterCode} [get]
// @Param masterCode path string true "Master code"
// @Produce json
// @Tags Products-Get
// @Security ApiKeyAuth
func (p *productControllerV2Impl) GetSizeChartByMasterCode(c *fiber.Ctx) error {
	masterCode := c.Params("master_code")
	formattedMasterCode := commonhelpers.ReplacePercent20WithSpace(masterCode)

	response, err := p.productService.GetSizeChartByMasterCode(formattedMasterCode)
	if err != nil {
		print(err.Error())
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// GetCoverImageByMasterCode godoc
// @Summary ดึงข้อมูล cover image โดยใช้ master code
// @Description Get cover image by master code
// @Router /v1/products/admin/get-cover-image/{masterCode} [get]
// @Param masterCode path string true "Master code"
// @Tags Products-Get
// @Produce json
// @Security ApiKeyAuth
func (p *productControllerV2Impl) GetCoverImageByMasterCode(c *fiber.Ctx) error {
	masterCode := c.Params("master_code")
	formattedMasterCode := commonhelpers.ReplacePercent20WithSpace(masterCode)

	response, err := p.productService.GetCoverImageByMasterCode(formattedMasterCode)
	if err != nil {
		print(err.Error())
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// CheckIsProductExistByMasterCode godoc
// @Summary เชคว่า product นี้มีอยู่ในระบบหรือไม่ โดยใช้ master code
// @Description Check is product exist by master code
// @Router /v1/products/admin/check-product-exist/{masterCode} [get]
// @Param masterCode path string true "Master code"
// @Produce json
// @Tags Products-Get
// @Security ApiKeyAuth
func (p *productControllerV2Impl) CheckIsProductExistByMasterCode(c *fiber.Ctx) error {
	masterCode := c.Params("master_code")
	formattedMasterCode := commonhelpers.ReplacePercent20WithSpace(masterCode)

	response, err := p.productService.CheckIsProductExistByMasterCode(formattedMasterCode)
	if err != nil {
		print(err.Error())
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}

// CreateProduct godoc
// @Summary สร้าง product
// @Description Create product
// @Router /v1/products/create-product [post]
// @Param name formData string true "Name"
// @Param product_code formData string true "Product code"
// @Param master_code formData string true "Master code"
// @Param color_code formData string false "Color code"
// @Param product_status formData string true "Product status" Enums(available, hidden, out_of_stock)
// @Param cover_image formData string false "Cover image"
// @Param front_image formData file false "Front image"
// @Param back_image formData file false "Back image"
// @Param price formData string true "Price"
// @Param product_group formData string false "Product group"
// @Param season formData string false "Season"
// @Param gender formData string true "Gender" Enums(male, female, unisex, kids)
// @Param product_class formData string false "Product class"
// @Param collection formData string false "Collection"
// @Param category formData string false "Category"
// @Param brand formData string false "Brand"
// @Param is_club formData boolean false "Is club"
// @Param remark formData string false "Remark"
// @Param launch_date formData string false "Launch date"
// @Param size_chart formData string false "Size chart"
// @Param current_supplier formData string false "Current supplier"
// @Param description formData string false "Description"
// @Param fabric_content formData string false "Fabric content"
// @Param fabric_type formData string false "Fabric type"
// @Param weight formData string false "Weight"
// @Param created_by_company formData string true "Created by company"
// @Param created_by formData string true "Created by"
// @Produce json
// @Tags Products-Create
// @Security ApiKeyAuth
func (p *productControllerV2Impl) CreateProduct(c *fiber.Ctx) error {
	req := new(productsdto.CreateProductRequest)

	frontImage, err := c.FormFile("front_image")
	if err != nil {
		frontImage = nil
	}

	backImage, err := c.FormFile("back_image")
	if err != nil {
		backImage = nil
	}
	sizeChartImage, err := c.FormFile("size_chart")
	if err != nil {
		sizeChartImage = nil
	}

	weight, err := strconv.ParseFloat(c.FormValue("weight"), 64)
	if err != nil {
		weight = 0
	}

	priceStr := c.FormValue("price")
	if priceStr == "" {
		priceStr = "0"
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantCreateProduct), err.Error()).Res()
	}

	weightStr := c.FormValue("weight")
	if weightStr == "" {
		weightStr = "0"
	}

	req.Name = c.FormValue("name")
	req.ProductCode = c.FormValue("product_code")
	req.MasterCode = c.FormValue("master_code")
	req.ColorCode = c.FormValue("color_code")
	req.ProductStatus = c.FormValue("product_status")
	req.CoverImage = c.FormValue("cover_image")
	req.FrontImage = frontImage
	req.BackImage = backImage
	req.Price = price
	req.ProductGroup = c.FormValue("product_group")
	req.Season = c.FormValue("season")
	req.Gender = c.FormValue("gender")
	req.ProductClass = c.FormValue("product_class")
	req.Collection = c.FormValue("collection")
	req.Category = c.FormValue("category")
	req.Brand = c.FormValue("brand")
	req.IsClub = c.FormValue("is_club") == "true"
	req.Remark = c.FormValue("remark")
	req.LaunchDate = c.FormValue("launch_date")
	req.EndOfLife = c.FormValue("end_of_life")
	req.SizeChart = sizeChartImage
	req.CurrentSupplier = c.FormValue("current_supplier")
	req.Description = c.FormValue("description")
	req.FabricContent = c.FormValue("fabric_content")
	req.FabricType = c.FormValue("fabric_type")
	req.Weight = weight
	req.CreatedByCompany = c.FormValue("created_by_company")
	req.CreatedBy = c.FormValue("created_by")
	req.UseAsPrimaryData = c.FormValue("use_as_primary_data") == "true"

	// request body parser
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(cantCreateProduct), err.Error()).Res()
	}

	response, err := p.productService.CreateProduct(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantCreateProduct), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, response).Res()
}
