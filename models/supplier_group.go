package models

import (
	"errors"

	"gorm.io/gorm"
)

// supplier_group is struct that represent data in Database
type SupplierGroup struct {
	gorm.Model
	GroupID int  `gorm:"column:group_id;default:" json:"group_id"`
	Status  bool `gorm:"column:status;default:1" json:"status"`
}

// SupplierGroup is interface that that model needs to implement
type SupplierGroupRepository interface {
	Create(supplier_group *SupplierGroup) error
	GetSingle(id int) (*SupplierGroup, error)
	GetAll() ([]SupplierGroup, error)
	Update(supplier_group *SupplierGroup) error
	Delete(id int) error
}

type supplier_groupRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSupplierGroupRepository(db *gorm.DB) SupplierGroupRepository {
	return &supplier_groupRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *supplier_groupRepository) Create(supplier_group *SupplierGroup) error {
	return r.db.Create(supplier_group).Error
}

// Function to get single instance of supplier_group
func (r *supplier_groupRepository) GetSingle(id int) (*SupplierGroup, error) {
	var supplier_group SupplierGroup
	err := r.db.First(&supplier_group, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &supplier_group, nil
}

// Function to get all instances of supplier_group
func (r *supplier_groupRepository) GetAll() ([]SupplierGroup, error) {
	var supplier_group []SupplierGroup
	err := r.db.Find(&supplier_group).Error
	return supplier_group, err
}

// Function to update existing instances of supplier_group
func (r *supplier_groupRepository) Update(supplier_group *SupplierGroup) error {
	return r.db.Save(supplier_group).Error
}

// Function to delete single instances of supplier_group
func (r *supplier_groupRepository) Delete(id int) error {
	result := r.db.Delete(&SupplierGroup{}, id)
	return result.Error
}
