package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// product_option is struct that represent data in Database
type ProductOption struct {
	gorm.Model
	ProductOptionID int `gorm:"column:product_option_id;default:" json:"product_option_id"`
ProductID int `gorm:"column:product_id;default:" json:"product_id"`
OptionID int `gorm:"column:option_id;default:" json:"option_id"`
Conditions string `gorm:"column:conditions;default:NULL" json:"conditions"`
Value string `gorm:"column:value;default:NULL" json:"value"`
Required bool `gorm:"column:required;default:1" json:"required"`
IsQuote bool `gorm:"column:is_quote;default:0" json:"is_quote"`

}

// ProductOption is interface that that model needs to implement
type ProductOptionRepository interface {
	Create(product_option *ProductOption) error
	GetSingle(id int) (*ProductOption, error)
	GetAll() ([]ProductOption, error)
	Update(product_option *ProductOption) error
	Delete(id int) error
}

type product_optionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewProductOptionRepository(db *gorm.DB) ProductOptionRepository {
	return &product_optionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_optionRepository) Create(product_option *ProductOption) error {
	return r.db.Create(product_option).Error
}


//Function to get single instance of product_option 
func (r *product_optionRepository) GetSingle(id int) (*ProductOption, error) {
	var product_option ProductOption
	err := r.db.First(&product_option, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_option, nil
}


//Function to get all instances of product_option 
func (r *product_optionRepository) GetAll() ([]ProductOption, error) {
	var product_option []ProductOption
	err := r.db.Find(&product_option).Error
	return product_option, err
}

//Function to update existing instances of product_option 
func (r *product_optionRepository) Update(product_option *ProductOption) error {
	return r.db.Save(product_option).Error
}

//Function to delete single instances of product_option 
func (r *product_optionRepository) Delete(id int) error {
	result := r.db.Delete(&ProductOption{}, id)
	return result.Error
}