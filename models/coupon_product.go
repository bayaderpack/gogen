package models

import (
	"errors"

	"gorm.io/gorm"
)

// coupon_product is struct that represent data in Database
type CouponProduct struct {
	gorm.Model
	CouponID  int `gorm:"column:coupon_id;default:" json:"coupon_id"`
	ProductID int `gorm:"column:product_id;default:" json:"product_id"`
}

// CouponProduct is interface that that model needs to implement
type CouponProductRepository interface {
	Create(coupon_product *CouponProduct) error
	GetSingle(id int) (*CouponProduct, error)
	GetAll() ([]CouponProduct, error)
	Update(coupon_product *CouponProduct) error
	Delete(id int) error
}

type coupon_productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewCouponProductRepository(db *gorm.DB) CouponProductRepository {
	return &coupon_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *coupon_productRepository) Create(coupon_product *CouponProduct) error {
	return r.db.Create(coupon_product).Error
}

// Function to get single instance of coupon_product
func (r *coupon_productRepository) GetSingle(id int) (*CouponProduct, error) {
	var coupon_product CouponProduct
	err := r.db.First(&coupon_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &coupon_product, nil
}

// Function to get all instances of coupon_product
func (r *coupon_productRepository) GetAll() ([]CouponProduct, error) {
	var coupon_product []CouponProduct
	err := r.db.Find(&coupon_product).Error
	return coupon_product, err
}

// Function to update existing instances of coupon_product
func (r *coupon_productRepository) Update(coupon_product *CouponProduct) error {
	return r.db.Save(coupon_product).Error
}

// Function to delete single instances of coupon_product
func (r *coupon_productRepository) Delete(id int) error {
	result := r.db.Delete(&CouponProduct{}, id)
	return result.Error
}
