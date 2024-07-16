package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// purchase is struct that represent data in Database
type Purchase struct {
	gorm.Model
	PurchaseID    int        `gorm:"column:purchase_id;default:" json:"purchase_id"`
	SupplierID    int        `gorm:"column:supplier_id;default:" json:"supplier_id"`
	AdminID       int        `gorm:"column:admin_id;default:" json:"admin_id"`
	InvoiceNumber string     `gorm:"column:invoice_number;default:" json:"invoice_number"`
	PurchaseDate  *time.Time `gorm:"column:purchase_date;default:current_timestamp()" json:"purchase_date"`
	Gross         float32    `gorm:"column:gross;default:" json:"gross"`
	Discount      float32    `gorm:"column:discount;default:" json:"discount"`
	Taxable       float32    `gorm:"column:taxable;default:" json:"taxable"`
	Tax           float32    `gorm:"column:tax;default:" json:"tax"`
	Net           float32    `gorm:"column:net;default:" json:"net"`
	Status        bool       `gorm:"column:status;default:1" json:"status"`
}

// Purchase is interface that that model needs to implement
type PurchaseRepository interface {
	Create(purchase *Purchase) error
	GetSingle(id int) (*Purchase, error)
	GetAll() ([]Purchase, error)
	Update(purchase *Purchase) error
	Delete(id int) error
}

type purchaseRepository struct {
	db *gorm.DB
}

// Create new instance
func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *purchaseRepository) Create(purchase *Purchase) error {
	return r.db.Create(purchase).Error
}

// Function to get single instance of purchase
func (r *purchaseRepository) GetSingle(id int) (*Purchase, error) {
	var purchase Purchase
	err := r.db.First(&purchase, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &purchase, nil
}

// Function to get all instances of purchase
func (r *purchaseRepository) GetAll() ([]Purchase, error) {
	var purchase []Purchase
	err := r.db.Find(&purchase).Error
	return purchase, err
}

// Function to update existing instances of purchase
func (r *purchaseRepository) Update(purchase *Purchase) error {
	return r.db.Save(purchase).Error
}

// Function to delete single instances of purchase
func (r *purchaseRepository) Delete(id int) error {
	result := r.db.Delete(&Purchase{}, id)
	return result.Error
}
