package entities

type Taxonomy struct {
	ID               string
	FromTagID        string
	ToTagID          string
	RelationshipKind string
	State            string
}
