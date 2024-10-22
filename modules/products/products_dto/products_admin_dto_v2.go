package productsdto

import "mime/multipart"

type CreateProductRequest struct {
	ProductFieldRequest
}

type CreateStockRequest struct {
	ProductId   string  `json:"product_id"`
	Size        string  `json:"size"`
	SizeRemark  string  `json:"size_remark"`
	Quantity    int     `json:"quantity"`
	PreQuantity int     `json:"pre_quantity"`
	Price       float64 `json:"price"`
	RrpPrice    float64 `json:"rrp_price"`
	UsdPrice    float64 `json:"usd_price"`
	Currency    string  `json:"currency"`
}

type UpdateStockRequest struct {
	StockId     string  `json:"stock_id"`
	Quantity    int     `json:"quantity"`
	PreQuantity int     `json:"pre_quantity"`
	Price       float64 `json:"price"`
	RrpPrice    float64 `json:"rrp_price"`
	SizeRemark  string  `json:"size_remark"`
	UsdPrice    float64 `json:"usd_price"`
	Currency    string  `json:"currency"`
}

type DecreaseStockRequest struct {
	StockId          string `json:"stock_id"`
	DecreaseQuantity int    `json:"decrease_quantity"`
}

type UploadCoverImageRequest struct {
	CoverImage *multipart.FileHeader `json:"cover_image"`
}

type UploadSizeChartRequest struct {
	SizeChart *multipart.FileHeader `json:"size_chart"`
}

type UpdateCoverImageRequest struct {
	MasterCode string `json:"master_code"`
	CoverImage string `json:"cover_image"`
}

type UpdateSizeChartRequest struct {
	MasterCode string                `json:"master_code"`
	SizeChart  *multipart.FileHeader `json:"size_chart"`
}

type UpdateMainProductDetailByMasterCodeRequest struct {
	MasterCode       string  `json:"master_code"`
	Name             string  `json:"name"`
	CoverImage       string  `json:"cover_image"`
	ProductGroup     string  `json:"product_group"`
	Season           string  `json:"season"`
	Gender           string  `json:"gender"`
	ProductStatus    string  `json:"product_status"`
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
	EditedBy         string  `json:"edited_by"`
}

type UpdateMasterCodeRequest struct {
	OldMasterCode string `json:"old_master_code"`
	NewMasterCode string `json:"new_master_code"`
	UpdatedBy     string `json:"updated_by"`
}

type UpdateProductVariantRequest struct {
	ProductId        string                `json:"product_id"`
	ProductCode      string                `json:"product_code"`
	ColorCode        string                `json:"color_code"`
	FrontImage       *multipart.FileHeader `json:"front_image"`
	BackImage        *multipart.FileHeader `json:"back_image"`
	Price            float64               `json:"price"`
	UseAsPrimaryData bool                  `json:"use_as_primary_data"`
}

type ProductFieldRequest struct {
	// Essential product information ------------
	Name          string                `json:"name"`
	ProductCode   string                `json:"product_code"`
	MasterCode    string                `json:"master_code"`
	ColorCode     string                `json:"color_code"`
	ProductStatus string                `json:"product_status"`
	CoverImage    string                `json:"cover_image"`
	FrontImage    *multipart.FileHeader `json:"front_image"`
	BackImage     *multipart.FileHeader `json:"back_image"`
	Price         float64               `json:"price"`
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
	// -- Specific field for service only (don't have in YPT biz table) ---------------
	UseAsPrimaryData bool `json:"use_as_primary_data"`
	// Extra product information ----------------
	Remark          string                `json:"remark"`
	LaunchDate      string                `json:"launch_date"`
	EndOfLife       string                `json:"end_of_life"`
	SizeChart       *multipart.FileHeader `json:"size_chart"`
	PackSize        string                `json:"pack_size"`
	CurrentSupplier string                `json:"current_supplier"`
	// product description ---------------------
	Description   string  `json:"description"`
	FabricContent string  `json:"fabric_content"`
	FabricType    string  `json:"fabric_type"`
	Weight        float64 `json:"weight"`
	// product history -------------------------
	CreatedByCompany string `json:"created_by_company"`
	CreatedBy        string `json:"created_by"`
}

type DeleteAllProductInMasterCodeRequest struct {
	MasterCode string `json:"master_code"`
}

type DeleteProductVariantRequest struct {
	ProductId string `json:"product_id"`
}

//stock json string

type MainProductData struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	MasterCode       string `json:"master_code"`
	CoverImage       string `json:"cover_image"`
	ProductStatus    string `json:"product_status"`
	ProductGroup     string `json:"product_group"`
	Season           string `json:"season"`
	Gender           string `json:"gender"`
	ProductClass     string `json:"product_class"`
	Collection       string `json:"collection"`
	Category         string `json:"category"`
	Brand            string `json:"brand"`
	IsClub           bool   `json:"is_club"`
	ClubName         string `json:"club_name"`
	Remark           string `json:"remark"`
	LaunchDate       string `json:"launch_date"`
	EndOfLife        string `json:"end_of_life"`
	SizeChart        string `json:"size_chart"`
	PackSize         string `json:"pack_size"`
	CurrentSupplier  string `json:"current_supplier"`
	Description      string `json:"description"`
	FabricContent    string `json:"fabric_content"`
	FabricType       string `json:"fabric_type"`
	Weight           int    `json:"weight"`
	CreatedByCompany string `json:"created_by_company"`
	CreatedBy        string `json:"created_by"`
	EditedBy         string `json:"edited_by"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type Stock struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
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

type ProductVariant struct {
	ProductID   string  `json:"product_id"`
	ProductCode string  `json:"product_code"`
	ColorCode   string  `json:"color_code"`
	ColorName   string  `json:"ColorName"`
	Color       string  `json:"Color"`
	FrontImage  string  `json:"front_image"`
	BackImage   string  `json:"back_image"`
	Stock       []Stock `json:"stock"`
}

type Price struct {
	ID          string  `json:"id"`
	ProductID   string  `json:"product_id"`
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

type Product struct {
	MainProductData MainProductData  `json:"main_product_data"`
	ProductVariants []ProductVariant `json:"product_varaints"`
	Price           []Price          `json:"price"`
	Color           string           `json:"color"`
	Quantity        map[string]int   `json:"quantity"`
	Name            string           `json:"name"`
	Type            string           `json:"type"`
}
