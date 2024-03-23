package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// coupon_history is struct that represent data in Database
type CouponHistory struct {
	gorm.Model
	CouponHistoryID int `gorm:"column:coupon_history_id;default:" json:"coupon_history_id"`
CouponID int `gorm:"column:coupon_id;default:" json:"coupon_id"`
OrderID int `gorm:"column:order_id;default:" json:"order_id"`
CustomerID int `gorm:"column:customer_id;default:" json:"customer_id"`
Amount float32 `gorm:"column:amount;default:" json:"amount"`

}

// CouponHistory is interface that that model needs to implement
type CouponHistoryRepository interface {
	Create(coupon_history *CouponHistory) error
	GetSingle(id int) (*CouponHistory, error)
	GetAll() ([]CouponHistory, error)
	Update(coupon_history *CouponHistory) error
	Delete(id int) error
}

type coupon_historyRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCouponHistoryRepository(db *gorm.DB) CouponHistoryRepository {
	return &coupon_historyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *coupon_historyRepository) Create(coupon_history *CouponHistory) error {
	return r.db.Create(coupon_history).Error
}


//Function to get single instance of coupon_history 
func (r *coupon_historyRepository) GetSingle(id int) (*CouponHistory, error) {
	var coupon_history CouponHistory
	err := r.db.First(&coupon_history, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &coupon_history, nil
}


//Function to get all instances of coupon_history 
func (r *coupon_historyRepository) GetAll() ([]CouponHistory, error) {
	var coupon_history []CouponHistory
	err := r.db.Find(&coupon_history).Error
	return coupon_history, err
}

//Function to update existing instances of coupon_history 
func (r *coupon_historyRepository) Update(coupon_history *CouponHistory) error {
	return r.db.Save(coupon_history).Error
}

//Function to delete single instances of coupon_history 
func (r *coupon_historyRepository) Delete(id int) error {
	result := r.db.Delete(&CouponHistory{}, id)
	return result.Error
}