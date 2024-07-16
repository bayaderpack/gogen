package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// quotation_products is struct that represent data in Database
type QuotationProducts struct {
	// gorm.Model
	QuotationProductID       int                        `gorm:"column:quotation_product_id;default:" json:"quotation_product_id"`
	ProductID                int                        `gorm:"column:product_id;default:0" json:"product_id"`
	QuotationID              int                        `gorm:"column:quotation_id;default:" json:"quotation_id"`
	Quotation                Quotation                  `gorm:"foreignKey:QuotationProductID;references:QuotationID"`
	QuotationProductsOptions []QuotationProductsOptions `json:"quotation_products_options" gorm:"-"`
	// QuotationProductsOptions QuotationProductsOptions `gorm:"foreignKey:QuotationProductID;references:QuotationProductsOptionsID"`
	Title               string    `gorm:"column:title;default:NULL" json:"title"`
	Length              float32   `gorm:"column:length;default:0.00" json:"length"`
	Width               float32   `gorm:"column:width;default:0.00" json:"width"`
	Height              float32   `gorm:"column:height;default:0.00" json:"height"`
	Quantity            int       `gorm:"column:quantity;default:" json:"quantity"`
	DieValue            float32   `gorm:"column:die_value;default:0.00" json:"die_value"`
	DieCost             float32   `gorm:"column:die_cost;default:0.00" json:"die_cost"`
	MagneticValue       int       `gorm:"column:magnetic_value;default:0" json:"magnetic_value"`
	MagneticCost        float32   `gorm:"column:magnetic_cost;default:0.00" json:"magnetic_cost"`
	Filling             float32   `gorm:"column:filling;default:0.00" json:"filling"`
	Miscellaneous       float32   `gorm:"column:miscellaneous;default:0.00" json:"miscellaneous"`
	MiscellaneousNotice string    `gorm:"column:miscellaneous_notice;default:NULL" json:"miscellaneous_notice"`
	RopeValue           string    `gorm:"column:rope_value;default:NULL" json:"rope_value"`
	RopeCost            float32   `gorm:"column:rope_cost;default:0.00" json:"rope_cost"`
	DesignerNotice      string    `gorm:"column:designer_notice;default:NULL" json:"designer_notice"`
	DesignerFiles       string    `gorm:"column:designer_files;default:NULL" json:"designer_files"`
	Comment             string    `gorm:"column:comment;default:NULL" json:"comment"`
	Operational         float32   `gorm:"column:operational;default:0.00" json:"operational"`
	Revenue             float32   `gorm:"column:revenue;default:0.00" json:"revenue"`
	Discount            float32   `gorm:"column:discount;default:0.00" json:"discount"`
	DiscountType        TypeValue `gorm:"column:discount_type;default:NULL" json:"discount_type"`
	RevenueType         TypeValue `gorm:"column:revenue_type;default:NULL" json:"revenue_type"`
	UnitPriceCommission int       `gorm:"column:unit_price_commission;default:NULL" json:"unit_price_commission"`
	CommissionSalesman  string    `gorm:"column:commission_salesman;default:NULL" json:"commission_salesman"`
	UnitPrice           float32   `gorm:"column:unit_price;default:0.0000" json:"unit_price"`
	Gross               float32   `gorm:"column:gross;default:0.0000" json:"gross"`
	Net                 float32   `gorm:"column:net;default:0.0000" json:"net"`
	SortOrder           int       `gorm:"column:sort_order;default:0" json:"sort_order"`
	Version             int       `gorm:"column:version;default:0" json:"version"`
}

// QuotationProducts is interface that that model needs to implement
type QuotationProductsRepository interface {
	Create(quotation_product *QuotationProducts) error
	GetSingle(id int) (*QuotationProducts, error)
	GetAll() ([]QuotationProducts, error)
	Update(quotation_product *QuotationProducts) error
	Delete(id int) error
}

type quotation_productRepository struct {
	db *gorm.DB
}

// Create new instance
func NewQuotationProductsRepository(db *gorm.DB) QuotationProductsRepository {
	return &quotation_productRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_productRepository) Create(quotation_product *QuotationProducts) error {
	return r.db.Create(quotation_product).Error
}

// Function to get single instance of quotation_product
func (r *quotation_productRepository) GetSingle(id int) (*QuotationProducts, error) {
	var quotation_product QuotationProducts
	var quotation_product_options []QuotationProductsOptions
	// .Where("quotation_product_id = ?", id).Association("Languages")
	// err := r.db.Joins("INNER JOIN quotation_products_options ON quotation_products_options.quotation_product_id = quotation_products.quotation_product_id").Where("quotation_products.quotation_product_id = ? ", id).Preload("Quotation").First(&quotation_product).Error
	err := r.db.Where("quotation_product_id = ?", id).Find(&quotation_product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // Record not found
		}
		return nil, err
	}
	err = r.db.Where("quotation_product_id = ?", id).Find(&quotation_product_options).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // Record not found
		}
		return nil, err
	}
	quotation_product.QuotationProductsOptions = quotation_product_options
	fmt.Printf("%+v", quotation_product)
	return &quotation_product, nil
}

// Function to get all instances of quotation_products
func (r *quotation_productRepository) GetAll() ([]QuotationProducts, error) {
	var quotation_products []QuotationProducts
	err := r.db.Find(&quotation_products).Error
	return quotation_products, err
}

// Function to update existing instances of quotation_product
func (r *quotation_productRepository) Update(quotation_product *QuotationProducts) error {
	return r.db.Save(quotation_product).Error
}

// Function to delete single instances of quotation_product
func (r *quotation_productRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationProducts{}, id)
	return result.Error
}
