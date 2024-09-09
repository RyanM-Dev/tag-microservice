package services

import (
	"errors"
	"fmt"
	"strings"
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"
)

var ErrNoTagExistsWithThisID = errors.New("no tag exists with this ID")

type TagService struct {
	TagRepo      repositories.TagRepository
	TaxonomyRepo repositories.TaxonomyRepository
}

func NewTagService(tagRepo repositories.TagRepository, taxonomyRepo repositories.TaxonomyRepository) *TagService {
	return &TagService{
		TagRepo:      tagRepo,
		TaxonomyRepo: taxonomyRepo,
	}
}
func (s *TagService) CreateTag(tag *entities.Tag) error {
	if tag.Title == "" || tag.Description == "" {
		return fmt.Errorf("tag title and description cannot be empty")
	}
	return s.TagRepo.Create(tag)
}

func (s *TagService) UpdateTag(tag *entities.Tag) error {
	return s.TagRepo.Update(tag)
}

func (s *TagService) DeleteTag(tagID uint) error {

	return s.TagRepo.Delete(tagID)
}

func (s *TagService) FindTagByID(tagID uint) (entities.Tag, error) {
	return s.TagRepo.FindByID(tagID)

}

func (s *TagService) MergeTags(fromTagID, toTagID uint) error {
	if err := s.TaxonomyRepo.UpdateTagReferences(fromTagID, toTagID); err != nil {
		return fmt.Errorf("failed to update tag references: %v", err)
	}
	newTag, err := s.TagRepo.FindByID(fromTagID)
	if err != nil {
		return err
	}
	newTag.ID = toTagID
	if err := s.TagRepo.Update(&newTag); err != nil {
		return fmt.Errorf("failed to merge tags: %v", err)
	}
	return nil
}

func (s *TagService) GetRelatedTagsByKey(key string) ([]entities.Tag, error) {
	var relatedTags []entities.Tag
	tag, err := s.TagRepo.FindByKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to find tag with key %s: %v", key, err)
	}
	taxonomies, err := s.TaxonomyRepo.FindTaxonomiesByTagID(tag.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find taxonomy with ID %d: %v", tag.ID, err)
	}
	for _, taxonomy := range taxonomies {
		var relatedTag entities.Tag
		var relatedTagID uint
		if taxonomy.FromTagID == tag.ID {
			relatedTagID = taxonomy.ToTagID
		} else if taxonomy.ToTagID == tag.ID {
			relatedTagID = taxonomy.FromTagID
		}
		relatedTag, err := s.TagRepo.FindByID(relatedTagID)
		if err != nil {
			return nil, err
		}
		relatedTags = append(relatedTags, relatedTag)
	}
	return relatedTags, nil
}

func (s *TagService) GetRelatedTagsByID(id uint) ([]entities.Tag, error) {
	var relatedTags []entities.Tag
	tag, err := s.TagRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find tag with id %d: %v", id, err)
	}
	taxonomies, err := s.TaxonomyRepo.FindTaxonomiesByTagID(tag.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find taxonomy with ID %d: %v", tag.ID, err)
	}
	for _, taxonomy := range taxonomies {
		var relatedTag entities.Tag
		var relatedTagID uint
		if taxonomy.FromTagID == tag.ID {
			relatedTagID = taxonomy.ToTagID
		} else if taxonomy.ToTagID == tag.ID {
			relatedTagID = taxonomy.FromTagID
		}
		relatedTag, err := s.TagRepo.FindByID(relatedTagID)
		if err != nil {
			return nil, err
		}
		relatedTags = append(relatedTags, relatedTag)
	}
	return relatedTags, nil
}

func (s *TagService) GetRelatedTagsByTitleAndKey(title, key string) ([]entities.Tag, error) {
	tag, err := s.TagRepo.FindByKey(key)
	var relatedTagsByTitle []entities.Tag
	if err != nil {
		return nil, fmt.Errorf("failed to get tag by key: %v", err)
	}
	relatedTags, err := s.GetRelatedTagsByID(tag.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get related tags by tag ID : %v", err)
	}
	for _, relatedTag := range relatedTags {
		if strings.Contains(strings.ToLower(relatedTag.Title), strings.ToLower(title)) {
			relatedTagsByTitle = append(relatedTagsByTitle, relatedTag)
		}
	}
	return relatedTagsByTitle, nil
}

func (s *TagService) GetAllTags() ([]entities.Tag, error) {
	tags, err := s.TagRepo.GetAllTags()
	if err != nil {
		return []entities.Tag{}, fmt.Errorf("failed to get tags: %v", err)
	}
	return tags, nil
}
