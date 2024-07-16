package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// job_order_products is struct that represent data in Database
type JobOrderProducts struct {
	gorm.Model
	JobOrderID int `gorm:"column:job_order_id;default:" json:"job_order_id"`
QuotationProductID int `gorm:"column:quotation_product_id;default:0" json:"quotation_product_id"`
EditedQuotationProductID int `gorm:"column:edited_quotation_product_id;default:NULL" json:"edited_quotation_product_id"`
Samples int `gorm:"column:samples;default:0" json:"samples"`
Printed bool `gorm:"column:printed;default:0" json:"printed"`
Priority bool `gorm:"column:priority;default:0" json:"priority"`
Design bool `gorm:"column:design;default:0" json:"design"`
Die bool `gorm:"column:die;default:0" json:"die"`
Note string `gorm:"column:note;default:NULL" json:"note"`
Name string `gorm:"column:name;default:NULL" json:"name"`
Quantity int `gorm:"column:quantity;default:0" json:"quantity"`
Price float32 `gorm:"column:price;default:0.0000" json:"price"`

}

// JobOrderProducts is interface that that model needs to implement
type JobOrderProductsRepository interface {
	Create(job_order_product *JobOrderProducts) error
	GetSingle(id int) (*JobOrderProducts, error)
	GetAll() ([]JobOrderProducts, error)
	Update(job_order_product *JobOrderProducts) error
	Delete(id int) error
}

type job_order_productRepository struct {
	db *gorm.DB
}


// Create new instance
func NewJobOrderProductsRepository(db *gorm.DB) JobOrderProductsRepository {
	return &job_order_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *job_order_productRepository) Create(job_order_product *JobOrderProducts) error {
	return r.db.Create(job_order_product).Error
}


//Function to get single instance of job_order_product 
func (r *job_order_productRepository) GetSingle(id int) (*JobOrderProducts, error) {
	var job_order_product JobOrderProducts
	err := r.db.First(&job_order_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &job_order_product, nil
}


//Function to get all instances of job_order_products 
func (r *job_order_productRepository) GetAll() ([]JobOrderProducts, error) {
	var job_order_products []JobOrderProducts
	err := r.db.Find(&job_order_products).Error
	return job_order_products, err
}

//Function to update existing instances of job_order_product 
func (r *job_order_productRepository) Update(job_order_product *JobOrderProducts) error {
	return r.db.Save(job_order_product).Error
}

//Function to delete single instances of job_order_product 
func (r *job_order_productRepository) Delete(id int) error {
	result := r.db.Delete(&JobOrderProducts{}, id)
	return result.Error
}