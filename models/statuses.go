package models

import (
	"errors"

	"gorm.io/gorm"
)

// statuses is struct that represent data in Database
type Statuses struct {
	gorm.Model
	StatusID int    `gorm:"column:status_id;default:" json:"status_id"`
	Language string `gorm:"column:language;default:" json:"language"`
	Text     string `gorm:"column:text;default:" json:"text"`
	Code     string `gorm:"column:code;default:" json:"code"`
}

// Statuses is interface that that model needs to implement
type StatusesRepository interface {
	Create(status *Statuses) error
	GetSingle(id int) (*Statuses, error)
	GetAll() ([]Statuses, error)
	Update(status *Statuses) error
	Delete(id int) error
}

type statusRepository struct {
	db *gorm.DB
}

// Create new instance
func NewStatusesRepository(db *gorm.DB) StatusesRepository {
	return &statusRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *statusRepository) Create(status *Statuses) error {
	return r.db.Create(status).Error
}

// Function to get single instance of status
func (r *statusRepository) GetSingle(id int) (*Statuses, error) {
	var status Statuses
	err := r.db.First(&status, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &status, nil
}

// Function to get all instances of statuses
func (r *statusRepository) GetAll() ([]Statuses, error) {
	var statuses []Statuses
	err := r.db.Find(&statuses).Error
	return statuses, err
}

// Function to update existing instances of status
func (r *statusRepository) Update(status *Statuses) error {
	return r.db.Save(status).Error
}

// Function to delete single instances of status
func (r *statusRepository) Delete(id int) error {
	result := r.db.Delete(&Statuses{}, id)
	return result.Error
}
