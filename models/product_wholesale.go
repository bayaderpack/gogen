package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_wholesale is struct that represent data in Database
type ProductWholesale struct {
	gorm.Model
	ProductID int     `gorm:"column:product_id;default:" json:"product_id"`
	GroupID   int     `gorm:"column:group_id;default:" json:"group_id"`
	Quantity  int     `gorm:"column:quantity;default:" json:"quantity"`
	Price     float32 `gorm:"column:price;default:0.00" json:"price"`
}

// ProductWholesale is interface that that model needs to implement
type ProductWholesaleRepository interface {
	Create(product_wholesale *ProductWholesale) error
	GetSingle(id int) (*ProductWholesale, error)
	GetAll() ([]ProductWholesale, error)
	Update(product_wholesale *ProductWholesale) error
	Delete(id int) error
}

type product_wholesaleRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductWholesaleRepository(db *gorm.DB) ProductWholesaleRepository {
	return &product_wholesaleRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_wholesaleRepository) Create(product_wholesale *ProductWholesale) error {
	return r.db.Create(product_wholesale).Error
}

// Function to get single instance of product_wholesale
func (r *product_wholesaleRepository) GetSingle(id int) (*ProductWholesale, error) {
	var product_wholesale ProductWholesale
	err := r.db.First(&product_wholesale, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_wholesale, nil
}

// Function to get all instances of product_wholesale
func (r *product_wholesaleRepository) GetAll() ([]ProductWholesale, error) {
	var product_wholesale []ProductWholesale
	err := r.db.Find(&product_wholesale).Error
	return product_wholesale, err
}

// Function to update existing instances of product_wholesale
func (r *product_wholesaleRepository) Update(product_wholesale *ProductWholesale) error {
	return r.db.Save(product_wholesale).Error
}

// Function to delete single instances of product_wholesale
func (r *product_wholesaleRepository) Delete(id int) error {
	result := r.db.Delete(&ProductWholesale{}, id)
	return result.Error
}
