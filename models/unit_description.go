package models

import (
	"errors"

	"gorm.io/gorm"
)

// unit_description is struct that represent data in Database
type UnitDescription struct {
	gorm.Model
	UnitID   int    `gorm:"column:unit_id;default:" json:"unit_id"`
	Language string `gorm:"column:language;default:" json:"language"`
	Title    string `gorm:"column:title;default:" json:"title"`
}

// UnitDescription is interface that that model needs to implement
type UnitDescriptionRepository interface {
	Create(unit_description *UnitDescription) error
	GetSingle(id int) (*UnitDescription, error)
	GetAll() ([]UnitDescription, error)
	Update(unit_description *UnitDescription) error
	Delete(id int) error
}

type unit_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewUnitDescriptionRepository(db *gorm.DB) UnitDescriptionRepository {
	return &unit_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *unit_descriptionRepository) Create(unit_description *UnitDescription) error {
	return r.db.Create(unit_description).Error
}

// Function to get single instance of unit_description
func (r *unit_descriptionRepository) GetSingle(id int) (*UnitDescription, error) {
	var unit_description UnitDescription
	err := r.db.First(&unit_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &unit_description, nil
}

// Function to get all instances of unit_description
func (r *unit_descriptionRepository) GetAll() ([]UnitDescription, error) {
	var unit_description []UnitDescription
	err := r.db.Find(&unit_description).Error
	return unit_description, err
}

// Function to update existing instances of unit_description
func (r *unit_descriptionRepository) Update(unit_description *UnitDescription) error {
	return r.db.Save(unit_description).Error
}

// Function to delete single instances of unit_description
func (r *unit_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&UnitDescription{}, id)
	return result.Error
}
