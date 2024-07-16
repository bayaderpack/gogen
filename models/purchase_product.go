package models

import (
	"errors"

	"gorm.io/gorm"
)

// purchase_product is struct that represent data in Database
type PurchaseProduct struct {
	gorm.Model
	PurchaseProductID int     `gorm:"column:purchase_product_id;default:" json:"purchase_product_id"`
	PurchaseID        int     `gorm:"column:purchase_id;default:" json:"purchase_id"`
	MaterialProductID int     `gorm:"column:material_product_id;default:" json:"material_product_id"`
	UnitID            int     `gorm:"column:unit_id;default:" json:"unit_id"`
	SupplierRefNumber string  `gorm:"column:supplier_ref_number;default:NULL" json:"supplier_ref_number"`
	Comment           string  `gorm:"column:comment;default:NULL" json:"comment"`
	TaxID             int     `gorm:"column:tax_id;default:" json:"tax_id"`
	UnitPrice         float32 `gorm:"column:unit_price;default:" json:"unit_price"`
	Quantity          int     `gorm:"column:quantity;default:" json:"quantity"`
}

// PurchaseProduct is interface that that model needs to implement
type PurchaseProductRepository interface {
	Create(purchase_product *PurchaseProduct) error
	GetSingle(id int) (*PurchaseProduct, error)
	GetAll() ([]PurchaseProduct, error)
	Update(purchase_product *PurchaseProduct) error
	Delete(id int) error
}

type purchase_productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewPurchaseProductRepository(db *gorm.DB) PurchaseProductRepository {
	return &purchase_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *purchase_productRepository) Create(purchase_product *PurchaseProduct) error {
	return r.db.Create(purchase_product).Error
}

// Function to get single instance of purchase_product
func (r *purchase_productRepository) GetSingle(id int) (*PurchaseProduct, error) {
	var purchase_product PurchaseProduct
	err := r.db.First(&purchase_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &purchase_product, nil
}

// Function to get all instances of purchase_product
func (r *purchase_productRepository) GetAll() ([]PurchaseProduct, error) {
	var purchase_product []PurchaseProduct
	err := r.db.Find(&purchase_product).Error
	return purchase_product, err
}

// Function to update existing instances of purchase_product
func (r *purchase_productRepository) Update(purchase_product *PurchaseProduct) error {
	return r.db.Save(purchase_product).Error
}

// Function to delete single instances of purchase_product
func (r *purchase_productRepository) Delete(id int) error {
	result := r.db.Delete(&PurchaseProduct{}, id)
	return result.Error
}
