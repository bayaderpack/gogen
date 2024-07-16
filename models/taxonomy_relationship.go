package models

import (
	
	"errors"
	"gorm.io/gorm"
)


// taxonomy_relationship is struct that represent data in Database
type TaxonomyRelationship struct {
	gorm.Model
	ObjectID int `gorm:"column:object_id;default:" json:"object_id"`
TaxonomyID int `gorm:"column:taxonomy_id;default:" json:"taxonomy_id"`

}

// TaxonomyRelationship is interface that that model needs to implement
type TaxonomyRelationshipRepository interface {
	Create(taxonomy_relationship *TaxonomyRelationship) error
	GetSingle(id int) (*TaxonomyRelationship, error)
	GetAll() ([]TaxonomyRelationship, error)
	Update(taxonomy_relationship *TaxonomyRelationship) error
	Delete(id int) error
}

type taxonomy_relationshipRepository struct {
	db *gorm.DB
}


// Create new instance
func NewTaxonomyRelationshipRepository(db *gorm.DB) TaxonomyRelationshipRepository {
	return &taxonomy_relationshipRepository{db: db}
}

// Function to create data in database it take model parameter
func (r *taxonomy_relationshipRepository) Create(taxonomy_relationship *TaxonomyRelationship) error {
	return r.db.Create(taxonomy_relationship).Error
}


//Function to get single instance of taxonomy_relationship 
func (r *taxonomy_relationshipRepository) GetSingle(id int) (*TaxonomyRelationship, error) {
	var taxonomy_relationship TaxonomyRelationship
	err := r.db.First(&taxonomy_relationship, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Record not found
		}
		return nil, err
	}
	return &taxonomy_relationship, nil
}


//Function to get all instances of taxonomy_relationship 
func (r *taxonomy_relationshipRepository) GetAll() ([]TaxonomyRelationship, error) {
	var taxonomy_relationship []TaxonomyRelationship
	err := r.db.Find(&taxonomy_relationship).Error
	return taxonomy_relationship, err
}

//Function to update existing instances of taxonomy_relationship 
func (r *taxonomy_relationshipRepository) Update(taxonomy_relationship *TaxonomyRelationship) error {
	return r.db.Save(taxonomy_relationship).Error
}

//Function to delete single instances of taxonomy_relationship 
func (r *taxonomy_relationshipRepository) Delete(id int) error {
	result := r.db.Delete(&TaxonomyRelationship{}, id)
	return result.Error
}