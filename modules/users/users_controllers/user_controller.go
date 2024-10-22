package userscontrollers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/natersland/b2b-e-commerce-api/configs"
	"github.com/natersland/b2b-e-commerce-api/modules/auth/pkg"
	commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"
	commonstring "github.com/natersland/b2b-e-commerce-api/modules/commons/common_string"
	userhelpers "github.com/natersland/b2b-e-commerce-api/modules/users/user_helpers"
	usersdto "github.com/natersland/b2b-e-commerce-api/modules/users/users_dto"
	usersservices "github.com/natersland/b2b-e-commerce-api/modules/users/users_services"
)

type userControllerErrorCode string

const (
	cantGetAllCustomers   userControllerErrorCode = "USER_ERR_001"
	cantGetCustomerDetail userControllerErrorCode = "USER_ERR_002"
	cantChangePassword    userControllerErrorCode = "USER_ERR_003"
	cantVerifyUser        userControllerErrorCode = "USER_ERR_004"
	cantBanUser           userControllerErrorCode = "USER_ERR_005"
)

type UserControllerInterface interface {
	GetAllCustomers(c *fiber.Ctx) error
	GetCustomerDetail(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	VerifyUser(c *fiber.Ctx) error
	BanUser(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}

func UserControllerImpl(config configs.ConfigInterface, userService usersservices.UserServiceInterface) UserControllerInterface {
	return &userControllerImpl{
		config:      config,
		userService: userService,
	}
}

type userControllerImpl struct {
	config      configs.ConfigInterface
	userService usersservices.UserServiceInterface
}

func (u *userControllerImpl) GetAllCustomers(c *fiber.Ctx) error {
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbXMiOnsiQ29tcGFueU5hbWUiOiIiLCJpZCI6ImU4N2I1MzhhLTE0YTUtNGFmZS1hZGM5LWE3MTA3ZDI5MzUzMSIsInJvbGUiOiJTdXBlckFkbWluIn0sImlzcyI6ImIyYi15ZWhwYXR0YW5hLWFwaSIsInN1YiI6ImFjY2Vzcy10b2tlbiIsImF1ZCI6WyJjdXN0b21lciIsImFkbWluIl0sImV4cCI6MTcyNTk5Nzc5OCwibmJmIjoxNzI0OTk3Nzk5LCJpYXQiOjE3MjQ5OTc3OTl9.bI8ov0s-M_Tw55WK5NI9hezMKVmoLlMP6arXsr-U3LQ"
	if token == "" {
		return commonresponse.NewResponse(c).Error(fiber.StatusUnauthorized, "Token is required", "Missing Authorization token").Res()
	}
	claims, _ := pkg.ParseToken(u.config.Jwt(), token)
	role := ""
	companyId := ""
	if claims == nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantVerifyUser), "Please Login Again").Res()
	}
	companyId = claims.Claims.CompanyName
	role = claims.Claims.Role

	result, err := u.userService.GetAllCustomers(role, companyId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetAllCustomers), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (u *userControllerImpl) GetCustomerDetail(c *fiber.Ctx) error {
	userId := c.Params("userId")

	result, err := u.userService.GetCustomerDetail(userId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantGetCustomerDetail), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (u *userControllerImpl) ChangePassword(c *fiber.Ctx) error {
	userId := c.Params("userId")
	fmt.Println("🚩userId controller", userId)

	// Parse request body
	req := new(usersdto.ChangePasswordRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, string(cantChangePassword), err.Error()).Res()
	}

	// Get newPassword from the parsed request body
	newPassword := req.NewPassword
	fmt.Println("🚩newPassword", newPassword)

	// Check if the password is valid
	if !userhelpers.CheckIsValidPassword(newPassword) {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, string(cantChangePassword), string(commonstring.PasswordMustBeValid)).Res()
	}

	// Call userService.ChangePassword with userId and newPassword
	result, err := u.userService.ChangePassword(userId, newPassword)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantChangePassword), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (u *userControllerImpl) VerifyUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	result, err := u.userService.VerifyUser(userId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantVerifyUser), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (u *userControllerImpl) BanUser(c *fiber.Ctx) error {
	userId := c.Params("userId")

	result, err := u.userService.BanUser(userId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(cantBanUser), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (u *userControllerImpl) UpdateCustomer(c *fiber.Ctx) error {
	req := new(usersdto.UpdateCustomerRequest)
	userId := c.Params("userId")

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusBadRequest, string(commonstring.CantUpdateCustomer), err.Error()).Res()
	}

	result, err := u.userService.UpdateCustomer(req, userId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(commonstring.CantUpdateCustomer), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()

}

func (u *userControllerImpl) DeleteCustomer(c *fiber.Ctx) error {
	userId := c.Params("userId")
	customerId := c.Params("customerId")

	result, err := u.userService.DeleteCustomer(userId, customerId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, string(commonstring.CantUpdateCustomer), err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()

}
