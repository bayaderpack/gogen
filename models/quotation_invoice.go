package models

import (
	"errors"

	"gorm.io/gorm"
)

// quotation_invoice is struct that represent data in Database
type QuotationInvoice struct {
	gorm.Model
	InvoiceID       int     `gorm:"column:invoice_id;default:" json:"invoice_id"`
	QuotationID     int     `gorm:"column:quotation_id;default:" json:"quotation_id"`
	Discount        float32 `gorm:"column:discount;default:" json:"discount"`
	Final           float32 `gorm:"column:final;default:" json:"final"`
	UpfrontPayment  int     `gorm:"column:upfront_payment;default:" json:"upfront_payment"`
	QuantityDiff    int     `gorm:"column:quantity_diff;default:" json:"quantity_diff"`
	DeliveryDays    int     `gorm:"column:delivery_days;default:" json:"delivery_days"`
	ValidDays       int     `gorm:"column:valid_days;default:" json:"valid_days"`
	AdminComment    string  `gorm:"column:admin_comment;default:" json:"admin_comment"`
	CustomerComment string  `gorm:"column:customer_comment;default:NULL" json:"customer_comment"`
	CustomerVisible bool    `gorm:"column:customer_visible;default:0" json:"customer_visible"`
	CustomerTotals  bool    `gorm:"column:customer_totals;default:1" json:"customer_totals"`
	CreatedBy       int     `gorm:"column:created_by;default:" json:"created_by"`
	ModifiedBy      int     `gorm:"column:modified_by;default:0" json:"modified_by"`
}

// QuotationInvoice is interface that that model needs to implement
type QuotationInvoiceRepository interface {
	Create(quotation_invoice *QuotationInvoice) error
	GetSingle(id int) (*QuotationInvoice, error)
	GetAll() ([]QuotationInvoice, error)
	Update(quotation_invoice *QuotationInvoice) error
	Delete(id int) error
}

type quotation_invoiceRepository struct {
	db *gorm.DB
}

// Create new instance
func NewQuotationInvoiceRepository(db *gorm.DB) QuotationInvoiceRepository {
	return &quotation_invoiceRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_invoiceRepository) Create(quotation_invoice *QuotationInvoice) error {
	return r.db.Create(quotation_invoice).Error
}

// Function to get single instance of quotation_invoice
func (r *quotation_invoiceRepository) GetSingle(id int) (*QuotationInvoice, error) {
	var quotation_invoice QuotationInvoice
	err := r.db.First(&quotation_invoice, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_invoice, nil
}

// Function to get all instances of quotation_invoice
func (r *quotation_invoiceRepository) GetAll() ([]QuotationInvoice, error) {
	var quotation_invoice []QuotationInvoice
	err := r.db.Find(&quotation_invoice).Error
	return quotation_invoice, err
}

// Function to update existing instances of quotation_invoice
func (r *quotation_invoiceRepository) Update(quotation_invoice *QuotationInvoice) error {
	return r.db.Save(quotation_invoice).Error
}

// Function to delete single instances of quotation_invoice
func (r *quotation_invoiceRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationInvoice{}, id)
	return result.Error
}
