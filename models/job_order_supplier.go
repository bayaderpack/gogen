package models

import (
	"errors"

	"gorm.io/gorm"
)

// job_order_supplier is struct that represent data in Database
type JobOrderSupplier struct {
	gorm.Model
	JobOrderID int  `gorm:"column:job_order_id;default:" json:"job_order_id"`
	SupplierID int  `gorm:"column:supplier_id;default:" json:"supplier_id"`
	Confirmed  bool `gorm:"column:confirmed;default:0" json:"confirmed"`
}

// JobOrderSupplier is interface that that model needs to implement
type JobOrderSupplierRepository interface {
	Create(job_order_supplier *JobOrderSupplier) error
	GetSingle(id int) (*JobOrderSupplier, error)
	GetAll() ([]JobOrderSupplier, error)
	Update(job_order_supplier *JobOrderSupplier) error
	Delete(id int) error
}

type job_order_supplierRepository struct {
	db *gorm.DB
}

// Create new instance
func NewJobOrderSupplierRepository(db *gorm.DB) JobOrderSupplierRepository {
	return &job_order_supplierRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *job_order_supplierRepository) Create(job_order_supplier *JobOrderSupplier) error {
	return r.db.Create(job_order_supplier).Error
}

// Function to get single instance of job_order_supplier
func (r *job_order_supplierRepository) GetSingle(id int) (*JobOrderSupplier, error) {
	var job_order_supplier JobOrderSupplier
	err := r.db.First(&job_order_supplier, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &job_order_supplier, nil
}

// Function to get all instances of job_order_supplier
func (r *job_order_supplierRepository) GetAll() ([]JobOrderSupplier, error) {
	var job_order_supplier []JobOrderSupplier
	err := r.db.Find(&job_order_supplier).Error
	return job_order_supplier, err
}

// Function to update existing instances of job_order_supplier
func (r *job_order_supplierRepository) Update(job_order_supplier *JobOrderSupplier) error {
	return r.db.Save(job_order_supplier).Error
}

// Function to delete single instances of job_order_supplier
func (r *job_order_supplierRepository) Delete(id int) error {
	result := r.db.Delete(&JobOrderSupplier{}, id)
	return result.Error
}
