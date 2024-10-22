package entities

type Company struct {
	Id                       string  `gorm:"column:id;primaryKey" json:"id"`
	CompanyCode              string  `gorm:"column:company_code" json:"company_code" validate:"required"`
	CompanyName              string  `gorm:"column:company_name" json:"company_name" validate:"required"`
	Currency                 string  `gorm:"column:currency" json:"currency"`
	MinimumCostAvoidShipping float64 `gorm:"column:minimum_cost_avoid_shipping" json:"minimum_cost_avoid_shipping"`
	Logo                     string  `gorm:"column:logo" json:"logo"`
}
