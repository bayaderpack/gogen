package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// customer is struct that represent data in Database
type Customer struct {
	gorm.Model
	CustomerID          int        `gorm:"column:customer_id;default:" json:"customer_id"`
	GroupID             int        `gorm:"column:group_id;default:" json:"group_id"`
	SellerID            int        `gorm:"column:seller_id;default:NULL" json:"seller_id"`
	Kid                 string     `gorm:"column:kid;default:" json:"kid"`
	Firstname           string     `gorm:"column:firstname;default:" json:"firstname"`
	Lastname            string     `gorm:"column:lastname;default:" json:"lastname"`
	Position            string     `gorm:"column:position;default:NULL" json:"position"`
	Email               string     `gorm:"column:email;default:" json:"email"`
	CountryCode         int        `gorm:"column:country_code;default:" json:"country_code"`
	Mobile              int        `gorm:"column:mobile;default:" json:"mobile"`
	Company             string     `gorm:"column:company;default:NULL" json:"company"`
	CompanyPhone        string     `gorm:"column:company_phone;default:" json:"company_phone"`
	Address             string     `gorm:"column:address;default:NULL" json:"address"`
	Location            int        `gorm:"column:location;default:NULL" json:"location"`
	Files               string     `gorm:"column:files;default:NULL" json:"files"`
	Registration        string     `gorm:"column:registration;default:NULL" json:"registration"`
	RegistrationExpires *time.Time `gorm:"column:registration_expires;default:NULL" json:"registration_expires"`
	Vat                 string     `gorm:"column:vat;default:NULL" json:"vat"`
	Password            string     `gorm:"column:password;default:" json:"password"`
	Newsletter          bool       `gorm:"column:newsletter;default:0" json:"newsletter"`
	Ip                  string     `gorm:"column:ip;default:" json:"ip"`
	Status              bool       `gorm:"column:status;default:0" json:"status"`
	ResetToken          string     `gorm:"column:reset_token;default:NULL" json:"reset_token"`
	ResetExpires        *time.Time `gorm:"column:reset_expires;default:NULL" json:"reset_expires"`
	Otp                 int        `gorm:"column:otp;default:NULL" json:"otp"`
}

// Customer is interface that that model needs to implement
type CustomerRepository interface {
	Create(customer *Customer) error
	GetSingle(id int) (*Customer, error)
	GetAll() ([]Customer, error)
	Update(customer *Customer) error
	Delete(id int) error
}

type customerRepository struct {
	db *gorm.DB
}

// Create new instance
func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *customerRepository) Create(customer *Customer) error {
	return r.db.Create(customer).Error
}

// Function to get single instance of customer
func (r *customerRepository) GetSingle(id int) (*Customer, error) {
	var customer Customer
	err := r.db.First(&customer, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &customer, nil
}

// Function to get all instances of customer
func (r *customerRepository) GetAll() ([]Customer, error) {
	var customer []Customer
	err := r.db.Find(&customer).Error
	return customer, err
}

// Function to update existing instances of customer
func (r *customerRepository) Update(customer *Customer) error {
	return r.db.Save(customer).Error
}

// Function to delete single instances of customer
func (r *customerRepository) Delete(id int) error {
	result := r.db.Delete(&Customer{}, id)
	return result.Error
}
