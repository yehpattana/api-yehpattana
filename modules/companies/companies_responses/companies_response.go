package companiesresponses

type CompanyStatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CompanyFieldResponse struct {
	Id                       string  `json:"id"`
	CompanyCode              string  `json:"company_code"`
	CompanyName              string  `json:"company_name"`
	Currency                 string  `json:"currency"`
	MinimumCostAvoidShipping float64 `json:"minimum_cost_avoid_shipping"`
	Logo                     string  `json:"logo"`
}

type GetCompaniesResponse struct {
	Status *CompanyStatusResponse  `json:"status"`
	Data   []*CompanyFieldResponse `json:"data"`
}

type GetCompanyDetailResponse struct {
	Status *CompanyStatusResponse `json:"status"`
	Data   *CompanyFieldResponse  `json:"data"`
}

type CreateCompanyResponse struct {
	Status *CompanyStatusResponse `json:"status"`
	Data   *CompanyFieldResponse  `json:"data"`
}

type EditCompanyResponse struct {
	Status *CompanyStatusResponse `json:"status"`
	Data   *CompanyFieldResponse  `json:"data"`
}

type DeleteCompanyResponse struct {
	Status *CompanyStatusResponse `json:"status"`
}
