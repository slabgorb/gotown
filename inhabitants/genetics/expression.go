package genetics

import "encoding/json"

// Expression is a set of traits to be applied to a genotype (i.e. Chromosome)
type Expression struct {
	Traits []Trait `json:"traits"`
}

// Add adds a Trait to an Expression
func (e *Expression) Add(trait Trait) {
	e.Traits = append(e.Traits, trait)
}

// GetTraits is a getter for the traits
func (e *Expression) GetTraits() []Trait {
	return e.Traits
}

// UnmarshalJSON implements json.Unmarshaler
func (e *Expression) UnmarshalJSON(data []byte) error {
	tmp := make(map[string][]Trait)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	for _, t := range tmp["traits"] {
		e.Add(t)
	}
	return nil
}
