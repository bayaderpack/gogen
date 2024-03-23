package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// event_assignee is struct that represent data in Database
type EventAssignee struct {
	gorm.Model
	EventID int `gorm:"column:event_id;default:" json:"event_id"`
AdminID int `gorm:"column:admin_id;default:" json:"admin_id"`

}

// EventAssignee is interface that that model needs to implement
type EventAssigneeRepository interface {
	Create(event_assignee *EventAssignee) error
	GetSingle(id int) (*EventAssignee, error)
	GetAll() ([]EventAssignee, error)
	Update(event_assignee *EventAssignee) error
	Delete(id int) error
}

type event_assigneeRepository struct {
	db *gorm.DB
}


// Create new instance
func NewEventAssigneeRepository(db *gorm.DB) EventAssigneeRepository {
	return &event_assigneeRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *event_assigneeRepository) Create(event_assignee *EventAssignee) error {
	return r.db.Create(event_assignee).Error
}


//Function to get single instance of event_assignee 
func (r *event_assigneeRepository) GetSingle(id int) (*EventAssignee, error) {
	var event_assignee EventAssignee
	err := r.db.First(&event_assignee, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &event_assignee, nil
}


//Function to get all instances of event_assignee 
func (r *event_assigneeRepository) GetAll() ([]EventAssignee, error) {
	var event_assignee []EventAssignee
	err := r.db.Find(&event_assignee).Error
	return event_assignee, err
}

//Function to update existing instances of event_assignee 
func (r *event_assigneeRepository) Update(event_assignee *EventAssignee) error {
	return r.db.Save(event_assignee).Error
}

//Function to delete single instances of event_assignee 
func (r *event_assigneeRepository) Delete(id int) error {
	result := r.db.Delete(&EventAssignee{}, id)
	return result.Error
}