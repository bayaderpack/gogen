package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_description is struct that represent data in Database
type ProductDescription struct {
	gorm.Model
	ProductID       int    `gorm:"column:product_id;default:" json:"product_id"`
	Language        string `gorm:"column:language;default:" json:"language"`
	Title           string `gorm:"column:title;default:" json:"title"`
	Description     string `gorm:"column:description;default:" json:"description"`
	MetaTitle       string `gorm:"column:meta_title;default:" json:"meta_title"`
	MetaDescription string `gorm:"column:meta_description;default:" json:"meta_description"`
	MetaKeywords    string `gorm:"column:meta_keywords;default:" json:"meta_keywords"`
	MediaID         int    `gorm:"column:media_id;default:" json:"media_id"`
}

// ProductDescription is interface that that model needs to implement
type ProductDescriptionRepository interface {
	Create(product_description *ProductDescription) error
	GetSingle(id int) (*ProductDescription, error)
	GetAll() ([]ProductDescription, error)
	Update(product_description *ProductDescription) error
	Delete(id int) error
}

type product_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductDescriptionRepository(db *gorm.DB) ProductDescriptionRepository {
	return &product_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_descriptionRepository) Create(product_description *ProductDescription) error {
	return r.db.Create(product_description).Error
}

// Function to get single instance of product_description
func (r *product_descriptionRepository) GetSingle(id int) (*ProductDescription, error) {
	var product_description ProductDescription
	err := r.db.First(&product_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_description, nil
}

// Function to get all instances of product_description
func (r *product_descriptionRepository) GetAll() ([]ProductDescription, error) {
	var product_description []ProductDescription
	err := r.db.Find(&product_description).Error
	return product_description, err
}

// Function to update existing instances of product_description
func (r *product_descriptionRepository) Update(product_description *ProductDescription) error {
	return r.db.Save(product_description).Error
}

// Function to delete single instances of product_description
func (r *product_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&ProductDescription{}, id)
	return result.Error
}
