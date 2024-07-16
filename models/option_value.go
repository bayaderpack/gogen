package models

import (
	"errors"

	"gorm.io/gorm"
)

// option_value is struct that represent data in Database
type OptionValue struct {
	gorm.Model
	OptionValueID int `gorm:"column:option_value_id;default:" json:"option_value_id"`
	OptionID      int `gorm:"column:option_id;default:" json:"option_id"`
	MediaID       int `gorm:"column:media_id;default:0" json:"media_id"`
	SortOrder     int `gorm:"column:sort_order;default:" json:"sort_order"`
}

// OptionValue is interface that that model needs to implement
type OptionValueRepository interface {
	Create(option_value *OptionValue) error
	GetSingle(id int) (*OptionValue, error)
	GetAll() ([]OptionValue, error)
	Update(option_value *OptionValue) error
	Delete(id int) error
}

type option_valueRepository struct {
	db *gorm.DB
}

// Create new instance
func NewOptionValueRepository(db *gorm.DB) OptionValueRepository {
	return &option_valueRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *option_valueRepository) Create(option_value *OptionValue) error {
	return r.db.Create(option_value).Error
}

// Function to get single instance of option_value
func (r *option_valueRepository) GetSingle(id int) (*OptionValue, error) {
	var option_value OptionValue
	err := r.db.First(&option_value, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &option_value, nil
}

// Function to get all instances of option_value
func (r *option_valueRepository) GetAll() ([]OptionValue, error) {
	var option_value []OptionValue
	err := r.db.Find(&option_value).Error
	return option_value, err
}

// Function to update existing instances of option_value
func (r *option_valueRepository) Update(option_value *OptionValue) error {
	return r.db.Save(option_value).Error
}

// Function to delete single instances of option_value
func (r *option_valueRepository) Delete(id int) error {
	result := r.db.Delete(&OptionValue{}, id)
	return result.Error
}
