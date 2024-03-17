package entity

type Goal struct {
	ID      string
	ActorID string
	Key     string

	// Relations
	Actor Actor // Belongs to
}
