package logcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	logservices "github.com/yehpattana/api-yehpattana/modules/log/log_services"
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
