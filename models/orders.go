package models

import (
	"errors"

	"gorm.io/gorm"
)

// orders is struct that represent data in Database
type Orders struct {
	gorm.Model
	OrderID         int     `gorm:"column:order_id;default:" json:"order_id"`
	InvoiceNo       string  `gorm:"column:invoice_no;default:" json:"invoice_no"`
	CustomerID      int     `gorm:"column:customer_id;default:" json:"customer_id"`
	Name            string  `gorm:"column:name;default:" json:"name"`
	CountryCode     int     `gorm:"column:country_code;default:" json:"country_code"`
	Mobile          int     `gorm:"column:mobile;default:" json:"mobile"`
	Country         string  `gorm:"column:country;default:" json:"country"`
	City            string  `gorm:"column:city;default:" json:"city"`
	Line1           string  `gorm:"column:line1;default:" json:"line1"`
	Location        int     `gorm:"column:location;default:" json:"location"`
	PaymentMethod   string  `gorm:"column:payment_method;default:" json:"payment_method"`
	ChargeID        string  `gorm:"column:charge_id;default:" json:"charge_id"`
	Total           float32 `gorm:"column:total;default:" json:"total"`
	StatusID        bool    `gorm:"column:status_id;default:0" json:"status_id"`
	Tracking        string  `gorm:"column:tracking;default:" json:"tracking"`
	TrackingCompany string  `gorm:"column:tracking_company;default:" json:"tracking_company"`
}

// Orders is interface that that model needs to implement
type OrdersRepository interface {
	Create(order *Orders) error
	GetSingle(id int) (*Orders, error)
	GetAll() ([]Orders, error)
	Update(order *Orders) error
	Delete(id int) error
}

type orderRepository struct {
	db *gorm.DB
}

// Create new instance
func NewOrdersRepository(db *gorm.DB) OrdersRepository {
	return &orderRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *orderRepository) Create(order *Orders) error {
	return r.db.Create(order).Error
}

// Function to get single instance of order
func (r *orderRepository) GetSingle(id int) (*Orders, error) {
	var order Orders
	err := r.db.First(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &order, nil
}

// Function to get all instances of orders
func (r *orderRepository) GetAll() ([]Orders, error) {
	var orders []Orders
	err := r.db.Find(&orders).Error
	return orders, err
}

// Function to update existing instances of order
func (r *orderRepository) Update(order *Orders) error {
	return r.db.Save(order).Error
}

// Function to delete single instances of order
func (r *orderRepository) Delete(id int) error {
	result := r.db.Delete(&Orders{}, id)
	return result.Error
}
