package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_option_value is struct that represent data in Database
type ProductOptionValue struct {
	gorm.Model
	ProductOptionID int     `gorm:"column:product_option_id;default:" json:"product_option_id"`
	OptionValueID   int     `gorm:"column:option_value_id;default:" json:"option_value_id"`
	MediaID         int     `gorm:"column:media_id;default:NULL" json:"media_id"`
	CarouselDisplay bool    `gorm:"column:carousel_display;default:0" json:"carousel_display"`
	Quantity        int     `gorm:"column:quantity;default:" json:"quantity"`
	Minimum         int     `gorm:"column:minimum;default:0" json:"minimum"`
	PricePrefix     string  `gorm:"column:price_prefix;default:" json:"price_prefix"`
	Price           float32 `gorm:"column:price;default:" json:"price"`
	Ocp             string  `gorm:"column:ocp;default:'{}'" json:"ocp"`
	Docp            string  `gorm:"column:docp;default:'[]'" json:"docp"`
}

// ProductOptionValue is interface that that model needs to implement
type ProductOptionValueRepository interface {
	Create(product_option_value *ProductOptionValue) error
	GetSingle(id int) (*ProductOptionValue, error)
	GetAll() ([]ProductOptionValue, error)
	Update(product_option_value *ProductOptionValue) error
	Delete(id int) error
}

type product_option_valueRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductOptionValueRepository(db *gorm.DB) ProductOptionValueRepository {
	return &product_option_valueRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_option_valueRepository) Create(product_option_value *ProductOptionValue) error {
	return r.db.Create(product_option_value).Error
}

// Function to get single instance of product_option_value
func (r *product_option_valueRepository) GetSingle(id int) (*ProductOptionValue, error) {
	var product_option_value ProductOptionValue
	err := r.db.First(&product_option_value, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_option_value, nil
}

// Function to get all instances of product_option_value
func (r *product_option_valueRepository) GetAll() ([]ProductOptionValue, error) {
	var product_option_value []ProductOptionValue
	err := r.db.Find(&product_option_value).Error
	return product_option_value, err
}

// Function to update existing instances of product_option_value
func (r *product_option_valueRepository) Update(product_option_value *ProductOptionValue) error {
	return r.db.Save(product_option_value).Error
}

// Function to delete single instances of product_option_value
func (r *product_option_valueRepository) Delete(id int) error {
	result := r.db.Delete(&ProductOptionValue{}, id)
	return result.Error
}
