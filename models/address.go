package models

import (
	"errors"

	"gorm.io/gorm"
)

// address is struct that represent data in Database
type Address struct {
	gorm.Model
	AddressID   int    `gorm:"column:address_id;default:" json:"address_id"`
	CustomerID  int    `gorm:"column:customer_id;default:" json:"customer_id"`
	Firstname   string `gorm:"column:firstname;default:" json:"firstname"`
	Lastname    string `gorm:"column:lastname;default:" json:"lastname"`
	CountryCode int    `gorm:"column:country_code;default:" json:"country_code"`
	Mobile      int    `gorm:"column:mobile;default:" json:"mobile"`
	Line1       string `gorm:"column:line1;default:" json:"line1"`
	Location    int    `gorm:"column:location;default:" json:"location"`
	Country     string `gorm:"column:country;default:" json:"country"`
	City        string `gorm:"column:city;default:" json:"city"`
	StateCode   string `gorm:"column:state_code;default:" json:"state_code"`
	PostCode    string `gorm:"column:post_code;default:" json:"post_code"`
	IsPrimary   bool   `gorm:"column:is_primary;default:0" json:"is_primary"`
}

// Address is interface that that model needs to implement
type AddressRepository interface {
	Create(address *Address) error
	GetSingle(id int) (*Address, error)
	GetAll() ([]Address, error)
	Update(address *Address) error
	Delete(id int) error
}

type addressRepository struct {
	db *gorm.DB
}

// Create new instance
func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *addressRepository) Create(address *Address) error {
	return r.db.Create(address).Error
}

// Function to get single instance of address
func (r *addressRepository) GetSingle(id int) (*Address, error) {
	var address Address
	err := r.db.First(&address, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &address, nil
}

// Function to get all instances of address
func (r *addressRepository) GetAll() ([]Address, error) {
	var address []Address
	err := r.db.Find(&address).Error
	return address, err
}

// Function to update existing instances of address
func (r *addressRepository) Update(address *Address) error {
	return r.db.Save(address).Error
}

// Function to delete single instances of address
func (r *addressRepository) Delete(id int) error {
	result := r.db.Delete(&Address{}, id)
	return result.Error
}
