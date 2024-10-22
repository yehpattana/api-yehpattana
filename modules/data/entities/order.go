package entities

import "time"

// type Order struct {
// 	OrderID        int       `gorm:"primaryKey;autoIncrement;column:order_id"`
// 	OrderDetail    []byte    `gorm:"column:order_detail;type:varbinary(MAX);not null"`
// 	Status         string    `gorm:"column:status" json:"status"`
// 	CustomerID     int       `gorm:"column:customer_id;not null"`
// 	CustomerDetail Customer  `gorm:"foreignKey:CustomerID"`
// 	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
// 	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`
// }

type Order struct {
	OrderID         int       `gorm:"primaryKey;autoIncrement;column:order_id"`
	OrderDetail     []byte    `gorm:"column:order_detail;type:varbinary(MAX);not null"`
	OrderNo         string    `gorm:"column:order_no" json:"order_no"`
	TotalAmount     string    `gorm:"column:total_amount" json:"total_amount"`
	ShippingAddress string    `gorm:"column:shipping_address" json:"shipping_address"`
	TrackingNo      string    `gorm:"column:tracking_no" json:"tracking_no"`
	PackingList     string    `gorm:"column:packing_list" json:"packing_list"`
	Status          string    `gorm:"column:status" json:"status"`
	PaymentStatus   string    `gorm:"column:payment_status" json:"payment_status"`
	PaymentID       string    `gorm:"column:payment_id" json:"payment_id"`
	CustomerID      int       `gorm:"column:customer_id;not null"`
	CustomerDetail  Customer  `gorm:"foreignKey:CustomerID;references:customer_id"` // Ensure the correct foreign key and reference
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
