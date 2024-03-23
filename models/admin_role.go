package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// admin_role is struct that represent data in Database
type AdminRole struct {
	gorm.Model
	AdminID int `gorm:"column:admin_id;default:" json:"admin_id"`
Role string `gorm:"column:role;default:" json:"role"`

}

// AdminRole is interface that that model needs to implement
type AdminRoleRepository interface {
	Create(admin_role *AdminRole) error
	GetSingle(id int) (*AdminRole, error)
	GetAll() ([]AdminRole, error)
	Update(admin_role *AdminRole) error
	Delete(id int) error
}

type admin_roleRepository struct {
	db *gorm.DB
}


// Create new instance
func NewAdminRoleRepository(db *gorm.DB) AdminRoleRepository {
	return &admin_roleRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *admin_roleRepository) Create(admin_role *AdminRole) error {
	return r.db.Create(admin_role).Error
}


//Function to get single instance of admin_role 
func (r *admin_roleRepository) GetSingle(id int) (*AdminRole, error) {
	var admin_role AdminRole
	err := r.db.First(&admin_role, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &admin_role, nil
}


//Function to get all instances of admin_role 
func (r *admin_roleRepository) GetAll() ([]AdminRole, error) {
	var admin_role []AdminRole
	err := r.db.Find(&admin_role).Error
	return admin_role, err
}

//Function to update existing instances of admin_role 
func (r *admin_roleRepository) Update(admin_role *AdminRole) error {
	return r.db.Save(admin_role).Error
}

//Function to delete single instances of admin_role 
func (r *admin_roleRepository) Delete(id int) error {
	result := r.db.Delete(&AdminRole{}, id)
	return result.Error
}