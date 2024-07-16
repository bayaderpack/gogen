package models

import (
	"errors"

	"gorm.io/gorm"
)

// print_description is struct that represent data in Database
type PrintDescription struct {
	gorm.Model
	PrintID  int    `gorm:"column:print_id;default:" json:"print_id"`
	Language string `gorm:"column:language;default:" json:"language"`
	Title    string `gorm:"column:title;default:" json:"title"`
}

// PrintDescription is interface that that model needs to implement
type PrintDescriptionRepository interface {
	Create(print_description *PrintDescription) error
	GetSingle(id int) (*PrintDescription, error)
	GetAll() ([]PrintDescription, error)
	Update(print_description *PrintDescription) error
	Delete(id int) error
}

type print_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewPrintDescriptionRepository(db *gorm.DB) PrintDescriptionRepository {
	return &print_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *print_descriptionRepository) Create(print_description *PrintDescription) error {
	return r.db.Create(print_description).Error
}

// Function to get single instance of print_description
func (r *print_descriptionRepository) GetSingle(id int) (*PrintDescription, error) {
	var print_description PrintDescription
	err := r.db.First(&print_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &print_description, nil
}

// Function to get all instances of print_description
func (r *print_descriptionRepository) GetAll() ([]PrintDescription, error) {
	var print_description []PrintDescription
	err := r.db.Find(&print_description).Error
	return print_description, err
}

// Function to update existing instances of print_description
func (r *print_descriptionRepository) Update(print_description *PrintDescription) error {
	return r.db.Save(print_description).Error
}

// Function to delete single instances of print_description
func (r *print_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&PrintDescription{}, id)
	return result.Error
}
