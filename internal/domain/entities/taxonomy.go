package entities

type Taxonomy struct {
	ID               uint
	FromTagID        uint
	ToTagID          uint
	RelationshipKind string
	State            string
}
