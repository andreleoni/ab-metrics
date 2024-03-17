package entity

type Experiment struct {
	ID     string
	Name   string
	Key    string
	Status string

	// Relations
	Variations []Variation // Has many
}

// LookupVariationByID searches for a variation in the experiment's variations by ID.
// It returns the found variation and a boolean indicating whether the variation was found.
func (e *Experiment) LookupVariationByID(variationID string) (Variation, bool) {
	for _, v := range e.Variations {
		if v.ID == variationID {
			return v, true
		}
	}
	return Variation{}, false
}
