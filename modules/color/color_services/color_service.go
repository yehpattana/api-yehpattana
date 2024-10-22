package colorservices

import (
	"github.com/yehpattana/api-yehpattana/configs"
	colordto "github.com/yehpattana/api-yehpattana/modules/color/color_dto"
	colorresponses "github.com/yehpattana/api-yehpattana/modules/color/color_response"
	"github.com/yehpattana/api-yehpattana/modules/data/repositories"
)

type ColorServiceInterface interface {
	GetAllColor() ([]*colorresponses.ColorResponse, error)
	CreateColor(req *colordto.ColorRequest) (*colorresponses.ColorCreateResponse, error)
	DeleteColor(req string) (*colorresponses.ColorCreateResponse, error)
}

func ColorServiceImpl(cfg configs.ConfigInterface, colorRepository repositories.ColorRepositoryInterface) ColorServiceInterface {
	return &colorServiceImpl{
		config:          cfg,
		colorRepository: colorRepository,
	}
}

type colorServiceImpl struct {
	config          configs.ConfigInterface
	colorRepository repositories.ColorRepositoryInterface
}

func (c *colorServiceImpl) GetAllColor() ([]*colorresponses.ColorResponse, error) {
	result, err := c.colorRepository.GetAllColor()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *colorServiceImpl) CreateColor(req *colordto.ColorRequest) (*colorresponses.ColorCreateResponse, error) {
	result, err := c.colorRepository.CreateColor(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (c *colorServiceImpl) DeleteColor(req string) (*colorresponses.ColorCreateResponse, error) {
	result, err := c.colorRepository.DeleteColor(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
