package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_favorite is struct that represent data in Database
type ProductFavorite struct {
	gorm.Model
	ProductID  int `gorm:"column:product_id;default:" json:"product_id"`
	CustomerID int `gorm:"column:customer_id;default:" json:"customer_id"`
}

// ProductFavorite is interface that that model needs to implement
type ProductFavoriteRepository interface {
	Create(product_favorite *ProductFavorite) error
	GetSingle(id int) (*ProductFavorite, error)
	GetAll() ([]ProductFavorite, error)
	Update(product_favorite *ProductFavorite) error
	Delete(id int) error
}

type product_favoriteRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductFavoriteRepository(db *gorm.DB) ProductFavoriteRepository {
	return &product_favoriteRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_favoriteRepository) Create(product_favorite *ProductFavorite) error {
	return r.db.Create(product_favorite).Error
}

// Function to get single instance of product_favorite
func (r *product_favoriteRepository) GetSingle(id int) (*ProductFavorite, error) {
	var product_favorite ProductFavorite
	err := r.db.First(&product_favorite, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_favorite, nil
}

// Function to get all instances of product_favorite
func (r *product_favoriteRepository) GetAll() ([]ProductFavorite, error) {
	var product_favorite []ProductFavorite
	err := r.db.Find(&product_favorite).Error
	return product_favorite, err
}

// Function to update existing instances of product_favorite
func (r *product_favoriteRepository) Update(product_favorite *ProductFavorite) error {
	return r.db.Save(product_favorite).Error
}

// Function to delete single instances of product_favorite
func (r *product_favoriteRepository) Delete(id int) error {
	result := r.db.Delete(&ProductFavorite{}, id)
	return result.Error
}
