package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// sample_quotation is struct that represent data in Database
type SampleQuotation struct {
	gorm.Model
	SampleID int `gorm:"column:sample_id;default:" json:"sample_id"`
QuotationID int `gorm:"column:quotation_id;default:" json:"quotation_id"`

}

// SampleQuotation is interface that that model needs to implement
type SampleQuotationRepository interface {
	Create(sample_quotation *SampleQuotation) error
	GetSingle(id int) (*SampleQuotation, error)
	GetAll() ([]SampleQuotation, error)
	Update(sample_quotation *SampleQuotation) error
	Delete(id int) error
}

type sample_quotationRepository struct {
	db *gorm.DB
}


// Create new instance
func NewSampleQuotationRepository(db *gorm.DB) SampleQuotationRepository {
	return &sample_quotationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sample_quotationRepository) Create(sample_quotation *SampleQuotation) error {
	return r.db.Create(sample_quotation).Error
}


//Function to get single instance of sample_quotation 
func (r *sample_quotationRepository) GetSingle(id int) (*SampleQuotation, error) {
	var sample_quotation SampleQuotation
	err := r.db.First(&sample_quotation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sample_quotation, nil
}


//Function to get all instances of sample_quotation 
func (r *sample_quotationRepository) GetAll() ([]SampleQuotation, error) {
	var sample_quotation []SampleQuotation
	err := r.db.Find(&sample_quotation).Error
	return sample_quotation, err
}

//Function to update existing instances of sample_quotation 
func (r *sample_quotationRepository) Update(sample_quotation *SampleQuotation) error {
	return r.db.Save(sample_quotation).Error
}

//Function to delete single instances of sample_quotation 
func (r *sample_quotationRepository) Delete(id int) error {
	result := r.db.Delete(&SampleQuotation{}, id)
	return result.Error
}