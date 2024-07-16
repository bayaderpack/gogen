package models

import (
	"errors"

	"gorm.io/gorm"
)

// supplier_group_description is struct that represent data in Database
type SupplierGroupDescription struct {
	gorm.Model
	GroupID  int    `gorm:"column:group_id;default:" json:"group_id"`
	Language string `gorm:"column:language;default:" json:"language"`
	Title    string `gorm:"column:title;default:" json:"title"`
}

// SupplierGroupDescription is interface that that model needs to implement
type SupplierGroupDescriptionRepository interface {
	Create(supplier_group_description *SupplierGroupDescription) error
	GetSingle(id int) (*SupplierGroupDescription, error)
	GetAll() ([]SupplierGroupDescription, error)
	Update(supplier_group_description *SupplierGroupDescription) error
	Delete(id int) error
}

type supplier_group_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSupplierGroupDescriptionRepository(db *gorm.DB) SupplierGroupDescriptionRepository {
	return &supplier_group_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *supplier_group_descriptionRepository) Create(supplier_group_description *SupplierGroupDescription) error {
	return r.db.Create(supplier_group_description).Error
}

// Function to get single instance of supplier_group_description
func (r *supplier_group_descriptionRepository) GetSingle(id int) (*SupplierGroupDescription, error) {
	var supplier_group_description SupplierGroupDescription
	err := r.db.First(&supplier_group_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &supplier_group_description, nil
}

// Function to get all instances of supplier_group_description
func (r *supplier_group_descriptionRepository) GetAll() ([]SupplierGroupDescription, error) {
	var supplier_group_description []SupplierGroupDescription
	err := r.db.Find(&supplier_group_description).Error
	return supplier_group_description, err
}

// Function to update existing instances of supplier_group_description
func (r *supplier_group_descriptionRepository) Update(supplier_group_description *SupplierGroupDescription) error {
	return r.db.Save(supplier_group_description).Error
}

// Function to delete single instances of supplier_group_description
func (r *supplier_group_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&SupplierGroupDescription{}, id)
	return result.Error
}
