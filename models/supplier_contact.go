package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// supplier_contact is struct that represent data in Database
type SupplierContact struct {
	gorm.Model
	SupplierID int `gorm:"column:supplier_id;default:" json:"supplier_id"`
Name string `gorm:"column:name;default:" json:"name"`
CountryCode int `gorm:"column:country_code;default:" json:"country_code"`
Mobile int `gorm:"column:mobile;default:" json:"mobile"`

}

// SupplierContact is interface that that model needs to implement
type SupplierContactRepository interface {
	Create(supplier_contact *SupplierContact) error
	GetSingle(id int) (*SupplierContact, error)
	GetAll() ([]SupplierContact, error)
	Update(supplier_contact *SupplierContact) error
	Delete(id int) error
}

type supplier_contactRepository struct {
	db *gorm.DB
}


// Create new instance
func NewSupplierContactRepository(db *gorm.DB) SupplierContactRepository {
	return &supplier_contactRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *supplier_contactRepository) Create(supplier_contact *SupplierContact) error {
	return r.db.Create(supplier_contact).Error
}


//Function to get single instance of supplier_contact 
func (r *supplier_contactRepository) GetSingle(id int) (*SupplierContact, error) {
	var supplier_contact SupplierContact
	err := r.db.First(&supplier_contact, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &supplier_contact, nil
}


//Function to get all instances of supplier_contact 
func (r *supplier_contactRepository) GetAll() ([]SupplierContact, error) {
	var supplier_contact []SupplierContact
	err := r.db.Find(&supplier_contact).Error
	return supplier_contact, err
}

//Function to update existing instances of supplier_contact 
func (r *supplier_contactRepository) Update(supplier_contact *SupplierContact) error {
	return r.db.Save(supplier_contact).Error
}

//Function to delete single instances of supplier_contact 
func (r *supplier_contactRepository) Delete(id int) error {
	result := r.db.Delete(&SupplierContact{}, id)
	return result.Error
}