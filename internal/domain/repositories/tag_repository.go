package repositories

import "tagMicroservice/internal/domain/entities"

type TagRepository interface {
	Create(tag entities.Tag) error
	Update(tag entities.Tag) error
	FindByID(id string) (entities.Tag, error)
	Merge(fromTagID, toTagID string) error
}
