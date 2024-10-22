package entities

import (
	"time"
)

type User struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"  validate:"required"` // TODO required
	Password  string    `gorm:"column:password" json:"password" validate:"required"`
	// Avatar    string    `gorm:"column:avatar" json:"avatar"`
	IsActived bool   `gorm:"column:is_actived" json:"is_actived"`
	Role      string `gorm:"column:role" json:"role"`
}

type Customer struct {
	CustomerID  int    `gorm:"primaryKey;autoIncrement;column:customer_id"`
	UserId      string `gorm:"foreignkey:user_id"`
	ContactName string `gorm:"column:contact_name" json:"contact_name"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	VatNumber   string `gorm:"column:vat_number;not null" json:"vat_number"` // TODO required
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
	Address     string `gorm:"column:address" json:"address"`
	Cap         string `gorm:"column:cap" json:"cap"`
	City        string `gorm:"column:city" json:"city"`
	Province    string `gorm:"column:province" json:"province"`
	Country     string `gorm:"column:country" json:"country"`
	Message     string `gorm:"column:message" json:"message"`
}

type Admin struct {
	UserId      string `gorm:"foreignkey:user_id"` // Foreign key
	UserName    string `gorm:"column:user_name" json:"user_name"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
}
