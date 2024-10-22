package repositories

import (
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
	logresponses "github.com/yehpattana/api-yehpattana/modules/log/log_response"
	"gorm.io/gorm"
)

type LogRepositoryInterface interface {
	GetLog() ([]*logresponses.LogResponse, error)
}

type logRepositoryImpl struct {
	DB *gorm.DB
}

func LogRepositoryImpl(db *gorm.DB) LogRepositoryInterface {
	return &logRepositoryImpl{
		DB: db,
	}
}

func (logRepository *logRepositoryImpl) GetLog() ([]*logresponses.LogResponse, error) {
	var log []entities.Log
	var logResponses []*logresponses.LogResponse

	result := logRepository.DB.Table("Log").Find(&log)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, c := range log {
		logResponses = append(logResponses, &logresponses.LogResponse{
			Id:          c.Id,
			EndPoint:    "",
			Description: c.Description,
			UpdatedBy:   c.UpdatedBy,
			CreatedAt:   c.CreatedAt,
		})
	}

	return logResponses, nil
}
