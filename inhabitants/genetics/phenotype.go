package genetics

import (
	"regexp"
	"sort"
)

// Expression is a set of traits to be applied to a genotype (i.e. Chromosome)
type Expression []Trait

func (e *Expression) Add(trait Trait) {
	*e = append(*e, trait)
}

type Variant struct {
	Name  string
	Match *regexp.Regexp
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
	Name     string
	Variants []*Variant
}

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
