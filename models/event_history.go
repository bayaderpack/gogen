package models

import (
	"errors"

	"gorm.io/gorm"
)

// event_history is struct that represent data in Database
type EventHistory struct {
	gorm.Model
	EventHistoryID int    `gorm:"column:event_history_id;default:" json:"event_history_id"`
	EventID        int    `gorm:"column:event_id;default:" json:"event_id"`
	CustomerID     int    `gorm:"column:customer_id;default:0" json:"customer_id"`
	AdminID        int    `gorm:"column:admin_id;default:0" json:"admin_id"`
	Message        string `gorm:"column:message;default:" json:"message"`
	Args           string `gorm:"column:args;default:'{}'" json:"args"`
}

// EventHistory is interface that that model needs to implement
type EventHistoryRepository interface {
	Create(event_history *EventHistory) error
	GetSingle(id int) (*EventHistory, error)
	GetAll() ([]EventHistory, error)
	Update(event_history *EventHistory) error
	Delete(id int) error
}

type event_historyRepository struct {
	db *gorm.DB
}

// Create new instance
func NewEventHistoryRepository(db *gorm.DB) EventHistoryRepository {
	return &event_historyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *event_historyRepository) Create(event_history *EventHistory) error {
	return r.db.Create(event_history).Error
}

// Function to get single instance of event_history
func (r *event_historyRepository) GetSingle(id int) (*EventHistory, error) {
	var event_history EventHistory
	err := r.db.First(&event_history, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &event_history, nil
}

// Function to get all instances of event_history
func (r *event_historyRepository) GetAll() ([]EventHistory, error) {
	var event_history []EventHistory
	err := r.db.Find(&event_history).Error
	return event_history, err
}

// Function to update existing instances of event_history
func (r *event_historyRepository) Update(event_history *EventHistory) error {
	return r.db.Save(event_history).Error
}

// Function to delete single instances of event_history
func (r *event_historyRepository) Delete(id int) error {
	result := r.db.Delete(&EventHistory{}, id)
	return result.Error
}
