package models

import (
	"errors"

	"gorm.io/gorm"
)

// sheet is struct that represent data in Database
type Sheet struct {
	gorm.Model
	SheetID       int     `gorm:"column:sheet_id;default:" json:"sheet_id"`
	ParentID      int     `gorm:"column:parent_id;default:0" json:"parent_id"`
	Weight        float32 `gorm:"column:weight;default:0.0000" json:"weight"`
	WeightUnit    string  `gorm:"column:weight_unit;default:NULL" json:"weight_unit"`
	Thickness     float32 `gorm:"column:thickness;default:0.0000" json:"thickness"`
	ThicknessUnit string  `gorm:"column:thickness_unit;default:NULL" json:"thickness_unit"`
	Price         float32 `gorm:"column:price;default:1.0000" json:"price"`
	MediaID       int     `gorm:"column:media_id;default:0" json:"media_id"`
	Status        bool    `gorm:"column:status;default:1" json:"status"`
	SortOrder     int     `gorm:"column:sort_order;default:" json:"sort_order"`
}

// Sheet is interface that that model needs to implement
type SheetRepository interface {
	Create(sheet *Sheet) error
	GetSingle(id int) (*Sheet, error)
	GetAll() ([]Sheet, error)
	Update(sheet *Sheet) error
	Delete(id int) error
}

type sheetRepository struct {
	db *gorm.DB
}

// Create new instance
func NewSheetRepository(db *gorm.DB) SheetRepository {
	return &sheetRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *sheetRepository) Create(sheet *Sheet) error {
	return r.db.Create(sheet).Error
}

// Function to get single instance of sheet
func (r *sheetRepository) GetSingle(id int) (*Sheet, error) {
	var sheet Sheet
	err := r.db.First(&sheet, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &sheet, nil
}

// Function to get all instances of sheet
func (r *sheetRepository) GetAll() ([]Sheet, error) {
	var sheet []Sheet
	err := r.db.Find(&sheet).Error
	return sheet, err
}

// Function to update existing instances of sheet
func (r *sheetRepository) Update(sheet *Sheet) error {
	return r.db.Save(sheet).Error
}

// Function to delete single instances of sheet
func (r *sheetRepository) Delete(id int) error {
	result := r.db.Delete(&Sheet{}, id)
	return result.Error
}
