package middlewareservices

import middlewarerepositories "github.com/natersland/b2b-e-commerce-api/middleware/middleware_repositories"

type MiddlewareServiceInterface interface {
	FindAccessToken(userId string, accessToken string) bool
}

func MiddlewareServiceImpl(middlewareRepository middlewarerepositories.MiddlewareRepositoryInterface) MiddlewareServiceInterface {
	return &middlewareServiceImpl{
		middlewareRepository: middlewareRepository,
	}
}

type middlewareServiceImpl struct {
	middlewareRepository middlewarerepositories.MiddlewareRepositoryInterface
}

func (m *middlewareServiceImpl) FindAccessToken(userId string, accessToken string) bool {
	return m.middlewareRepository.FindAccessToken(userId, accessToken)
}