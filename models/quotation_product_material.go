package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// quotation_product_material is struct that represent data in Database
type QuotationProductMaterial struct {
	gorm.Model
	MaterialID int `gorm:"column:material_id;default:" json:"material_id"`
QuotationProductID int `gorm:"column:quotation_product_id;default:" json:"quotation_product_id"`
Title string `gorm:"column:title;default:NULL" json:"title"`
PerSheet float32 `gorm:"column:per_sheet;default:1.0" json:"per_sheet"`
TotalSheet int `gorm:"column:total_sheet;default:0" json:"total_sheet"`
Waste int `gorm:"column:waste;default:0" json:"waste"`
SheetSize string `gorm:"column:sheet_size;default:NULL" json:"sheet_size"`
SheetPrice float32 `gorm:"column:sheet_price;default:1.00" json:"sheet_price"`
SheetInfo string `gorm:"column:sheet_info;default:NULL" json:"sheet_info"`
PrintMachineID int `gorm:"column:print_machine_id;default:0" json:"print_machine_id"`
PrintID int `gorm:"column:print_id;default:0" json:"print_id"`
LaminationMachineID int `gorm:"column:lamination_machine_id;default:0" json:"lamination_machine_id"`
LaminationID int `gorm:"column:lamination_id;default:0" json:"lamination_id"`
LaminationTypeID int `gorm:"column:lamination_type_id;default:0" json:"lamination_type_id"`
LaminationWidth string `gorm:"column:lamination_width;default:NULL" json:"lamination_width"`
Plate float32 `gorm:"column:plate;default:0.00" json:"plate"`
MergingMachineID int `gorm:"column:merging_machine_id;default:0" json:"merging_machine_id"`
MergingFaces int `gorm:"column:merging_faces;default:0" json:"merging_faces"`
CreasingMachineID int `gorm:"column:creasing_machine_id;default:0" json:"creasing_machine_id"`
CreasingID int `gorm:"column:creasing_id;default:0" json:"creasing_id"`
Foil bool `gorm:"column:foil;default:0" json:"foil"`
FoilWidth float32 `gorm:"column:foil_width;default:0.00" json:"foil_width"`
FoilLength float32 `gorm:"column:foil_length;default:0.00" json:"foil_length"`
Spotuv bool `gorm:"column:spotuv;default:0" json:"spotuv"`
StickID int `gorm:"column:stick_id;default:0" json:"stick_id"`
StickPoints int `gorm:"column:stick_points;default:0" json:"stick_points"`
BagstickID int `gorm:"column:bagstick_id;default:0" json:"bagstick_id"`
BagstickType string `gorm:"column:bagstick_type;default:NULL" json:"bagstick_type"`
BagstickBasesheet int `gorm:"column:bagstick_basesheet;default:0" json:"bagstick_basesheet"`
BindingID int `gorm:"column:binding_id;default:0" json:"binding_id"`
BindingPieces int `gorm:"column:binding_pieces;default:0" json:"binding_pieces"`
FoamPerSheet float32 `gorm:"column:foam_per_sheet;default:0.0" json:"foam_per_sheet"`
FoamSheetPrice float32 `gorm:"column:foam_sheet_price;default:0.00" json:"foam_sheet_price"`
UnitPrice float32 `gorm:"column:unit_price;default:0.0000" json:"unit_price"`
Total float32 `gorm:"column:total;default:0.0000" json:"total"`

}

// QuotationProductMaterial is interface that that model needs to implement
type QuotationProductMaterialRepository interface {
	Create(quotation_product_material *QuotationProductMaterial) error
	GetSingle(id int) (*QuotationProductMaterial, error)
	GetAll() ([]QuotationProductMaterial, error)
	Update(quotation_product_material *QuotationProductMaterial) error
	Delete(id int) error
}

type quotation_product_materialRepository struct {
	db *gorm.DB
}


// Create new instance
func NewQuotationProductMaterialRepository(db *gorm.DB) QuotationProductMaterialRepository {
	return &quotation_product_materialRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *quotation_product_materialRepository) Create(quotation_product_material *QuotationProductMaterial) error {
	return r.db.Create(quotation_product_material).Error
}


//Function to get single instance of quotation_product_material 
func (r *quotation_product_materialRepository) GetSingle(id int) (*QuotationProductMaterial, error) {
	var quotation_product_material QuotationProductMaterial
	err := r.db.First(&quotation_product_material, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &quotation_product_material, nil
}


//Function to get all instances of quotation_product_material 
func (r *quotation_product_materialRepository) GetAll() ([]QuotationProductMaterial, error) {
	var quotation_product_material []QuotationProductMaterial
	err := r.db.Find(&quotation_product_material).Error
	return quotation_product_material, err
}

//Function to update existing instances of quotation_product_material 
func (r *quotation_product_materialRepository) Update(quotation_product_material *QuotationProductMaterial) error {
	return r.db.Save(quotation_product_material).Error
}

//Function to delete single instances of quotation_product_material 
func (r *quotation_product_materialRepository) Delete(id int) error {
	result := r.db.Delete(&QuotationProductMaterial{}, id)
	return result.Error
}