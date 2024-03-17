package entity

type Actor struct {
	ID          string
	VariationID string
	Identifier  string

	// Relations
	Variation Variation // Belongs to
}
