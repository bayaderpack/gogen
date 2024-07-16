package models

import (
	"time"
	"errors"
	"gorm.io/gorm"
)


// sample is struct that represent data in Database
type Sample struct {
	gorm.Model
	SampleID int `gorm:"column:sample_id;default:" json:"sample_id"`
CustomerID int `gorm:"column:customer_id;default:NULL" json:"customer_id"`
Viewed bool `gorm:"column:viewed;default:0" json:"viewed"`
Deadline *time.Time `gorm:"column:deadline;default:current_timestamp()" json:"deadline"`

}

// Sample is interface that that model needs to implement
type SampleRepository interface {
	Create(sample *Sample) error
	GetSingle(id int) (*Sample, error)
	GetAll() ([]Sample, error)
	Update(sample *Sample) error
	Delete(id int) error
}

type sampleRepository struct {
	db *gorm.DB
}


// Create new instance
func NewSampleRepository(db *gorm.DB) SampleRepository {
	return &sampleRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sampleRepository) Create(sample *Sample) error {
	return r.db.Create(sample).Error
}


//Function to get single instance of sample 
func (r *sampleRepository) GetSingle(id int) (*Sample, error) {
	var sample Sample
	err := r.db.First(&sample, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sample, nil
}


//Function to get all instances of sample 
func (r *sampleRepository) GetAll() ([]Sample, error) {
	var sample []Sample
	err := r.db.Find(&sample).Error
	return sample, err
}

//Function to update existing instances of sample 
func (r *sampleRepository) Update(sample *Sample) error {
	return r.db.Save(sample).Error
}

//Function to delete single instances of sample 
func (r *sampleRepository) Delete(id int) error {
	result := r.db.Delete(&Sample{}, id)
	return result.Error
}