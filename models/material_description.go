package models

import (
	"errors"

	"gorm.io/gorm"
)

// material_description is struct that represent data in Database
type MaterialDescription struct {
	gorm.Model
	MaterialID int    `gorm:"column:material_id;default:" json:"material_id"`
	Language   string `gorm:"column:language;default:" json:"language"`
	Title      string `gorm:"column:title;default:" json:"title"`
}

// MaterialDescription is interface that that model needs to implement
type MaterialDescriptionRepository interface {
	Create(material_description *MaterialDescription) error
	GetSingle(id int) (*MaterialDescription, error)
	GetAll() ([]MaterialDescription, error)
	Update(material_description *MaterialDescription) error
	Delete(id int) error
}

type material_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMaterialDescriptionRepository(db *gorm.DB) MaterialDescriptionRepository {
	return &material_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *material_descriptionRepository) Create(material_description *MaterialDescription) error {
	return r.db.Create(material_description).Error
}

// Function to get single instance of material_description
func (r *material_descriptionRepository) GetSingle(id int) (*MaterialDescription, error) {
	var material_description MaterialDescription
	err := r.db.First(&material_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material_description, nil
}

// Function to get all instances of material_description
func (r *material_descriptionRepository) GetAll() ([]MaterialDescription, error) {
	var material_description []MaterialDescription
	err := r.db.Find(&material_description).Error
	return material_description, err
}

// Function to update existing instances of material_description
func (r *material_descriptionRepository) Update(material_description *MaterialDescription) error {
	return r.db.Save(material_description).Error
}

// Function to delete single instances of material_description
func (r *material_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&MaterialDescription{}, id)
	return result.Error
}
