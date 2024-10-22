package repositories

import (
	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	sizedto "github.com/natersland/b2b-e-commerce-api/modules/size/size_dto"
	sizeresponses "github.com/natersland/b2b-e-commerce-api/modules/size/size_response"
	"gorm.io/gorm"
)

type SizeRepositoryInterface interface {
	GetAllSize() ([]*sizeresponses.SizeResponse, error)
	CreateSize(req *sizedto.SizeRequest) (*sizeresponses.SizeCreateResponse, error)
	DeleteSize(req string) (*sizeresponses.SizeCreateResponse, error)
}

type sizeRepositoryImpl struct {
	DB *gorm.DB
}

func SizeRepositoryImpl(db *gorm.DB) SizeRepositoryInterface {
	return &sizeRepositoryImpl{
		DB: db,
	}
}

func (sizeRepository *sizeRepositoryImpl) GetAllSize() ([]*sizeresponses.SizeResponse, error) {
	var size []entities.Size
	var sizeResponses []*sizeresponses.SizeResponse

	result := sizeRepository.DB.Table("Size").Find(&size)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, c := range size {
		sizeResponses = append(sizeResponses, &sizeresponses.SizeResponse{
			ID:   c.Id,
			Size: c.Size,
		})
	}

	return sizeResponses, nil
}

func (sizeRepository *sizeRepositoryImpl) CreateSize(req *sizedto.SizeRequest) (*sizeresponses.SizeCreateResponse, error) {
	size := &entities.Size{
		Size: req.Size,
	}

	result := sizeRepository.DB.Create(size)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sizeresponses.SizeCreateResponse{
		Success: true,
		Message: "Create Size successful.",
	}, nil
}
func (sizeRepository *sizeRepositoryImpl) DeleteSize(req string) (*sizeresponses.SizeCreateResponse, error) {

	result := sizeRepository.DB.Delete(&entities.Size{}, req)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sizeresponses.SizeCreateResponse{
		Success: true,
		Message: "Delete Size successful.",
	}, nil
}
