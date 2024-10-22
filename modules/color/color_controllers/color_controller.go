package colorcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natersland/b2b-e-commerce-api/configs"
	colordto "github.com/natersland/b2b-e-commerce-api/modules/color/color_dto"
	colorservices "github.com/natersland/b2b-e-commerce-api/modules/color/color_services"
	commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"
)

type ColorControllerInterface interface {
	GetAllColor(c *fiber.Ctx) error
	CreateColor(c *fiber.Ctx) error
	DeleteColor(c *fiber.Ctx) error
}

type colorControllerImpl struct {
	config        configs.ConfigInterface
	colorServices colorservices.ColorServiceInterface
}

func ColorControllerImpl(config configs.ConfigInterface, colorServices colorservices.ColorServiceInterface) ColorControllerInterface {
	return &colorControllerImpl{
		config:        config,
		colorServices: colorServices,
	}
}

func (config *colorControllerImpl) GetAllColor(c *fiber.Ctx) error {
	result, err := config.colorServices.GetAllColor()
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllColor", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *colorControllerImpl) CreateColor(c *fiber.Ctx) error {
	req := new(colordto.ColorRequest)

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "CreateColor", err.Error()).Res()
	}
	result, err := config.colorServices.CreateColor(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "CreateColor", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *colorControllerImpl) DeleteColor(c *fiber.Ctx) error {

	colorId := c.Params("colorId")
	result, err := config.colorServices.DeleteColor(colorId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "DeleteColor", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
