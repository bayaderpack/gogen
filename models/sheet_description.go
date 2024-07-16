package models

import (
	"errors"

	"gorm.io/gorm"
)

// sheet_description is struct that represent data in Database
type SheetDescription struct {
	gorm.Model
	SheetID  int    `gorm:"column:sheet_id;default:" json:"sheet_id"`
	Language string `gorm:"column:language;default:" json:"language"`
	Title    string `gorm:"column:title;default:" json:"title"`
	Subtitle string `gorm:"column:subtitle;default:NULL" json:"subtitle"`
}

// SheetDescription is interface that that model needs to implement
type SheetDescriptionRepository interface {
	Create(sheet_description *SheetDescription) error
	GetSingle(id int) (*SheetDescription, error)
	GetAll() ([]SheetDescription, error)
	Update(sheet_description *SheetDescription) error
	Delete(id int) error
}

type sheet_descriptionRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSheetDescriptionRepository(db *gorm.DB) SheetDescriptionRepository {
	return &sheet_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sheet_descriptionRepository) Create(sheet_description *SheetDescription) error {
	return r.db.Create(sheet_description).Error
}

// Function to get single instance of sheet_description
func (r *sheet_descriptionRepository) GetSingle(id int) (*SheetDescription, error) {
	var sheet_description SheetDescription
	err := r.db.First(&sheet_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sheet_description, nil
}

// Function to get all instances of sheet_description
func (r *sheet_descriptionRepository) GetAll() ([]SheetDescription, error) {
	var sheet_description []SheetDescription
	err := r.db.Find(&sheet_description).Error
	return sheet_description, err
}

// Function to update existing instances of sheet_description
func (r *sheet_descriptionRepository) Update(sheet_description *SheetDescription) error {
	return r.db.Save(sheet_description).Error
}

// Function to delete single instances of sheet_description
func (r *sheet_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&SheetDescription{}, id)
	return result.Error
}
