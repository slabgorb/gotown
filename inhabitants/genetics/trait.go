package genetics

import (
	"encoding/json"
	"fmt"
	"sort"
)

// Trait models an individual genetic trait, such as eye color.
type Trait struct {
	Name     string     `json:"name"`
	Variants []*Variant `json:"variants"`
}

func (t *Trait) GetName() string {
	return t.Name
}

// UnmarshalJSON implements json.Unmarshaler
func (t *Trait) UnmarshalJSON(data []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	name := fmt.Sprintf("%v", tmp["name"])
	vmaps, ok := tmp["variants"].([]interface{})
	if !ok {
		return fmt.Errorf("Could not parse variants for trait %s", name)
	}
	variants := []*Variant{}
	for _, vraw := range vmaps {
		v := vraw.(map[string]interface{})
		variant, err := NewVariant(fmt.Sprintf("%v", v["name"]), fmt.Sprintf("%v", v["match"]))
		if err != nil {
			return err
		}
		variants = append(variants, variant)
	}
	t.Name = name
	t.Variants = variants
	return nil
}

// Expression is the genetic expression of a Trait
func (t Trait) Expression(s string) (string, int) {
	m := t.matches(s)
	pl := make(pairList, len(m))
	i := 0
	for k, v := range m {
		pl[i] = pair{key: k, value: v, index: i}
		i++
	}
	sort.Sort(pl)
	return pl[len(m)-1].key, pl[len(m)-1].value

}

// NewTrait instantiates a Trait
func NewTrait(name string, variants []*Variant) Trait {
	return Trait{Name: name, Variants: variants}
}

func (t Trait) matches(s string) map[string]int {
	results := make(map[string]int)
	for _, variant := range t.Variants {
		if _, ok := results[variant.Name]; !ok {
			results[variant.Name] = 0
		}
		results[variant.Name] += variant.Matches(s)
	}
	return results
}

type pair struct {
	index int
	key   string
	value int
}

type pairList []pair

func (p pairList) Len() int { return len(p) }
func (p pairList) Less(i, j int) bool {
	if p[i].value == p[j].value { // tiebreaker
		return p[i].index < p[j].index
	}
	return p[i].value < p[j].value
}
func (p pairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
