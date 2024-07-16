package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// setting is struct that represent data in Database
type Setting struct {
	gorm.Model
	SettingID int `gorm:"column:setting_id;default:" json:"setting_id"`
Code string `gorm:"column:code;default:" json:"code"`
KeyID string `gorm:"column:key_id;default:" json:"key_id"`
Value string `gorm:"column:value;default:" json:"value"`
MediaID int `gorm:"column:media_id;default:0" json:"media_id"`

}

// Setting is interface that that model needs to implement
type SettingRepository interface {
	Create(setting *Setting) error
	GetSingle(id int) (*Setting, error)
	GetAll() ([]Setting, error)
	Update(setting *Setting) error
	Delete(id int) error
}

type settingRepository struct {
	db *gorm.DB
}


// Create new instance
func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *settingRepository) Create(setting *Setting) error {
	return r.db.Create(setting).Error
}


//Function to get single instance of setting 
func (r *settingRepository) GetSingle(id int) (*Setting, error) {
	var setting Setting
	err := r.db.First(&setting, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &setting, nil
}


//Function to get all instances of setting 
func (r *settingRepository) GetAll() ([]Setting, error) {
	var setting []Setting
	err := r.db.Find(&setting).Error
	return setting, err
}

//Function to update existing instances of setting 
func (r *settingRepository) Update(setting *Setting) error {
	return r.db.Save(setting).Error
}

//Function to delete single instances of setting 
func (r *settingRepository) Delete(id int) error {
	result := r.db.Delete(&Setting{}, id)
	return result.Error
}