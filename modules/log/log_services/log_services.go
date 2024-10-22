package logservices

import (
	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
	logresponses "github.com/yehpattana/api-yehpattana/modules/log/log_response"
)

type LogServiceInterface interface {
	GetLog() ([]*logresponses.LogResponse, error)
}

func LogServiceImpl(cfg configs.ConfigInterface, logRepository repositories.LogRepositoryInterface) LogServiceInterface {
	return &logServiceImpl{
		config:        cfg,
		logRepository: logRepository,
	}
}

type logServiceImpl struct {
	config        configs.ConfigInterface
	logRepository repositories.LogRepositoryInterface
}

func (c *logServiceImpl) GetLog() ([]*logresponses.LogResponse, error) {
	result, err := c.logRepository.GetLog()
	if err != nil {
		return nil, err
	}

	return result, nil
}
