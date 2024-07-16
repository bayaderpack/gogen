package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// customer_group_description is struct that represent data in Database
type CustomerGroupDescription struct {
	gorm.Model
	GroupID int `gorm:"column:group_id;default:" json:"group_id"`
Language string `gorm:"column:language;default:" json:"language"`
Title string `gorm:"column:title;default:" json:"title"`

}

// CustomerGroupDescription is interface that that model needs to implement
type CustomerGroupDescriptionRepository interface {
	Create(customer_group_description *CustomerGroupDescription) error
	GetSingle(id int) (*CustomerGroupDescription, error)
	GetAll() ([]CustomerGroupDescription, error)
	Update(customer_group_description *CustomerGroupDescription) error
	Delete(id int) error
}

type customer_group_descriptionRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCustomerGroupDescriptionRepository(db *gorm.DB) CustomerGroupDescriptionRepository {
	return &customer_group_descriptionRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *customer_group_descriptionRepository) Create(customer_group_description *CustomerGroupDescription) error {
	return r.db.Create(customer_group_description).Error
}


//Function to get single instance of customer_group_description 
func (r *customer_group_descriptionRepository) GetSingle(id int) (*CustomerGroupDescription, error) {
	var customer_group_description CustomerGroupDescription
	err := r.db.First(&customer_group_description, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &customer_group_description, nil
}


//Function to get all instances of customer_group_description 
func (r *customer_group_descriptionRepository) GetAll() ([]CustomerGroupDescription, error) {
	var customer_group_description []CustomerGroupDescription
	err := r.db.Find(&customer_group_description).Error
	return customer_group_description, err
}

//Function to update existing instances of customer_group_description 
func (r *customer_group_descriptionRepository) Update(customer_group_description *CustomerGroupDescription) error {
	return r.db.Save(customer_group_description).Error
}

//Function to delete single instances of customer_group_description 
func (r *customer_group_descriptionRepository) Delete(id int) error {
	result := r.db.Delete(&CustomerGroupDescription{}, id)
	return result.Error
}