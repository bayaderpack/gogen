package models

import (
	"errors"

	"gorm.io/gorm"
)

// print is struct that represent data in Database
type Print struct {
	gorm.Model
	PrintID   int     `gorm:"column:print_id;default:" json:"print_id"`
	Qmf       float32 `gorm:"column:qmf;default:" json:"qmf"`
	Rmf       float32 `gorm:"column:rmf;default:" json:"rmf"`
	Raf       string  `gorm:"column:raf;default:NULL" json:"raf"`
	Status    bool    `gorm:"column:status;default:1" json:"status"`
	SortOrder int     `gorm:"column:sort_order;default:0" json:"sort_order"`
}

// Print is interface that that model needs to implement
type PrintRepository interface {
	Create(print *Print) error
	GetSingle(id int) (*Print, error)
	GetAll() ([]Print, error)
	Update(print *Print) error
	Delete(id int) error
}

type printRepository struct {
	db *gorm.DB
}

// Create new instance
func NewPrintRepository(db *gorm.DB) PrintRepository {
	return &printRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *printRepository) Create(print *Print) error {
	return r.db.Create(print).Error
}

// Function to get single instance of print
func (r *printRepository) GetSingle(id int) (*Print, error) {
	var print Print
	err := r.db.First(&print, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &print, nil
}

// Function to get all instances of print
func (r *printRepository) GetAll() ([]Print, error) {
	var print []Print
	err := r.db.Find(&print).Error
	return print, err
}

// Function to update existing instances of print
func (r *printRepository) Update(print *Print) error {
	return r.db.Save(print).Error
}

// Function to delete single instances of print
func (r *printRepository) Delete(id int) error {
	result := r.db.Delete(&Print{}, id)
	return result.Error
}
