package models

import (
	"errors"

	"gorm.io/gorm"
)

// gallery is struct that represent data in Database
type Gallery struct {
	gorm.Model
	GalleryID int  `gorm:"column:gallery_id;default:" json:"gallery_id"`
	Status    bool `gorm:"column:status;default:1" json:"status"`
}

// Gallery is interface that that model needs to implement
type GalleryRepository interface {
	Create(gallery *Gallery) error
	GetSingle(id int) (*Gallery, error)
	GetAll() ([]Gallery, error)
	Update(gallery *Gallery) error
	Delete(id int) error
}

type galleryRepository struct {
	db *gorm.DB
}

// Create new instance
func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &galleryRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *galleryRepository) Create(gallery *Gallery) error {
	return r.db.Create(gallery).Error
}

// Function to get single instance of gallery
func (r *galleryRepository) GetSingle(id int) (*Gallery, error) {
	var gallery Gallery
	err := r.db.First(&gallery, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &gallery, nil
}

// Function to get all instances of gallery
func (r *galleryRepository) GetAll() ([]Gallery, error) {
	var gallery []Gallery
	err := r.db.Find(&gallery).Error
	return gallery, err
}

// Function to update existing instances of gallery
func (r *galleryRepository) Update(gallery *Gallery) error {
	return r.db.Save(gallery).Error
}

// Function to delete single instances of gallery
func (r *galleryRepository) Delete(id int) error {
	result := r.db.Delete(&Gallery{}, id)
	return result.Error
}
