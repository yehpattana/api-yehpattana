package productsresponses

import commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"

type GetAllProductsResponse struct {
	Status   commonresponse.CommonResponse    `json:"status"`
	Products []ProductEssentialDetailResponse `json:"products"`
}

type GetAllProductByCompanyResponse struct {
	Status   commonresponse.CommonResponse         `json:"status"`
	Products []ProductEssentialDetailAdminResponse `json:"products"`
}

type GetProductDetailResponse struct {
	Status          commonresponse.CommonResponse    `json:"status"`
	MainProductData *MainProductDetailResponse       `json:"main_product_data"`
	ProductVaraints []*ProductVaraintWebViewResponse `json:"product_varaints"` // Correct field name
}

type ProductVaraintWebViewResponse struct {
	ProductId   string                  `json:"product_id"`
	ProductCode string                  `json:"product_code"`
	ColorCode   string                  `json:"color_code"`
	ColorName   string                  `json:name_color`
	Color       string                  `json:color`
	FrontImage  string                  `json:"front_image"`
	BackImage   string                  `json:"back_image"`
	Stock       []*StockWebviewResponse `json:"stock"`
}

type StockWebviewResponse struct {
	Id          string  `json:"id"`
	ProductId   string  `json:"product_id"`
	Size        string  `json:"size"`
	SizeRemark  string  `json:"size_remark"`
	Quantity    int     `json:"quantity"`
	PreQuantity int     `json:"pre_quantity"`
	Price       float64 `json:"price"`
	RrpPrice    float64 `json:"rrp_price"`
	Currency    string  `json:"currency"`
	ItemStatus  string  `json:"item_status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductEssentialDetailResponse struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	MasterCode       string   `json:"master_code"`
	ProductStatus    string   `json:"product_status"`
	CoverImage       string   `json:"cover_image"`
	Colors           []string `json:"colors"`
	SizeRange        string   `json:"size_range"`
	CreatedByCompany string   `json:"created_by_company"`
	Collection       string   `json:"collection"`
	Category         string   `json:"category"`
}

type ProductResponse struct {
	// Essential product information ------------
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	ProductCode   string  `json:"product_code"`
	MasterCode    string  `json:"master_code"`
	ColorCode     string  `json:"color_code"`
	ProductStatus string  `json:"product_status"`
	CoverImage    string  `json:"cover_image"`
	Price         float32 `json:"price"`
	// Extra product information ----------------
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
	Weight          float32 `json:"weight"`
	// product log -------------------------
	CreatedBy        string `json:"created_by"`
	EditedBy         string `json:"edited_by"`
	CreatedByCompany string `json:"created_by_company"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type StockResponse struct {
	Id          string `json:"id"`
	ProductId   string `json:"product_id"`
	Size        string `json:"size"`
	SizeRemark  string `json:"size_remark"`
	Quantity    int    `json:"quantity"`
	PreQuantity int    `json:"pre_quantity"`
	ItemStatus  string `json:"item_status"`
}
