package repositories

import "tagMicroservice/internal/domain/entities"

type TaxonomyRepository interface {
	Create(taxonomy *entities.Taxonomy) error
	Update(taxonomy *entities.Taxonomy) error
	Delete(taxonomy *entities.Taxonomy) error
	FindByID(id uint) (*entities.Taxonomy, error)
	FindByKey(key string) ([]entities.Taxonomy, error)
	SearchByTitle(title string) ([]entities.Taxonomy, error)
}
