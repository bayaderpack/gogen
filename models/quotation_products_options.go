package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// quotation_products_options is struct that represent data in Database
type QuotationProductsOptions struct {
	gorm.Model
	QuotationProductOptionID int `gorm:"column:quotation_product_option_id;default:" json:"quotation_product_option_id"`
QuotationProductID int `gorm:"column:quotation_product_id;default:" json:"quotation_product_id"`
OptionID int `gorm:"column:option_id;default:" json:"option_id"`
OptionValueID int `gorm:"column:option_value_id;default:" json:"option_value_id"`
Type string `gorm:"column:type;default:" json:"type"`
Title string `gorm:"column:title;default:" json:"title"`
Value string `gorm:"column:value;default:" json:"value"`

}

// QuotationProductsOptions is interface that that model needs to implement
type QuotationProductsOptionsRepository interface {
	Create(quotation_products_option *QuotationProductsOptions) error
	GetSingle(id int) (*QuotationProductsOptions, error)
	GetAll() ([]QuotationProductsOptions, error)
	Update(quotation_products_option *QuotationProductsOptions) error
	Delete(id int) error
}

type quotation_products_optionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewQuotationProductsOptionsRepository(db *gorm.DB) QuotationProductsOptionsRepository {
	return &quotation_products_optionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_products_optionRepository) Create(quotation_products_option *QuotationProductsOptions) error {
	return r.db.Create(quotation_products_option).Error
}


//Function to get single instance of quotation_products_option 
func (r *quotation_products_optionRepository) GetSingle(id int) (*QuotationProductsOptions, error) {
	var quotation_products_option QuotationProductsOptions
	err := r.db.First(&quotation_products_option, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_products_option, nil
}


//Function to get all instances of quotation_products_options 
func (r *quotation_products_optionRepository) GetAll() ([]QuotationProductsOptions, error) {
	var quotation_products_options []QuotationProductsOptions
	err := r.db.Find(&quotation_products_options).Error
	return quotation_products_options, err
}

//Function to update existing instances of quotation_products_option 
func (r *quotation_products_optionRepository) Update(quotation_products_option *QuotationProductsOptions) error {
	return r.db.Save(quotation_products_option).Error
}

//Function to delete single instances of quotation_products_option 
func (r *quotation_products_optionRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationProductsOptions{}, id)
	return result.Error
}