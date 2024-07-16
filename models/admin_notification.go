package models

import (
	"errors"

	"gorm.io/gorm"
)

// admin_notification is struct that represent data in Database
type AdminNotification struct {
	gorm.Model
	AdminID  int    `gorm:"column:admin_id;default:" json:"admin_id"`
	ObjectID int    `gorm:"column:object_id;default:" json:"object_id"`
	Type     string `gorm:"column:type;default:" json:"type"`
	Seen     bool   `gorm:"column:seen;default:0" json:"seen"`
}

// AdminNotification is interface that that model needs to implement
type AdminNotificationRepository interface {
	Create(admin_notification *AdminNotification) error
	GetSingle(id int) (*AdminNotification, error)
	GetAll() ([]AdminNotification, error)
	Update(admin_notification *AdminNotification) error
	Delete(id int) error
}

type admin_notificationRepository struct {
	db *gorm.DB
}

// Create new instance
func NewAdminNotificationRepository(db *gorm.DB) AdminNotificationRepository {
	return &admin_notificationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *admin_notificationRepository) Create(admin_notification *AdminNotification) error {
	return r.db.Create(admin_notification).Error
}

// Function to get single instance of admin_notification
func (r *admin_notificationRepository) GetSingle(id int) (*AdminNotification, error) {
	var admin_notification AdminNotification
	err := r.db.First(&admin_notification, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &admin_notification, nil
}

// Function to get all instances of admin_notification
func (r *admin_notificationRepository) GetAll() ([]AdminNotification, error) {
	var admin_notification []AdminNotification
	err := r.db.Find(&admin_notification).Error
	return admin_notification, err
}

// Function to update existing instances of admin_notification
func (r *admin_notificationRepository) Update(admin_notification *AdminNotification) error {
	return r.db.Save(admin_notification).Error
}

// Function to delete single instances of admin_notification
func (r *admin_notificationRepository) Delete(id int) error {
	result := r.db.Delete(&AdminNotification{}, id)
	return result.Error
}
