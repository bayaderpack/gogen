package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// customer_group is struct that represent data in Database
type CustomerGroup struct {
	gorm.Model
	GroupID int `gorm:"column:group_id;default:" json:"group_id"`
Status bool `gorm:"column:status;default:" json:"status"`
AuthSetting string `gorm:"column:auth_setting;default:NULL" json:"auth_setting"`
JobSetting string `gorm:"column:job_setting;default:NULL" json:"job_setting"`

}

// CustomerGroup is interface that that model needs to implement
type CustomerGroupRepository interface {
	Create(customer_group *CustomerGroup) error
	GetSingle(id int) (*CustomerGroup, error)
	GetAll() ([]CustomerGroup, error)
	Update(customer_group *CustomerGroup) error
	Delete(id int) error
}

type customer_groupRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCustomerGroupRepository(db *gorm.DB) CustomerGroupRepository {
	return &customer_groupRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *customer_groupRepository) Create(customer_group *CustomerGroup) error {
	return r.db.Create(customer_group).Error
}


//Function to get single instance of customer_group 
func (r *customer_groupRepository) GetSingle(id int) (*CustomerGroup, error) {
	var customer_group CustomerGroup
	err := r.db.First(&customer_group, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &customer_group, nil
}


//Function to get all instances of customer_group 
func (r *customer_groupRepository) GetAll() ([]CustomerGroup, error) {
	var customer_group []CustomerGroup
	err := r.db.Find(&customer_group).Error
	return customer_group, err
}

//Function to update existing instances of customer_group 
func (r *customer_groupRepository) Update(customer_group *CustomerGroup) error {
	return r.db.Save(customer_group).Error
}

//Function to delete single instances of customer_group 
func (r *customer_groupRepository) Delete(id int) error {
	result := r.db.Delete(&CustomerGroup{}, id)
	return result.Error
}