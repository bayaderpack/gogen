package models

import (
	"errors"

	"gorm.io/gorm"
)

// supplier is struct that represent data in Database
type Supplier struct {
	gorm.Model
	SupplierID  int    `gorm:"column:supplier_id;default:" json:"supplier_id"`
	GroupID     int    `gorm:"column:group_id;default:" json:"group_id"`
	Name        string `gorm:"column:name;default:" json:"name"`
	ContactName string `gorm:"column:contact_name;default:NULL" json:"contact_name"`
	CountryCode int    `gorm:"column:country_code;default:NULL" json:"country_code"`
	Mobile      int    `gorm:"column:mobile;default:NULL" json:"mobile"`
	City        string `gorm:"column:city;default:" json:"city"`
	Address     string `gorm:"column:address;default:" json:"address"`
	Location    int    `gorm:"column:location;default:" json:"location"`
}

// Supplier is interface that that model needs to implement
type SupplierRepository interface {
	Create(supplier *Supplier) error
	GetSingle(id int) (*Supplier, error)
	GetAll() ([]Supplier, error)
	Update(supplier *Supplier) error
	Delete(id int) error
}

type supplierRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *supplierRepository) Create(supplier *Supplier) error {
	return r.db.Create(supplier).Error
}

// Function to get single instance of supplier
func (r *supplierRepository) GetSingle(id int) (*Supplier, error) {
	var supplier Supplier
	err := r.db.First(&supplier, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &supplier, nil
}

// Function to get all instances of supplier
func (r *supplierRepository) GetAll() ([]Supplier, error) {
	var supplier []Supplier
	err := r.db.Find(&supplier).Error
	return supplier, err
}

// Function to update existing instances of supplier
func (r *supplierRepository) Update(supplier *Supplier) error {
	return r.db.Save(supplier).Error
}

// Function to delete single instances of supplier
func (r *supplierRepository) Delete(id int) error {
	result := r.db.Delete(&Supplier{}, id)
	return result.Error
}
