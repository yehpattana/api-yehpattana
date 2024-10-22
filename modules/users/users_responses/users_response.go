package usersresponses

type GetAllCustomersResponse struct {
	Customers []CustomerResponse `json:"customers"`
}

type CustomerResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	IsActived   bool   `json:"is_actived"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ContactName string `json:"contact_name"`
	CompanyName string `json:"company_name"`
	CustomerId  string `json:"customer_id"`
	VatNumber   string `json:"vat_number"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Cap         string `json:"cap"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	Message     string `json:"message"`
}

type GetCustomerDetialResponse struct {
	Email       string `json:"email"`
	ContactName string `json:"contact_name"`
	CompanyName string `json:"company_name"`
	VatNumber   string `json:"vat_number"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Cap         string `json:"cap"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	Message     string `json:"message"`
}

type GetUserIdByEmailResponse struct {
	UserId string `json:"user_id"`
}

type ChangePasswordResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type VerifyUserResponse struct {
	Succuess bool   `json:"success"`
	Message  string `json:"message"`
}

type BanUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type UpdateUserResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	IsActived   bool   `json:"is_actived"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	ContactName string `json:"contact_name"`
	CompanyName string `json:"company_name"`
	VatNumber   string `json:"vat_number"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Cap         string `json:"cap"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	Message     string `json:"message"`
}
