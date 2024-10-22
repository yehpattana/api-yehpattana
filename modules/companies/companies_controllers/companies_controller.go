package companiescontrollers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/auth/pkg"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	companiesdto "github.com/yehpattana/api-yehpattana/modules/companies/companies_dto"
	companiesservices "github.com/yehpattana/api-yehpattana/modules/companies/companies_services"
)

type CompaniesControllerInterface interface {
	GetCompanies(c *fiber.Ctx) error
	GetCompanyDetail(c *fiber.Ctx) error
	CreateCompany(c *fiber.Ctx) error
	EditCompany(c *fiber.Ctx) error
	DeleteCompany(c *fiber.Ctx) error
}

func CompaniesControllerImpl(cfg configs.ConfigInterface, companiesService companiesservices.CompaniesServiceInterface) CompaniesControllerInterface {
	return &companiesController{
		config:           cfg,
		companiesService: companiesService,
	}
}

type companiesController struct {
	config           configs.ConfigInterface
	companiesService companiesservices.CompaniesServiceInterface
}

// GetCompanies godoc
// @Summary Get companies
// @Description Get companies
// @Router /v1/companies [get]
// @Produces json
// @Tags Companies
// @Security ApiKeyAuth
func (u *companiesController) GetCompanies(c *fiber.Ctx) error {
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	if token == "" {
		return commonresponse.NewResponse(c).Error(fiber.StatusUnauthorized, "Token is required", "Missing Authorization token").Res()
	}
	claims, _ := pkg.ParseToken(u.config.Jwt(), token)
	role := ""
	companyId := ""
	if claims != nil {
		companyId = claims.Claims.CompanyName
		role = claims.Claims.Role
	}
	result, err := u.companiesService.GetAllCompanies(role, companyId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch companies." + err.Error(),
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Companies fetched successfully.",
		"data":    result.Data,
	})

}

// GetCompanyDetail godoc
// @Summary Get company detail
// @Description Get company detail
// @Router /v1/companies/{companyId} [get]
// @Param companyId path string true "Company ID"
// @Produces json
// @Tags Companies
// @Security ApiKeyAuth
func (u *companiesController) GetCompanyDetail(c *fiber.Ctx) error {
	req := new(companiesdto.GetCompanyDetailRequest)
	companyId := c.Params("companyId")

	req.CompanyId = companyId

	result, err := u.companiesService.GetCompanyById(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "Failed to fetch company detail.", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()

}

// CreateCompany godoc
// @Summary Create a company
// @Description Create a company
// @Router /v1/companies/create-company [post]
// @Param company_code formData string true "Company Code"
// @Param company_name formData string true "Company Name"
// @Param logo formData file false "Company Logo"
// @Produces json
// @Tags Companies
// @Security ApiKeyAuth
func (u *companiesController) CreateCompany(c *fiber.Ctx) error {
	req := new(companiesdto.CreateCompanyRequest)

	// assign logo file to request but not require field
	logo, err := c.FormFile("logo")
	if err != nil {
		req.Logo = nil
	}

	req.Logo = logo

	// request body parser
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Company registration failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	result, err := u.companiesService.CreateCompany(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Company registration failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Company registration successful.",
		"data":    result.Data,
	})
}

// EditCompany godoc
// @Summary Edit a company
// @Description Edit a company
// @Router /v1/companies/edit-company [patch]
// @Param company_id formData string true "Company ID"
// @Param company_code formData string false "Company Code"
// @Param company_name formData string false "Company Name"
// @Param logo formData file false "Company Logo"
// @Produces json
// @Tags Companies
// @Security ApiKeyAuth
func (u *companiesController) EditCompany(c *fiber.Ctx) error {
	req := new(companiesdto.EditCompanyRequest)

	// assign logo file to request but not require field
	logo, err := c.FormFile("logo")
	if err != nil {
		req.Logo = nil
	}

	req.Logo = logo

	// request body parser
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Company update failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	result, err := u.companiesService.EditCompanyById(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Company update failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Company update successful.",
		"data":    result.Data,
	})
}

// DeleteCompany godoc
// @Summary Delete a company
// @Description Delete a company
// @Router /v1/companies/delete-company [delete]
// @Param company_id formData string true "Company ID"
// @Produces json
// @Tags Companies
// @Security ApiKeyAuth
func (u *companiesController) DeleteCompany(c *fiber.Ctx) error {
	req := new(companiesdto.DeleteCompanyRequest)

	// request body parser
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Company deletion failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	result, err := u.companiesService.DeleteCompany(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Company deletion failed." + err.Error(),
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": result.Status.Message,
	})
}
