package models

import (
	"errors"

	"gorm.io/gorm"
)

// media is struct that represent data in Database
type Media struct {
	gorm.Model
	MediaID     int    `gorm:"column:media_id;default:" json:"media_id"`
	Title       string `gorm:"column:title;default:" json:"title"`
	Url         string `gorm:"column:url;default:" json:"url"`
	Destination string `gorm:"column:destination;default:" json:"destination"`
	Mimetype    string `gorm:"column:mimetype;default:" json:"mimetype"`
	IsImage     bool   `gorm:"column:is_image;default:1" json:"is_image"`
	UploadedBy  int    `gorm:"column:uploaded_by;default:" json:"uploaded_by"`
}

// Media is interface that that model needs to implement
type MediaRepository interface {
	Create(media *Media) error
	GetSingle(id int) (*Media, error)
	GetAll() ([]Media, error)
	Update(media *Media) error
	Delete(id int) error
}

type mediaRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *mediaRepository) Create(media *Media) error {
	return r.db.Create(media).Error
}

// Function to get single instance of media
func (r *mediaRepository) GetSingle(id int) (*Media, error) {
	var media Media
	err := r.db.First(&media, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &media, nil
}

// Function to get all instances of media
func (r *mediaRepository) GetAll() ([]Media, error) {
	var media []Media
	err := r.db.Find(&media).Error
	return media, err
}

// Function to update existing instances of media
func (r *mediaRepository) Update(media *Media) error {
	return r.db.Save(media).Error
}

// Function to delete single instances of media
func (r *mediaRepository) Delete(id int) error {
	result := r.db.Delete(&Media{}, id)
	return result.Error
}
