package models

import (
	"errors"

	"gorm.io/gorm"
)

// quotation_product_relation is struct that represent data in Database
type QuotationProductRelation struct {
	gorm.Model
	QuotationID        int `gorm:"column:quotation_id;default:" json:"quotation_id"`
	QuotationProductID int `gorm:"column:quotation_product_id;default:" json:"quotation_product_id"`
}

// QuotationProductRelation is interface that that model needs to implement
type QuotationProductRelationRepository interface {
	Create(quotation_product_relation *QuotationProductRelation) error
	GetSingle(id int) (*QuotationProductRelation, error)
	GetAll() ([]QuotationProductRelation, error)
	Update(quotation_product_relation *QuotationProductRelation) error
	Delete(id int) error
}

type quotation_product_relationRepository struct {
	db *gorm.DB
}

// Create new instance
func NewQuotationProductRelationRepository(db *gorm.DB) QuotationProductRelationRepository {
	return &quotation_product_relationRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_product_relationRepository) Create(quotation_product_relation *QuotationProductRelation) error {
	return r.db.Create(quotation_product_relation).Error
}

// Function to get single instance of quotation_product_relation
func (r *quotation_product_relationRepository) GetSingle(id int) (*QuotationProductRelation, error) {
	var quotation_product_relation QuotationProductRelation
	err := r.db.First(&quotation_product_relation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_product_relation, nil
}

// Function to get all instances of quotation_product_relation
func (r *quotation_product_relationRepository) GetAll() ([]QuotationProductRelation, error) {
	var quotation_product_relation []QuotationProductRelation
	err := r.db.Find(&quotation_product_relation).Error
	return quotation_product_relation, err
}

// Function to update existing instances of quotation_product_relation
func (r *quotation_product_relationRepository) Update(quotation_product_relation *QuotationProductRelation) error {
	return r.db.Save(quotation_product_relation).Error
}

// Function to delete single instances of quotation_product_relation
func (r *quotation_product_relationRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationProductRelation{}, id)
	return result.Error
}
