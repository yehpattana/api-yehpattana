package entities

type UserCredentialCheck struct {
	Id          string `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"` // TODO check case when is null
	Password    string `json:"password"`
	IsActived   bool   `json:"is_active"`
	Role        string `json:"role"`
	CompanyName string `json:"company_name"`
}

type UserClaims struct {
	CompanyName string `json:"company`
	Id          string `json:"id"`
	Role        string `json:"role"`
	Email       string `json:"email"` // TODO check case when is null
}

type Oauth struct {
	Id     string `json:"id" gorm:"id"`
	UserId string `json:"user_id" gorm:"user_id"`
}
