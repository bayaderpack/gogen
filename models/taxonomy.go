package models

import (
	"errors"

	"gorm.io/gorm"
)

// taxonomy is struct that represent data in Database
type Taxonomy struct {
	gorm.Model
	TaxonomyID int         `gorm:"column:taxonomy_id;default:" json:"taxonomy_id"`
	ParentID   int         `gorm:"column:parent_id;default:0" json:"parent_id"`
	MediaID    int         `gorm:"column:media_id;default:0" json:"media_id"`
	SeoMediaID int         `gorm:"column:seo_media_id;default:0" json:"seo_media_id"`
	BannerID   int         `gorm:"column:banner_id;default:0" json:"banner_id"`
	Type       ProductEnum `gorm:"column:type;default:" json:"type"`
	Slug       string      `gorm:"column:slug;default:" json:"slug"`
	Filterable bool        `gorm:"column:filterable;default:1" json:"filterable"`
	SortOrder  int         `gorm:"column:sort_order;default:0" json:"sort_order"`
	Status     bool        `gorm:"column:status;default:1" json:"status"`
}

// Taxonomy is interface that that model needs to implement
type TaxonomyRepository interface {
	Create(taxonomy *Taxonomy) error
	GetSingle(id int) (*Taxonomy, error)
	GetAll() ([]Taxonomy, error)
	Update(taxonomy *Taxonomy) error
	Delete(id int) error
}

type taxonomyRepository struct {
	db *gorm.DB
}

// Create new instance
func NewTaxonomyRepository(db *gorm.DB) TaxonomyRepository {
	return &taxonomyRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *taxonomyRepository) Create(taxonomy *Taxonomy) error {
	return r.db.Create(taxonomy).Error
}

// Function to get single instance of taxonomy
func (r *taxonomyRepository) GetSingle(id int) (*Taxonomy, error) {
	var taxonomy Taxonomy
	err := r.db.First(&taxonomy, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &taxonomy, nil
}

// Function to get all instances of taxonomy
func (r *taxonomyRepository) GetAll() ([]Taxonomy, error) {
	var taxonomy []Taxonomy
	err := r.db.Find(&taxonomy).Error
	return taxonomy, err
}

// Function to update existing instances of taxonomy
func (r *taxonomyRepository) Update(taxonomy *Taxonomy) error {
	return r.db.Save(taxonomy).Error
}

// Function to delete single instances of taxonomy
func (r *taxonomyRepository) Delete(id int) error {
	result := r.db.Delete(&Taxonomy{}, id)
	return result.Error
}
