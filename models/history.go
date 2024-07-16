package models

import (
	"errors"

	"gorm.io/gorm"
)

// history is struct that represent data in Database
type History struct {
	HistoryID    int    `gorm:"column:history_id;default:" json:"history_id"`
	StatusID     int    `gorm:"column:status_id;default:" json:"status_id"`
	OrderID      int    `gorm:"column:order_id;default:0" json:"order_id"`
	QuotationID  int    `gorm:"column:quotation_id;default:0" json:"quotation_id"`
	SampleID     int    `gorm:"column:sample_id;default:0" json:"sample_id"`
	JobID        int    `gorm:"column:job_id;default:0" json:"job_id"`
	CustomerID   int    `gorm:"column:customer_id;default:0" json:"customer_id"`
	AdminID      int    `gorm:"column:admin_id;default:0" json:"admin_id"`
	Comment      string `gorm:"column:comment;default:" json:"comment"`
	Notified     string `gorm:"column:notified;default:NULL" json:"notified"`
	ShowCustomer bool   `gorm:"column:show_customer;default:0" json:"show_customer"`
	Uploads      string `gorm:"column:uploads;default:NULL" json:"uploads"`
}

// History is interface that that model needs to implement
type HistoryRepository interface {
	Create(history *History) error
	GetSingle(id int) (*History, error)
	GetAll() ([]History, error)
	Update(history *History) error
	Delete(id int) error
}

type historyRepository struct {
	db *gorm.DB
}

// Create new instance
func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *historyRepository) Create(history *History) error {
	return r.db.Create(history).Error
}

// Function to get single instance of history
func (r *historyRepository) GetSingle(id int) (*History, error) {
	var history History
	err := r.db.Table("history").First(&history, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &history, nil
}

// Function to get all instances of history
func (r *historyRepository) GetAll() ([]History, error) {
	var history []History
	err := r.db.Table("history").Find(&history).Error
	return history, err
}

// Function to update existing instances of history
func (r *historyRepository) Update(history *History) error {
	return r.db.Save(history).Error
}

// Function to delete single instances of history
func (r *historyRepository) Delete(id int) error {
	result := r.db.Delete(&History{}, id)
	return result.Error
}
