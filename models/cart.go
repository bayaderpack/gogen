package models

import (
	"errors"

	"gorm.io/gorm"
)

// cart is struct that represent data in Database
type Cart struct {
	gorm.Model
	ItemID     int    `gorm:"column:item_id;default:" json:"item_id"`
	ProductID  int    `gorm:"column:product_id;default:" json:"product_id"`
	CustomerID string `gorm:"column:customer_id;default:" json:"customer_id"`
	Quantity   int    `gorm:"column:quantity;default:" json:"quantity"`
	Options    string `gorm:"column:options;default:" json:"options"`
	ItemType   bool   `gorm:"column:item_type;default:1" json:"item_type"`
}

// Cart is interface that that model needs to implement
type CartRepository interface {
	Create(cart *Cart) error
	GetSingle(id int) (*Cart, error)
	GetAll() ([]Cart, error)
	Update(cart *Cart) error
	Delete(id int) error
}

type cartRepository struct {
	db *gorm.DB
}

// Create new instance
func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *cartRepository) Create(cart *Cart) error {
	return r.db.Create(cart).Error
}

// Function to get single instance of cart
func (r *cartRepository) GetSingle(id int) (*Cart, error) {
	var cart Cart
	err := r.db.First(&cart, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &cart, nil
}

// Function to get all instances of cart
func (r *cartRepository) GetAll() ([]Cart, error) {
	var cart []Cart
	err := r.db.Find(&cart).Error
	return cart, err
}

// Function to update existing instances of cart
func (r *cartRepository) Update(cart *Cart) error {
	return r.db.Save(cart).Error
}

// Function to delete single instances of cart
func (r *cartRepository) Delete(id int) error {
	result := r.db.Delete(&Cart{}, id)
	return result.Error
}
