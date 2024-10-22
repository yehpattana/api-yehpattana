package configmenucontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	configmenudto "github.com/yehpattana/api-yehpattana/modules/config_menu/config_menu_dto"
	configmenuservices "github.com/yehpattana/api-yehpattana/modules/config_menu/config_menu_services"
)

type ConfigMenuControllerInterface interface {
	GetAllConfigMenu(c *fiber.Ctx) error
	CreateConfigMenu(c *fiber.Ctx) error
	UpdateConfigMenu(c *fiber.Ctx) error
	DeleteConfigMenu(c *fiber.Ctx) error
	GetAllConfigSubMenu(c *fiber.Ctx) error
	CreateConfigSubMenu(c *fiber.Ctx) error
	UpdateConfigSubMenu(c *fiber.Ctx) error
	DeleteConfigSubMenu(c *fiber.Ctx) error
}

type configMenuControllerImpl struct {
	config            configs.ConfigInterface
	configMenuService configmenuservices.ConfigMenuServiceInterface
}

func ConfigMenuControllerImpl(config configs.ConfigInterface, configMenuService configmenuservices.ConfigMenuServiceInterface) ConfigMenuControllerInterface {
	return &configMenuControllerImpl{
		config:            config,
		configMenuService: configMenuService,
	}
}

func (config *configMenuControllerImpl) GetAllConfigMenu(c *fiber.Ctx) error {
	companyId := c.Params("companyId")
	result, err := config.configMenuService.GetAllConfigMenu(companyId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllConfigMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) CreateConfigMenu(c *fiber.Ctx) error {
	req := new(configmenudto.ConfigMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "CreateConfigMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.CreateConfigMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "CreateConfigMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) UpdateConfigMenu(c *fiber.Ctx) error {
	req := new(configmenudto.UpdateConfigMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "UpdateConfigMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.UpdateConfigMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "UpdateConfigMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) DeleteConfigMenu(c *fiber.Ctx) error {
	req := new(configmenudto.UpdateConfigMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "DeleteConfigMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.DeleteConfigMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "DeleteConfigMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

// subMenu
func (config *configMenuControllerImpl) GetAllConfigSubMenu(c *fiber.Ctx) error {
	companyId := c.Params("companyId")
	result, err := config.configMenuService.GetAllConfigSubMenu(companyId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllConfigSubMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) CreateConfigSubMenu(c *fiber.Ctx) error {
	req := new(configmenudto.ConfigSubMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "CreateConfigSubMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.CreateConfigSubMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "CreateConfigSubMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) UpdateConfigSubMenu(c *fiber.Ctx) error {
	req := new(configmenudto.UpdateConfigSubMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "UpdateConfigMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.UpdateConfigSubMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "UpdateConfigMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *configMenuControllerImpl) DeleteConfigSubMenu(c *fiber.Ctx) error {
	req := new(configmenudto.UpdateConfigSubMenuRequest)
	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "DeleteConfigSubMenu", err.Error()).Res()
	}
	result, err := config.configMenuService.DeleteConfigSubMenu(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "DeleteConfigSubMenu", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
