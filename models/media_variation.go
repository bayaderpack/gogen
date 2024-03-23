package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// media_variation is struct that represent data in Database
type MediaVariation struct {
	gorm.Model
	MediaID int `gorm:"column:media_id;default:" json:"media_id"`
Url string `gorm:"column:url;default:" json:"url"`
Webp string `gorm:"column:webp;default:" json:"webp"`
Code string `gorm:"column:code;default:" json:"code"`

}

// MediaVariation is interface that that model needs to implement
type MediaVariationRepository interface {
	Create(media_variation *MediaVariation) error
	GetSingle(id int) (*MediaVariation, error)
	GetAll() ([]MediaVariation, error)
	Update(media_variation *MediaVariation) error
	Delete(id int) error
}

type media_variationRepository struct {
	db *gorm.DB
}


// Create new instance
func NewMediaVariationRepository(db *gorm.DB) MediaVariationRepository {
	return &media_variationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *media_variationRepository) Create(media_variation *MediaVariation) error {
	return r.db.Create(media_variation).Error
}


//Function to get single instance of media_variation 
func (r *media_variationRepository) GetSingle(id int) (*MediaVariation, error) {
	var media_variation MediaVariation
	err := r.db.First(&media_variation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &media_variation, nil
}


//Function to get all instances of media_variation 
func (r *media_variationRepository) GetAll() ([]MediaVariation, error) {
	var media_variation []MediaVariation
	err := r.db.Find(&media_variation).Error
	return media_variation, err
}

//Function to update existing instances of media_variation 
func (r *media_variationRepository) Update(media_variation *MediaVariation) error {
	return r.db.Save(media_variation).Error
}

//Function to delete single instances of media_variation 
func (r *media_variationRepository) Delete(id int) error {
	result := r.db.Delete(&MediaVariation{}, id)
	return result.Error
}