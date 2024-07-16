package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// gallery_description is struct that represent data in Database
type GalleryDescription struct {
	gorm.Model
	GalleryID int `gorm:"column:gallery_id;default:" json:"gallery_id"`
Language string `gorm:"column:language;default:" json:"language"`
Title string `gorm:"column:title;default:" json:"title"`
Description string `gorm:"column:description;default:" json:"description"`

}

// GalleryDescription is interface that that model needs to implement
type GalleryDescriptionRepository interface {
	Create(gallery_description *GalleryDescription) error
	GetSingle(id int) (*GalleryDescription, error)
	GetAll() ([]GalleryDescription, error)
	Update(gallery_description *GalleryDescription) error
	Delete(id int) error
}

type gallery_descriptionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewGalleryDescriptionRepository(db *gorm.DB) GalleryDescriptionRepository {
	return &gallery_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *gallery_descriptionRepository) Create(gallery_description *GalleryDescription) error {
	return r.db.Create(gallery_description).Error
}


//Function to get single instance of gallery_description 
func (r *gallery_descriptionRepository) GetSingle(id int) (*GalleryDescription, error) {
	var gallery_description GalleryDescription
	err := r.db.First(&gallery_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &gallery_description, nil
}


//Function to get all instances of gallery_description 
func (r *gallery_descriptionRepository) GetAll() ([]GalleryDescription, error) {
	var gallery_description []GalleryDescription
	err := r.db.Find(&gallery_description).Error
	return gallery_description, err
}

//Function to update existing instances of gallery_description 
func (r *gallery_descriptionRepository) Update(gallery_description *GalleryDescription) error {
	return r.db.Save(gallery_description).Error
}

//Function to delete single instances of gallery_description 
func (r *gallery_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&GalleryDescription{}, id)
	return result.Error
}