package repositories

import "tagMicroservice/internal/domain/entities"

type TagRepository interface {
	Create(tag entities.Tag) error
	Update(tag entities.Tag) error
	Delete(tag entities.Tag) error
	FindByID(id uint) (*entities.Tag, error)
	Merge(fromTagID, toTagID uint) error
}
