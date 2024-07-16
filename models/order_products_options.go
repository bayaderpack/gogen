package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// order_products_options is struct that represent data in Database
type OrderProductsOptions struct {
	gorm.Model
	OrderProductOptionID int `gorm:"column:order_product_option_id;default:" json:"order_product_option_id"`
OrderProductID int `gorm:"column:order_product_id;default:" json:"order_product_id"`
OptionID int `gorm:"column:option_id;default:" json:"option_id"`
OptionValueID int `gorm:"column:option_value_id;default:" json:"option_value_id"`
Type string `gorm:"column:type;default:" json:"type"`
OpPrefix string `gorm:"column:op_prefix;default:" json:"op_prefix"`
OpAmount float32 `gorm:"column:op_amount;default:" json:"op_amount"`
Title string `gorm:"column:title;default:" json:"title"`
Value string `gorm:"column:value;default:" json:"value"`

}

// OrderProductsOptions is interface that that model needs to implement
type OrderProductsOptionsRepository interface {
	Create(order_products_option *OrderProductsOptions) error
	GetSingle(id int) (*OrderProductsOptions, error)
	GetAll() ([]OrderProductsOptions, error)
	Update(order_products_option *OrderProductsOptions) error
	Delete(id int) error
}

type order_products_optionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewOrderProductsOptionsRepository(db *gorm.DB) OrderProductsOptionsRepository {
	return &order_products_optionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *order_products_optionRepository) Create(order_products_option *OrderProductsOptions) error {
	return r.db.Create(order_products_option).Error
}


//Function to get single instance of order_products_option 
func (r *order_products_optionRepository) GetSingle(id int) (*OrderProductsOptions, error) {
	var order_products_option OrderProductsOptions
	err := r.db.First(&order_products_option, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &order_products_option, nil
}


//Function to get all instances of order_products_options 
func (r *order_products_optionRepository) GetAll() ([]OrderProductsOptions, error) {
	var order_products_options []OrderProductsOptions
	err := r.db.Find(&order_products_options).Error
	return order_products_options, err
}

//Function to update existing instances of order_products_option 
func (r *order_products_optionRepository) Update(order_products_option *OrderProductsOptions) error {
	return r.db.Save(order_products_option).Error
}

//Function to delete single instances of order_products_option 
func (r *order_products_optionRepository) Delete(id int) error {
	result := r.db.Delete(&OrderProductsOptions{}, id)
	return result.Error
}