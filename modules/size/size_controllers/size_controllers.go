package sizecontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	sizedto "github.com/yehpattana/api-yehpattana/modules/size/size_dto"
	sizeservices "github.com/yehpattana/api-yehpattana/modules/size/size_services"
)

type SizeControllerInterface interface {
	GetAllSize(c *fiber.Ctx) error
	CreateSize(c *fiber.Ctx) error
	DeleteSize(c *fiber.Ctx) error
}

type sizeControllerImpl struct {
	config       configs.ConfigInterface
	sizeServices sizeservices.SizeServiceInterface
}

func SizeControllerImpl(config configs.ConfigInterface, sizeServices sizeservices.SizeServiceInterface) SizeControllerInterface {
	return &sizeControllerImpl{
		config:       config,
		sizeServices: sizeServices,
	}
}

func (config *sizeControllerImpl) GetAllSize(c *fiber.Ctx) error {
	result, err := config.sizeServices.GetAllSize()
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllSize", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *sizeControllerImpl) CreateSize(c *fiber.Ctx) error {
	req := new(sizedto.SizeRequest)

	if err := c.BodyParser(req); err != nil {
		return commonresponse.NewResponse(c).Error(fiber.ErrBadRequest.Code, "CreateSize", err.Error()).Res()
	}
	result, err := config.sizeServices.CreateSize(req)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "CreateSize", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}

func (config *sizeControllerImpl) DeleteSize(c *fiber.Ctx) error {
	sizeId := c.Params("sizeId")
	result, err := config.sizeServices.DeleteSize(sizeId)
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "DeleteSize", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
