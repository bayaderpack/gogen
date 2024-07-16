package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// machine is struct that represent data in Database
type Machine struct {
	gorm.Model
	MachineID int `gorm:"column:machine_id;default:" json:"machine_id"`
Title string `gorm:"column:title;default:" json:"title"`
PrintFactors string `gorm:"column:print_factors;default:" json:"print_factors"`
SpotuvFactors string `gorm:"column:spotuv_factors;default:" json:"spotuv_factors"`
MergingFactors string `gorm:"column:merging_factors;default:" json:"merging_factors"`
LaminationBaseprice float32 `gorm:"column:lamination_baseprice;default:" json:"lamination_baseprice"`
PlatePrice float32 `gorm:"column:plate_price;default:" json:"plate_price"`
BpsFactor int `gorm:"column:bps_factor;default:" json:"bps_factor"`
Status bool `gorm:"column:status;default:1" json:"status"`
SortOrder int `gorm:"column:sort_order;default:0" json:"sort_order"`

}

// Machine is interface that that model needs to implement
type MachineRepository interface {
	Create(machine *Machine) error
	GetSingle(id int) (*Machine, error)
	GetAll() ([]Machine, error)
	Update(machine *Machine) error
	Delete(id int) error
}

type machineRepository struct {
	db *gorm.DB
}


// Create new instance
func NewMachineRepository(db *gorm.DB) MachineRepository {
	return &machineRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *machineRepository) Create(machine *Machine) error {
	return r.db.Create(machine).Error
}


//Function to get single instance of machine 
func (r *machineRepository) GetSingle(id int) (*Machine, error) {
	var machine Machine
	err := r.db.First(&machine, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &machine, nil
}


//Function to get all instances of machine 
func (r *machineRepository) GetAll() ([]Machine, error) {
	var machine []Machine
	err := r.db.Find(&machine).Error
	return machine, err
}

//Function to update existing instances of machine 
func (r *machineRepository) Update(machine *Machine) error {
	return r.db.Save(machine).Error
}

//Function to delete single instances of machine 
func (r *machineRepository) Delete(id int) error {
	result := r.db.Delete(&Machine{}, id)
	return result.Error
}