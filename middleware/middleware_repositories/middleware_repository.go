package middlewarerepositories

import (
	"github.com/natersland/b2b-e-commerce-api/middleware"
	"gorm.io/gorm"
)

type MiddlewareRepositoryInterface interface {
	FindAccessToken(userId string, accessToken string) bool
}

func MiddlewareRepositoryImpl(db *gorm.DB) MiddlewareRepositoryInterface {
	return &middlewareRepositoryImpl{
		DB: db,
	}
}

type middlewareRepositoryImpl struct {
	*gorm.DB
}

func (middlewareRepository *middlewareRepositoryImpl) FindAccessToken(userId string, accessToken string) bool {
	var count int64

	query := new(middleware.FindAccessToken)
	result := middlewareRepository.DB.Table("oauth").Where("user_id = ? AND access_token = ?", userId, accessToken).Scan(query)
	if result.Error != nil || count == 0 {
		return false
	}
	return true
}
