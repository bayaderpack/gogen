package models

import (
	"time"
	"errors"
	"gorm.io/gorm"
)


// product_special is struct that represent data in Database
type ProductSpecial struct {
	gorm.Model
	ProductID int `gorm:"column:product_id;default:" json:"product_id"`
Price float32 `gorm:"column:price;default:0.00" json:"price"`
Deadline *time.Time `gorm:"column:deadline;default:" json:"deadline"`
Status bool `gorm:"column:status;default:1" json:"status"`

}

// ProductSpecial is interface that that model needs to implement
type ProductSpecialRepository interface {
	Create(product_special *ProductSpecial) error
	GetSingle(id int) (*ProductSpecial, error)
	GetAll() ([]ProductSpecial, error)
	Update(product_special *ProductSpecial) error
	Delete(id int) error
}

type product_specialRepository struct {
	db *gorm.DB
}


// Create new instance
func NewProductSpecialRepository(db *gorm.DB) ProductSpecialRepository {
	return &product_specialRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *product_specialRepository) Create(product_special *ProductSpecial) error {
	return r.db.Create(product_special).Error
}


//Function to get single instance of product_special 
func (r *product_specialRepository) GetSingle(id int) (*ProductSpecial, error) {
	var product_special ProductSpecial
	err := r.db.First(&product_special, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product_special, nil
}


//Function to get all instances of product_special 
func (r *product_specialRepository) GetAll() ([]ProductSpecial, error) {
	var product_special []ProductSpecial
	err := r.db.Find(&product_special).Error
	return product_special, err
}

//Function to update existing instances of product_special 
func (r *product_specialRepository) Update(product_special *ProductSpecial) error {
	return r.db.Save(product_special).Error
}

//Function to delete single instances of product_special 
func (r *product_specialRepository) Delete(id int) error {
	result := r.db.Delete(&ProductSpecial{}, id)
	return result.Error
}