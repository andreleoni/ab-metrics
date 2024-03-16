package repository

type variation struct {
	percentage int
	key        string
}

var experiments = map[string][]variation{
	"new_checkout_button": {
		{80, "control"},
		{20, "a"},
		{20, "b"},
	},
}

type Experiment struct{}

func NewExperiment() Experiment {
	return Experiment{}
}

func (Experiment) Get() map[string][]variation {
	return experiments
}
