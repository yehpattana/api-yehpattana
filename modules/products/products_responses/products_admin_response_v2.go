package productsresponses

import (
	commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"
	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
)

type GetAllProductsAdminResponse struct {
	Status   commonresponse.CommonResponse         `json:"status"`
	Products []ProductEssentialDetailAdminResponse `json:"products"`
}

type GetProductDetailAdminResponse struct {
	Status          *commonresponse.CommonResponse `json:"status"`
	MainProductData *MainProductDetailResponse     `json:"main_product_data"`
	ProductVaraints []*ProductVaraintAdminResponse `json:"product_varaints"` // Correct field name
}

type CreatedProductAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductAdminResponse          `json:"data"`
}

type CreatedStockAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   *StockAdminResponse           `json:"data"`
}

type UpdateStockResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   StockAdminResponse            `json:"data"`
}

type DecreaseStockResponse struct {
	Status          *commonresponse.CommonResponse `json:"status"`
	CurrentQuantity int                            `json:"current_quantity"`
	Data            *StockAdminResponse            `json:"data"`
}

type UploadedCoverImageAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type UploadedSizeChartAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type UpdatedCoverImageAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type UpdatedSizeChartAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type GetCoverImageByMasterCodeResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type GetSizeChartByMasterCodeResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductImageResponse          `json:"data"`
}

type UpdatedProductAdminResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   ProductAdminResponse          `json:"data"`
}

type UpdateMasterCodeResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
	Data   []entities.Product            `json:"data"`
}

type UpdateProductVaraintResponse struct {
	Status *commonresponse.CommonResponse     `json:"status"`
	Data   *ProductVaraintAfterUpdateResponse `json:"data"`
}

type CheckIsProductExistByMasterCodeResponse struct {
	Status         commonresponse.CommonResponse `json:"status"`
	IsProductExist bool                          `json:"is_product_exist"`
}

type DeleteAllProductInMasterCodeResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
}

type DeleteProductVariantResponse struct {
	Status commonresponse.CommonResponse `json:"status"`
}

type ProductEssentialDetailAdminResponse struct {
	Id               string                           `json:"id"`
	Name             string                           `json:"name"`
	ProductCode      string                           `json:"product_code"`
	ProductGroup     string                           `json:"product_group"`
	MasterCode       string                           `json:"master_code"`
	ProductStatus    string                           `json:"product_status"`
	CoverImage       string                           `json:"cover_image"`
	Currency         string                           `json:"currency"`
	Price            float64                          `json:"price"`
	RrpPrice         float64                          `json:"rrp_price"`
	UseAsPrimaryData bool                             `json:"use_as_primary_data"`
	LaunchDate       string                           `json:"launch_date"`
	EndOfLife        string                           `json:"end_of_life"`
	CreatedByCompany string                           `json:"created_by_company"`
	CreatedAt        string                           `json:"created_at"`
	CreatedBy        string                           `json:"created_by"`
	UpdatedAt        string                           `json:"updated_at"`
	EditedBy         string                           `json:"edited_by"`
	Collection       string                           `json:"collection"`
	Category         string                           `json:"category"`
	Colors           []string                         `json:"colors"`
	Gender           string                           `json:"gender"`
	Remark           string                           `json:"remark"`
	Variants         []*ProductVaraintWebViewResponse `json:"variant"`
}

type ProductsResponse struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	ProductCode      string  `json:"product_code"`
	MasterCode       string  `json:"master_code"`
	ProductStatus    string  `json:"product_status"`
	CoverImage       string  `json:"cover_image"`
	Price            float32 `json:"price"`
	UseAsPrimaryData bool    `json:"use_as_primary_data"`
	LaunchDate       string  `json:"launch_date"`
	EndOfLife        string  `json:"end_of_life"`
	CreatedByCompany string  `json:"created_by_company"`
	CreatedAt        string  `json:"created_at"`
	CreatedBy        string  `json:"created_by"`
	UpdatedAt        string  `json:"updated_at"`
	EditedBy         string  `json:"edited_by"`
	Collection       string  `json:"collection"`
	Category         string  `json:"category"`
	Varaint          []*ProductVaraintAdminResponse
}

type ProductAdminResponse struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	ProductCode   string  `json:"product_code"`
	MasterCode    string  `json:"master_code"`
	ColorCode     string  `json:"color_code"`
	ProductStatus string  `json:"product_status"`
	CoverImage    string  `json:"cover_image"`
	FrontImage    string  `json:"front_image"`
	BackImage     string  `json:"back_image"`
	Price         float64 `json:"price"`
	// -- Specific field for service only (don't have in YPT biz table) ---------------
	UseAsPrimaryData bool `json:"use_as_primary_data"`
	// Filter Product Information ---------------
	ProductGroup string `json:"product_group"`
	Season       string `json:"season"`
	Gender       string `json:"gender"`
	ProductClass string `json:"product_class"` // change name from classification
	Collection   string `json:"collection"`
	Category     string `json:"category"`
	Brand        string `json:"brand"`
	IsClub       bool   `json:"is_club"` // is club or general product
	ClubName     string `json:"club_name"`
	// Extra product information ----------------
	Remark          string  `json:"remark"`
	LaunchDate      string  `json:"launch_date"`
	EndOfLife       string  `json:"end_of_life"`
	SizeChart       string  `json:"size_chart"` // image url
	PackSize        string  `json:"pack_size"`
	CurrentSupplier string  `json:"current_supplier"`
	Description     string  `json:"description"`
	FabricContent   string  `json:"fabric_content"`
	FabricType      string  `json:"fabric_type"`
	Weight          float64 `json:"weight"`
	// product log -------------------------
	CreatedBy        string `json:"created_by"`
	EditedBy         string `json:"edited_by"`
	CreatedByCompany string `json:"created_by_company"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type ProductDetailAdminResponse struct {
	// Essential product information ------------
	Id            string `json:"id"`
	Name          string `json:"name"`
	MasterCode    string `json:"master_code"`
	CoverImage    string `json:"cover_image"`
	ProductStatus string `json:"product_status"`
	// filter product information ---------------
	ProductGroup string `json:"product_group"`
	Season       string `json:"season"`
	Gender       string `json:"gender"`
	ProductClass string `json:"product_class"`
	Collection   string `json:"collection"`
	Category     string `json:"category"`
	Brand        string `json:"brand"`
	IsClub       bool   `json:"is_club"`
	ClubName     string `json:"club_name"`
	// Extra product information ----------------
	Remark          string  `json:"remark"`
	LaunchDate      string  `json:"launch_date"`
	EndOfLife       string  `json:"end_of_life"`
	SizeChart       string  `json:"size_chart"`
	PackSize        string  `json:"pack_size"`
	CurrentSupplier string  `json:"current_supplier"`
	Description     string  `json:"description"`
	FabricContent   string  `json:"fabric_content"`
	FabricType      string  `json:"fabric_type"`
	Weight          float64 `json:"weight"`
	// product log -------------------------
	CreatedByCompany string                           `json:"created_by_company"`
	CreatedBy        string                           `json:"created_by"`
	EditedBy         string                           `json:"edited_by"`
	CreatedAt        string                           `json:"created_at"`
	UpdatedAt        string                           `json:"updated_at"`
	ProductItem      []ProductDetailItemAdminResponse `json:"product_item"`
}

type ProductDetailItemAdminResponse struct {
	ProductCode      string               `json:"product_code"`
	ColorCode        string               `json:"color_code"`
	FrontImage       string               `json:"front_image"`
	BackImage        string               `json:"back_image"`
	Price            float64              `json:"price"`
	UseAsPrimaryData bool                 `json:"use_as_primary_data"`
	Stock            []StockAdminResponse `json:"stock"`
}

type StockAdminResponse struct {
	Id          string  `json:"id"`
	ProductId   string  `json:"product_id"`
	Size        string  `json:"size"`
	SizeRemark  string  `json:"size_remark"`
	Quantity    int     `json:"quantity"`
	PreQuantity int     `json:"pre_quantity"`
	Price       float64 `json:"price"`
	RrpPrice    float64 `json:"rrp_price"`
	UsdPrice    float64 `json:"usd_price"`
	Currency    string  `json:"currency"`
	ItemStatus  string  `json:"item_status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductImageResponse struct {
	MasterCode string `json:"master_code"`
	Image      string `json:"image"`
}

type MainProductDetailResponse struct {
	Id               string  `json:"id"`
	Name             string  `json:"name"`
	MasterCode       string  `json:"master_code"`
	CoverImage       string  `json:"cover_image"`
	ProductStatus    string  `json:"product_status"`
	ProductGroup     string  `json:"product_group"`
	Season           string  `json:"season"`
	Gender           string  `json:"gender"`
	ProductClass     string  `json:"product_class"`
	Collection       string  `json:"collection"`
	Category         string  `json:"category"`
	Brand            string  `json:"brand"`
	IsClub           bool    `json:"is_club"`
	ClubName         string  `json:"club_name"`
	Remark           string  `json:"remark"`
	LaunchDate       string  `json:"launch_date"`
	EndOfLife        string  `json:"end_of_life"`
	SizeChart        string  `json:"size_chart"`
	PackSize         string  `json:"pack_size"`
	CurrentSupplier  string  `json:"current_supplier"`
	Description      string  `json:"description"`
	FabricContent    string  `json:"fabric_content"`
	FabricType       string  `json:"fabric_type"`
	Weight           float64 `json:"weight"`
	CreatedByCompany string  `json:"created_by_company"`
	CreatedBy        string  `json:"created_by"`
	EditedBy         string  `json:"edited_by"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type ProductVaraintAdminResponse struct {
	ProductId        string                `json:"product_id"`
	ProductCode      string                `json:"product_code"`
	ColorCode        string                `json:"color_code"`
	FrontImage       string                `json:"front_image"`
	BackImage        string                `json:"back_image"`
	Price            float64               `json:"price"`
	UseAsPrimaryData bool                  `json:"use_as_primary_data"`
	Stock            []*StockAdminResponse `json:"stock"`
}

type ProductVaraintAfterUpdateResponse struct {
	ProductId        string  `json:"product_id"`
	MasterCode       string  `json:"master_code"`
	ProductCode      string  `json:"product_code"`
	ColorCode        string  `json:"color_code"`
	FrontImage       string  `json:"front_image"`
	BackImage        string  `json:"back_image"`
	Price            float64 `json:"price"`
	UseAsPrimaryData bool    `json:"use_as_primary_data"`
}
