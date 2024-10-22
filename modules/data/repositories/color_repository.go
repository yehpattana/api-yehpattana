package repositories

import (
	colordto "github.com/natersland/b2b-e-commerce-api/modules/color/color_dto"
	colorresponses "github.com/natersland/b2b-e-commerce-api/modules/color/color_response"
	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	"gorm.io/gorm"
)

type ColorRepositoryInterface interface {
	GetAllColor() ([]*colorresponses.ColorResponse, error)
	CreateColor(req *colordto.ColorRequest) (*colorresponses.ColorCreateResponse, error)
	DeleteColor(req string) (*colorresponses.ColorCreateResponse, error)
}

type colorRepositoryImpl struct {
	DB *gorm.DB
}

func ColorRepositoryImpl(db *gorm.DB) ColorRepositoryInterface {
	return &colorRepositoryImpl{
		DB: db,
	}
}

func (colorRepository *colorRepositoryImpl) GetAllColor() ([]*colorresponses.ColorResponse, error) {
	var color []entities.Colors
	var colorResponses []*colorresponses.ColorResponse

	result := colorRepository.DB.Table("Colors").Find(&color)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, c := range color {
		colorResponses = append(colorResponses, &colorresponses.ColorResponse{
			ID:       c.Id,
			Name:     c.Name,
			Code:     c.Code,
			CodeName: c.CodeName,
		})
	}

	return colorResponses, nil
}

func (colorRepository *colorRepositoryImpl) CreateColor(req *colordto.ColorRequest) (*colorresponses.ColorCreateResponse, error) {
	color := &entities.Colors{
		Name:     req.Name,
		Code:     req.Code,
		CodeName: req.CodeName,
	}

	result := colorRepository.DB.Create(color)
	if result.Error != nil {
		return nil, result.Error
	}

	return &colorresponses.ColorCreateResponse{
		Success: true,
		Message: "Create Color successful.",
	}, nil
}
func (colorRepository *colorRepositoryImpl) DeleteColor(req string) (*colorresponses.ColorCreateResponse, error) {

	result := colorRepository.DB.Delete(&entities.Colors{}, req)
	if result.Error != nil {
		return &colorresponses.ColorCreateResponse{
			Success: false,
			Message: "Color in use",
		}, nil
	}

	return &colorresponses.ColorCreateResponse{
		Success: true,
		Message: "Delete Color successful.",
	}, nil
}
