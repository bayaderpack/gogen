package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// quotation_assignee is struct that represent data in Database
type QuotationAssignee struct {
	gorm.Model
	QuotationID int `gorm:"column:quotation_id;default:" json:"quotation_id"`
AdminID int `gorm:"column:admin_id;default:" json:"admin_id"`

}

// QuotationAssignee is interface that that model needs to implement
type QuotationAssigneeRepository interface {
	Create(quotation_assignee *QuotationAssignee) error
	GetSingle(id int) (*QuotationAssignee, error)
	GetAll() ([]QuotationAssignee, error)
	Update(quotation_assignee *QuotationAssignee) error
	Delete(id int) error
}

type quotation_assigneeRepository struct {
	db *gorm.DB
}


// Create new instance
func NewQuotationAssigneeRepository(db *gorm.DB) QuotationAssigneeRepository {
	return &quotation_assigneeRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_assigneeRepository) Create(quotation_assignee *QuotationAssignee) error {
	return r.db.Create(quotation_assignee).Error
}


//Function to get single instance of quotation_assignee 
func (r *quotation_assigneeRepository) GetSingle(id int) (*QuotationAssignee, error) {
	var quotation_assignee QuotationAssignee
	err := r.db.First(&quotation_assignee, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_assignee, nil
}


//Function to get all instances of quotation_assignee 
func (r *quotation_assigneeRepository) GetAll() ([]QuotationAssignee, error) {
	var quotation_assignee []QuotationAssignee
	err := r.db.Find(&quotation_assignee).Error
	return quotation_assignee, err
}

//Function to update existing instances of quotation_assignee 
func (r *quotation_assigneeRepository) Update(quotation_assignee *QuotationAssignee) error {
	return r.db.Save(quotation_assignee).Error
}

//Function to delete single instances of quotation_assignee 
func (r *quotation_assigneeRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationAssignee{}, id)
	return result.Error
}