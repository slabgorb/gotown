package genetics

import (
	"encoding/json"
	"regexp"
)

// Variant is the abstraction of a genetic variant
type Variant struct {
	Name  string
	Match *regexp.Regexp
}

type serializableVariant struct {
	Name  string `json:"name"`
	Match string `json:"match"`
}

// MarshalJSON implements json.Marshaler
func (v *Variant) MarshalJSON() ([]byte, error) {
	sv := &serializableVariant{
		Name:  v.Name,
		Match: v.Match.String(),
	}
	return json.Marshal(sv)
}

// UnmarshalJSON implements json.Unmarshaler
func (v *Variant) UnmarshalJSON(data []byte) error {
	sv := &serializableVariant{}
	json.Unmarshal(data, sv)
	v.Name = sv.Name
	r, err := regexp.Compile(sv.Match)
	if err != nil {
		return err
	}
	v.Match = r
	return nil
}

// NewVariant creates a new Variant struct
func NewVariant(name, match string) (*Variant, error) {
	r, err := regexp.Compile(match)
	if err != nil {
		return nil, err
	}
	return &Variant{Name: name, Match: r}, nil

}

// Matches checks a variant against a string
func (v *Variant) Matches(s string) int {
	matches := v.Match.FindAllStringIndex(s, -1)
	return len(matches)
}
