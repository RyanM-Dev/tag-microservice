package usecases

import (
	"tagMicroservice/internal/domain/entities"
	"tagMicroservice/internal/domain/services"
)

type TagUsecaseImp struct {
	tagService      services.TagService
	taxonomyService services.TaxonomyService
}

func NewTagUsecases(tagService services.TagService, taxonomyService services.TaxonomyService) TagUsecase {
	return &TagUsecaseImp{tagService: tagService, taxonomyService: taxonomyService}
}

func (use *TagUsecaseImp) CreateTag(tag *entities.Tag) error {
	return use.tagService.CreateTag(tag)

}

func (use *TagUsecaseImp) UpdateTag(tag *entities.Tag) error {
	return use.tagService.UpdateTag(tag)
}

func (use *TagUsecaseImp) DeleteTag(tag *entities.Tag) error {
	return use.tagService.DeleteTag(tag)
}

func (use *TagUsecaseImp) GetTagByID(tagID uint) (*entities.Tag, error) {
	tag, err := use.tagService.FindTagByID(tagID)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (use *TagUsecaseImp) ApproveTag(tagID uint) error {
	tag, err := use.tagService.FindTagByID(tagID)
	if err != nil {
		return err
	}
	tag.State = true
	if err = use.tagService.UpdateTag(&tag); err != nil {
		return err
	}
	return nil
}

func (use *TagUsecaseImp) RejectTag(tagID uint) error {
	tag, err := use.tagService.FindTagByID(tagID)
	if err != nil {
		return err
	}
	tag.State = false
	if err = use.tagService.UpdateTag(&tag); err != nil {
		return err
	}
	return nil

}

func (use *TagUsecaseImp) MergeTags(fromTagID, toTagID uint) error {
	if err := use.tagService.MergeTags(fromTagID, toTagID); err != nil {
		return err
	}
	return nil
}

func (use *TagUsecaseImp) AddTaxonomy(fromTagID, toTagID uint, relationshipKind string, state bool) error {
	taxonomy := entities.Taxonomy{
		FromTagID:        fromTagID,
		ToTagID:          toTagID,
		RelationshipKind: relationshipKind,
		State:            state,
	}
	if err := use.taxonomyService.CreateTaxonomy(&taxonomy); err != nil {
		return err
	}
	return nil
}

func (use *TagUsecaseImp) SetTaxonomy(taxonomyID uint, relationshipKind string) error {
	return use.taxonomyService.SetRelationshipKind(taxonomyID, relationshipKind)
}

func (use *TagUsecaseImp) GetRelatedTagsByKey(key string) ([]entities.Tag, error) {
	relatedTagsByKeys, err := use.tagService.GetRelatedTagsByKey(key)

	if err != nil {
		return nil, err
	}

	return relatedTagsByKeys, nil
}

func (use *TagUsecaseImp) GetRelatedTagsByID(tagID uint) ([]entities.Tag, error) {
	relatedTagsByID, err := use.tagService.GetRelatedTagsByID(tagID)
	if err != nil {
		return nil, err
	}
	return relatedTagsByID, nil
}

func (use *TagUsecaseImp) GetRelatedTagsByTitleAndKey(title, key string) ([]entities.Tag, error) {
	relatedTagsByTitleAndKey, err := use.tagService.GetRelatedTagsByTitleAndKey(title, key)
	if err != nil {
		return nil, err
	}
	return relatedTagsByTitleAndKey, nil
}
