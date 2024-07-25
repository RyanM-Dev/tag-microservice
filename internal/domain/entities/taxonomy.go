package entities

type Taxonomy struct {
	ID               uint
	FromTagID        uint
	ToTagID          uint
	RelationshipKind string //inclusion,key_value,synonym,antonym
	State            string
}
