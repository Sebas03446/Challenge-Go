package validation

import (
	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/go-playground/validator/v10"
)

type Rectangle struct {
	X0 float64
	Y0 float64
	X2 float64
	Y2 float64
}
type Point struct {
	X float64
	Y float64
}

func CountryValidation(property models.Property, validate *validator.Validate) string {
	rect := Rectangle{-99.296741, 19.296134, -98.916339, 19.661237}
	point := Point{property.Location.Longitude, property.Location.Latitude}
	if point.X >= rect.X0 && point.X <= rect.X2 && point.Y >= rect.Y0 && point.Y <= rect.Y2 {
		if property.Pricing.SalePrice >= 1000000 && property.Pricing.SalePrice <= 15000000 {
			return "ACTIVE"
		}
		return "INVALID"
	}
	if property.Pricing.SalePrice >= 50000000 && property.Pricing.SalePrice <= 350000000 {
		return "INACTIVE"
	}

	return "INVALID"

}
