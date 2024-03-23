package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// quotation_product_material_sheets is struct that represent data in Database
type QuotationProductMaterialSheets struct {
	gorm.Model
	MaterialID int `gorm:"column:material_id;default:" json:"material_id"`
SheetID int `gorm:"column:sheet_id;default:" json:"sheet_id"`

}

// QuotationProductMaterialSheets is interface that that model needs to implement
type QuotationProductMaterialSheetsRepository interface {
	Create(quotation_product_material_sheet *QuotationProductMaterialSheets) error
	GetSingle(id int) (*QuotationProductMaterialSheets, error)
	GetAll() ([]QuotationProductMaterialSheets, error)
	Update(quotation_product_material_sheet *QuotationProductMaterialSheets) error
	Delete(id int) error
}

type quotation_product_material_sheetRepository struct {
	db *gorm.DB
}


// Create new instance
func NewQuotationProductMaterialSheetsRepository(db *gorm.DB) QuotationProductMaterialSheetsRepository {
	return &quotation_product_material_sheetRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_product_material_sheetRepository) Create(quotation_product_material_sheet *QuotationProductMaterialSheets) error {
	return r.db.Create(quotation_product_material_sheet).Error
}


//Function to get single instance of quotation_product_material_sheet 
func (r *quotation_product_material_sheetRepository) GetSingle(id int) (*QuotationProductMaterialSheets, error) {
	var quotation_product_material_sheet QuotationProductMaterialSheets
	err := r.db.First(&quotation_product_material_sheet, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_product_material_sheet, nil
}


//Function to get all instances of quotation_product_material_sheets 
func (r *quotation_product_material_sheetRepository) GetAll() ([]QuotationProductMaterialSheets, error) {
	var quotation_product_material_sheets []QuotationProductMaterialSheets
	err := r.db.Find(&quotation_product_material_sheets).Error
	return quotation_product_material_sheets, err
}

//Function to update existing instances of quotation_product_material_sheet 
func (r *quotation_product_material_sheetRepository) Update(quotation_product_material_sheet *QuotationProductMaterialSheets) error {
	return r.db.Save(quotation_product_material_sheet).Error
}

//Function to delete single instances of quotation_product_material_sheet 
func (r *quotation_product_material_sheetRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationProductMaterialSheets{}, id)
	return result.Error
}