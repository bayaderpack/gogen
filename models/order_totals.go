package models

import (
	"errors"

	"gorm.io/gorm"
)

// order_totals is struct that represent data in Database
type OrderTotals struct {
	gorm.Model
	OrderTotalID int     `gorm:"column:order_total_id;default:" json:"order_total_id"`
	OrderID      int     `gorm:"column:order_id;default:" json:"order_id"`
	Text         string  `gorm:"column:text;default:" json:"text"`
	Code         string  `gorm:"column:code;default:" json:"code"`
	Value        float32 `gorm:"column:value;default:" json:"value"`
	SortOrder    bool    `gorm:"column:sort_order;default:0" json:"sort_order"`
}

// OrderTotals is interface that that model needs to implement
type OrderTotalsRepository interface {
	Create(order_total *OrderTotals) error
	GetSingle(id int) (*OrderTotals, error)
	GetAll() ([]OrderTotals, error)
	Update(order_total *OrderTotals) error
	Delete(id int) error
}

type order_totalRepository struct {
	db *gorm.DB
}

// Create new instance
func NewOrderTotalsRepository(db *gorm.DB) OrderTotalsRepository {
	return &order_totalRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *order_totalRepository) Create(order_total *OrderTotals) error {
	return r.db.Create(order_total).Error
}

// Function to get single instance of order_total
func (r *order_totalRepository) GetSingle(id int) (*OrderTotals, error) {
	var order_total OrderTotals
	err := r.db.First(&order_total, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &order_total, nil
}

// Function to get all instances of order_totals
func (r *order_totalRepository) GetAll() ([]OrderTotals, error) {
	var order_totals []OrderTotals
	err := r.db.Find(&order_totals).Error
	return order_totals, err
}

// Function to update existing instances of order_total
func (r *order_totalRepository) Update(order_total *OrderTotals) error {
	return r.db.Save(order_total).Error
}

// Function to delete single instances of order_total
func (r *order_totalRepository) Delete(id int) error {
	result := r.db.Delete(&OrderTotals{}, id)
	return result.Error
}
