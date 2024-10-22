package entities

type Product struct {
	// Essential product information ------------
	Id            string  `gorm:"column:id;primaryKey" json:"id"`
	Name          string  `gorm:"column:name" json:"name" validate:"required"`
	ProductCode   string  `gorm:"column:product_code" json:"product_code" validate:"required"`
	MasterCode    string  `gorm:"column:master_code" json:"master_code" validate:"required"`
	ColorCode     string  `gorm:"column:color_code" json:"color_code" validate:"required"`
	ProductStatus string  `gorm:"column:product_status" json:"product_status" validate:"required,oneof=available, hidden, out_of_stock"`
	CoverImage    string  `gorm:"column:cover_image" json:"cover_image"`
	FrontImage    string  `gorm:"column:front_image" json:"front_image"`
	BackImage     string  `gorm:"column:back_image" json:"back_image"`
	Price         float64 `gorm:"column:price" json:"price" validate:"required,number"`
	// -- Specific field for service only (don't have in YPT biz table) ---------------
	UseAsPrimaryData bool `gorm:"use_as_primary_data" json:"use_as_primary_data"`
	// Filter Product Information ---------------
	ProductGroup string `gorm:"column:product_group" json:"product_group"`
	Season       string `gorm:"column:season" json:"season"`
	Gender       string `gorm:"column:gender" json:"gender" validate:"required,oneof=male,female,unisex,kids"`
	ProductClass string `gorm:"column:product_class" json:"product_class"` // change name from classification
	Collection   string `gorm:"column:collection" json:"collection"`
	Category     string `gorm:"column:category" json:"category"`
	Brand        string `gorm:"column:brand" json:"brand"`
	IsClub       bool   `gorm:"column:is_club" json:"is_club"` // is club or general product
	ClubName     string `gorm:"column:club_name" json:"club_name"`
	// Extra product information ----------------
	Remark          string  `gorm:"column:remark" json:"remark"`
	LaunchDate      string  `gorm:"column:launch_date" json:"launch_date"`
	EndOfLife       string  `gorm:"column:end_of_life" json:"end_of_life"`
	SizeChart       string  `gorm:"column:size_chart" json:"size_chart"` // image url
	PackSize        string  `gorm:"column:pack_size" json:"pack_size"`
	CurrentSupplier string  `gorm:"column:current_supplier" json:"current_supplier"`
	Description     string  `gorm:"column:description" json:"description"`
	FabricContent   string  `gorm:"column:fabric_content" json:"fabric_content"`
	FabricType      string  `gorm:"column:fabric_type" json:"fabric_type"`
	Weight          float64 `gorm:"column:weight" json:"weight"`
	// product log -------------------------
	CreatedByCompany string `gorm:"column:created_by_company" json:"created_by_company" validate:"required"`
	CreatedBy        string `gorm:"column:created_by" json:"created_by" validate:"required"`
	EditedBy         string `gorm:"column:edited_by" json:"edited_by" validate:"required"`
	CreatedAt        string `gorm:"column:created_at" json:"created_at" validate:"required"`
	UpdatedAt        string `gorm:"column:updated_at" json:"updated_at" validate:"required"`
}

type Stock struct {
	Id          string  `gorm:"column:id;primaryKey" json:"id"`
	ProductId   string  `gorm:"column:product_id" json:"product_id" validate:"required"`
	Size        string  `gorm:"column:size" json:"size" validate:"required"`
	SizeRemark  string  `gorm:"column:size_remark" json:"size_remark"`
	Quantity    int     `gorm:"column:quantity" json:"quantity" validate:"required,number"`
	PreQuantity int     `gorm:"column:pre_quantity" json:"pre_quantity" validate:"required,number"`
	Price       float64 `gorm:"column:price" json:"price" validate:"required,number"`
	RrpPrice    float64 `gorm:"column:rrp_price" json:"rrp_price" validate:"required,number"`
	UsdPrice    float64 `gorm:"column:usd_price" json:"usd_price"`
	Currency    string  `gorm:"column:currency" json:"currency"`
	ItemStatus  string  `gorm:"column:item_status" json:"item_status" validate:"required,oneof=available, out_of_stock"`
	CreatedAt   string  `gorm:"column:created_at" json:"created_at" validate:"required"`
	UpdatedAt   string  `gorm:"column:updated_at" json:"updated_at" validate:"required"`
}

type ProductCoverImage struct {
	Id         string `gorm:"column:id;primaryKey" json:"id"`
	MasterCode string `gorm:"column:master_code" json:"master_code" validate:"required"`
	Image      string `gorm:"column:image" json:"image" validate:"required"`
}

type SizeChartImage struct {
	Id         string `gorm:"column:id;primaryKey" json:"id"`
	MasterCode string `gorm:"column:master_code" json:"master_code" validate:"required"`
	Image      string `gorm:"column:image" json:"image" validate:"required"`
}
