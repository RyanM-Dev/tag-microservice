package usecases

import "tagMicroservice/internal/domain/entities"

type TagUsecase interface {
	CreateTag(tag *entities.Tag) error
	UpdateTag(tag *entities.Tag) error
	DeleteTag(tagID uint) error
	GetTagByID(tagID uint) (*entities.Tag, error)
	ApproveTag(tagID uint) error
	RejectTag(tagID uint) error
	MergeTags(fromTagID, toTagID uint) error
	AddTaxonomy(fromTagID, toTagID uint, relationshipKind string, state bool) error
	SetTaxonomy(taxonomyID uint, relationshipKind string) error
	GetRelatedTagsByKey(key string) ([]entities.Tag, error)
	GetRelatedTagsByID(tagID uint) ([]entities.Tag, error)
	GetRelatedTagsByTitleAndKey(title, key string) ([]entities.Tag, error)
}
