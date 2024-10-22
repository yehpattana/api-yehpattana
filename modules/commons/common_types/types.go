package commontypes

type UserRole string

const (
	CustomerRole   UserRole = "Customer"
	AdminRole      UserRole = "Admin"
	SuperAdminRole UserRole = "SuperAdmin"
)

type ProductStatus string

const (
	AvailableProductStatus ProductStatus = "Available"
	HiddenProductStatus    ProductStatus = "Hidden"
	SoldOutProductStatus   ProductStatus = "Sold Out"
)
