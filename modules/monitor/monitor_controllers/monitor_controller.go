package monitorcontrollers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yehpattana/api-yehpattana/configs"
	commonresponse "github.com/yehpattana/api-yehpattana/modules/commons/common_response"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
)

type MonitorControllerInterface interface {
	HealthCheck(c *fiber.Ctx) error
}

func MonitorController(config configs.ConfigInterface) MonitorControllerInterface {
	return &monitorController{
		config: config,
	}
}

type monitorController struct {
	config configs.ConfigInterface
}

// HealthCheck godoc
// @Summary เชคว่า service ทำงานได้หรือไม่
// @Description เชคว่า service ทำงานได้หรือไม่
// @Router /v1 [get]
// @Produce json
// @Tags Monitor
func (m *monitorController) HealthCheck(c *fiber.Ctx) error {
	res := &entities.Monitor{
		Name:    m.config.Service().Name(),
		Version: m.config.Service().Version(),
	}
	return commonresponse.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
