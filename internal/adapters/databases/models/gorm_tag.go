package models

import (
	"tagMicroservice/internal/domain/entities"

	"gorm.io/gorm"
)

type GormTag struct {
	gorm.Model
	Title       string `gorm:"unique;not null"`
	Description string
	Image       string
	Key         string `gorm:"unique;not null"`
	State       bool   `gorm:"not null"`
}

func (g GormTag) ToDomain() entities.Tag {
	return entities.Tag{
		ID:          g.ID,
		Title:       g.Title,
		Description: g.Description,
		Image:       g.Image,
		Key:         g.Key,
		State:       g.State,
	}
}

func GormTagFromDomain(tag entities.Tag) GormTag {
	return GormTag{
		Model:       gorm.Model{ID: tag.ID},
		Title:       tag.Title,
		Description: tag.Description,
		Image:       tag.Image,
		Key:         tag.Key,
		State:       tag.State,
	}
}
