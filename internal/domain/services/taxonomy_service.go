package services

import (
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"
)

type TaxonomyService struct {
	taxonomyRepo repositories.TaxonomyRepository
}

func NewTaxonomyService(taxonomyRepo repositories.TaxonomyRepository) *TaxonomyService {
	return &TaxonomyService{taxonomyRepo: taxonomyRepo}
}

func (s *TaxonomyService) CreateTaxonomy(taxonomy *entities.Taxonomy) error {
	return s.taxonomyRepo.Create(taxonomy)
}

func (s *TaxonomyService) GetTaxonomy(id uint) (*entities.Taxonomy, error) {
	return s.taxonomyRepo.FindByID(id)
}

func (s *TaxonomyService) UpdateTaxonomy(taxonomy *entities.Taxonomy) error {
	return s.taxonomyRepo.Update(taxonomy)
}

func (s *TaxonomyService) DeleteTaxonomy(taxonomy *entities.Taxonomy) error {
	return s.taxonomyRepo.Delete(taxonomy)
}

func (s *TaxonomyService) GetTaxonomiesByKey(key string) ([]entities.Taxonomy, error) {
	return s.taxonomyRepo.FindByKey(key)
}

func (s *TaxonomyService) SearchTaxonomiesByTitle(title string) ([]entities.Taxonomy, error) {
	return s.taxonomyRepo.SearchByTitle(title)
}
