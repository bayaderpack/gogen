package models

import (
	"time"
	"errors"
	"gorm.io/gorm"
)


// coupon is struct that represent data in Database
type Coupon struct {
	gorm.Model
	CouponID int `gorm:"column:coupon_id;default:" json:"coupon_id"`
Title string `gorm:"column:title;default:" json:"title"`
Code string `gorm:"column:code;default:" json:"code"`
Type TypeValue `gorm:"column:type;default:" json:"type"`
Amount float32 `gorm:"column:amount;default:0.00" json:"amount"`
MinTotal float32 `gorm:"column:min_total;default:0.00" json:"min_total"`
DateStart *time.Time `gorm:"column:date_start;default:current_timestamp()" json:"date_start"`
DateEnd *time.Time `gorm:"column:date_end;default:current_timestamp()" json:"date_end"`
TotalLimit int `gorm:"column:total_limit;default:0" json:"total_limit"`
CustomerLimit int `gorm:"column:customer_limit;default:0" json:"customer_limit"`
Status bool `gorm:"column:status;default:1" json:"status"`

}

// Coupon is interface that that model needs to implement
type CouponRepository interface {
	Create(coupon *Coupon) error
	GetSingle(id int) (*Coupon, error)
	GetAll() ([]Coupon, error)
	Update(coupon *Coupon) error
	Delete(id int) error
}

type couponRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *couponRepository) Create(coupon *Coupon) error {
	return r.db.Create(coupon).Error
}


//Function to get single instance of coupon 
func (r *couponRepository) GetSingle(id int) (*Coupon, error) {
	var coupon Coupon
	err := r.db.First(&coupon, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &coupon, nil
}


//Function to get all instances of coupon 
func (r *couponRepository) GetAll() ([]Coupon, error) {
	var coupon []Coupon
	err := r.db.Find(&coupon).Error
	return coupon, err
}

//Function to update existing instances of coupon 
func (r *couponRepository) Update(coupon *Coupon) error {
	return r.db.Save(coupon).Error
}

//Function to delete single instances of coupon 
func (r *couponRepository) Delete(id int) error {
	result := r.db.Delete(&Coupon{}, id)
	return result.Error
}