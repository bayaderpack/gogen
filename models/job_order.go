package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// job_order is struct that represent data in Database
type JobOrder struct {
	gorm.Model
	JobOrderID          int        `gorm:"column:job_order_id;default:" json:"job_order_id"`
	Type                bool       `gorm:"column:type;default:1" json:"type"`
	Deadline            *time.Time `gorm:"column:deadline;default:" json:"deadline"`
	AuthorizationLetter string     `gorm:"column:authorization_letter;default:NULL" json:"authorization_letter"`
	AgencyLetter        string     `gorm:"column:agency_letter;default:NULL" json:"agency_letter"`
	Uuid                string     `gorm:"column:uuid;default:NULL" json:"uuid"`
	Signature           string     `gorm:"column:signature;default:NULL" json:"signature"`
	Stamp               string     `gorm:"column:stamp;default:NULL" json:"stamp"`
	SignedDate          *time.Time `gorm:"column:signed_date;default:NULL" json:"signed_date"`
	PrintedBy           int        `gorm:"column:printed_by;default:NULL" json:"printed_by"`
}

// JobOrder is interface that that model needs to implement
type JobOrderRepository interface {
	Create(job_order *JobOrder) error
	GetSingle(id int) (*JobOrder, error)
	GetAll() ([]JobOrder, error)
	Update(job_order *JobOrder) error
	Delete(id int) error
}

type job_orderRepository struct {
	db *gorm.DB
}

// Create new instance
func NewJobOrderRepository(db *gorm.DB) JobOrderRepository {
	return &job_orderRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *job_orderRepository) Create(job_order *JobOrder) error {
	return r.db.Create(job_order).Error
}

// Function to get single instance of job_order
func (r *job_orderRepository) GetSingle(id int) (*JobOrder, error) {
	var job_order JobOrder
	err := r.db.First(&job_order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &job_order, nil
}

// Function to get all instances of job_order
func (r *job_orderRepository) GetAll() ([]JobOrder, error) {
	var job_order []JobOrder
	err := r.db.Find(&job_order).Error
	return job_order, err
}

// Function to update existing instances of job_order
func (r *job_orderRepository) Update(job_order *JobOrder) error {
	return r.db.Save(job_order).Error
}

// Function to delete single instances of job_order
func (r *job_orderRepository) Delete(id int) error {
	result := r.db.Delete(&JobOrder{}, id)
	return result.Error
}
