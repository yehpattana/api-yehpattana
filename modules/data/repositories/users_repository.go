package repositories

import (
	"fmt"
	"time"

	"github.com/yehpattana/api-yehpattana/modules/auth/helpers"
	companiesdto "github.com/yehpattana/api-yehpattana/modules/companies/companies_dto"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	usersdto "github.com/yehpattana/api-yehpattana/modules/users/users_dto"
	usersresponses "github.com/yehpattana/api-yehpattana/modules/users/users_responses"
	"gorm.io/gorm"
)

type UsersRepositoryInterface interface {
	FindAllCustomers(role string, companyId string) ([]*usersresponses.CustomerResponse, error)
	GetCustomerDetail(userId string) (*usersresponses.GetCustomerDetialResponse, error)
	GetUserIdByEmail(userEmail string) (*usersresponses.GetUserIdByEmailResponse, error)
	ChangePassword(userId string, newPassword string) (*usersresponses.ChangePasswordResponse, error)
	VerifyUser(userId string) (*usersresponses.VerifyUserResponse, error)
	BanUser(userId string) (*usersresponses.BanUserResponse, error)
	UpdateCustomer(userRequest *usersdto.UpdateCustomerRequest, userId string) (*usersresponses.UpdateUserResponse, error)
	DeleteCustomer(userId string, customerId string) (*usersresponses.BanUserResponse, error)
}

func UsersRepositoryImpl(db *gorm.DB) UsersRepositoryInterface {
	return &usersRepositoryImpl{
		DB: db,
	}
}

type usersRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *usersRepositoryImpl) FindAllCustomers(role string, companyId string) ([]*usersresponses.CustomerResponse, error) {
	var customers []*usersresponses.CustomerResponse

	result := userRepository.DB.Table("Users").Where("role = ?", "customer").Find(&customers)
	// Fetch basic customer details
	if result.Error != nil {
		return nil, result.Error
	}

	// Fetch additional details (CompanyName, PhoneNumber) for each customer
	for _, customer := range customers {
		var customerDetails struct {
			CustomerId  string
			ContactName string
			CompanyName string
			PhoneNumber string
			Address     string
			Cap         string
			VatNumber   string
			City        string
			Province    string
			Country     string
			Message     string
		}
		// Fetch CompanyName and PhoneNumber from Customers table based on customer ID
		err := userRepository.DB.Table("Customers").Select("customer_id,company_name,contact_name, phone_number, address, vat_number, city, province, country, message, cap").
			Where("user_id = ?", customer.Id).Scan(&customerDetails).Error
		if err != nil {
			return nil, err
		}

		// Assign fetched details to the respective customer
		customer.CustomerId = customerDetails.CustomerId
		customer.ContactName = customerDetails.ContactName
		customer.CompanyName = customerDetails.CompanyName
		customer.PhoneNumber = customerDetails.PhoneNumber
		customer.Address = customerDetails.Address
		customer.Cap = customerDetails.Cap
		customer.VatNumber = customerDetails.VatNumber
		customer.City = customerDetails.City
		customer.Province = customerDetails.Province
		customer.Country = customerDetails.Country
		customer.Message = customerDetails.Message
	}
	if role != "SuperAdmin" {
		var filterCustomer []*usersresponses.CustomerResponse
		companyDetail, _ := CompaniesRepositoryImpl(userRepository.DB).GetCompanyById(&companiesdto.GetCompanyDetailRequest{CompanyId: companyId})
		for _, c := range customers {
			if companyDetail.Data != nil && c.CompanyName == companyDetail.Data.CompanyName {
				filterCustomer = append(filterCustomer, c)
			}
		}
		return filterCustomer, nil
	}

	return customers, nil
}

func (userRepository *usersRepositoryImpl) GetCustomerDetail(userId string) (*usersresponses.GetCustomerDetialResponse, error) {
	var customer usersresponses.GetCustomerDetialResponse
	var user entities.User
	result := userRepository.DB.Table("Customers").Where("user_id = ?", userId).Find(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	userData := userRepository.DB.Table("Users").Where("id = ?", userId).Find(&customer)
	if userData.Error != nil {
		return nil, userData.Error
	}

	if err := userRepository.DB.Table("Users").Select("email").Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("error fetching user details: %v", err)
	}

	// Assigning the email to the customer response
	customer.Email = user.Email

	return &usersresponses.GetCustomerDetialResponse{
		Email:       customer.Email,
		ContactName: customer.ContactName,
		CompanyName: customer.CompanyName,
		VatNumber:   customer.VatNumber,
		PhoneNumber: customer.PhoneNumber,
		Address:     customer.Address,
		Cap:         customer.Cap,
		City:        customer.City,
		Province:    customer.Province,
		Country:     customer.Country,
		Message:     customer.Message,
	}, nil
}
func (userRepository *usersRepositoryImpl) GetUserIdByEmail(userEmail string) (*usersresponses.GetUserIdByEmailResponse, error) {
	var user entities.User

	userData := userRepository.DB.Table("Users").Where("email = ?", userEmail).Find(&user)
	if userData.Error != nil {
		return nil, userData.Error
	}

	return &usersresponses.GetUserIdByEmailResponse{
		UserId: user.Id,
	}, nil
}

func (userRepository *usersRepositoryImpl) ChangePassword(userId string, newPassword string) (*usersresponses.ChangePasswordResponse, error) {
	hashedPassword, err := helpers.BcryptHashingPassword(newPassword)
	if err != nil {
		return nil, err

	}

	result := userRepository.DB.Table("Users").Where("id = ?", userId).Update("password", hashedPassword)
	if result.Error != nil {
		return nil, result.Error
	}

	return &usersresponses.ChangePasswordResponse{
		Success: true,
		Message: "Password has been changed",
	}, nil
}

func (userRepository *usersRepositoryImpl) VerifyUser(userId string) (*usersresponses.VerifyUserResponse, error) {
	fmt.Println("🚩userId", userId)
	result := userRepository.DB.Table("Users").Where("id = ?", userId).Update("is_actived", true)
	if result.Error != nil {
		return nil, result.Error
	}

	return &usersresponses.VerifyUserResponse{
		Succuess: true,
		Message:  "User has been verified",
	}, nil
}

func (userRepository *usersRepositoryImpl) BanUser(userId string) (*usersresponses.BanUserResponse, error) {
	result := userRepository.DB.Table("Users").Where("id = ?", userId).Update("is_actived", false)
	if result.Error != nil {
		return nil, result.Error
	}

	return &usersresponses.BanUserResponse{
		Success: true,
		Message: "User has been banned",
	}, nil
}

func (userRepository *usersRepositoryImpl) UpdateCustomer(userRequest *usersdto.UpdateCustomerRequest, userId string) (*usersresponses.UpdateUserResponse, error) {
	var user entities.User
	userResult := userRepository.DB.Table("Users").Where("id = ?", userId).Updates(map[string]interface{}{
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	})

	if userResult.Error != nil {
		return nil, userResult.Error
	}

	// fetch user data and assign to user
	userData := userRepository.DB.Table("Users").Where("id = ?", userId).Find(&user)
	if userData.Error != nil {
		return nil, userData.Error
	}

	// TODO implement upload avatar image

	customerResult := userRepository.DB.Table("Customers").Where("user_id = ?", userId).Updates(map[string]interface{}{
		"contact_name": userRequest.ContactName,
		"company_name": userRequest.CompanyName,
		"vat_number":   userRequest.VatNumber,
		"phone_number": userRequest.PhoneNumber,
		"address":      userRequest.Address,
		"cap":          userRequest.Cap,
		"city":         userRequest.City,
		"province":     userRequest.Province,
		"country":      userRequest.Country,
		"message":      userRequest.Message,
	})

	if customerResult.Error != nil {
		return nil, customerResult.Error
	}

	return &usersresponses.UpdateUserResponse{
		Id:          userId,
		Email:       user.Email,
		IsActived:   user.IsActived,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
		ContactName: userRequest.ContactName,
		CompanyName: userRequest.CompanyName,
		VatNumber:   userRequest.VatNumber,
		PhoneNumber: userRequest.PhoneNumber,
		Address:     userRequest.Address,
		Cap:         userRequest.Cap,
		City:        userRequest.City,
		Province:    userRequest.Province,
		Country:     userRequest.Country,
		Message:     userRequest.Message,
	}, nil
}

func (userRepository *usersRepositoryImpl) DeleteCustomer(userId string, customerId string) (*usersresponses.BanUserResponse, error) {
	result := userRepository.DB.Table("Orders").Where("customer_id = ?", customerId).Delete(&entities.Order{})
	if result.Error != nil {
		return nil, result.Error
	}
	result = userRepository.DB.Table("Customers").Where("user_id = ?", userId).Delete(&entities.Customer{})
	if result.Error != nil {
		return nil, result.Error
	}
	result = userRepository.DB.Table("Oauth").Where("user_id = ?", userId).Delete(&entities.Oauth{})
	if result.Error != nil {
		return nil, result.Error
	}
	result = userRepository.DB.Table("Users").Where("id = ?", userId).Delete(&entities.User{})
	if result.Error != nil {
		return nil, result.Error
	}

	return &usersresponses.BanUserResponse{
		Success: true,
		Message: "User has been delete",
	}, nil
}
