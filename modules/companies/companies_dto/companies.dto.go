package companiesdto

import "mime/multipart"

type CreateCompanyRequest struct {
	CompanyCode              string                `db:"company_code" json:"company_code" form:"company_code"`
	CompanyName              string                `db:"company_name" json:"company_name" form:"company_name"`
	Currency                 string                `db:"currency" json:"currency" form:"currency"`
	MinimumCostAvoidShipping float64               `db:"minimum_cost_avoid_shipping" json:"minimum_cost_avoid_shipping" form:"minimum_cost_avoid_shipping"`
	Logo                     *multipart.FileHeader `db:"logo" json:"logo" form:"logo"`
}

type GetCompanyDetailRequest struct {
	CompanyId string `json:"company_id" form:"company_id"`
}

type EditCompanyRequest struct {
	CompanyId                string                `json:"company_id" form:"company_id"`
	CompanyCode              string                `json:"company_code" form:"company_code"`
	CompanyName              string                `json:"company_name" form:"company_name"`
	Currency                 string                `json:"currency" form:"currency"`
	MinimumCostAvoidShipping float64               `db:"minimum_cost_avoid_shipping" json:"minimum_cost_avoid_shipping" form:"minimum_cost_avoid_shipping"`
	Logo                     *multipart.FileHeader `json:"logo" form:"logo"`
}

type DeleteCompanyRequest struct {
	CompanyId string `json:"company_id" form:"company_id"`
}
