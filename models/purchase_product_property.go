package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// purchase_product_property is struct that represent data in Database
type PurchaseProductProperty struct {
	gorm.Model
	PurchaseProductID int `gorm:"column:purchase_product_id;default:" json:"purchase_product_id"`
MaterialPropertyID int `gorm:"column:material_property_id;default:" json:"material_property_id"`

}

// PurchaseProductProperty is interface that that model needs to implement
type PurchaseProductPropertyRepository interface {
	Create(purchase_product_property *PurchaseProductProperty) error
	GetSingle(id int) (*PurchaseProductProperty, error)
	GetAll() ([]PurchaseProductProperty, error)
	Update(purchase_product_property *PurchaseProductProperty) error
	Delete(id int) error
}

type purchase_product_propertyRepository struct {
	db *gorm.DB
}


// Create new instance
func NewPurchaseProductPropertyRepository(db *gorm.DB) PurchaseProductPropertyRepository {
	return &purchase_product_propertyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *purchase_product_propertyRepository) Create(purchase_product_property *PurchaseProductProperty) error {
	return r.db.Create(purchase_product_property).Error
}


//Function to get single instance of purchase_product_property 
func (r *purchase_product_propertyRepository) GetSingle(id int) (*PurchaseProductProperty, error) {
	var purchase_product_property PurchaseProductProperty
	err := r.db.First(&purchase_product_property, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &purchase_product_property, nil
}


//Function to get all instances of purchase_product_property 
func (r *purchase_product_propertyRepository) GetAll() ([]PurchaseProductProperty, error) {
	var purchase_product_property []PurchaseProductProperty
	err := r.db.Find(&purchase_product_property).Error
	return purchase_product_property, err
}

//Function to update existing instances of purchase_product_property 
func (r *purchase_product_propertyRepository) Update(purchase_product_property *PurchaseProductProperty) error {
	return r.db.Save(purchase_product_property).Error
}

//Function to delete single instances of purchase_product_property 
func (r *purchase_product_propertyRepository) Delete(id int) error {
	result := r.db.Delete(&PurchaseProductProperty{}, id)
	return result.Error
}