package ordersresponses

type OrderResponse struct {
	OrderID         int            `json:"order_id"`
	OrderDetail     string         `json:"order_detail"`
	OrderNo         string         `json:"order_no"`
	ShippingAddress string         `json:"shipping_address"`
	TrackingNo      string         `json:"tracking_no"`
	PackingList     string         `json:"packing_list"`
	TotalAmount     string         `json:"total_amount"`
	CustomerDetail  CustomerDetail `json:"customer_detail"`
	Status          string         `json:"status"`
	PaymentStatus   string         `json:"payment_status"`
	PaymentID       string         `json:"payment_id"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}

type OrderCreateResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CustomerDetail struct {
	CustomerID   int    `json:"customer_id"`
	ContractName string `json:"contract_name"`
	CompanyName  string `json:"company_name"`
}
