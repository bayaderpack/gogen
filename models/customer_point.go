package models

import (
	"errors"

	"gorm.io/gorm"
)

// customer_point is struct that represent data in Database
type CustomerPoint struct {
	gorm.Model
	CustomerID int `gorm:"column:customer_id;default:" json:"customer_id"`
	OrderID    int `gorm:"column:order_id;default:" json:"order_id"`
	Points     int `gorm:"column:points;default:" json:"points"`
}

// CustomerPoint is interface that that model needs to implement
type CustomerPointRepository interface {
	Create(customer_point *CustomerPoint) error
	GetSingle(id int) (*CustomerPoint, error)
	GetAll() ([]CustomerPoint, error)
	Update(customer_point *CustomerPoint) error
	Delete(id int) error
}

type customer_pointRepository struct {
	db *gorm.DB
}

// Create new instance
func NewCustomerPointRepository(db *gorm.DB) CustomerPointRepository {
	return &customer_pointRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *customer_pointRepository) Create(customer_point *CustomerPoint) error {
	return r.db.Create(customer_point).Error
}

// Function to get single instance of customer_point
func (r *customer_pointRepository) GetSingle(id int) (*CustomerPoint, error) {
	var customer_point CustomerPoint
	err := r.db.First(&customer_point, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &customer_point, nil
}

// Function to get all instances of customer_point
func (r *customer_pointRepository) GetAll() ([]CustomerPoint, error) {
	var customer_point []CustomerPoint
	err := r.db.Find(&customer_point).Error
	return customer_point, err
}

// Function to update existing instances of customer_point
func (r *customer_pointRepository) Update(customer_point *CustomerPoint) error {
	return r.db.Save(customer_point).Error
}

// Function to delete single instances of customer_point
func (r *customer_pointRepository) Delete(id int) error {
	result := r.db.Delete(&CustomerPoint{}, id)
	return result.Error
}
