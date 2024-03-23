package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// material_product_description is struct that represent data in Database
type MaterialProductDescription struct {
	gorm.Model
	MaterialProductID int `gorm:"column:material_product_id;default:" json:"material_product_id"`
Language string `gorm:"column:language;default:" json:"language"`
Title string `gorm:"column:title;default:" json:"title"`

}

// MaterialProductDescription is interface that that model needs to implement
type MaterialProductDescriptionRepository interface {
	Create(material_product_description *MaterialProductDescription) error
	GetSingle(id int) (*MaterialProductDescription, error)
	GetAll() ([]MaterialProductDescription, error)
	Update(material_product_description *MaterialProductDescription) error
	Delete(id int) error
}

type material_product_descriptionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewMaterialProductDescriptionRepository(db *gorm.DB) MaterialProductDescriptionRepository {
	return &material_product_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *material_product_descriptionRepository) Create(material_product_description *MaterialProductDescription) error {
	return r.db.Create(material_product_description).Error
}


//Function to get single instance of material_product_description 
func (r *material_product_descriptionRepository) GetSingle(id int) (*MaterialProductDescription, error) {
	var material_product_description MaterialProductDescription
	err := r.db.First(&material_product_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material_product_description, nil
}


//Function to get all instances of material_product_description 
func (r *material_product_descriptionRepository) GetAll() ([]MaterialProductDescription, error) {
	var material_product_description []MaterialProductDescription
	err := r.db.Find(&material_product_description).Error
	return material_product_description, err
}

//Function to update existing instances of material_product_description 
func (r *material_product_descriptionRepository) Update(material_product_description *MaterialProductDescription) error {
	return r.db.Save(material_product_description).Error
}

//Function to delete single instances of material_product_description 
func (r *material_product_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&MaterialProductDescription{}, id)
	return result.Error
}