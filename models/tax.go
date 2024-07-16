package models

import (
	"errors"

	"gorm.io/gorm"
)

// tax is struct that represent data in Database
type Tax struct {
	gorm.Model
	TaxID  int     `gorm:"column:tax_id;default:" json:"tax_id"`
	Value  float32 `gorm:"column:value;default:" json:"value"`
	Title  string  `gorm:"column:title;default:" json:"title"`
	Status bool    `gorm:"column:status;default:1" json:"status"`
}

// Tax is interface that that model needs to implement
type TaxRepository interface {
	Create(tax *Tax) error
	GetSingle(id int) (*Tax, error)
	GetAll() ([]Tax, error)
	Update(tax *Tax) error
	Delete(id int) error
}

type taxRepository struct {
	db *gorm.DB
}

// Create new instance
func NewTaxRepository(db *gorm.DB) TaxRepository {
	return &taxRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *taxRepository) Create(tax *Tax) error {
	return r.db.Create(tax).Error
}

// Function to get single instance of tax
func (r *taxRepository) GetSingle(id int) (*Tax, error) {
	var tax Tax
	err := r.db.First(&tax, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &tax, nil
}

// Function to get all instances of tax
func (r *taxRepository) GetAll() ([]Tax, error) {
	var tax []Tax
	err := r.db.Find(&tax).Error
	return tax, err
}

// Function to update existing instances of tax
func (r *taxRepository) Update(tax *Tax) error {
	return r.db.Save(tax).Error
}

// Function to delete single instances of tax
func (r *taxRepository) Delete(id int) error {
	result := r.db.Delete(&Tax{}, id)
	return result.Error
}
