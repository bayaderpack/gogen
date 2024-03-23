package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// order_products is struct that represent data in Database
type OrderProducts struct {
	gorm.Model
	OrderProductID int `gorm:"column:order_product_id;default:" json:"order_product_id"`
OrderID int `gorm:"column:order_id;default:" json:"order_id"`
ProductID int `gorm:"column:product_id;default:" json:"product_id"`
Title string `gorm:"column:title;default:" json:"title"`
Quantity int `gorm:"column:quantity;default:" json:"quantity"`
Price float32 `gorm:"column:price;default:0.0000" json:"price"`
Total float32 `gorm:"column:total;default:0.0000" json:"total"`
Tax float32 `gorm:"column:tax;default:0.0000" json:"tax"`
PriceCode string `gorm:"column:price_code;default:" json:"price_code"`

}

// OrderProducts is interface that that model needs to implement
type OrderProductsRepository interface {
	Create(order_product *OrderProducts) error
	GetSingle(id int) (*OrderProducts, error)
	GetAll() ([]OrderProducts, error)
	Update(order_product *OrderProducts) error
	Delete(id int) error
}

type order_productRepository struct {
	db *gorm.DB
}


// Create new instance
func NewOrderProductsRepository(db *gorm.DB) OrderProductsRepository {
	return &order_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *order_productRepository) Create(order_product *OrderProducts) error {
	return r.db.Create(order_product).Error
}


//Function to get single instance of order_product 
func (r *order_productRepository) GetSingle(id int) (*OrderProducts, error) {
	var order_product OrderProducts
	err := r.db.First(&order_product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &order_product, nil
}


//Function to get all instances of order_products 
func (r *order_productRepository) GetAll() ([]OrderProducts, error) {
	var order_products []OrderProducts
	err := r.db.Find(&order_products).Error
	return order_products, err
}

//Function to update existing instances of order_product 
func (r *order_productRepository) Update(order_product *OrderProducts) error {
	return r.db.Save(order_product).Error
}

//Function to delete single instances of order_product 
func (r *order_productRepository) Delete(id int) error {
	result := r.db.Delete(&OrderProducts{}, id)
	return result.Error
}