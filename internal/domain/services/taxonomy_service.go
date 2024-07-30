package services

import (
	"errors"
	"fmt"
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"
)

type TaxonomyService struct {
	TagRepo      repositories.TagRepository
	TaxonomyRepo repositories.TaxonomyRepository
}

func NewTaxonomyService(tagRepo repositories.TagRepository, taxonomyRepo repositories.TaxonomyRepository) *TagService {
	return &TagService{
		TagRepo:      tagRepo,
		TaxonomyRepo: taxonomyRepo,
	}
}
func (s *TaxonomyService) CreateTaxonomy(taxonomy *entities.Taxonomy) error {
	if taxonomy.FromTagID == 0 || taxonomy.ToTagID == 0 {
		return errors.New("from tag ID and to tag ID must be specified")
	}
	if taxonomy.RelationshipKind == "" {
		return errors.New("relationship must be specified")
	}
	_, err := s.TagRepo.FindByID(taxonomy.FromTagID)
	if err != nil {
		return fmt.Errorf("failed to get from tag ID: %v", err)
	}
	_, err = s.TagRepo.FindByID(taxonomy.ToTagID)
	if err != nil {
		return fmt.Errorf("failed to get to tag ID: %v", err)
	}
	if err := s.TaxonomyRepo.Create(taxonomy); err != nil {
		return fmt.Errorf("failed to create taxonomy: %v", err)
	}

	return nil
}

func (s *TaxonomyService) SetRelationshipKind(taxonomyID uint, kind string) error {
	taxonomy, err := s.TaxonomyRepo.FindByID(taxonomyID)
	if err != nil {
		return fmt.Errorf("failed to get taxonomy: %v", err)
	}
	taxonomy.RelationshipKind = kind
	if err := s.TaxonomyRepo.Update(&taxonomy); err != nil {
		return fmt.Errorf("failed to set relationship kind: %v", err)
	}
	return nil
}

func (s *TaxonomyService) GetTaxonomyByID(id uint) (entities.Taxonomy, error) {
	return s.TaxonomyRepo.FindByID(id)
}

func (s *TaxonomyService) UpdateTaxonomy(taxonomy *entities.Taxonomy) error {
	return s.TaxonomyRepo.Update(taxonomy)
}

func (s *TaxonomyService) DeleteTaxonomy(taxonomy *entities.Taxonomy) error {
	return s.TaxonomyRepo.Delete(taxonomy)
}
