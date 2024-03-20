package entity

type Actor struct {
	ID          string
	VariationID string
	Identifier  string
	CreatedAt   string

	// Relations
	Variation Variation // Belongs to
}
