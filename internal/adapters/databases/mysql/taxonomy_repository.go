package mysql

import (
	"fmt"
	"tagMicroservice/internal/adapters/databases/models"
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"

	"gorm.io/gorm"
)

type TaxonomyRepository struct {
	gormDB *gorm.DB
}

func NewTaxonomyRepository(gormDB *gorm.DB) repositories.TaxonomyRepository {
	return &TaxonomyRepository{gormDB: gormDB}
}

func (r *TaxonomyRepository) Create(taxonomy *entities.Taxonomy) error {
	gormTaxonomy := models.GormTaxonomyFromDomain(*taxonomy)
	if err := mysqlDB.db.Create(&gormTaxonomy).Error; err != nil {
		return fmt.Errorf("error creating taxonomy: %v", err)
	}
	return nil
}

func (r *TaxonomyRepository) Update(taxonomy *entities.Taxonomy) error {
	var existingGormTaxonomy models.GormTaxonomy
	if err := mysqlDB.db.First(&existingGormTaxonomy, "id = ?", taxonomy.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("taxonomy %v not found: %v", taxonomy, err)
		}
	}
	updatedGormTaxonomy := models.GormTaxonomyFromDomain(*taxonomy)
	existingGormTaxonomy.FromTagID = updatedGormTaxonomy.FromTagID
	existingGormTaxonomy.ToTagID = updatedGormTaxonomy.ToTagID
	existingGormTaxonomy.RelationshipKind = updatedGormTaxonomy.RelationshipKind
	existingGormTaxonomy.State = updatedGormTaxonomy.State

	// Save the updated entity back to the database
	if err := mysqlDB.db.Save(&existingGormTaxonomy).Error; err != nil {
		return fmt.Errorf("error updating taxonomy: %v", err)
	}
	return nil
}

func (r *TaxonomyRepository) Delete(taxonomy *entities.Taxonomy) error {
	deletedGormTaxonomy := models.GormTaxonomyFromDomain(*taxonomy)
	if err := mysqlDB.db.Delete(&deletedGormTaxonomy).Error; err != nil {
		return fmt.Errorf("error deleting taxonomy: %v", err)
	}
	return nil
}

func (r *TaxonomyRepository) FindByID(id uint) (entities.Taxonomy, error) {
	var existingGormTaxonomy models.GormTaxonomy
	if err := mysqlDB.db.First(&existingGormTaxonomy, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.Taxonomy{}, fmt.Errorf("taxonomy id %v was not found: %v", id, err)
		}
		return entities.Taxonomy{}, fmt.Errorf("failed to find taxonomy: %v", err)
	}
	existingTaxonomy := existingGormTaxonomy.ToDomain()
	return existingTaxonomy, nil
}

func (r *TaxonomyRepository) SetRelationship(taxonomyID uint, relationship string) error {
	taxonomy, err := r.FindByID(taxonomyID)
	if err != nil {
		return fmt.Errorf("failed to find taxonomy: %v", err)
	}
	taxonomy.RelationshipKind = relationship
	if err := r.Update(&taxonomy); err != nil {
		return fmt.Errorf("failed to update relationship: %v", err)
	}
	return nil
}

func (r *TaxonomyRepository) UpdateTagReferences(fromTagID, toTagID uint) error {

	if err := mysqlDB.db.Model(&models.GormTaxonomy{}).Where("from_tag_id = ?", fromTagID).Update("from_tag_id", toTagID).Error; err != nil {
		return fmt.Errorf("failed to update tag references")
	}
	if err := mysqlDB.db.Model(&models.GormTaxonomy{}).Where("to_tag_id = ?", fromTagID).Update("to_tag_id", toTagID).Error; err != nil {
		return fmt.Errorf("failed to update tag references")
	}
	return nil
}

func (r *TaxonomyRepository) FindTaxonomiesByTagID(tagID uint) ([]entities.Taxonomy, error) {
	var taxonomies []entities.Taxonomy
	var gormTaxonomies []models.GormTaxonomy

	if err := mysqlDB.db.Where("from_tag_id = ? ", tagID).Or("to_tag_id = ?", tagID).Find(&gormTaxonomies).Error; err != nil {
		return []entities.Taxonomy{}, fmt.Errorf("failed to find taxonomy by key: %v", err)
	}
	for _, gormTaxonomy := range gormTaxonomies {
		taxonomy := gormTaxonomy.ToDomain()
		taxonomies = append(taxonomies, taxonomy)
	}
	return taxonomies, nil

}
