package responses

import (
	"time"
)

type AuthStatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CustomerRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AdminRegisterResponse struct {
	Success           bool               `json:"success"`
	Message           string             `json:"message"`
	AdminDataResponse *AdminDataResponse `json:"admin_data"`
}

type AdminDataResponse struct {
	UserId      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsActived   bool      `json:"is_actived"`
	Role        string    `json:"role"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type CustomerPassportResponse struct {
	UserId        string             `json:"user_id"`
	Email         string             `json:"email"`
	Roles         string             `json:"roles"`
	CompanyName   string             `json:"company_name"`
	Logo          string             `json:"logo"`
	Token         *UserTokenResponse `json:"token"`
	CustomerID    int                `json:"customer_id"`
	ResetPassword bool               `json:"reset_password"`
}

type AdminPassportResponse struct {
	UserId   string             `json:"user_id"`
	UserName string             `json:"user_name"`
	Email    string             `json:"email"`
	Roles    string             `json:"roles"`
	Token    *UserTokenResponse `json:"token"`
}

type SignOutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
