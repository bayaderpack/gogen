package models

import (
	"errors"

	"gorm.io/gorm"
)

// material_property is struct that represent data in Database
type MaterialProperty struct {
	gorm.Model
	MaterialPropertyID int    `gorm:"column:material_property_id;default:" json:"material_property_id"`
	ParentID           int    `gorm:"column:parent_id;default:0" json:"parent_id"`
	Comment            string `gorm:"column:comment;default:NULL" json:"comment"`
}

// MaterialProperty is interface that that model needs to implement
type MaterialPropertyRepository interface {
	Create(material_property *MaterialProperty) error
	GetSingle(id int) (*MaterialProperty, error)
	GetAll() ([]MaterialProperty, error)
	Update(material_property *MaterialProperty) error
	Delete(id int) error
}

type material_propertyRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMaterialPropertyRepository(db *gorm.DB) MaterialPropertyRepository {
	return &material_propertyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *material_propertyRepository) Create(material_property *MaterialProperty) error {
	return r.db.Create(material_property).Error
}

// Function to get single instance of material_property
func (r *material_propertyRepository) GetSingle(id int) (*MaterialProperty, error) {
	var material_property MaterialProperty
	err := r.db.First(&material_property, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material_property, nil
}

// Function to get all instances of material_property
func (r *material_propertyRepository) GetAll() ([]MaterialProperty, error) {
	var material_property []MaterialProperty
	err := r.db.Find(&material_property).Error
	return material_property, err
}

// Function to update existing instances of material_property
func (r *material_propertyRepository) Update(material_property *MaterialProperty) error {
	return r.db.Save(material_property).Error
}

// Function to delete single instances of material_property
func (r *material_propertyRepository) Delete(id int) error {
	result := r.db.Delete(&MaterialProperty{}, id)
	return result.Error
}
