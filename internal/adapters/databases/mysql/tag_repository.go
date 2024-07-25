package mysql

import (
	"fmt"
	"log"
	"tagMicroservice/internal/adapters/databases/models"
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"

	"gorm.io/gorm"
)

type TagRepository struct{}

func NewTagRepository() repositories.TagRepository {
	return &TagRepository{}
}

func (r *TagRepository) Create(tag *entities.Tag) error {
	gormTag := models.GormTagFromDomain(*tag)
	if err := mysqlDB.db.Create(&gormTag).Error; err != nil {
		log.Println("failed to create tag", err)
	}
	tag.ID = gormTag.ID
	return nil
}

func (r *TagRepository) Update(tag *entities.Tag) error {
	gormTag := models.GormTagFromDomain(*tag)
	if err := mysqlDB.db.Save(&gormTag).Error; err != nil {
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
	if err := mysqlDB.db.Save(&gormTag).Error; err != nil {
		return fmt.Errorf("failed to update tag state: %v", err)
	}
	return nil
}

func (r *TagRepository) Delete(tag *entities.Tag) error {
	gormTag := models.GormTagFromDomain(*tag)
	if err := mysqlDB.db.Delete(&gormTag).Error; err != nil {
		log.Println("failed to delete tag", err)
	}
	return nil
}

func (r *TagRepository) FindByID(id uint) (entities.Tag, error) {
	var gormTag models.GormTag
	if err := mysqlDB.db.Where("id=?", id).First(&gormTag).Error; err != nil {
		log.Println("failed to find tag", err)
	}
	tag := gormTag.ToDomain()
	return tag, nil
}

func (r *TagRepository) FindByKey(key string) (entities.Tag, error) {
	var gormTag models.GormTag
	if err := mysqlDB.db.Where("key=?", key).First(&gormTag).Error; err != nil {
		return entities.Tag{}, fmt.Errorf("failed to find tag by provided key: %v", err)
	}
	tag := gormTag.ToDomain()
	return tag, nil
}

// func (r *TagRepository) Merge(fromTagID, toTagID uint) error {
// 	_, err := r.FindByID(fromTagID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return fmt.Errorf("fromTagID %d not found : %v", fromTagID, err)
// 		}
// 		return err
// 	}
// 	_, err = r.FindByID(fromTagID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return fmt.Errorf("ToTagID %d not found : %v", fromTagID, err)
// 		}
// 		return err
// 	}

// 	var fromTaxonomies []models.GormTaxonomy
// 	if err := mysqlDB.db.Where("from_tag = ?", fromTagID).Find(fromTaxonomies).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return services.ErrNoTagExistsWithThisID
// 		}
// 		return fmt.Errorf("failed to find taxonomy from tag %d : %v", fromTagID, err)
// 	}
// 	for _, taxonomy := range fromTaxonomies {
// 		taxonomy.FromTagID = toTagID
// 		if err := mysqlDB.db.Create(&taxonomy).Error; err != nil {
// 			return fmt.Errorf("failed to create new merged taxonomy from tag %d : %v", fromTagID, err)

// 		}
// 	}
// 	var toTaxonomies []models.GormTaxonomy
// 	if err := mysqlDB.db.Where("to_tag = ?", fromTagID).Find(toTaxonomies).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return services.ErrNoTagExistsWithThisID
// 		}
// 		return fmt.Errorf("failed to find taxonomy from tag %d : %v", toTagID, err)
// 	}
// 	for _, taxonomy := range toTaxonomies {
// 		taxonomy.ToTagID = toTagID
// 		if err := mysqlDB.db.Create(&taxonomy).Error; err != nil {
// 			return fmt.Errorf("failed to create new merged taxonomy from tag %d : %v", fromTagID, err)

// 		}
// 	}
// 	return nil
// }
