package colorresponses

type ColorResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	CodeName string `json:"codeName"`
}

type ColorCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ColorDetail struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
