package models

import (
	"errors"

	"gorm.io/gorm"
)

// options is struct that represent data in Database
type Options struct {
	gorm.Model
	OptionID   int    `gorm:"column:option_id;default:" json:"option_id"`
	Type       string `gorm:"column:type;default:" json:"type"`
	Status     bool   `gorm:"column:status;default:1" json:"status"`
	SortOrder  int    `gorm:"column:sort_order;default:" json:"sort_order"`
	Filterable bool   `gorm:"column:filterable;default:" json:"filterable"`
}

// Options is interface that that model needs to implement
type OptionsRepository interface {
	Create(option *Options) error
	GetSingle(id int) (*Options, error)
	GetAll() ([]Options, error)
	Update(option *Options) error
	Delete(id int) error
}

type optionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewOptionsRepository(db *gorm.DB) OptionsRepository {
	return &optionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *optionRepository) Create(option *Options) error {
	return r.db.Create(option).Error
}

// Function to get single instance of option
func (r *optionRepository) GetSingle(id int) (*Options, error) {
	var option Options
	err := r.db.First(&option, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &option, nil
}

// Function to get all instances of options
func (r *optionRepository) GetAll() ([]Options, error) {
	var options []Options
	err := r.db.Find(&options).Error
	return options, err
}

// Function to update existing instances of option
func (r *optionRepository) Update(option *Options) error {
	return r.db.Save(option).Error
}

// Function to delete single instances of option
func (r *optionRepository) Delete(id int) error {
	result := r.db.Delete(&Options{}, id)
	return result.Error
}
