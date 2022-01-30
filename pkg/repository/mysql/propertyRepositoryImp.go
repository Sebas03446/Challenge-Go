package mysql

import (
	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"gorm.io/gorm"
)

type propertyImpl struct {
	db *gorm.DB
}

func NewQueryBuilder(db *gorm.DB) *propertyImpl {
	return &propertyImpl{
		db: db,
	}
}
func (p *propertyImpl) GetAll() ([]models.Property, error) {
	var properties []models.Property
	err := p.db.Find(&properties).Error
	return properties, err
}
func (p *propertyImpl) GetById(id int) (models.Property, error) {
	var property models.Property
	err := p.db.First(&property, id).Error
	return property, err
}
func (p *propertyImpl) Create(property models.Property) (models.Property, error) {
	err := p.db.Create(&property).Error
	return property, err
}
func (p *propertyImpl) Update(property models.Property) error {
	err := p.db.Save(&property).Error
	return err
}
