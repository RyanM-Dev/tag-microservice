package repositories

import "tagMicroservice/internal/domain/entities"

type TagRepository interface {
	Create(tag *entities.Tag) error
	Update(tag *entities.Tag) error
	Delete(tag *entities.Tag) error
	FindByID(id uint) (entities.Tag, error)
	FindByKey(key string) (entities.Tag, error)
	UpdateTagState(tagID uint, accepted bool) error
	// Merge(fromTagID, toTagID uint) error
}
