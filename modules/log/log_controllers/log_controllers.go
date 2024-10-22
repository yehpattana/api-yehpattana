package logcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natersland/b2b-e-commerce-api/configs"
	commonresponse "github.com/natersland/b2b-e-commerce-api/modules/commons/common_response"
	logservices "github.com/natersland/b2b-e-commerce-api/modules/log/log_services"
)

type LogControllerInterface interface {
	GetLog(c *fiber.Ctx) error
}

type logControllerImpl struct {
	config      configs.ConfigInterface
	logServices logservices.LogServiceInterface
}

func LogControllerImpl(config configs.ConfigInterface, logServices logservices.LogServiceInterface) LogControllerInterface {
	return &logControllerImpl{
		config:      config,
		logServices: logServices,
	}
}

func (config *logControllerImpl) GetLog(c *fiber.Ctx) error {
	result, err := config.logServices.GetLog()
	if err != nil {
		return commonresponse.NewResponse(c).Error(fiber.StatusInternalServerError, "GetAllSize", err.Error()).Res()
	}

	return commonresponse.NewResponse(c).Success(fiber.StatusOK, result).Res()
}
