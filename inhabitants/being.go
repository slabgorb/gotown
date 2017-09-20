package inhabitants

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/slabgorb/gotown/inhabitants/genetics"
)

type Name struct {
	GivenName  string   `json:"given_name"`
	FamilyName string   `json:"family_name"`
	Other      []string `json:"other_name"`
	Display    string   `json:"display_name"`
}

func (n *Name) OtherNames() string {
	return strings.Join(n.Other, " ")
}

func NewName(fullName string) *Name {
	name := &Name{Display: fullName}
	names := strings.Split(fullName, " ")
	if len(names) > 0 {
		name.GivenName = names[0]
	}
	if len(names) > 1 {
		name.FamilyName = names[1]
	}
	if len(names) > 2 {
		name.Other = names[2:]
	}
	return name
}

type members []*Being

func (m members) Strings() []string {
	var out []string
	for _, b := range m {
		out = append(out, b.String())
	}
	return out
}

func (m members) String() string {
	return strings.Join(m.Strings(), ", ")
}

type Being struct {
	*Name      `json:"name"`
	*Species   `json:"species"`
	Parents    members              `json:"parents"`
	Children   members              `json:"children"`
	Spouses    members              `json:"spouses"`
	Age        int                  `json:"age"`
	Sex        Gender               `json:"gender"`
	Dead       bool                 `json:"dead"`
	Chromosome *genetics.Chromosome `json:"chromosome"`
}

func (b *Being) genderedParent(gender Gender) *Being {
	for _, b := range b.Parents {
		if b.Sex == gender {
			return b
		}
	}
	return nil
}

func (b *Being) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Expression map[string]string `json:"expression"`
		Age        int               `json"age"`
		Sex        string            `json:"sex"`
		Species    string            `json:"species"`
		Parents    []string          `json:"parents"`
		Children   []string          `json:"children"`
		Spouses    []string          `json:"spouses"`
		Living     bool              `json:"alive"`
	}{
		Expression: b.Expression(),
		Age:        b.Age,
		Sex:        b.Sex.String(),
		Species:    b.Species.String(),
		Parents:    b.Parents.Strings(),
		Children:   b.Children.Strings(),
		Spouses:    b.Spouses.Strings(),
		Living:     !b.Dead,
	})
}

func (b *Being) Father() *Being {
	return b.genderedParent(Male)
}

func (b *Being) Mother() *Being {
	return b.genderedParent(Female)
}

func (b *Being) Randomize() error {
	if b.Species == nil {
		return fmt.Errorf("Cannot randomize a being without a species")
	}
	b.RandomizeChromosome()
	b.RandomizeGender()
	b.RandomizeName()
	b.RandomizeAge(-1)
	return nil
}

func (b *Being) RandomizeAge(slot int) {
	genders := b.GetGenders()
	b.Age = genders[b.Sex].RandomAge(slot)
}

func (b *Being) RandomizeGender() {
	possibleGenders := []Gender{}
	for g := range b.GetGenders() {
		possibleGenders = append(possibleGenders, g)
	}
	b.Sex = possibleGenders[randomizer.Intn(len(possibleGenders))]
}

func (b *Being) RandomizeName() {
	genders := b.GetGenders()
	b.Name = genders[b.Sex].NameStrategy(b)
}

func (b *Being) RandomizeChromosome() {
	b.Chromosome = genetics.RandomChromosome(20)
}

func (b *Being) Express(e genetics.Expression) map[string]string {
	return b.Chromosome.Express(e)
}

func (b *Being) Expression() map[string]string {
	return b.Express(*b.Species.Expression)
}

// Reproduce creates new Being objects from the 'parent' beings
func (b *Being) Reproduce(with *Being) ([]*Being, error) {
	if with == nil && b.Sex != Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	child := &Being{Species: b.Species, Age: 0}

	child.Parents = members{b, with}
	child.Randomize()
	b.Children = append(b.Children, child)
	with.Children = append(with.Children, child)

	return b.Children, nil
}

func (b *Being) Die() {
	b.Dead = true
}

func (b *Being) String() string {
	return strings.Trim(b.Name.Display, " ")
}

func (b *Being) Alive() bool {
	return !b.Dead
}
