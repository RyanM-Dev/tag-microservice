package mysql

import (
	"fmt"
	"log"
	"tagMicroservice/internal/adapters/databases/models"
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"

	"gorm.io/gorm"
)

type TagRepository struct {
	gormDB *gorm.DB
}

func NewTagRepository(gormDB *gorm.DB) repositories.TagRepository {
	return &TagRepository{gormDB: gormDB}
}

func (r *TagRepository) Create(tag *entities.Tag) error {
	gormTag := models.GormTagFromDomain(*tag)
	if err := r.gormDB.Create(&gormTag).Error; err != nil {
		log.Println("failed to create tag", err)
	}
	tag.ID = gormTag.ID
	return nil
}

func (r *TagRepository) Update(tag *entities.Tag) error {
	gormTag := models.GormTagFromDomain(*tag)
	var databaseGormTag models.GormTag

	if err := r.gormDB.First(&databaseGormTag, tag.ID).Error; err != nil {
		return fmt.Errorf("failed to find tag to be updated: %v", err)
	}
	gormTag.CreatedAt = databaseGormTag.CreatedAt
	if err := r.gormDB.Save(&gormTag).Error; err != nil {
		log.Println("failed to update tag", err)
	}
	return nil
}
func (r *TagRepository) UpdateTagState(tagID uint, accepted bool) error {
	gormTag, err := r.FindByID(tagID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("no tag with ID %d exists: %v", tagID, err)
		}
		return err
	}
	gormTag.State = accepted
	if err := r.gormDB.Save(&gormTag).Error; err != nil {
		return fmt.Errorf("failed to update tag state: %v", err)
	}
	return nil
}

func (r *TagRepository) Delete(tagID uint) error {
	tag, err := r.FindByID(tagID)
	if err != nil {
		return err
	}
	gormTag := models.GormTagFromDomain(tag)
	if err := r.gormDB.Delete(&gormTag).Error; err != nil {
		log.Println("failed to delete tag", err)
	}
	return nil
}

func (r *TagRepository) FindByID(id uint) (entities.Tag, error) {
	var gormTag models.GormTag
	if err := r.gormDB.Where("id=?", id).First(&gormTag).Error; err != nil {
		log.Println("failed to find tag", err)
	}
	tag := gormTag.ToDomain()
	return tag, nil
}

func (r *TagRepository) FindByKey(key string) (entities.Tag, error) {
	var gormTag models.GormTag
	if err := r.gormDB.Where(" `key` = ?", key).First(&gormTag).Error; err != nil {
		return entities.Tag{}, fmt.Errorf("failed to find tag by provided key: %v", err)
	}
	tag := gormTag.ToDomain()
	return tag, nil
}

func (r *TagRepository) GetAllTags() ([]entities.Tag, error) {
	var gormTags []models.GormTag
	if err := r.gormDB.Find(&gormTags).Error; err != nil {
		return []entities.Tag{}, fmt.Errorf("failed to find tags %v", err)
	}
	var tags []entities.Tag
	for _, gormTag := range gormTags {
		tag := gormTag.ToDomain()
		tags = append(tags, tag)
	}
	return tags, nil
}
