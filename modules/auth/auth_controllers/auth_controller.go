package authcontrollers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	authservices "github.com/yehpattana/api-yehpattana/modules/auth/auth_services"
	"github.com/yehpattana/api-yehpattana/modules/auth/dto"
	"github.com/yehpattana/api-yehpattana/modules/auth/helpers"
	"github.com/yehpattana/api-yehpattana/modules/auth/pkg"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	commonstring "github.com/yehpattana/api-yehpattana/modules/commons/common_string"
)

type authControllerErrorCode string

const (
	signUpCustomerError     authControllerErrorCode = "AUTH_ERR_001"
	signUpAdminError        authControllerErrorCode = "AUTH_ERR_002"
	signUpError             authControllerErrorCode = "AUTH_ERR_003"
	signInCustomerError     authControllerErrorCode = "AUTH_ERR_004"
	signInAdminError        authControllerErrorCode = "AUTH_ERR_005"
	signInError             authControllerErrorCode = "AUTH_ERR_006"
	signOutError            authControllerErrorCode = "AUTH_ERR_007"
	generateAdminTokenError authControllerErrorCode = "AUTH_ERR_008"
)

type AuthControllerInterface interface {
	SignUpCustomer(c *fiber.Ctx) error
	SignUpAdmin(c *fiber.Ctx) error
	SignInCustomer(c *fiber.Ctx) error
	SignInAdmin(c *fiber.Ctx) error
	RefreshCustomerPassport(c *fiber.Ctx) error
	RefreshAdminPassport(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
	GenerateAdminToken(c *fiber.Ctx) error
}

func AuthController(config configs.ConfigInterface, authService authservices.AuthServiceInterface) AuthControllerInterface {
	return &authController{
		config:      config,
		authService: authService,
	}
}

type authController struct {
	config      configs.ConfigInterface
	authService authservices.AuthServiceInterface
}

func (u *authController) signUp(req interface{}, signUpFunc func(dto interface{}) (interface{}, error), c *fiber.Ctx, errorContext string) error {
	// request body parser
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signUpError), err.Error()).Res()
	}

	// Extract email validation
	validateEmail := func(email string) error {
		if !helpers.CheckIsValidEmailPattern(email) {
			return fmt.Errorf(commonstring.InvalidEmailPatternError)
		}
		return nil
	}

	// Email validation for the request
	switch req := req.(type) {
	case dto.CustomerRegisterRequest:
		if err := validateEmail(req.User.Email); err != nil {
			return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signUpCustomerError), err.Error()).Res()
		}
	case dto.AdminRegisterRequest:
		if err := validateEmail(req.Email); err != nil {
			return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signUpAdminError), err.Error()).Res()
		}
	}

	// Error message for the response
	errorMessage := ""
	switch req.(type) {
	case dto.CustomerRegisterRequest:
		errorMessage = string(signUpCustomerError)
	case dto.AdminRegisterRequest:
		errorMessage = string(signUpAdminError)
	}

	// create a new user
	result, err := signUpFunc(req)
	if err != nil {
		var code = fiber.ErrInternalServerError.Code
		if err.Error() == commonstring.EmailAlreadyExistsError {
			code = fiber.ErrBadRequest.Code
		}
		return commonresponse.NewResponse(c).Error(code, errorMessage, err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusCreated, result).Res()

}

func (u *authController) SignUpCustomer(c *fiber.Ctx) error {
	req := new(dto.CustomerRegisterRequest)
	return u.signUp(req, func(dto interface{}) (interface{}, error) {
		// Call CreateCustomer with the req
		return u.authService.CreateCustomer(req)
	}, c, string(signUpCustomerError))
}

// SignUpAdmin godoc
// @Summary Sign up an admin
// @Description Sign up an admin
// @Router /v1/auth/signup/admin [post]
// @Param email formData string true "Email"
// @Param user_name formData string true "User Name"
// @Tags auth
// @Produce json
func (u *authController) SignUpAdmin(c *fiber.Ctx) error {
	req := new(dto.AdminRegisterRequest)

	req.Email = c.FormValue("email")
	req.UserName = c.FormValue("user_name")
	req.CompanyName = c.FormValue("company_name")

	return u.signUp(req, func(dto interface{}) (interface{}, error) {
		// Call CreateAdmin with the req
		return u.authService.CreateAdmin(req)
	}, c, string(signUpAdminError))
}

func (u *authController) SignInCustomer(c *fiber.Ctx) error {
	req := new(dto.UserCredentialRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInError), err.Error()).Res()
	}

	passport, err := u.authService.SignInCustomer(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrUnauthorized.Code, string(signInCustomerError), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (u *authController) SignInAdmin(c *fiber.Ctx) error {
	req := new(dto.UserCredentialRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInError), err.Error()).Res()
	}

	passport, err := u.authService.SignInAdmin(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrUnauthorized.Code, string(signInAdminError), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (u *authController) RefreshCustomerPassport(c *fiber.Ctx) error {
	req := new(dto.UserRefreshCredentialRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInError), err.Error()).Res()
	}

	passport, err := u.authService.RefreshCustomerPassport(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrUnauthorized.Code, string(signInCustomerError), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (u *authController) RefreshAdminPassport(c *fiber.Ctx) error {
	req := new(dto.UserRefreshCredentialRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, string(signInError), err.Error()).Res()
	}

	passport, err := u.authService.RefreshAdminPassport(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrUnauthorized.Code, string(signInAdminError), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (u *authController) SignOut(c *fiber.Ctx) error {
	req := new(dto.UserRemoveCredentialRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signOutError),
			err.Error()).Res()
	}

	_, err := u.authService.DeleteOauth(req.OauthId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(
			fiber.ErrUnauthorized.Code,
			string(signOutError),
			err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, nil).Res()
}

func (u *authController) GenerateAdminToken(c *fiber.Ctx) error {
	adminToken, err := pkg.NewYptAuth(
		u.config.Jwt(),
		pkg.Admin,
		nil,
	)

	if err != nil {
		return commonresponse.NewResponse(c).Error(
			fiber.ErrUnauthorized.Code,
			string(generateAdminTokenError),
			err.Error(),
		).Res()
	}

	return commonresponse.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Token string `json:"token"`
		}{
			Token: adminToken.SignToken(),
		},
	).Res()
}
