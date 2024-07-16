package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_compare is struct that represent data in Database
type ProductCompare struct {
	gorm.Model
	CustomerID string `gorm:"column:customer_id;default:" json:"customer_id"`
	ProductID  int    `gorm:"column:product_id;default:" json:"product_id"`
}

// ProductCompare is interface that that model needs to implement
type ProductCompareRepository interface {
	Create(product_compare *ProductCompare) error
	GetSingle(id int) (*ProductCompare, error)
	GetAll() ([]ProductCompare, error)
	Update(product_compare *ProductCompare) error
	Delete(id int) error
}

type product_compareRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductCompareRepository(db *gorm.DB) ProductCompareRepository {
	return &product_compareRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_compareRepository) Create(product_compare *ProductCompare) error {
	return r.db.Create(product_compare).Error
}

// Function to get single instance of product_compare
func (r *product_compareRepository) GetSingle(id int) (*ProductCompare, error) {
	var product_compare ProductCompare
	err := r.db.First(&product_compare, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_compare, nil
}

// Function to get all instances of product_compare
func (r *product_compareRepository) GetAll() ([]ProductCompare, error) {
	var product_compare []ProductCompare
	err := r.db.Find(&product_compare).Error
	return product_compare, err
}

// Function to update existing instances of product_compare
func (r *product_compareRepository) Update(product_compare *ProductCompare) error {
	return r.db.Save(product_compare).Error
}

// Function to delete single instances of product_compare
func (r *product_compareRepository) Delete(id int) error {
	result := r.db.Delete(&ProductCompare{}, id)
	return result.Error
}
