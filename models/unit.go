package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// unit is struct that represent data in Database
type Unit struct {
	gorm.Model
	UnitID int `gorm:"column:unit_id;default:" json:"unit_id"`
SingleFactor int `gorm:"column:single_factor;default:1" json:"single_factor"`
Status bool `gorm:"column:status;default:1" json:"status"`

}

// Unit is interface that that model needs to implement
type UnitRepository interface {
	Create(unit *Unit) error
	GetSingle(id int) (*Unit, error)
	GetAll() ([]Unit, error)
	Update(unit *Unit) error
	Delete(id int) error
}

type unitRepository struct {
	db *gorm.DB
}


// Create new instance
func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *unitRepository) Create(unit *Unit) error {
	return r.db.Create(unit).Error
}


//Function to get single instance of unit 
func (r *unitRepository) GetSingle(id int) (*Unit, error) {
	var unit Unit
	err := r.db.First(&unit, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &unit, nil
}


//Function to get all instances of unit 
func (r *unitRepository) GetAll() ([]Unit, error) {
	var unit []Unit
	err := r.db.Find(&unit).Error
	return unit, err
}

//Function to update existing instances of unit 
func (r *unitRepository) Update(unit *Unit) error {
	return r.db.Save(unit).Error
}

//Function to delete single instances of unit 
func (r *unitRepository) Delete(id int) error {
	result := r.db.Delete(&Unit{}, id)
	return result.Error
}