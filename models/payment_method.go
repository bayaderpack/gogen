package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// payment_method is struct that represent data in Database
type PaymentMethod struct {
	gorm.Model
	PaymentID int `gorm:"column:payment_id;default:" json:"payment_id"`
Title string `gorm:"column:title;default:" json:"title"`
Status bool `gorm:"column:status;default:" json:"status"`
IsPrimary bool `gorm:"column:is_primary;default:" json:"is_primary"`
Configuration string `gorm:"column:configuration;default:NULL" json:"configuration"`

}

// PaymentMethod is interface that that model needs to implement
type PaymentMethodRepository interface {
	Create(payment_method *PaymentMethod) error
	GetSingle(id int) (*PaymentMethod, error)
	GetAll() ([]PaymentMethod, error)
	Update(payment_method *PaymentMethod) error
	Delete(id int) error
}

type payment_methodRepository struct {
	db *gorm.DB
}


// Create new instance
func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &payment_methodRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *payment_methodRepository) Create(payment_method *PaymentMethod) error {
	return r.db.Create(payment_method).Error
}


//Function to get single instance of payment_method 
func (r *payment_methodRepository) GetSingle(id int) (*PaymentMethod, error) {
	var payment_method PaymentMethod
	err := r.db.First(&payment_method, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &payment_method, nil
}


//Function to get all instances of payment_method 
func (r *payment_methodRepository) GetAll() ([]PaymentMethod, error) {
	var payment_method []PaymentMethod
	err := r.db.Find(&payment_method).Error
	return payment_method, err
}

//Function to update existing instances of payment_method 
func (r *payment_methodRepository) Update(payment_method *PaymentMethod) error {
	return r.db.Save(payment_method).Error
}

//Function to delete single instances of payment_method 
func (r *payment_methodRepository) Delete(id int) error {
	result := r.db.Delete(&PaymentMethod{}, id)
	return result.Error
}