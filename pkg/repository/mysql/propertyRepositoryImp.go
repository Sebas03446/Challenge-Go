package mysql

import (
	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/morkid/paginate"
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
func (p *propertyImpl) GetByParams(params map[string]interface{}) (paginate.Page, error) {
	pg := paginate.New(&paginate.Config{
		DefaultSize: 10,
	})
	var model *gorm.DB
	var properties []models.Property
	x0, errL := params["x0"]
	y0, errLa := params["y0"]
	x1, errLb := params["x1"]
	y1, errLc := params["y1"]
	status, errS := params["status"]
	if errL && errLa && errLb && errLc && errS {
		if status == "ACTIVE" || status == "INACTIVE" {
			model = p.db.Where("longitude BETWEEN ? AND ? AND latitude BETWEEN ? AND ? AND status = ?", x0, x1, y0, y1, status).Find(&properties)
		} else {
			model = p.db.Where("longitude BETWEEN ? AND ? AND latitude BETWEEN ? AND ?", x0, x1, y0, y1).Find(&properties)
		}
	} else if errL && errLa && errLb && errLc {
		model = p.db.Where("longitude BETWEEN ? AND ? AND latitude BETWEEN ? AND ?", x0, x1, y0, y1).Find(&properties)
	} else {
		if errS && (status == "ACTIVE" || status == "INACTIVE") {
			model = p.db.Where("status=?", status).Find(&properties)
		} else {
			model = p.db.Find(&properties)
		}
	}
	paginated := pg.With(model).Request(params).Response(&properties)
	return paginated, model.Error
}
func (p *propertyImpl) Create(property models.Property) (models.Property, error) {
	err := p.db.Create(&property).Error
	return property, err
}
func (p *propertyImpl) Update(property models.Property) error {
	err := p.db.Save(&property).Error
	return err
}
