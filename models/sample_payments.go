package models

import (
	"errors"

	"gorm.io/gorm"
)

// sample_payments is struct that represent data in Database
type SamplePayments struct {
	gorm.Model
	SampleID  int     `gorm:"column:sample_id;default:" json:"sample_id"`
	AdminID   int     `gorm:"column:admin_id;default:" json:"admin_id"`
	Amount    float32 `gorm:"column:amount;default:" json:"amount"`
	Remaining float32 `gorm:"column:remaining;default:0.0000" json:"remaining"`
	Note      string  `gorm:"column:note;default:NULL" json:"note"`
}

// SamplePayments is interface that that model needs to implement
type SamplePaymentsRepository interface {
	Create(sample_payment *SamplePayments) error
	GetSingle(id int) (*SamplePayments, error)
	GetAll() ([]SamplePayments, error)
	Update(sample_payment *SamplePayments) error
	Delete(id int) error
}

type sample_paymentRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSamplePaymentsRepository(db *gorm.DB) SamplePaymentsRepository {
	return &sample_paymentRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sample_paymentRepository) Create(sample_payment *SamplePayments) error {
	return r.db.Create(sample_payment).Error
}

// Function to get single instance of sample_payment
func (r *sample_paymentRepository) GetSingle(id int) (*SamplePayments, error) {
	var sample_payment SamplePayments
	err := r.db.First(&sample_payment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sample_payment, nil
}

// Function to get all instances of sample_payments
func (r *sample_paymentRepository) GetAll() ([]SamplePayments, error) {
	var sample_payments []SamplePayments
	err := r.db.Find(&sample_payments).Error
	return sample_payments, err
}

// Function to update existing instances of sample_payment
func (r *sample_paymentRepository) Update(sample_payment *SamplePayments) error {
	return r.db.Save(sample_payment).Error
}

// Function to delete single instances of sample_payment
func (r *sample_paymentRepository) Delete(id int) error {
	result := r.db.Delete(&SamplePayments{}, id)
	return result.Error
}
