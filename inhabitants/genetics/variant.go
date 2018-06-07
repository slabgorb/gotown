package genetics

import "strings"

// Variant is the abstraction of a genetic variant
type Variant struct {
	Name  string `json:"name"`
	Match string `json:"match"`
}

// NewVariant creates a new Variant struct
func NewVariant(name, match string) (*Variant, error) {
	return &Variant{Name: name, Match: match}, nil
}

// Matches checks a variant against a string
func (v *Variant) Matches(s string) int {
	count := 0
	matchers := strings.Split(v.Match, ",")
	for _, m := range matchers {
		count += strings.Count(s, m)
	}
	return count
}
