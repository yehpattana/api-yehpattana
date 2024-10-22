package logresponses

type LogResponse struct {
	Id          int    `json:"id"`
	EndPoint    string `json:"end_point"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedBy   string `json:"updated_by"`
}
