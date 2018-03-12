package inhabitants

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/timeline"
)

// Name is the name of a being.
type Name struct {
	GivenName  string   `json:"given_name"`
	FamilyName string   `json:"family_name"`
	Other      []string `json:"other_name"`
	Display    string   `json:"display_name"`
}

// OtherNames returns any other names a being may have as a space-separated list
func (n *Name) OtherNames() string {
	return strings.Join(n.Other, " ")
}

// NewName tries valiantly to create a formal name from a string
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

// Members is a set of Being
type Members []*Being

// Strings gets all the beings in the Members slice and maps them to their
// string representation.
func (m Members) Strings() []string {
	var out []string
	for _, b := range m {
		out = append(out, b.String())
	}
	return out
}

// String returns the strings of all the Beings in the slice and joins them with
// commas.
func (m Members) String() string {
	return strings.Join(m.Strings(), ", ")
}

// Being represents any being, like a human, a vampire, whatever.
type Being struct {
	*Name
	*Species
	*Culture
	Parents    Members
	Children   Members              `json:"children"`
	Spouses    Members              `json:"spouses"`
	Sex        Gender               `json:"gender"`
	Dead       bool                 `json:"dead"`
	Chromosome *genetics.Chromosome `json:"chromosome"`
	Chronology *timeline.Chronology
}

// NewBeing initializes a being
func NewBeing(s *Species, c *Culture) *Being {
	return &Being{
		Species:    s,
		Culture:    c,
		Chronology: timeline.NewChronology(),
		Chromosome: genetics.RandomChromosome(30),
	}
}

func (b *Being) genderedParent(gender Gender) *Being {
	for _, b := range b.Parents {
		if b.Sex == gender {
			return b
		}
	}
	return nil
}

// MarshalJSON implements json.marshaler
func (b *Being) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Expression map[string]string         `json:"expression"`
		Age        int                       `json:"age"`
		Sex        string                    `json:"sex"`
		Species    string                    `json:"species"`
		Parents    []string                  `json:"parents"`
		Children   []string                  `json:"children"`
		Spouses    []string                  `json:"spouses"`
		Living     bool                      `json:"alive"`
		Events     map[int][]*timeline.Event `json:"events"`
		Culture    *Culture                  `json:"culture"`
	}{
		Expression: b.Expression(),
		Age:        b.Age(),
		Sex:        b.Sex.String(),
		Species:    b.Species.String(),
		Parents:    b.Parents.Strings(),
		Children:   b.Children.Strings(),
		Spouses:    b.Spouses.Strings(),
		Living:     !b.Dead,
		Events:     b.Chronology.Events,
		Culture:    b.Culture,
	})
}

// Father returns a male parent of the Being
func (b *Being) Father() *Being {
	return b.genderedParent(Male)
}

// Mother returns a female parent of the Being
func (b *Being) Mother() *Being {
	return b.genderedParent(Female)
}

// Randomize scrambles a Being randomly
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

// RandomizeAge sets the being age to a random number, based on the passed-in
// demographic slot.
func (b *Being) RandomizeAge(slot int) {
	b.Chronology.CurrentYear = b.Species.RandomAge(slot)
}

// RandomizeGender randomizes the Being's gender based on the possible genders
// the species exposes.
func (b *Being) RandomizeGender() {
	b.Sex = b.Species.Genders[randomizer.Intn(len(b.Species.Genders))]
}

// RandomizeName creates a new random name based on the being's culture.
func (b *Being) RandomizeName() {
	b.Name = NameStrategies[b.Culture.NameStrategies[b.Sex]](b)
}

// RandomizeChromosome randomizes the being's chromosome.
func (b *Being) RandomizeChromosome() {
	b.Chromosome = genetics.RandomChromosome(20)
}

//
func (b *Being) Express(e genetics.Expression) map[string]string {
	return b.Chromosome.Express(e)
}

func (b *Being) Expression() map[string]string {
	return b.Express(*b.Species.Expression)
}

// Marry marries two beings together. Marry does not check whether the beings
// are compatible marriage partners based on cultural settings, it is up to the
// caller to make sure they should be candidates.
func (b *Being) Marry(with *Being) {
	b.Spouses = append(b.Spouses, with)
	with.Spouses = append(with.Spouses, b)
	message := fmt.Sprintf("%s got married to %s", b.String(), with.String())
	b.Chronology.AddCurrentEvent(message)
	message = fmt.Sprintf("%s got married to %s", with.String(), b.String())
	with.Chronology.AddCurrentEvent(message)
}

// IsParentOf returns true of the receiver is the parent of the passed in being
func (b *Being) IsParentOf(with *Being) bool {
	for _, c := range b.Children {
		if c == with {
			return true
		}
	}
	return false
}

// IsChildOf returns true if the receiver being is a child of the passed in
// being
func (b *Being) IsChildOf(with *Being) bool {
	for _, c := range with.Children {
		if c == b {
			return true
		}
	}
	return false
}

// Siblings gets all siblings (half and full) of the receiver
func (b *Being) Siblings() Members {
	children := make(map[string]*Being)
	sibs := Members{}
	for _, p := range b.Parents {
		for _, c := range p.Children {
			children[fmt.Sprintf("%p", c)] = c
		}
	}
	for _, s := range children {
		if s != b {
			sibs = append(sibs, s)
		}
	}
	return sibs
}

// Piblings returns aunts and uncles of the receiver
func (b *Being) Piblings() Members {
	parentSiblings := Members{}
	for _, p := range b.Parents {
		parentSiblings = append(parentSiblings, p.Siblings()...)
	}
	return parentSiblings
}

func (b *Being) Cousins() Members {
	piblings := b.Piblings()
	cousins := Members{}
	for _, p := range piblings {
		cousins = append(cousins, p.Children...)
	}
	return cousins

}

// Niblings returns nieces and nephews of the receiver
func (b *Being) Niblings() Members {
	siblings := b.Siblings()
	niblings := Members{}
	for _, s := range siblings {
		niblings = append(niblings, s.Children...)
	}
	return niblings
}

// IsSiblingOf checks to see if the receiver is a sibling of the passed in being
func (b *Being) IsSiblingOf(with *Being) bool {
	siblings := b.Siblings()
	for _, s := range siblings {
		if s == with {
			return true
		}
	}
	return false
}

// IsCloseRelativeOf returns true if the receiver is a close relative of the
// passed in being
func (b *Being) IsCloseRelativeOf(with *Being) bool {
	close := false
	close = close || b.IsChildOf(with)
	close = close || b.IsParentOf(with)
	close = close || b.IsSiblingOf(with)
	return close
}

// Reproduce creates new Being objects from the 'parent' beings
func (b *Being) Reproduce(with *Being) ([]*Being, error) {
	if with == nil && b.Sex != Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	child := &Being{Species: b.Species, Chronology: timeline.NewChronology(), Culture: b.Culture}

	child.Parents = Members{b, with}
	child.Randomize()
	b.Children = append(b.Children, child)
	with.Children = append(with.Children, child)
	b.Chronology.AddCurrentEvent(fmt.Sprintf("%s had a child %s with %s", b, child, with))
	with.Chronology.AddCurrentEvent(fmt.Sprintf("%s had a child %s with %s", with, child, b))
	return b.Children, nil
}

func (b *Being) Age() int {
	return b.Chronology.CurrentYear
}

// Die makes the being dead.
func (b *Being) Die(explanation ...string) {
	if len(explanation) == 0 {
		explanation = append(explanation, "unknown causes")
	}
	b.Dead = true
	b.Chronology.AddCurrentEvent(fmt.Sprintf("Died from %s", explanation[0]))
	b.Chronology.Freeze()
}

// String returns the string representation of the being.
func (b *Being) String() string {
	return strings.Trim(b.Name.Display, " ")
}

// Alive returns whether this being is currently alive
func (b *Being) Alive() bool {
	return !b.Dead
}
