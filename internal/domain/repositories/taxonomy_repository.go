package repositories

import "tagMicroservice/internal/domain/entities"

type TaxonomyRepository interface {
	Create(taxonomy entities.Taxonomy) error
	FindRelatedTags(tagID string) ([]entities.Tag, error)
	FindByKey(key string) ([]entities.Tag, error)
}
