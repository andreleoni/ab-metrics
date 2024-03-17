package entity

import (
	"time"
)

type Event struct {
	ID        string
	ActorID   string
	GoalID    string
	Timestamp time.Time

	// Relations
	Actor Actor // Belongs to
	Goal  Goal  // Belongs to
}
