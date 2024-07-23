package repositories

import "tagMicroservice/internal/domain/entities"

type TagRepository interface {
	CreateTag(tag entities.Tag) error
	UpdateTag(tag entities.Tag) error
	FindByID(id string) (entities.Tag, error)
	Merge(fromTagID, toTagID string) error
}
