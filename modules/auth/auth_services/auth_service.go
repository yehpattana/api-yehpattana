package authservices

import (
	"fmt"

	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/auth/dto"
	"github.com/yehpattana/api-yehpattana/modules/auth/helpers"
	"github.com/yehpattana/api-yehpattana/modules/auth/pkg"
	"github.com/yehpattana/api-yehpattana/modules/auth/responses"
	commonhelpers "github.com/yehpattana/api-yehpattana/modules/commons/common_helpers"
	commonimages "github.com/yehpattana/api-yehpattana/modules/commons/common_images"
	commonstring "github.com/yehpattana/api-yehpattana/modules/commons/common_string"
	commontypes "github.com/yehpattana/api-yehpattana/modules/commons/common_types"
	companiesdto "github.com/yehpattana/api-yehpattana/modules/companies/companies_dto"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
)

type AuthServiceInterface interface {
	CreateCustomer(req *dto.CustomerRegisterRequest) (*responses.CustomerRegisterResponse, error)
	CreateAdmin(req *dto.AdminRegisterRequest) (*responses.AdminRegisterResponse, error)
	SignInCustomer(req *dto.UserCredentialRequest) (*responses.CustomerPassportResponse, error)
	SignInAdmin(req *dto.UserCredentialRequest) (*responses.AdminPassportResponse, error)
	RefreshCustomerPassport(req *dto.UserRefreshCredentialRequest) (*responses.CustomerPassportResponse, error)
	RefreshAdminPassport(req *dto.UserRefreshCredentialRequest) (*responses.AdminPassportResponse, error)
	DeleteOauth(oauthId string) (*responses.SignOutResponse, error)
}

func AuthServiceImpl(cfg configs.ConfigInterface, authRepository repositories.AuthRepositoryInterface, companyRepository repositories.CompaniesRepositoryInterface) AuthServiceInterface {
	return &userServiceImpl{
		config:            cfg,
		authRepository:    authRepository,
		companyRepository: companyRepository,
	}
}

type userServiceImpl struct {
	config            configs.ConfigInterface
	authRepository    repositories.AuthRepositoryInterface
	companyRepository repositories.CompaniesRepositoryInterface
}

func (u *userServiceImpl) CreateCustomer(req *dto.CustomerRegisterRequest) (*responses.CustomerRegisterResponse, error) {
	result, err := u.authRepository.CreateCustomer(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userServiceImpl) CreateAdmin(req *dto.AdminRegisterRequest) (*responses.AdminRegisterResponse, error) {
	result, err := u.authRepository.CreateAdmin(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userServiceImpl) SignInCustomer(req *dto.UserCredentialRequest) (*responses.CustomerPassportResponse, error) {
	// find the user by email
	user, err := u.authRepository.FindOneUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	// find user company by customer id
	customer, err := u.authRepository.FindCustomerDataById(user.Id)
	if err != nil {
		return nil, err
	}
	company, err := u.companyRepository.GetCompanyByName(&companiesdto.GetCompanyDetailRequest{CompanyId: customer.CompanyName})
	if err != nil {
		return nil, err
	}
	userCompanyLogo := ""
	userCompanyName := customer.UserId

	isCompanyNameUUID := commonhelpers.CheckIsValidUUID(userCompanyName)

	if isCompanyNameUUID {
		print("company name is UUID")
		// Find the user company by company ID
		companyDetailRequest := &companiesdto.GetCompanyDetailRequest{
			CompanyId: customer.CompanyName, // company name is the company ID
		}
		company, err := u.companyRepository.GetCompanyByName(companyDetailRequest)

		if err != nil {
			userCompanyName = user.CompanyName
			userCompanyLogo = commonimages.DefaultCompanyLogoImage
		}

		userCompanyName = company.Data.CompanyName
		userCompanyLogo = company.Data.Logo
	} else {
		print("company name is not UUID")
		userCompanyName = customer.CompanyName
	}

	println("userCompanyLogo", userCompanyLogo)

	// Set default values if they are empty
	if userCompanyName == "" {
		userCompanyName = commonstring.UnknowCompanyName
	}

	if userCompanyLogo == "" {
		userCompanyLogo = commonimages.DefaultCompanyLogoImage
	}

	// check is account active
	if !user.IsActived {
		return nil, fmt.Errorf(commonstring.AccountIsNotActiveOrBanned)
	}

	// check role is valid
	if user.Role != string(commontypes.CustomerRole) {
		return nil, fmt.Errorf(commonstring.UserRoleIsNotValid)
	}

	// check the password
	if helpers.IsHashedPassword(user.Password) {
		// If it's hashed, compare it with the provided password
		err = helpers.BcryptComparePassword(user.Password, req.Password)
		if err != nil {
			return nil, err
		}
	} else {
		// If it's not hashed, hash the provided password and compare it with the stored password
		if user.Password != req.Password {
			return nil, fmt.Errorf(commonstring.IncorrectPasswordOrEmail)
		}
	}

	// sign the token
	accessToken, err := pkg.NewYptAuth(u.config.Jwt(), pkg.AccessToken, &entities.UserClaims{
		Id:          user.Id,
		Role:        user.Role,
		CompanyName: company.Data.Id,
		Email:       user.Email,
	})

	refreshToken, err := pkg.NewYptAuth(u.config.Jwt(), pkg.RefreshToken, &entities.UserClaims{
		Id:          user.Id,
		Role:        user.Role,
		CompanyName: company.Data.Id,
		Email:       user.Email,
	})

	resetPassword := req.Password == "WelcomeToYmt!"
	// set the user to the passport
	passport := &responses.CustomerPassportResponse{
		UserId:      user.Id,
		Email:       user.Email,
		Roles:       user.Role,
		CompanyName: userCompanyName,
		Logo:        userCompanyLogo,
		Token: &responses.UserTokenResponse{
			UserId:       user.Id,
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken.SignToken(),
		},
		CustomerID:    customer.CustomerID,
		ResetPassword: resetPassword,
	}

	if err := u.authRepository.InsertCustomerOauth(passport); err != nil {
		return nil, err
	}

	return passport, nil
}

func (u *userServiceImpl) SignInAdmin(req *dto.UserCredentialRequest) (*responses.AdminPassportResponse, error) {
	// find the user by email
	user, err := u.authRepository.FindOneUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	admin, err := u.authRepository.FindAdminDataById(user.Id)
	if err != nil {
		return nil, err
	}

	// check is account active
	if !user.IsActived {
		return nil, fmt.Errorf(commonstring.AccountIsNotActiveOrBanned)
	}

	// check role is valid
	if user.Role != string(commontypes.AdminRole) && user.Role != string(commontypes.SuperAdminRole) {
		return nil, fmt.Errorf(commonstring.UserRoleIsNotValid)
	}

	// check the password
	if helpers.IsHashedPassword(user.Password) {
		// If it's hashed, compare it with the provided password
		err = helpers.BcryptComparePassword(user.Password, req.Password)
		if err != nil {
			return nil, fmt.Errorf(commonstring.IncorrectPasswordOrEmail)
		}
	}

	// handle the case where customer is nil by assigning a default value
	companyName := ""
	if admin != nil {
		companyName = admin.CompanyName
	}
	// sign the token
	accessToken, _ := pkg.NewYptAuth(u.config.Jwt(), pkg.AccessToken, &entities.UserClaims{
		Id:          user.Id,
		Role:        user.Role,
		CompanyName: companyName,
	})

	refreshToken, _ := pkg.NewYptAuth(u.config.Jwt(), pkg.RefreshToken, &entities.UserClaims{
		Id:          user.Id,
		Role:        user.Role,
		CompanyName: companyName,
	})

	// set the user to the passport
	passport := &responses.AdminPassportResponse{
		UserId:   user.Id,
		UserName: admin.UserName,
		Email:    user.Email,
		Roles:    user.Role,
		Token: &responses.UserTokenResponse{
			UserId:       admin.UserId,
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken.SignToken(),
		},
	}

	if err := u.authRepository.InsertAdminOauth(passport); err != nil {
		return nil, err
	}

	return passport, nil
}

func (u *userServiceImpl) RefreshCustomerPassport(req *dto.UserRefreshCredentialRequest) (*responses.CustomerPassportResponse, error) {
	claims, err := pkg.ParseToken(u.config.Jwt(), req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// find the oauth by refresh token
	oauth, err := u.authRepository.FindOneOauthByRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	//Find Profile
	profile, err := u.authRepository.FindOneUserById(oauth.UserId)
	if err != nil {
		return nil, err
	}

	// sign the token
	newClaims := &entities.UserClaims{
		Id:   profile.Id,
		Role: profile.Role,
	}

	accessToken, err := pkg.NewYptAuth(
		u.config.Jwt(),
		pkg.AccessToken,
		newClaims,
	)

	refreshToken := pkg.RepeatToken(
		u.config.Jwt(),
		newClaims,
		claims.ExpiresAt.Unix(),
	)

	// set the user to the passport
	passport := &responses.CustomerPassportResponse{
		UserId: profile.Id,
		Email:  profile.Email,
		Roles:  profile.Role,
		Token: &responses.UserTokenResponse{
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken,
		},
	}

	// update the oauth
	if err := u.authRepository.UpdateOauth(passport.Token); err != nil {
		return nil, err
	}

	return passport, nil
}

func (u *userServiceImpl) RefreshAdminPassport(req *dto.UserRefreshCredentialRequest) (*responses.AdminPassportResponse, error) {
	claims, err := pkg.ParseToken(u.config.Jwt(), req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// find the oauth by refresh token
	oauth, err := u.authRepository.FindOneOauthByRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	//Find Profile
	profile, err := u.authRepository.FindOneUserById(oauth.UserId)
	if err != nil {
		return nil, err
	}

	// sign the token
	newClaims := &entities.UserClaims{
		Id:   profile.Id,
		Role: profile.Role,
	}

	accessToken, err := pkg.NewYptAuth(
		u.config.Jwt(),
		pkg.AccessToken,
		newClaims,
	)

	refreshToken := pkg.RepeatToken(
		u.config.Jwt(),
		newClaims,
		claims.ExpiresAt.Unix(),
	)

	// set the user to the passport
	passport := &responses.AdminPassportResponse{
		UserId: profile.Id,
		Email:  profile.Email,
		Roles:  profile.Role,
		Token: &responses.UserTokenResponse{
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken,
		},
	}

	// update the oauth
	if err := u.authRepository.UpdateOauth(passport.Token); err != nil {
		return nil, err
	}

	return passport, nil
}

func (u *userServiceImpl) DeleteOauth(oauthId string) (*responses.SignOutResponse, error) {
	if err := u.authRepository.DeleteOauth(oauthId); err != nil {
		return nil, err
	}

	responses := &responses.SignOutResponse{
		Success: true,
		Message: commonstring.SignOutSuccessfully,
	}

	return responses, nil

}
