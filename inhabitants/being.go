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

type Being struct {
	*Name      `json:"name"`
	*Species   `json:"species"`
	Parents    map[Gender]*Being    `json:"parents"`
	Children   []*Being             `json:"children"`
	Age        int                  `json:"age"`
	Sex        Gender               `json:"gender"`
	Dead       bool                 `json:"dead"`
	Chromosome *genetics.Chromosome `json:"chromosome"`
}

func (b *Being) genderedParent(gender Gender) *Being {
	if parent, ok := b.Parents[gender]; ok {
		return parent
	}
	return nil
}

func (b *Being) MarshalJSON() ([]byte, error) {
	type Alias Being
	return json.Marshal(&struct {
		Expression map[string]string `json:"expression"`
		*Alias
		Living bool `json:"alive"`
	}{
		Expression: b.Expression(),
		Alias:      (*Alias)(b),
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
	possibleGenders := []Gender{}
	genders := b.GetGenders()
	for g := range b.GetGenders() {
		possibleGenders = append(possibleGenders, g)
	}
	//runtime.Breakpoint()
	b.Sex = possibleGenders[randomizer.Intn(len(possibleGenders))]
	b.RandomizeName()
	b.Age = genders[b.Sex].RandomAge(-1)
	return nil
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

	child.Parents = map[Gender]*Being{
		b.Sex:    b,
		with.Sex: with,
	}
	child.Randomize()
	b.Children = append(b.Children, child)

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
