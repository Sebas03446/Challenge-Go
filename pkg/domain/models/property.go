package models

import (
	"time"
)

type Iproperty interface {
	GetId() uint
	GetTitle() string
	GetDescription() string
	GetSalePrice() float64
	GetAdministrativeFee() float64
	GetPropertyType() string
	GetBedrooms() int
	GetBathrooms() int
	GetGarages() int
	GetArea() float64
	GetImage() []Images
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetStatus() string
	SetId(uint)
	SetTitle(string)
	SetDescription(string)
	SetSalePrice(float64)
	SetAdministrativeFee(float64)
	SetPropertyType(string)
	SetBedrooms(int)
	SetBathrooms(int)
	SetGarages(int)
	SetArea(float64)
	SetImage([]Images)
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	SetStatus(string)
} /* voy aqui, me falta hacer las validaciones*/

type Location struct {
	Latitude  float64 `validate:"required"`
	Longitude float64 `validate:"required"`
}
type Images struct {
	ID      uint `gorm:"primary_key"`
	Address string
}
type Pricing struct {
	SalePrice         float64 `validate:"required"`
	AdministrativeFee float64
}
type Property struct {
	ID           uint   `gorm:"primary_key"`
	Title        string `validate:"required"`
	Description  string
	Location     Location `gorm:"embedded"`
	Pricing      Pricing  `gorm:"embedded"`
	PropertyType string   `validate:"required"`
	Bedrooms     int      `validate:"required"`
	Bathrooms    int      `validate:"required"`
	ParkingSpots int
	Area         float64  `validate:"required"`
	Image        []Images `gorm:"many2many:property_image;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Status       string
}
type House struct {
	Property
}
type Apartment struct {
	Property
}

func (p *Property) GetId() uint {
	return p.ID
}
func (p *Property) GetTitle() string {
	return p.Title
}
func (p *Property) GetDescription() string {
	return p.Description
}
func (p *Property) GetSalePrice() float64 {
	return p.Pricing.SalePrice
}
func (p *Property) GetAdministrativeFee() float64 {
	return p.Pricing.AdministrativeFee
}
func (p *Property) GetPropertyType() string {
	return p.PropertyType
}
func (p *Property) GetBedrooms() int {
	return p.Bedrooms
}
func (p *Property) GetBathrooms() int {
	return p.Bathrooms
}
func (p *Property) GetGarages() int {
	return p.ParkingSpots
}
func (p *Property) GetArea() float64 {
	return p.Area
}
func (p *Property) GetImage() []Images {
	return p.Image
}
func (p *Property) GetCreatedAt() time.Time {
	return p.CreatedAt
}
func (p *Property) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}
func (p *Property) GetStatus() string {
	return p.Status
}
func (p *Property) SetId(id uint) {
	p.ID = id
}
func (p *Property) SetTitle(title string) {
	p.Title = title
}
func (p *Property) SetDescription(description string) {
	p.Description = description
}
func (p *Property) SetSalePrice(salePrice float64) {
	p.Pricing.SalePrice = salePrice
}
func (p *Property) SetAdministrativeFee(administrativeFee float64) {
	p.Pricing.AdministrativeFee = administrativeFee
}
func (p *Property) SetPropertyType(propertyType string) {
	p.PropertyType = propertyType
}
func (p *Property) SetBedrooms(bedrooms int) {
	p.Bedrooms = bedrooms
}
func (p *Property) SetBathrooms(bathrooms int) {
	p.Bathrooms = bathrooms
}
func (p *Property) SetGarages(garages int) {
	p.ParkingSpots = garages
}
func (p *Property) SetArea(area float64) {
	p.Area = area
}
func (p *Property) SetImage(image []Images) {
	p.Image = image
}
func (p *Property) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}
func (p *Property) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}
func (p *Property) SetStatus(status string) {
	p.Status = status
}
