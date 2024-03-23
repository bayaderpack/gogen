package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// quotation is struct that represent data in Database
type Quotation struct {
	gorm.Model
	QuotationID int `gorm:"column:quotation_id;default:" json:"quotation_id"`
CustomerID int `gorm:"column:customer_id;default:0" json:"customer_id"`
Number string `gorm:"column:number;default:'0'" json:"number"`
Viewed bool `gorm:"column:viewed;default:0" json:"viewed"`

}

// Quotation is interface that that model needs to implement
type QuotationRepository interface {
	Create(quotation *Quotation) error
	GetSingle(id int) (*Quotation, error)
	GetAll() ([]Quotation, error)
	Update(quotation *Quotation) error
	Delete(id int) error
}

type quotationRepository struct {
	db *gorm.DB
}


// Create new instance
func NewQuotationRepository(db *gorm.DB) QuotationRepository {
	return &quotationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotationRepository) Create(quotation *Quotation) error {
	return r.db.Create(quotation).Error
}


//Function to get single instance of quotation 
func (r *quotationRepository) GetSingle(id int) (*Quotation, error) {
	var quotation Quotation
	err := r.db.First(&quotation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation, nil
}


//Function to get all instances of quotation 
func (r *quotationRepository) GetAll() ([]Quotation, error) {
	var quotation []Quotation
	err := r.db.Find(&quotation).Error
	return quotation, err
}

//Function to update existing instances of quotation 
func (r *quotationRepository) Update(quotation *Quotation) error {
	return r.db.Save(quotation).Error
}

//Function to delete single instances of quotation 
func (r *quotationRepository) Delete(id int) error {
	result := r.db.Delete(&Quotation{}, id)
	return result.Error
}