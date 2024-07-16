package models

import (
	"errors"

	"gorm.io/gorm"
)

// media_favorite is struct that represent data in Database
type MediaFavorite struct {
	gorm.Model
	MediaID    int    `gorm:"column:media_id;default:" json:"media_id"`
	CustomerID string `gorm:"column:customer_id;default:" json:"customer_id"`
}

// MediaFavorite is interface that that model needs to implement
type MediaFavoriteRepository interface {
	Create(media_favorite *MediaFavorite) error
	GetSingle(id int) (*MediaFavorite, error)
	GetAll() ([]MediaFavorite, error)
	Update(media_favorite *MediaFavorite) error
	Delete(id int) error
}

type media_favoriteRepository struct {
	db *gorm.DB
}

// Create new instance
func NewMediaFavoriteRepository(db *gorm.DB) MediaFavoriteRepository {
	return &media_favoriteRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *media_favoriteRepository) Create(media_favorite *MediaFavorite) error {
	return r.db.Create(media_favorite).Error
}

// Function to get single instance of media_favorite
func (r *media_favoriteRepository) GetSingle(id int) (*MediaFavorite, error) {
	var media_favorite MediaFavorite
	err := r.db.First(&media_favorite, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &media_favorite, nil
}

// Function to get all instances of media_favorite
func (r *media_favoriteRepository) GetAll() ([]MediaFavorite, error) {
	var media_favorite []MediaFavorite
	err := r.db.Find(&media_favorite).Error
	return media_favorite, err
}

// Function to update existing instances of media_favorite
func (r *media_favoriteRepository) Update(media_favorite *MediaFavorite) error {
	return r.db.Save(media_favorite).Error
}

// Function to delete single instances of media_favorite
func (r *media_favoriteRepository) Delete(id int) error {
	result := r.db.Delete(&MediaFavorite{}, id)
	return result.Error
}
