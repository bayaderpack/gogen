package models

import (
	"time"
	"errors"
	"gorm.io/gorm"
)


// admin is struct that represent data in Database
type Admin struct {
	gorm.Model
	AdminID int `gorm:"column:admin_id;default:" json:"admin_id"`
Kid string `gorm:"column:kid;default:" json:"kid"`
Username string `gorm:"column:username;default:" json:"username"`
Firstname string `gorm:"column:firstname;default:" json:"firstname"`
Lastname string `gorm:"column:lastname;default:" json:"lastname"`
Email string `gorm:"column:email;default:" json:"email"`
CountryCode int `gorm:"column:country_code;default:" json:"country_code"`
Mobile int `gorm:"column:mobile;default:" json:"mobile"`
Password string `gorm:"column:password;default:" json:"password"`
Ip string `gorm:"column:ip;default:'0'" json:"ip"`
Lng string `gorm:"column:lng;default:'ar'" json:"lng"`
Status bool `gorm:"column:status;default:0" json:"status"`
ResetToken string `gorm:"column:reset_token;default:NULL" json:"reset_token"`
Otp int `gorm:"column:otp;default:NULL" json:"otp"`
ResetExpires *time.Time `gorm:"column:reset_expires;default:NULL" json:"reset_expires"`

}

// Admin is interface that that model needs to implement
type AdminRepository interface {
	Create(admin *Admin) error
	GetSingle(id int) (*Admin, error)
	GetAll() ([]Admin, error)
	Update(admin *Admin) error
	Delete(id int) error
}

type adminRepository struct {
	db *gorm.DB
}


// Create new instance
func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *adminRepository) Create(admin *Admin) error {
	return r.db.Create(admin).Error
}


//Function to get single instance of admin 
func (r *adminRepository) GetSingle(id int) (*Admin, error) {
	var admin Admin
	err := r.db.First(&admin, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &admin, nil
}


//Function to get all instances of admin 
func (r *adminRepository) GetAll() ([]Admin, error) {
	var admin []Admin
	err := r.db.Find(&admin).Error
	return admin, err
}

//Function to update existing instances of admin 
func (r *adminRepository) Update(admin *Admin) error {
	return r.db.Save(admin).Error
}

//Function to delete single instances of admin 
func (r *adminRepository) Delete(id int) error {
	result := r.db.Delete(&Admin{}, id)
	return result.Error
}