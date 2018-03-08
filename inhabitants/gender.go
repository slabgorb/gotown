package inhabitants

import "encoding/json"

// Gender represents a sexual gender
type Gender string

// Gender enum
const (
	Asexual Gender = "neuter"
	Male    Gender = "male"
	Female  Gender = "female"
)

// MarshalText implements encoding.TextMarshaler
func (g Gender) MarshalText() ([]byte, error) {
	return []byte(g.String()), nil
}

// String implements fmt.Stringer
func (g Gender) String() string {
	return string(g)
}

// MarshalJSON implements json.Marshaler
func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// UnmarshalJSON implements json.Unmarshaler
func (g *Gender) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*g = Gender(s)
	return nil

}
