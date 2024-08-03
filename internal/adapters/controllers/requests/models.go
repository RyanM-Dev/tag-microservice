package requests

import "tagMicroservice/internal/domain/entities"

type CreateTagReq struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Key         string `json:"key,omitempty"`
	State       bool   `json:"state"`
}

type UpdateTagReq struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Key         string `json:"key,omitempty"`
	State       bool   `json:"state"`
}

type CreateTaxonomyReq struct {
	FromTagID        uint   `json:"from_tag_id" validate:"required"`
	ToTagID          uint   `json:"to_tag_id" validate:"required"`
	RelationshipKind string `json:"relationship_kind" validate:"required"`
	State            string `json:"state"`
}

type UpdateTaxonomyReq struct {
	ID               uint   `json:"id" validate:"required"`
	FromTagID        uint   `json:"from_tag_id" validate:"required"`
	ToTagID          uint   `json:"to_tag_id" validate:"required"`
	RelationshipKind string `json:"relationship_kind" validate:"required"`
	State            string `json:"state"`
}
type TagIDReq struct {
	ID uint `json:"id" validate:"required"`
}

type MergeTagsReq struct {
	FromTagID uint `json:"from_tag_id" validate:"required"`
	ToTagID   uint `json:"to_tag_id" validate:"required"`
}

type AddTaxonomyReq struct {
	FromTagID        uint   `json:"from_tag_id" validate:"required"`
	ToTagID          uint   `json:"to_tag_id" validate:"required"`
	RelationshipKind string `json:"relationship_kind" validate:"required"`
	State            bool   `json:"state" validate:"required"`
}
type SetTaxonomyReq struct {
	TaxonomyID       uint   `json:"taxonomy_id" validate:"required"`
	RelationshipKind string `json:"relationship_kind" validate:"required" `
}
type GetRelatedTagsByKeyReq struct {
	Key string `json:"key" validate:"required"`
}
type GetRelatedTagsByTitleAndKeyReq struct {
	Title string `json:"title" validate:"required"`
	Key   string `json:"key" validate:"required:"`
}

func CreateTagReqToTagEntity(req CreateTagReq) *entities.Tag {
	return &entities.Tag{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Key:         req.Key,
		State:       req.State,
	}
}

func UpdateTagReqToTagEntity(req UpdateTagReq) *entities.Tag {
	return &entities.Tag{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Key:         req.Key,
		State:       req.State,
	}
}
