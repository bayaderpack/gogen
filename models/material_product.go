package models

import (
	"errors"

	"gorm.io/gorm"
)

// material_product is struct that represent data in Database
type MaterialProduct struct {
	gorm.Model
	MaterialProductID int     `gorm:"column:material_product_id;default:" json:"material_product_id"`
	MaterialID        int     `gorm:"column:material_id;default:" json:"material_id"`
	CalculationMethod bool    `gorm:"column:calculation_method;default:" json:"calculation_method"`
	StaticPrice       float32 `gorm:"column:static_price;default:0.0000" json:"static_price"`
}

// MaterialProduct is interface that that model needs to implement
type MaterialProductRepository interface {
	Create(material_product *MaterialProduct) error
	GetSingle(id int) (*MaterialProduct, error)
	GetAll() ([]MaterialProduct, error)
	Update(material_product *MaterialProduct) error
	Delete(id int) error
}

type material_productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMaterialProductRepository(db *gorm.DB) MaterialProductRepository {
	return &material_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *material_productRepository) Create(material_product *MaterialProduct) error {
	return r.db.Create(material_product).Error
}

// Function to get single instance of material_product
func (r *material_productRepository) GetSingle(id int) (*MaterialProduct, error) {
	var material_product MaterialProduct
	err := r.db.First(&material_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material_product, nil
}

// Function to get all instances of material_product
func (r *material_productRepository) GetAll() ([]MaterialProduct, error) {
	var material_product []MaterialProduct
	err := r.db.Find(&material_product).Error
	return material_product, err
}

// Function to update existing instances of material_product
func (r *material_productRepository) Update(material_product *MaterialProduct) error {
	return r.db.Save(material_product).Error
}

// Function to delete single instances of material_product
func (r *material_productRepository) Delete(id int) error {
	result := r.db.Delete(&MaterialProduct{}, id)
	return result.Error
}
