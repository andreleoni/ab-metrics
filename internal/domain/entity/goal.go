package entity

type Goal struct {
	ID        string
	ActorID   string
	Key       string
	CreatedAt string

	// Relations
	Actor Actor // Belongs to
}
