package service

import (
	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/morkid/paginate"
)

type IProperty interface {
	GetAll() ([]models.Property, error)
	GetById(id int) (models.Property, error)
	GetByParams(params map[string]interface{}) (paginate.Page, error)
	Create(property models.Property) (models.Property, error)
	Update(property models.Property) error
}
type Service struct {
	repository IProperty
}

func NewService(repository IProperty) *Service {
	return &Service{repository: repository}
}
func (s *Service) GetAll() ([]models.Property, error) {
	return s.repository.GetAll()
}
func (s *Service) GetById(id int) (models.Property, error) {
	return s.repository.GetById(id)
}
func (s *Service) GetByParams(params map[string]interface{}) (paginate.Page, error) {
	return s.repository.GetByParams(params)
}
func (s *Service) Create(property models.Property) (models.Property, error) {
	return s.repository.Create(property)
}
func (s *Service) Update(property models.Property) error {
	return s.repository.Update(property)
}
