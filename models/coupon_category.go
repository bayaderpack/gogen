package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// coupon_category is struct that represent data in Database
type CouponCategory struct {
	gorm.Model
	CouponID int `gorm:"column:coupon_id;default:" json:"coupon_id"`
TaxonomyID int `gorm:"column:taxonomy_id;default:" json:"taxonomy_id"`

}

// CouponCategory is interface that that model needs to implement
type CouponCategoryRepository interface {
	Create(coupon_category *CouponCategory) error
	GetSingle(id int) (*CouponCategory, error)
	GetAll() ([]CouponCategory, error)
	Update(coupon_category *CouponCategory) error
	Delete(id int) error
}

type coupon_categoryRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCouponCategoryRepository(db *gorm.DB) CouponCategoryRepository {
	return &coupon_categoryRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *coupon_categoryRepository) Create(coupon_category *CouponCategory) error {
	return r.db.Create(coupon_category).Error
}


//Function to get single instance of coupon_category 
func (r *coupon_categoryRepository) GetSingle(id int) (*CouponCategory, error) {
	var coupon_category CouponCategory
	err := r.db.First(&coupon_category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &coupon_category, nil
}


//Function to get all instances of coupon_category 
func (r *coupon_categoryRepository) GetAll() ([]CouponCategory, error) {
	var coupon_category []CouponCategory
	err := r.db.Find(&coupon_category).Error
	return coupon_category, err
}

//Function to update existing instances of coupon_category 
func (r *coupon_categoryRepository) Update(coupon_category *CouponCategory) error {
	return r.db.Save(coupon_category).Error
}

//Function to delete single instances of coupon_category 
func (r *coupon_categoryRepository) Delete(id int) error {
	result := r.db.Delete(&CouponCategory{}, id)
	return result.Error
}