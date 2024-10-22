package sizeservices

import (
	"github.com/natersland/b2b-e-commerce-api/configs"
	"github.com/natersland/b2b-e-commerce-api/modules/data/repositories"
	sizedto "github.com/natersland/b2b-e-commerce-api/modules/size/size_dto"
	sizeresponses "github.com/natersland/b2b-e-commerce-api/modules/size/size_response"
)

type SizeServiceInterface interface {
	GetAllSize() ([]*sizeresponses.SizeResponse, error)
	CreateSize(req *sizedto.SizeRequest) (*sizeresponses.SizeCreateResponse, error)
	DeleteSize(req string) (*sizeresponses.SizeCreateResponse, error)
}

func SizeServiceImpl(cfg configs.ConfigInterface, sizeRepository repositories.SizeRepositoryInterface) SizeServiceInterface {
	return &sizeServiceImpl{
		config:         cfg,
		sizeRepository: sizeRepository,
	}
}

type sizeServiceImpl struct {
	config         configs.ConfigInterface
	sizeRepository repositories.SizeRepositoryInterface
}

func (c *sizeServiceImpl) GetAllSize() ([]*sizeresponses.SizeResponse, error) {
	result, err := c.sizeRepository.GetAllSize()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *sizeServiceImpl) CreateSize(req *sizedto.SizeRequest) (*sizeresponses.SizeCreateResponse, error) {
	result, err := c.sizeRepository.CreateSize(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *sizeServiceImpl) DeleteSize(req string) (*sizeresponses.SizeCreateResponse, error) {
	result, err := c.sizeRepository.DeleteSize(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
