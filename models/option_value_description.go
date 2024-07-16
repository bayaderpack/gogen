package models

import (
	"errors"

	"gorm.io/gorm"
)

// option_value_description is struct that represent data in Database
type OptionValueDescription struct {
	gorm.Model
	OptionValueID int    `gorm:"column:option_value_id;default:" json:"option_value_id"`
	Language      string `gorm:"column:language;default:" json:"language"`
	Title         string `gorm:"column:title;default:" json:"title"`
	Description   string `gorm:"column:description;default:NULL" json:"description"`
}

// OptionValueDescription is interface that that model needs to implement
type OptionValueDescriptionRepository interface {
	Create(option_value_description *OptionValueDescription) error
	GetSingle(id int) (*OptionValueDescription, error)
	GetAll() ([]OptionValueDescription, error)
	Update(option_value_description *OptionValueDescription) error
	Delete(id int) error
}

type option_value_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewOptionValueDescriptionRepository(db *gorm.DB) OptionValueDescriptionRepository {
	return &option_value_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *option_value_descriptionRepository) Create(option_value_description *OptionValueDescription) error {
	return r.db.Create(option_value_description).Error
}

// Function to get single instance of option_value_description
func (r *option_value_descriptionRepository) GetSingle(id int) (*OptionValueDescription, error) {
	var option_value_description OptionValueDescription
	err := r.db.First(&option_value_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &option_value_description, nil
}

// Function to get all instances of option_value_description
func (r *option_value_descriptionRepository) GetAll() ([]OptionValueDescription, error) {
	var option_value_description []OptionValueDescription
	err := r.db.Find(&option_value_description).Error
	return option_value_description, err
}

// Function to update existing instances of option_value_description
func (r *option_value_descriptionRepository) Update(option_value_description *OptionValueDescription) error {
	return r.db.Save(option_value_description).Error
}

// Function to delete single instances of option_value_description
func (r *option_value_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&OptionValueDescription{}, id)
	return result.Error
}
