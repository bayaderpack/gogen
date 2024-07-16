package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// product is struct that represent data in Database
type Product struct {
	gorm.Model
	ProductID   int        `gorm:"column:product_id;default:" json:"product_id"`
	Sku         string     `gorm:"column:sku;default:" json:"sku"`
	ItemCode    string     `gorm:"column:item_code;default:" json:"item_code"`
	Quantity    int        `gorm:"column:quantity;default:0" json:"quantity"`
	PerCarton   int        `gorm:"column:per_carton;default:" json:"per_carton"`
	Price       float32    `gorm:"column:price;default:0.0000" json:"price"`
	Points      int        `gorm:"column:points;default:0" json:"points"`
	TaxID       int        `gorm:"column:tax_id;default:" json:"tax_id"`
	Weight      float32    `gorm:"column:weight;default:0.00" json:"weight"`
	Subtract    int        `gorm:"column:subtract;default:1" json:"subtract"`
	Minimum     int        `gorm:"column:minimum;default:1" json:"minimum"`
	Maximum     int        `gorm:"column:maximum;default:0" json:"maximum"`
	FixedRange  int        `gorm:"column:fixed_range;default:0" json:"fixed_range"`
	AvailableAt *time.Time `gorm:"column:available_at;default:current_timestamp()" json:"available_at"`
	Status      bool       `gorm:"column:status;default:1" json:"status"`
	Sold        int        `gorm:"column:sold;default:0" json:"sold"`
	Slug        string     `gorm:"column:slug;default:" json:"slug"`
	Recyclable  bool       `gorm:"column:recyclable;default:0" json:"recyclable"`
	Sustainable bool       `gorm:"column:sustainable;default:0" json:"sustainable"`
	Compostable bool       `gorm:"column:compostable;default:0" json:"compostable"`
	Dae         string     `gorm:"column:dae;default:NULL" json:"dae"`
}

// Product is interface that that model needs to implement
type ProductRepository interface {
	Create(product *Product) error
	GetSingle(id int) (*Product, error)
	GetAll() ([]Product, error)
	Update(product *Product) error
	Delete(id int) error
}

type productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *productRepository) Create(product *Product) error {
	return r.db.Create(product).Error
}

// Function to get single instance of product
func (r *productRepository) GetSingle(id int) (*Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &product, nil
}

// Function to get all instances of product
func (r *productRepository) GetAll() ([]Product, error) {
	var product []Product
	err := r.db.Find(&product).Error
	return product, err
}

// Function to update existing instances of product
func (r *productRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

// Function to delete single instances of product
func (r *productRepository) Delete(id int) error {
	result := r.db.Delete(&Product{}, id)
	return result.Error
}
