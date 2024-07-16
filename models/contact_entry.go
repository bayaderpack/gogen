package models

import (
	"errors"

	"gorm.io/gorm"
)

// contact_entry is struct that represent data in Database
type ContactEntry struct {
	gorm.Model
	EntryID     int    `gorm:"column:entry_id;default:" json:"entry_id"`
	CustomerID  string `gorm:"column:customer_id;default:" json:"customer_id"`
	Name        string `gorm:"column:name;default:" json:"name"`
	CountryCode string `gorm:"column:country_code;default:" json:"country_code"`
	Mobile      int    `gorm:"column:mobile;default:" json:"mobile"`
	Email       string `gorm:"column:email;default:" json:"email"`
	Company     string `gorm:"column:company;default:" json:"company"`
	Content     string `gorm:"column:content;default:" json:"content"`
	Know        string `gorm:"column:know;default:" json:"know"`
	Ip          string `gorm:"column:ip;default:" json:"ip"`
	Viewed      bool   `gorm:"column:viewed;default:1" json:"viewed"`
}

// ContactEntry is interface that that model needs to implement
type ContactEntryRepository interface {
	Create(contact_entry *ContactEntry) error
	GetSingle(id int) (*ContactEntry, error)
	GetAll() ([]ContactEntry, error)
	Update(contact_entry *ContactEntry) error
	Delete(id int) error
}

type contact_entryRepository struct {
	db *gorm.DB
}

// Create new instance
func NewContactEntryRepository(db *gorm.DB) ContactEntryRepository {
	return &contact_entryRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *contact_entryRepository) Create(contact_entry *ContactEntry) error {
	return r.db.Create(contact_entry).Error
}

// Function to get single instance of contact_entry
func (r *contact_entryRepository) GetSingle(id int) (*ContactEntry, error) {
	var contact_entry ContactEntry
	err := r.db.First(&contact_entry, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &contact_entry, nil
}

// Function to get all instances of contact_entry
func (r *contact_entryRepository) GetAll() ([]ContactEntry, error) {
	var contact_entry []ContactEntry
	err := r.db.Find(&contact_entry).Error
	return contact_entry, err
}

// Function to update existing instances of contact_entry
func (r *contact_entryRepository) Update(contact_entry *ContactEntry) error {
	return r.db.Save(contact_entry).Error
}

// Function to delete single instances of contact_entry
func (r *contact_entryRepository) Delete(id int) error {
	result := r.db.Delete(&ContactEntry{}, id)
	return result.Error
}
