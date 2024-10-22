package ordersdto

import "mime/multipart"

type OrderRequest struct {
	OrderNo         string `db:"column:order_no" json:"order_no" form:"OrderNo"`
	CompanyName     string `json:"company_name" form:"companyName"`
	ShippingAddress string `json:"shipping_address" form:"shipping_address"`
	TrackingNo      string `json:"tracking_no" form:"tracking_no"`
	OrderDetail     string `db:"column:order_detail" json:"order_detail" form:"orderDetail"`
	CustomerID      int    `db:"column:customer_id;not null" json:"customer_id" form:"customerId"`
	TotalAmount     string `db:"column:total_amount" json:"total_amount" form:"totalAmount"`
}
type UploadPackingListRequest struct {
	File *multipart.FileHeader `db:"column:packing_list" json:"file" form:"File"`
}
type OrderTrackingRequest struct {
	OrderNo    string `db:"column:order_no" json:"order_no" form:"OrderNo"`
	TrackingNo string `json:"tracking_no" form:"tracking_no"`
	Status     string `json:"status" form:"status"`
}

type OrderPaymentRequest struct {
	OrderNo       string `db:"column:order_no" json:"order_no" form:"OrderNo"`
	PaymentStatus string `json:"payment_status" form:"payment_status"`
	PaymentID     string `json:"payment_id" form:"payment_id"`
}
