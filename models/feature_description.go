package models

import (
	"errors"

	"gorm.io/gorm"
)

// feature_description is struct that represent data in Database
type FeatureDescription struct {
	gorm.Model
	FeatureID   int    `gorm:"column:feature_id;default:" json:"feature_id"`
	Language    string `gorm:"column:language;default:" json:"language"`
	Description string `gorm:"column:description;default:" json:"description"`
}

// FeatureDescription is interface that that model needs to implement
type FeatureDescriptionRepository interface {
	Create(feature_description *FeatureDescription) error
	GetSingle(id int) (*FeatureDescription, error)
	GetAll() ([]FeatureDescription, error)
	Update(feature_description *FeatureDescription) error
	Delete(id int) error
}

type feature_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewFeatureDescriptionRepository(db *gorm.DB) FeatureDescriptionRepository {
	return &feature_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *feature_descriptionRepository) Create(feature_description *FeatureDescription) error {
	return r.db.Create(feature_description).Error
}

// Function to get single instance of feature_description
func (r *feature_descriptionRepository) GetSingle(id int) (*FeatureDescription, error) {
	var feature_description FeatureDescription
	err := r.db.First(&feature_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &feature_description, nil
}

// Function to get all instances of feature_description
func (r *feature_descriptionRepository) GetAll() ([]FeatureDescription, error) {
	var feature_description []FeatureDescription
	err := r.db.Find(&feature_description).Error
	return feature_description, err
}

// Function to update existing instances of feature_description
func (r *feature_descriptionRepository) Update(feature_description *FeatureDescription) error {
	return r.db.Save(feature_description).Error
}

// Function to delete single instances of feature_description
func (r *feature_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&FeatureDescription{}, id)
	return result.Error
}
