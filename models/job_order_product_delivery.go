package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// job_order_product_delivery is struct that represent data in Database
type JobOrderProductDelivery struct {
	gorm.Model
	JobOrderProductID int        `gorm:"column:job_order_product_id;default:" json:"job_order_product_id"`
	Quantity          int        `gorm:"column:quantity;default:" json:"quantity"`
	Boxes             int        `gorm:"column:boxes;default:" json:"boxes"`
	Filling           int        `gorm:"column:filling;default:" json:"filling"`
	LastFilling       int        `gorm:"column:last_filling;default:0" json:"last_filling"`
	Note              string     `gorm:"column:note;default:" json:"note"`
	DeliveryDate      *time.Time `gorm:"column:delivery_date;default:current_timestamp()" json:"delivery_date"`
	AdminID           int        `gorm:"column:admin_id;default:" json:"admin_id"`
}

// JobOrderProductDelivery is interface that that model needs to implement
type JobOrderProductDeliveryRepository interface {
	Create(job_order_product_delivery *JobOrderProductDelivery) error
	GetSingle(id int) (*JobOrderProductDelivery, error)
	GetAll() ([]JobOrderProductDelivery, error)
	Update(job_order_product_delivery *JobOrderProductDelivery) error
	Delete(id int) error
}

type job_order_product_deliveryRepository struct {
	db *gorm.DB
}

// Create new instance
func NewJobOrderProductDeliveryRepository(db *gorm.DB) JobOrderProductDeliveryRepository {
	return &job_order_product_deliveryRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *job_order_product_deliveryRepository) Create(job_order_product_delivery *JobOrderProductDelivery) error {
	return r.db.Create(job_order_product_delivery).Error
}

// Function to get single instance of job_order_product_delivery
func (r *job_order_product_deliveryRepository) GetSingle(id int) (*JobOrderProductDelivery, error) {
	var job_order_product_delivery JobOrderProductDelivery
	err := r.db.First(&job_order_product_delivery, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &job_order_product_delivery, nil
}

// Function to get all instances of job_order_product_delivery
func (r *job_order_product_deliveryRepository) GetAll() ([]JobOrderProductDelivery, error) {
	var job_order_product_delivery []JobOrderProductDelivery
	err := r.db.Find(&job_order_product_delivery).Error
	return job_order_product_delivery, err
}

// Function to update existing instances of job_order_product_delivery
func (r *job_order_product_deliveryRepository) Update(job_order_product_delivery *JobOrderProductDelivery) error {
	return r.db.Save(job_order_product_delivery).Error
}

// Function to delete single instances of job_order_product_delivery
func (r *job_order_product_deliveryRepository) Delete(id int) error {
	result := r.db.Delete(&JobOrderProductDelivery{}, id)
	return result.Error
}
