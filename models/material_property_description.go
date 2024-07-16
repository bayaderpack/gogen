package models

import (
	"errors"

	"gorm.io/gorm"
)

// material_property_description is struct that represent data in Database
type MaterialPropertyDescription struct {
	gorm.Model
	MaterialPropertyID int    `gorm:"column:material_property_id;default:" json:"material_property_id"`
	Language           string `gorm:"column:language;default:" json:"language"`
	Title              string `gorm:"column:title;default:" json:"title"`
}

// MaterialPropertyDescription is interface that that model needs to implement
type MaterialPropertyDescriptionRepository interface {
	Create(material_property_description *MaterialPropertyDescription) error
	GetSingle(id int) (*MaterialPropertyDescription, error)
	GetAll() ([]MaterialPropertyDescription, error)
	Update(material_property_description *MaterialPropertyDescription) error
	Delete(id int) error
}

type material_property_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMaterialPropertyDescriptionRepository(db *gorm.DB) MaterialPropertyDescriptionRepository {
	return &material_property_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *material_property_descriptionRepository) Create(material_property_description *MaterialPropertyDescription) error {
	return r.db.Create(material_property_description).Error
}

// Function to get single instance of material_property_description
func (r *material_property_descriptionRepository) GetSingle(id int) (*MaterialPropertyDescription, error) {
	var material_property_description MaterialPropertyDescription
	err := r.db.First(&material_property_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material_property_description, nil
}

// Function to get all instances of material_property_description
func (r *material_property_descriptionRepository) GetAll() ([]MaterialPropertyDescription, error) {
	var material_property_description []MaterialPropertyDescription
	err := r.db.Find(&material_property_description).Error
	return material_property_description, err
}

// Function to update existing instances of material_property_description
func (r *material_property_descriptionRepository) Update(material_property_description *MaterialPropertyDescription) error {
	return r.db.Save(material_property_description).Error
}

// Function to delete single instances of material_property_description
func (r *material_property_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&MaterialPropertyDescription{}, id)
	return result.Error
}
