package usersdto

type GetUserDetailRequest struct {
	UserId string `json:"id"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password"`
}

type UpdateCustomerRequest struct {
	ContactName string `db:"contact_name" json:"contact_name" form:"contact_name"`
	CompanyName string `db:"company_name" json:"company_name" form:"company_name"`
	VatNumber   string `db:"vat_number" json:"vat_number" form:"vat_number"`
	PhoneNumber string `db:"phone_number" json:"phone_number" form:"phone_number"`
	Address     string `db:"address" json:"address" form:"address"`
	Cap         string `db:"cap" json:"cap" form:"cap"`
	City        string `db:"city" json:"city" form:"city"`
	Province    string `db:"province" json:"province" form:"province"`
	Country     string `db:"country" json:"country" form:"country"`
	Message     string `db:"message" json:"message" form:"message"`
}
