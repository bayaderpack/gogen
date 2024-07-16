package models

import (
	"errors"

	"gorm.io/gorm"
)

// sample_assignee is struct that represent data in Database
type SampleAssignee struct {
	gorm.Model
	SampleID int `gorm:"column:sample_id;default:" json:"sample_id"`
	AdminID  int `gorm:"column:admin_id;default:" json:"admin_id"`
}

// SampleAssignee is interface that that model needs to implement
type SampleAssigneeRepository interface {
	Create(sample_assignee *SampleAssignee) error
	GetSingle(id int) (*SampleAssignee, error)
	GetAll() ([]SampleAssignee, error)
	Update(sample_assignee *SampleAssignee) error
	Delete(id int) error
}

type sample_assigneeRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSampleAssigneeRepository(db *gorm.DB) SampleAssigneeRepository {
	return &sample_assigneeRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sample_assigneeRepository) Create(sample_assignee *SampleAssignee) error {
	return r.db.Create(sample_assignee).Error
}

// Function to get single instance of sample_assignee
func (r *sample_assigneeRepository) GetSingle(id int) (*SampleAssignee, error) {
	var sample_assignee SampleAssignee
	err := r.db.First(&sample_assignee, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sample_assignee, nil
}

// Function to get all instances of sample_assignee
func (r *sample_assigneeRepository) GetAll() ([]SampleAssignee, error) {
	var sample_assignee []SampleAssignee
	err := r.db.Find(&sample_assignee).Error
	return sample_assignee, err
}

// Function to update existing instances of sample_assignee
func (r *sample_assigneeRepository) Update(sample_assignee *SampleAssignee) error {
	return r.db.Save(sample_assignee).Error
}

// Function to delete single instances of sample_assignee
func (r *sample_assigneeRepository) Delete(id int) error {
	result := r.db.Delete(&SampleAssignee{}, id)
	return result.Error
}
