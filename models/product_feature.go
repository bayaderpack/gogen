package models

import (
	"errors"

	"gorm.io/gorm"
)

// product_feature is struct that represent data in Database
type ProductFeature struct {
	gorm.Model
	FeatureID int `gorm:"column:feature_id;default:" json:"feature_id"`
	ProductID int `gorm:"column:product_id;default:" json:"product_id"`
}

// ProductFeature is interface that that model needs to implement
type ProductFeatureRepository interface {
	Create(product_feature *ProductFeature) error
	GetSingle(id int) (*ProductFeature, error)
	GetAll() ([]ProductFeature, error)
	Update(product_feature *ProductFeature) error
	Delete(id int) error
}

type product_featureRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductFeatureRepository(db *gorm.DB) ProductFeatureRepository {
	return &product_featureRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_featureRepository) Create(product_feature *ProductFeature) error {
	return r.db.Create(product_feature).Error
}

// Function to get single instance of product_feature
func (r *product_featureRepository) GetSingle(id int) (*ProductFeature, error) {
	var product_feature ProductFeature
	err := r.db.First(&product_feature, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_feature, nil
}

// Function to get all instances of product_feature
func (r *product_featureRepository) GetAll() ([]ProductFeature, error) {
	var product_feature []ProductFeature
	err := r.db.Find(&product_feature).Error
	return product_feature, err
}

// Function to update existing instances of product_feature
func (r *product_featureRepository) Update(product_feature *ProductFeature) error {
	return r.db.Save(product_feature).Error
}

// Function to delete single instances of product_feature
func (r *product_featureRepository) Delete(id int) error {
	result := r.db.Delete(&ProductFeature{}, id)
	return result.Error
}
