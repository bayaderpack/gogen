package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// material is struct that represent data in Database
type Material struct {
	gorm.Model
	MaterialID int `gorm:"column:material_id;default:" json:"material_id"`
Quotable bool `gorm:"column:quotable;default:" json:"quotable"`

}

// Material is interface that that model needs to implement
type MaterialRepository interface {
	Create(material *Material) error
	GetSingle(id int) (*Material, error)
	GetAll() ([]Material, error)
	Update(material *Material) error
	Delete(id int) error
}

type materialRepository struct {
	db *gorm.DB
}


// Create new instance
func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *materialRepository) Create(material *Material) error {
	return r.db.Create(material).Error
}


//Function to get single instance of material 
func (r *materialRepository) GetSingle(id int) (*Material, error) {
	var material Material
	err := r.db.First(&material, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &material, nil
}


//Function to get all instances of material 
func (r *materialRepository) GetAll() ([]Material, error) {
	var material []Material
	err := r.db.Find(&material).Error
	return material, err
}

//Function to update existing instances of material 
func (r *materialRepository) Update(material *Material) error {
	return r.db.Save(material).Error
}

//Function to delete single instances of material 
func (r *materialRepository) Delete(id int) error {
	result := r.db.Delete(&Material{}, id)
	return result.Error
}