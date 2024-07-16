package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_media is struct that represent data in Database
type ProductMedia struct {
	gorm.Model
	ProductMediaID int `gorm:"column:product_media_id;default:" json:"product_media_id"`
	ProductID      int `gorm:"column:product_id;default:" json:"product_id"`
	MediaID        int `gorm:"column:media_id;default:" json:"media_id"`
}

// ProductMedia is interface that that model needs to implement
type ProductMediaRepository interface {
	Create(product_media *ProductMedia) error
	GetSingle(id int) (*ProductMedia, error)
	GetAll() ([]ProductMedia, error)
	Update(product_media *ProductMedia) error
	Delete(id int) error
}

type product_mediaRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductMediaRepository(db *gorm.DB) ProductMediaRepository {
	return &product_mediaRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_mediaRepository) Create(product_media *ProductMedia) error {
	return r.db.Create(product_media).Error
}

// Function to get single instance of product_media
func (r *product_mediaRepository) GetSingle(id int) (*ProductMedia, error) {
	var product_media ProductMedia
	err := r.db.First(&product_media, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_media, nil
}

// Function to get all instances of product_media
func (r *product_mediaRepository) GetAll() ([]ProductMedia, error) {
	var product_media []ProductMedia
	err := r.db.Find(&product_media).Error
	return product_media, err
}

// Function to update existing instances of product_media
func (r *product_mediaRepository) Update(product_media *ProductMedia) error {
	return r.db.Save(product_media).Error
}

// Function to delete single instances of product_media
func (r *product_mediaRepository) Delete(id int) error {
	result := r.db.Delete(&ProductMedia{}, id)
	return result.Error
}
