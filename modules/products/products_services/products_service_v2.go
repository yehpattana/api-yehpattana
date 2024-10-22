package productsservices

import (
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
	productsdto "github.com/yehpattana/api-yehpattana/modules/products/products_dto"
	productsresponses "github.com/yehpattana/api-yehpattana/modules/products/products_responses"
)

type ProductServiceV2Interface interface {
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
	DeleteProductVaraint(productId string) (*productsresponses.DeleteProductVariantResponse, error)

	// Common side
	GetAllProductByCompany(companyId string) (*productsresponses.GetAllProductByCompanyResponse, error)

	// Web view side
	GetAllProducts() (*productsresponses.GetAllProductsResponse, error)
	GetProductDetail(masterCode string) (*productsresponses.GetProductDetailResponse, error)
}

func ProductServiceV2Impl(
	cfg configs.ConfigInterface,
	productsRepository repositories.ProductsRepositoryV2Interface,
) ProductServiceV2Interface {
	return &productServiceV2Impl{
		config:             cfg,
		productsRepository: productsRepository,
	}
}

type productServiceV2Impl struct {
	config             configs.ConfigInterface
	productsRepository repositories.ProductsRepositoryV2Interface
}

func (p *productServiceV2Impl) GetAllProductByCompany(companyId string) (*productsresponses.GetAllProductByCompanyResponse, error) {
	result, err := p.productsRepository.GetAllProductByCompany(companyId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productServiceV2Impl) DeleteProductVaraint(productId string) (*productsresponses.DeleteProductVariantResponse, error) {
	result, err := p.productsRepository.DeleteProductVariant(productId)
	if err != nil {
		return &productsresponses.DeleteProductVariantResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to delete product varaint",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) DeleteAllProductInMasterCode(masterCode string) (*productsresponses.DeleteAllProductInMasterCodeResponse, error) {
	result, err := p.productsRepository.DeleteAllProductInMasterCode(masterCode)
	if err != nil {
		return &productsresponses.DeleteAllProductInMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to delete all product in master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateProductVaraint(req *productsdto.UpdateProductVariantRequest) (*productsresponses.UpdateProductVaraintResponse, error) {
	result, err := p.productsRepository.UpdateProductVaraint(req)
	if err != nil {
		return &productsresponses.UpdateProductVaraintResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update product varaint",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateMasterCode(req *productsdto.UpdateMasterCodeRequest) (*productsresponses.UpdateMasterCodeResponse, error) {
	result, err := p.productsRepository.UpdateMasterCode(req)
	if err != nil {
		return &productsresponses.UpdateMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) CheckIsProductExistByMasterCode(masterCode string) (*productsresponses.CheckIsProductExistByMasterCodeResponse, error) {
	result, err := p.productsRepository.CheckIsProductExistByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.CheckIsProductExistByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to check if product exist by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) CreateProduct(req *productsdto.CreateProductRequest) (*productsresponses.CreatedProductAdminResponse, error) {
	result, err := p.productsRepository.CreateProduct(req)
	if err != nil {
		return &productsresponses.CreatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to create product",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) CreateStock(req *productsdto.CreateStockRequest) (*productsresponses.CreatedStockAdminResponse, error) {
	result, err := p.productsRepository.CreateStock(req)
	if err != nil {
		return &productsresponses.CreatedStockAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to create stock",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateStock(req *productsdto.UpdateStockRequest) (*productsresponses.UpdateStockResponse, error) {
	result, err := p.productsRepository.UpdateStock(req)
	if err != nil {
		return &productsresponses.UpdateStockResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update stock",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) DecreaseStock(req *productsdto.DecreaseStockRequest) (*productsresponses.DecreaseStockResponse, error) {
	result, err := p.productsRepository.DecreaseStock(req)
	if err != nil {
		return &productsresponses.DecreaseStockResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to decrease stock",
				ErrorMessage: err.Error(),
			},
			Data: nil,
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetAllProducts() (*productsresponses.GetAllProductsResponse, error) {
	result, err := p.productsRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetAllProductsAdmin(role string, companyId string) (*productsresponses.GetAllProductsAdminResponse, error) {
	result, err := p.productsRepository.GetAllProductsAdmin(role, companyId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetCoverImageByMasterCode(masterCode string) (*productsresponses.GetCoverImageByMasterCodeResponse, error) {
	result, err := p.productsRepository.GetCoverImageByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetCoverImageByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to get cover image by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetProductDetail(masterCode string) (*productsresponses.GetProductDetailResponse, error) {
	result, err := p.productsRepository.GetProductDetail(masterCode)
	if err != nil {
		return &productsresponses.GetProductDetailResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to get product detail by product id",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetProductDetailAdmin(productId string) (*productsresponses.GetProductDetailAdminResponse, error) {
	result, err := p.productsRepository.GetProductDetailAdmin(productId)
	if err != nil {
		return &productsresponses.GetProductDetailAdminResponse{
			Status: &commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to get product detail by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) GetSizeChartByMasterCode(masterCode string) (*productsresponses.GetSizeChartByMasterCodeResponse, error) {
	result, err := p.productsRepository.GetSizeChartByMasterCode(masterCode)
	if err != nil {
		return &productsresponses.GetSizeChartByMasterCodeResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to get size chart by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateCoverImageByMasterCode(req *productsdto.UpdateCoverImageRequest) (*productsresponses.UpdatedCoverImageAdminResponse, error) {
	result, err := p.productsRepository.UpdateCoverImageByMasterCode(req)
	if err != nil {
		return &productsresponses.UpdatedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update cover image by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateSizeChartByMasterCode(req *productsdto.UpdateSizeChartRequest) (*productsresponses.UpdatedSizeChartAdminResponse, error) {
	result, err := p.productsRepository.UpdateSizeChartByMasterCode(req)
	if err != nil {
		return &productsresponses.UpdatedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update size chart by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UpdateMainProductDetailByMasterCode(req *productsdto.UpdateMainProductDetailByMasterCodeRequest) (*productsresponses.UpdatedProductAdminResponse, error) {
	result, err := p.productsRepository.UpdateMainProductDetailByMasterCode(req)
	if err != nil {
		return &productsresponses.UpdatedProductAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to update main product detail by master code",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UploadCoverImage(req *productsdto.UploadCoverImageRequest) (*productsresponses.UploadedCoverImageAdminResponse, error) {
	result, err := p.productsRepository.UploadCoverImage(req)
	if err != nil {
		return &productsresponses.UploadedCoverImageAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to upload cover image",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}

func (p *productServiceV2Impl) UploadSizeChart(req *productsdto.UploadSizeChartRequest) (*productsresponses.UploadedSizeChartAdminResponse, error) {
	result, err := p.productsRepository.UploadSizeChart(req)
	if err != nil {
		return &productsresponses.UploadedSizeChartAdminResponse{
			Status: commonresponse.CommonResponse{
				Status:       false,
				Message:      "SERVICE: Failed to upload size chart",
				ErrorMessage: err.Error(),
			},
		}, err
	}

	return result, nil
}
