package response

import "tagMicroservice/internal/domain/entities"

type TagRes struct {
	ID          uint   `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Key         string `json:"key,omitempty"`
	State       bool   `json:"state"`
}

func DomainToTagRes(tag entities.Tag) *TagRes {
	return &TagRes{
		ID:          tag.ID,
		Title:       tag.Title,
		Description: tag.Description,
		Image:       tag.Image,
		Key:         tag.Key,
		State:       tag.State,
	}
}
