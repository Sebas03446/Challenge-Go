package validation

import (
	"errors"

	"github.com/Sebas03446/Challenge-Go/pkg/domain/models"
	"github.com/go-playground/validator/v10"
)

func ChoicePropertyType(property models.Property, validate *validator.Validate) []string {
	if property.PropertyType == "HOUSE" {
		err := HouseValidation(property, validate)
		if err != nil {
			return err
		}
	} else if property.PropertyType == "APARTMENT" {
		err := ApartmentValidation(property, validate)
		if err != nil {
			return err
		}
	} else {
		return []string{"Property type must be HOUSE or APARTMENT"}
	}
	return nil
}

func HouseValidation(property models.Property, validate *validator.Validate) []string {
	var errs []error
	err := validate.Var(property.Bedrooms, "min=1,max=14")
	if err != nil {
		errs = append(errs, errors.New("bedrooms must be between 1 and 14"))
	}
	err = validate.Var(property.Bathrooms, "min=1,max=12")
	if err != nil {
		errs = append(errs, errors.New("bathrooms must be between 1 and 12"))
	}
	err = validate.Var(property.Area, "min=50,max=3000")
	if err != nil {
		errs = append(errs, errors.New("area must be between 50 and 3000"))
	}
	if len(errs) > 0 {
		return errToString(errs)
	}
	return nil
}
func ApartmentValidation(property models.Property, validate *validator.Validate) []string {
	var errs []error
	err := validate.Var(property.Bedrooms, "min=1,max=6")
	if err != nil {
		errs = append(errs, errors.New("bedrooms must be between 1 and 6"))
	}
	err = validate.Var(property.Bathrooms, "min=1,max=4")
	if err != nil {
		errs = append(errs, errors.New("bathrooms must be between 1 and 4"))
	}
	err = validate.Var(property.Area, "min=40,max=400")
	if err != nil {
		errs = append(errs, errors.New("area must be between 40 and 400"))
	}
	if len(errs) > 0 {
		return errToString(errs)
	}
	return nil
}
func errToString(err []error) []string {
	var errs []string
	for _, error := range err {
		errs = append(errs, error.Error())
	}
	return errs
}
