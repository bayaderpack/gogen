package models

import (
	"errors"

	"gorm.io/gorm"
)

// taxonomy_description is struct that represent data in Database
type TaxonomyDescription struct {
	gorm.Model
	TaxonomyID      int    `gorm:"column:taxonomy_id;default:" json:"taxonomy_id"`
	Language        string `gorm:"column:language;default:" json:"language"`
	Title           string `gorm:"column:title;default:NULL" json:"title"`
	Description     string `gorm:"column:description;default:NULL" json:"description"`
	MetaTitle       string `gorm:"column:meta_title;default:NULL" json:"meta_title"`
	MetaDescription string `gorm:"column:meta_description;default:NULL" json:"meta_description"`
	MetaKeywords    string `gorm:"column:meta_keywords;default:NULL" json:"meta_keywords"`
}

// TaxonomyDescription is interface that that model needs to implement
type TaxonomyDescriptionRepository interface {
	Create(taxonomy_description *TaxonomyDescription) error
	GetSingle(id int) (*TaxonomyDescription, error)
	GetAll() ([]TaxonomyDescription, error)
	Update(taxonomy_description *TaxonomyDescription) error
	Delete(id int) error
}

type taxonomy_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewTaxonomyDescriptionRepository(db *gorm.DB) TaxonomyDescriptionRepository {
	return &taxonomy_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *taxonomy_descriptionRepository) Create(taxonomy_description *TaxonomyDescription) error {
	return r.db.Create(taxonomy_description).Error
}

// Function to get single instance of taxonomy_description
func (r *taxonomy_descriptionRepository) GetSingle(id int) (*TaxonomyDescription, error) {
	var taxonomy_description TaxonomyDescription
	err := r.db.First(&taxonomy_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &taxonomy_description, nil
}

// Function to get all instances of taxonomy_description
func (r *taxonomy_descriptionRepository) GetAll() ([]TaxonomyDescription, error) {
	var taxonomy_description []TaxonomyDescription
	err := r.db.Find(&taxonomy_description).Error
	return taxonomy_description, err
}

// Function to update existing instances of taxonomy_description
func (r *taxonomy_descriptionRepository) Update(taxonomy_description *TaxonomyDescription) error {
	return r.db.Save(taxonomy_description).Error
}

// Function to delete single instances of taxonomy_description
func (r *taxonomy_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&TaxonomyDescription{}, id)
	return result.Error
}
