package models

import (
	"errors"

	"gorm.io/gorm"
)

// language is struct that represent data in Database
type Language struct {
	gorm.Model
	LanguageID int    `gorm:"column:language_id;default:" json:"language_id"`
	Language   string `gorm:"column:language;default:" json:"language"`
	Code       string `gorm:"column:code;default:" json:"code"`
	Text       string `gorm:"column:text;default:NULL" json:"text"`
	IsPrimary  int    `gorm:"column:is_primary;default:" json:"is_primary"`
	Status     int    `gorm:"column:status;default:" json:"status"`
}

// Language is interface that that model needs to implement
type LanguageRepository interface {
	Create(language *Language) error
	GetSingle(id int) (*Language, error)
	GetAll() ([]Language, error)
	Update(language *Language) error
	Delete(id int) error
}

type languageRepository struct {
	db *gorm.DB
}

// Create new instance
func NewLanguageRepository(db *gorm.DB) LanguageRepository {
	return &languageRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *languageRepository) Create(language *Language) error {
	return r.db.Create(language).Error
}

// Function to get single instance of language
func (r *languageRepository) GetSingle(id int) (*Language, error) {
	var language Language
	err := r.db.First(&language, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &language, nil
}

// Function to get all instances of language
func (r *languageRepository) GetAll() ([]Language, error) {
	var language []Language
	err := r.db.Find(&language).Error
	return language, err
}

// Function to update existing instances of language
func (r *languageRepository) Update(language *Language) error {
	return r.db.Save(language).Error
}

// Function to delete single instances of language
func (r *languageRepository) Delete(id int) error {
	result := r.db.Delete(&Language{}, id)
	return result.Error
}
