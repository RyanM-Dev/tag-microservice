package models

import (
	"tagMicroservice/internal/domain/entities"

	"gorm.io/gorm"
)

type GormTaxonomy struct {
	gorm.Model
	FromTagID        uint   `gorm:"not null"`
	ToTagID          uint   `gorm:"not null"`
	RelationshipKind string `gorm:"not null"`
	State            string `gorm:"not null"`
}

func (g GormTaxonomy) ToDomain() entities.Taxonomy {
	return entities.Taxonomy{
		ID:               g.ID,
		FromTagID:        g.FromTagID,
		ToTagID:          g.ToTagID,
		RelationshipKind: g.RelationshipKind,
		State:            g.State,
	}
}

func TaxonomyFromDomain(taxonomy entities.Taxonomy) GormTaxonomy {
	return GormTaxonomy{
		Model:            gorm.Model{ID: taxonomy.ID},
		FromTagID:        taxonomy.FromTagID,
		ToTagID:          taxonomy.ToTagID,
		RelationshipKind: taxonomy.RelationshipKind,
		State:            taxonomy.State,
	}
}
