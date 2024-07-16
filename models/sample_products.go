package models

import (
	"errors"

	"gorm.io/gorm"
)

// sample_products is struct that represent data in Database
type SampleProducts struct {
	gorm.Model
	SampleID                 int    `gorm:"column:sample_id;default:" json:"sample_id"`
	QuotationProductID       int    `gorm:"column:quotation_product_id;default:NULL" json:"quotation_product_id"`
	EditedQuotationProductID int    `gorm:"column:edited_quotation_product_id;default:NULL" json:"edited_quotation_product_id"`
	Priority                 bool   `gorm:"column:priority;default:0" json:"priority"`
	Design                   bool   `gorm:"column:design;default:0" json:"design"`
	Die                      bool   `gorm:"column:die;default:0" json:"die"`
	Quantity                 int    `gorm:"column:quantity;default:0" json:"quantity"`
	Note                     string `gorm:"column:note;default:NULL" json:"note"`
}

// SampleProducts is interface that that model needs to implement
type SampleProductsRepository interface {
	Create(sample_product *SampleProducts) error
	GetSingle(id int) (*SampleProducts, error)
	GetAll() ([]SampleProducts, error)
	Update(sample_product *SampleProducts) error
	Delete(id int) error
}

type sample_productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSampleProductsRepository(db *gorm.DB) SampleProductsRepository {
	return &sample_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sample_productRepository) Create(sample_product *SampleProducts) error {
	return r.db.Create(sample_product).Error
}

// Function to get single instance of sample_product
func (r *sample_productRepository) GetSingle(id int) (*SampleProducts, error) {
	var sample_product SampleProducts
	err := r.db.First(&sample_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sample_product, nil
}

// Function to get all instances of sample_products
func (r *sample_productRepository) GetAll() ([]SampleProducts, error) {
	var sample_products []SampleProducts
	err := r.db.Find(&sample_products).Error
	return sample_products, err
}

// Function to update existing instances of sample_product
func (r *sample_productRepository) Update(sample_product *SampleProducts) error {
	return r.db.Save(sample_product).Error
}

// Function to delete single instances of sample_product
func (r *sample_productRepository) Delete(id int) error {
	result := r.db.Delete(&SampleProducts{}, id)
	return result.Error
}
