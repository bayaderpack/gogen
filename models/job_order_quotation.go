package models

import (
	"errors"

	"gorm.io/gorm"
)

// job_order_quotation is struct that represent data in Database
type JobOrderQuotation struct {
	gorm.Model
	JobOrderID  int `gorm:"column:job_order_id;default:" json:"job_order_id"`
	QuotationID int `gorm:"column:quotation_id;default:" json:"quotation_id"`
}

// JobOrderQuotation is interface that that model needs to implement
type JobOrderQuotationRepository interface {
	Create(job_order_quotation *JobOrderQuotation) error
	GetSingle(id int) (*JobOrderQuotation, error)
	GetAll() ([]JobOrderQuotation, error)
	Update(job_order_quotation *JobOrderQuotation) error
	Delete(id int) error
}

type job_order_quotationRepository struct {
	db *gorm.DB
}

// Create new instance
func NewJobOrderQuotationRepository(db *gorm.DB) JobOrderQuotationRepository {
	return &job_order_quotationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *job_order_quotationRepository) Create(job_order_quotation *JobOrderQuotation) error {
	return r.db.Create(job_order_quotation).Error
}

// Function to get single instance of job_order_quotation
func (r *job_order_quotationRepository) GetSingle(id int) (*JobOrderQuotation, error) {
	var job_order_quotation JobOrderQuotation
	err := r.db.First(&job_order_quotation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &job_order_quotation, nil
}

// Function to get all instances of job_order_quotation
func (r *job_order_quotationRepository) GetAll() ([]JobOrderQuotation, error) {
	var job_order_quotation []JobOrderQuotation
	err := r.db.Find(&job_order_quotation).Error
	return job_order_quotation, err
}

// Function to update existing instances of job_order_quotation
func (r *job_order_quotationRepository) Update(job_order_quotation *JobOrderQuotation) error {
	return r.db.Save(job_order_quotation).Error
}

// Function to delete single instances of job_order_quotation
func (r *job_order_quotationRepository) Delete(id int) error {
	result := r.db.Delete(&JobOrderQuotation{}, id)
	return result.Error
}
