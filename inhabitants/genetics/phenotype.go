package genetics

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
)

// Expression is a set of traits to be applied to a genotype (i.e. Chromosome)
type Expression struct {
	Traits []Trait `json:"traits"`
}

func (e *Expression) Add(trait Trait) {
	e.Traits = append(e.Traits, trait)
}

func LoadExpression(r io.Reader) (Expression, error) {
	exp := Expression{}
	err := json.NewDecoder(r).Decode(&exp)
	return exp, err
}

func (e *Expression) UnmarshalJSON(data []byte) error {
	tmp := make(map[string][]Trait)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	for _, t := range tmp["traits"] {
		fmt.Println(t.Variants[0])
		e.Add(t)
	}
	return nil
}

type Variant struct {
	Name  string         `json:"name"`
	Match *regexp.Regexp `json:"match"`
}

// NewVariant creates a new Variant struct
func NewVariant(name, match string) (*Variant, error) {
	r, err := regexp.Compile(match)
	if err != nil {
		return nil, err
	}
	return &Variant{Name: name, Match: r}, nil

}

func (v *Variant) Matches(s string) int {
	matches := v.Match.FindAllStringIndex(s, -1)
	return len(matches)
}

// Trait models an individual genetic trait, such as eye color.
type Trait struct {
	Name     string     `json:"name"`
	Variants []*Variant `json:"variants"`
}

// UnmarshalJSON unmarshals the JSON. Yep.
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
