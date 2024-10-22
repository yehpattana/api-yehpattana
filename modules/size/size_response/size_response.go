package sizeresponses

type SizeResponse struct {
	ID   int    `json:"id"`
	Size string `json:"size"`
}

type SizeCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
