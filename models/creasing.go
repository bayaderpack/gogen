package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// creasing is struct that represent data in Database
type Creasing struct {
	gorm.Model
	CreasingID int `gorm:"column:creasing_id;default:" json:"creasing_id"`
Title string `gorm:"column:title;default:" json:"title"`
Pricing string `gorm:"column:pricing;default:" json:"pricing"`
Status bool `gorm:"column:status;default:1" json:"status"`
SortOrder int `gorm:"column:sort_order;default:0" json:"sort_order"`

}

// Creasing is interface that that model needs to implement
type CreasingRepository interface {
	Create(creasing *Creasing) error
	GetSingle(id int) (*Creasing, error)
	GetAll() ([]Creasing, error)
	Update(creasing *Creasing) error
	Delete(id int) error
}

type creasingRepository struct {
	db *gorm.DB
}


// Create new instance
func NewCreasingRepository(db *gorm.DB) CreasingRepository {
	return &creasingRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *creasingRepository) Create(creasing *Creasing) error {
	return r.db.Create(creasing).Error
}


//Function to get single instance of creasing 
func (r *creasingRepository) GetSingle(id int) (*Creasing, error) {
	var creasing Creasing
	err := r.db.First(&creasing, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &creasing, nil
}


//Function to get all instances of creasing 
func (r *creasingRepository) GetAll() ([]Creasing, error) {
	var creasing []Creasing
	err := r.db.Find(&creasing).Error
	return creasing, err
}

//Function to update existing instances of creasing 
func (r *creasingRepository) Update(creasing *Creasing) error {
	return r.db.Save(creasing).Error
}

//Function to delete single instances of creasing 
func (r *creasingRepository) Delete(id int) error {
	result := r.db.Delete(&Creasing{}, id)
	return result.Error
}