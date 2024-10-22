package dto

type UserRegisterRequest struct {
	Email string `db:"email" json:"email" form:"email"` // required

}

type CustomerRegisterRequest struct {
	User        *UserRegisterRequest // required
	ContactName string               `db:"contact_name" json:"contact_name" form:"contact_name"` // required
	CompanyName string               `db:"company_name" json:"company_name" form:"company_name"` // required
	VatNumber   string               `db:"vat_number" json:"vat_number" form:"vat_number"`       // required
	PhoneNumber string               `db:"phone_number" json:"phone_number" form:"phone_number"`
	Address     string               `db:"address" json:"address" form:"address"`    // required
	Cap         string               `db:"cap" json:"cap" form:"cap"`                // required
	City        string               `db:"city" json:"city" form:"city"`             // required
	Province    string               `db:"province" json:"province" form:"province"` // required
	Country     string               `db:"country" json:"country" form:"country"`    // required
	Message     string               `db:"message" json:"message" form:"message"`
}

type AdminRegisterRequest struct {
	Email       string `db:"email" json:"email" form:"email"`
	UserName    string `db:"user_name" json:"user_name" form:"user_name"`
	CompanyName string `db:"company_name" json:"company_name" form:"company_name"`
	Password    string `db:"password" json:"password" form:"password"`
}

type UserCredentialRequest struct {
	Email    string `db:"email" json:"email" form:"email"`          // required
	Password string `db:"password" json:"password" form:"password"` // required
}

type UserRefreshCredentialRequest struct {
	RefreshToken string `db:"refresh_token" json:"refresh_token" form:"refresh_token"` // required
}

type UserRemoveCredentialRequest struct {
	OauthId string `json:"oauth_id" form:"oauth_id"`
}
