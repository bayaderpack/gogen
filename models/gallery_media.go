package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// gallery_media is struct that represent data in Database
type GalleryMedia struct {
	gorm.Model
	GalleryID int `gorm:"column:gallery_id;default:" json:"gallery_id"`
MediaID int `gorm:"column:media_id;default:" json:"media_id"`

}

// GalleryMedia is interface that that model needs to implement
type GalleryMediaRepository interface {
	Create(gallery_media *GalleryMedia) error
	GetSingle(id int) (*GalleryMedia, error)
	GetAll() ([]GalleryMedia, error)
	Update(gallery_media *GalleryMedia) error
	Delete(id int) error
}

type gallery_mediaRepository struct {
	db *gorm.DB
}


// Create new instance
func NewGalleryMediaRepository(db *gorm.DB) GalleryMediaRepository {
	return &gallery_mediaRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *gallery_mediaRepository) Create(gallery_media *GalleryMedia) error {
	return r.db.Create(gallery_media).Error
}


//Function to get single instance of gallery_media 
func (r *gallery_mediaRepository) GetSingle(id int) (*GalleryMedia, error) {
	var gallery_media GalleryMedia
	err := r.db.First(&gallery_media, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &gallery_media, nil
}


//Function to get all instances of gallery_media 
func (r *gallery_mediaRepository) GetAll() ([]GalleryMedia, error) {
	var gallery_media []GalleryMedia
	err := r.db.Find(&gallery_media).Error
	return gallery_media, err
}

//Function to update existing instances of gallery_media 
func (r *gallery_mediaRepository) Update(gallery_media *GalleryMedia) error {
	return r.db.Save(gallery_media).Error
}

//Function to delete single instances of gallery_media 
func (r *gallery_mediaRepository) Delete(id int) error {
	result := r.db.Delete(&GalleryMedia{}, id)
	return result.Error
}