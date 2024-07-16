package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// product_view is struct that represent data in Database
type ProductView struct {
	gorm.Model
	ProductID int `gorm:"column:product_id;default:" json:"product_id"`
CustomerID string `gorm:"column:customer_id;default:" json:"customer_id"`

}

// ProductView is interface that that model needs to implement
type ProductViewRepository interface {
	Create(product_view *ProductView) error
	GetSingle(id int) (*ProductView, error)
	GetAll() ([]ProductView, error)
	Update(product_view *ProductView) error
	Delete(id int) error
}

type product_viewRepository struct {
	db *gorm.DB
}


// Create new instance
func NewProductViewRepository(db *gorm.DB) ProductViewRepository {
	return &product_viewRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_viewRepository) Create(product_view *ProductView) error {
	return r.db.Create(product_view).Error
}


//Function to get single instance of product_view 
func (r *product_viewRepository) GetSingle(id int) (*ProductView, error) {
	var product_view ProductView
	err := r.db.First(&product_view, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_view, nil
}


//Function to get all instances of product_view 
func (r *product_viewRepository) GetAll() ([]ProductView, error) {
	var product_view []ProductView
	err := r.db.Find(&product_view).Error
	return product_view, err
}

//Function to update existing instances of product_view 
func (r *product_viewRepository) Update(product_view *ProductView) error {
	return r.db.Save(product_view).Error
}

//Function to delete single instances of product_view 
func (r *product_viewRepository) Delete(id int) error {
	result := r.db.Delete(&ProductView{}, id)
	return result.Error
}