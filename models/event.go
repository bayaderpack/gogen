package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// event is struct that represent data in Database
type Event struct {
	gorm.Model
	EventID     int        `gorm:"column:event_id;default:" json:"event_id"`
	Title       string     `gorm:"column:title;default:" json:"title"`
	Name        string     `gorm:"column:name;default:NULL" json:"name"`
	Start       *time.Time `gorm:"column:start;default:" json:"start"`
	End         *time.Time `gorm:"column:end;default:" json:"end"`
	CountryCode int        `gorm:"column:country_code;default:" json:"country_code"`
	Mobile      int        `gorm:"column:mobile;default:" json:"mobile"`
	Company     string     `gorm:"column:company;default:" json:"company"`
	CompanyDesc string     `gorm:"column:company_desc;default:NULL" json:"company_desc"`
	Address     string     `gorm:"column:address;default:" json:"address"`
	Location    int        `gorm:"column:location;default:" json:"location"`
	Type        bool       `gorm:"column:type;default:2" json:"type"`
	Status      bool       `gorm:"column:status;default:1" json:"status"`
}

// Event is interface that that model needs to implement
type EventRepository interface {
	Create(event *Event) error
	GetSingle(id int) (*Event, error)
	GetAll() ([]Event, error)
	Update(event *Event) error
	Delete(id int) error
}

type eventRepository struct {
	db *gorm.DB
}

// Create new instance
func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *eventRepository) Create(event *Event) error {
	return r.db.Create(event).Error
}

// Function to get single instance of event
func (r *eventRepository) GetSingle(id int) (*Event, error) {
	var event Event
	err := r.db.First(&event, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &event, nil
}

// Function to get all instances of event
func (r *eventRepository) GetAll() ([]Event, error) {
	var event []Event
	err := r.db.Find(&event).Error
	return event, err
}

// Function to update existing instances of event
func (r *eventRepository) Update(event *Event) error {
	return r.db.Save(event).Error
}

// Function to delete single instances of event
func (r *eventRepository) Delete(id int) error {
	result := r.db.Delete(&Event{}, id)
	return result.Error
}
