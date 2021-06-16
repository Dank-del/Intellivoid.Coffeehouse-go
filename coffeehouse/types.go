package coffeehouse

type Generalization struct {
	ID             string         `json:"id"`
	Size           int            `json:"size"`
	TopLabel       string         `json:"top_label"`
	TopProbability float64        `json:"top_probability"`
	Probabilities  *Probabilities `json:"probabilities"`
}

type Probabilities struct {
	Label                 string    `json:"label"`
	CalculatedProbability float64   `json:"calculated_probability"`
	CurrentPointer        int       `json:"current_pointer"`
	Probabilities         []float64 `json:"probabilities"`
}

type Error struct {
	ErrorCode int    `json:"error_code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}
