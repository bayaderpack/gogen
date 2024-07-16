package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// option_description is struct that represent data in Database
type OptionDescription struct {
	gorm.Model
	OptionID int `gorm:"column:option_id;default:" json:"option_id"`
Language string `gorm:"column:language;default:" json:"language"`
Title string `gorm:"column:title;default:" json:"title"`

}

// OptionDescription is interface that that model needs to implement
type OptionDescriptionRepository interface {
	Create(option_description *OptionDescription) error
	GetSingle(id int) (*OptionDescription, error)
	GetAll() ([]OptionDescription, error)
	Update(option_description *OptionDescription) error
	Delete(id int) error
}

type option_descriptionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewOptionDescriptionRepository(db *gorm.DB) OptionDescriptionRepository {
	return &option_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *option_descriptionRepository) Create(option_description *OptionDescription) error {
	return r.db.Create(option_description).Error
}


//Function to get single instance of option_description 
func (r *option_descriptionRepository) GetSingle(id int) (*OptionDescription, error) {
	var option_description OptionDescription
	err := r.db.First(&option_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &option_description, nil
}


//Function to get all instances of option_description 
func (r *option_descriptionRepository) GetAll() ([]OptionDescription, error) {
	var option_description []OptionDescription
	err := r.db.Find(&option_description).Error
	return option_description, err
}

//Function to update existing instances of option_description 
func (r *option_descriptionRepository) Update(option_description *OptionDescription) error {
	return r.db.Save(option_description).Error
}

//Function to delete single instances of option_description 
func (r *option_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&OptionDescription{}, id)
	return result.Error
}