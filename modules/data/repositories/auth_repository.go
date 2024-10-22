package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/yehpattana/api-yehpattana/modules/auth/dto"
	"github.com/yehpattana/api-yehpattana/modules/auth/helpers"
	"github.com/yehpattana/api-yehpattana/modules/auth/responses"
	commonhelpers "github.com/yehpattana/api-yehpattana/modules/commons/common_helpers"
	commontypes "github.com/yehpattana/api-yehpattana/modules/commons/common_types"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	CreateCustomer(req *dto.CustomerRegisterRequest) (*responses.CustomerRegisterResponse, error)
	CreateAdmin(req *dto.AdminRegisterRequest) (*responses.AdminRegisterResponse, error)
	FindOneUserByEmail(email string) (*entities.UserCredentialCheck, error)
	FindOneUserById(userId string) (*entities.User, error)
	FindCustomerDataById(userId string) (*entities.Customer, error)
	FindAdminDataById(userId string) (*entities.Admin, error)
	InsertCustomerOauth(req *responses.CustomerPassportResponse) error
	InsertAdminOauth(req *responses.AdminPassportResponse) error
	FindOneOauthByRefreshToken(refreshToken string) (*entities.Oauth, error)
	UpdateOauth(req *responses.UserTokenResponse) error
	DeleteOauth(oauthId string) error
}

func AuthRepositoryImpl(db *gorm.DB) AuthRepositoryInterface {
	return &authRepositoryImpl{
		DB: db,
	}
}

type authRepositoryImpl struct {
	*gorm.DB
}

func (authRepository *authRepositoryImpl) FindCustomerDataById(userId string) (*entities.Customer, error) {
	customer := new(entities.Customer)
	if err := authRepository.DB.Table("Customers").Where("user_id = ?", userId).First(customer).Error; err != nil {
		return nil, fmt.Errorf("get customer profile failed: %v", err)
	}
	return customer, nil
}

func (authRepository *authRepositoryImpl) CreateCustomer(req *dto.CustomerRegisterRequest) (*responses.CustomerRegisterResponse, error) {
	// check company name is valid or not
	// isCompanyNameValid := commonhelpers.CheckIsValidUUID(req.CompanyName)

	// if !isCompanyNameValid {
	// 	return nil, fmt.Errorf("company name is not valid: company name must be company id (UUID) only")
	// }

	// Generate and hash the password
	generatePassword := "WelcomeToYmt!"
	// No hashing password before user verification

	// Create and save User
	user := &entities.User{
		Id:        commonhelpers.GenerateUUID(),
		Email:     req.User.Email,
		Password:  generatePassword,
		IsActived: true,
		Role:      string(commontypes.CustomerRole),
	}

	// Save the User to the database
	result := authRepository.DB.Table("Users").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retrieve the User from the database to ensure it's saved or to refresh any additional fields
	var savedUser entities.User
	if err := authRepository.DB.Table("Users").Where("id = ?", user.Id).First(&savedUser).Error; err != nil {
		return nil, err
	}

	// create and save Customer with UserID
	// TODO add avatar
	customer := &entities.Customer{
		UserId:      savedUser.Id,
		ContactName: req.ContactName,
		CompanyName: req.CompanyName, // company name must be the company id only
		VatNumber:   req.VatNumber,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Cap:         req.Cap,
		City:        req.City,
		Province:    req.Province,
		Country:     req.Country,
		Message:     req.Message,
	}

	// Save the Customer to the database
	result = authRepository.DB.Table("Customers").Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &responses.CustomerRegisterResponse{
		Success: true,
		Message: "Registration successful. Please wait for the admin to verify your account.",
	}, nil
}

func (authRepository *authRepositoryImpl) CreateAdmin(req *dto.AdminRegisterRequest) (*responses.AdminRegisterResponse, error) {
	role := string(commontypes.SuperAdminRole)
	// Generate and hash the password
	generatePassword := helpers.GeneratePassword(12)
	hashedPassword, err := helpers.BcryptHashingPassword(generatePassword)
	if req.Password != "" {
		role = string(commontypes.AdminRole)
		hashedPassword = req.Password
	}

	if err != nil {
		return nil, err
	}

	// Create and save User
	user := &entities.User{
		Id:        commonhelpers.GenerateUUID(),
		Email:     req.Email,
		Password:  hashedPassword,
		IsActived: true,
		Role:      role,
	}

	// Save the User to the database
	result := authRepository.DB.Table("Users").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retrieve the User from the database to ensure it's saved or to refresh any additional fields
	var savedUser entities.User
	if err := authRepository.DB.Table("Users").Where("id = ?", user.Id).First(&savedUser).Error; err != nil {
		return nil, err
	}

	var SavedAdmin entities.Admin
	if err := authRepository.DB.Table("Admins").Where("user_id = ?", savedUser.Id).First(&SavedAdmin).Error; err == nil {
		return nil, fmt.Errorf("admin already exists")
	}

	// create and save Admin with UserID
	admin := &entities.Admin{
		UserId:      savedUser.Id,
		UserName:    req.UserName,
		CompanyName: req.CompanyName,
	}

	// Save the Admin to the database
	result = authRepository.DB.Table("Admins").Create(admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return &responses.AdminRegisterResponse{
		Success: true,
		Message: "Registration successful. Your account is",
		AdminDataResponse: &responses.AdminDataResponse{
			UserId:      savedUser.Id,
			UserName:    req.UserName,
			Email:       savedUser.Email,
			Password:    generatePassword,
			IsActived:   savedUser.IsActived,
			Role:        savedUser.Role,
			CreatedDate: savedUser.CreatedAt,
			UpdatedDate: savedUser.UpdatedAt,
		},
	}, nil

}

func (authRepository *authRepositoryImpl) FindOneUserByEmail(email string) (*entities.UserCredentialCheck, error) {
	user := new(entities.UserCredentialCheck)
	if err := authRepository.DB.Table("Users").Where("email = ?", email).First(user).Error; err != nil {
		return nil, err // If an error occurs, return the error
	}

	return user, nil
}

func (authRepository *authRepositoryImpl) FindOneUserById(userId string) (*entities.User, error) {
	user := new(entities.User)
	if err := authRepository.DB.Table("Users").Where("id = ?", userId).First(user).Error; err != nil {
		return nil, fmt.Errorf("get profile failed: %v", err)
	}
	return user, nil
}

func (authRepository *authRepositoryImpl) FindAdminDataById(userId string) (*entities.Admin, error) {
	admin := new(entities.Admin)
	if err := authRepository.DB.Table("Admins").Where("user_id = ?", userId).First(admin).Error; err != nil {
		return nil, fmt.Errorf("get admin profile failed: %v", err)
	}
	return admin, nil
}

func (authRepository *authRepositoryImpl) InsertCustomerOauth(req *responses.CustomerPassportResponse) error {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	uuid := commonhelpers.GenerateUUID()

	oauth := &responses.UserTokenResponse{
		Id:           uuid,
		UserId:       req.UserId,
		RefreshToken: req.Token.RefreshToken,
		AccessToken:  req.Token.AccessToken,
	}
	if err := authRepository.DB.WithContext(ctx).Table("oauth").Create(oauth).Error; err != nil {
		return fmt.Errorf("insert oauth customer failed: %v", err)
	}
	// Set the generated ID back to the request object
	req.Token.Id = oauth.Id
	req.UserId = oauth.UserId
	return nil
}

func (authRepository *authRepositoryImpl) InsertAdminOauth(req *responses.AdminPassportResponse) error {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	uuid := commonhelpers.GenerateUUID()

	oauth := &responses.UserTokenResponse{
		Id:           uuid,
		UserId:       req.UserId,
		RefreshToken: req.Token.RefreshToken,
		AccessToken:  req.Token.AccessToken,
	}
	if err := authRepository.DB.WithContext(ctx).Table("oauth").Create(oauth).Error; err != nil {
		return fmt.Errorf("insert oauth admin failed: %v", err)
	}
	// Set the generated ID back to the request object
	req.Token.Id = oauth.Id
	req.UserId = oauth.UserId
	return nil
}

func (authRepository *authRepositoryImpl) FindOneOauthByRefreshToken(refreshToken string) (*entities.Oauth, error) {
	oauth := new(entities.Oauth)
	if err := authRepository.DB.Table("oauth").Where("refresh_token = ?", refreshToken).First(oauth).Error; err != nil {
		return nil, fmt.Errorf("find oauth by refresh token failed: %v", err)
	}
	return oauth, nil
}

func (authRepository *authRepositoryImpl) UpdateOauth(req *responses.UserTokenResponse) error {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	oauth := &responses.UserTokenResponse{
		UserId:       req.UserId,
		RefreshToken: req.RefreshToken,
		AccessToken:  req.AccessToken,
	}
	if err := authRepository.DB.WithContext(ctx).Table("oauth").Where("user_id = ?", req.UserId).Updates(oauth).Error; err != nil {
		return fmt.Errorf("update oauth customer failed: %v", err)
	}
	return nil
}

func (authRepository *authRepositoryImpl) DeleteOauth(oauthId string) error {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	if err := authRepository.DB.WithContext(ctx).Table("oauth").Where("Id = ?", oauthId).Delete(&entities.Oauth{}).Error; err != nil {
		return fmt.Errorf("delete oauth customer failed: %v", err)
	}
	return nil
}
