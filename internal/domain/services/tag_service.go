package services

import (
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/repositories"
)

type TagService struct {
	TagRepo repositories.TagRepository
}

func NewTagService(repo repositories.TagRepository) *TagService {
	return &TagService{TagRepo: repo}
}

func (s *TagService) CreateTag(tag entities.Tag) error {
	return s.TagRepo.Create(tag)
}

func (s *TagService) UpdateTag(tag entities.Tag) error {
	return s.TagRepo.Update(tag)
}

func (s *TagService) DeleteTag(tag entities.Tag) error {
	return s.TagRepo.Delete(tag)
}

func (s *TagService) FindTagByID(tagID string) (*entities.Tag, error) {
	return s.TagRepo.FindByID(tagID)

}

func (s *TagService) MergeTags(fromTagID, toTagID string) error {
	return s.TagRepo.Merge(fromTagID, toTagID)
}
