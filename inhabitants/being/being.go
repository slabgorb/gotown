package being

import (
	//"encoding/json"
	"fmt"
	"strings"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/timeline"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
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
	Name       *inhabitants.Name `json:"name"`
	Species    inhabitants.Specieser
	Parents    Members
	Children   Members              `json:"children"`
	Spouses    Members              `json:"spouses"`
	Gender     inhabitants.Gender   `json:"gender"`
	Dead       bool                 `json:"dead"`
	Chromosome *genetics.Chromosome `json:"chromosome"`
	Chronology *timeline.Chronology
}

// New initializes a being
func New(s inhabitants.Specieser) *Being {
	return &Being{
		Species:    s,
		Chronology: timeline.NewChronology(),
		Chromosome: genetics.RandomChromosome(30),
		Spouses:    make(Members, 1),
		Children:   make(Members, 1),
	}
}

func (b *Being) genderedParent(gender inhabitants.Gender) *Being {
	for _, b := range b.Parents {
		if b.Sex() == gender {
			return b
		}
	}
	return nil
}

func (b *Being) History() *timeline.Chronology {
	return b.Chronology
}

// MarshalJSON implements json.marshaler
// func (b *Being) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(&struct {
// 		Expression map[string]string         `json:"expression"`
// 		Chromosome *genetics.Chromosome      `json:"chromosome"`
// 		Age        int                       `json:"age"`
// 		Sex        string                    `json:"sex"`
// 		Species    string                    `json:"species"`
// 		Parents    []string                  `json:"parents"`
// 		Children   []string                  `json:"children"`
// 		Spouses    []string                  `json:"spouses"`
// 		Living     bool                      `json:"alive"`
// 		Events     map[int][]*timeline.Event `json:"events"`
// 		Culture    string                    `json:"culture"`
// 		Name       *Name                     `json:"name"`
// 	}{
// 		Expression: b.Expression(),
// 		Age:        b.Age(),
// 		Sex:        b.Sex.String(),
// 		Species:    b.Species.String(),
// 		Parents:    b.Parents.Strings(),
// 		Children:   b.Children.Strings(),
// 		Spouses:    b.Spouses.Strings(),
// 		Living:     !b.Dead,
// 		Events:     b.Chronology.Events,
// 		Culture:    b.Culture,
// 		Name:       b.Name,
// 	})
// }

func (b *Being) GetName() *inhabitants.Name {
	return b.Name
}

// Father returns a male parent of the Being
func (b *Being) Father() *Being {
	return b.genderedParent(inhabitants.Male)
}

// Mother returns a female parent of the Being
func (b *Being) Mother() *Being {
	return b.genderedParent(inhabitants.Female)
}

func (b *Being) SetAge(age int) {
	b.Chronology.CurrentYear = age
}

// Randomize scrambles a Being randomly
func (b *Being) Randomize(c inhabitants.Cultured) error {
	if b.Species == nil {
		return fmt.Errorf("Cannot randomize a being without a species")
	}
	b.RandomizeChromosome()
	b.RandomizeGender()
	b.RandomizeName(c)
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
	b.Gender = b.Species.GetGenders()[randomizer.Intn(len(b.Species.GetGenders()))]
}

// RandomizeName creates a new random name based on the being's culture.
func (b *Being) RandomizeName(c inhabitants.Cultured) {
	b.Name = c.RandomName(b.Sex())
}

// RandomizeChromosome randomizes the being's chromosome.
func (b *Being) RandomizeChromosome() {
	b.Chromosome = genetics.RandomChromosome(20)
}

// Express is the being's chromosome's expression
func (b *Being) Express(e inhabitants.Expresser) map[string]string {
	return b.Chromosome.Express(e)
}

// Expression returns the genetic expression of the being's chromosome in the
// context of the being's species.
func (b *Being) Expression() map[string]string {
	return b.Express(b.Species.Expression())
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
func (b *Being) IsParentOf(with inhabitants.Relatable) bool {
	for _, c := range b.Children {
		if inhabitants.Relatable(c) == with {
			return true
		}
	}
	return false
}

// IsChildOf returns true if the receiver being is a child of the passed in
// being
func (b *Being) IsChildOf(with inhabitants.Relatable) bool {
	for _, c := range with.GetChildren() {
		if c == b {
			return true
		}
	}
	return false
}

func (b *Being) Sex() inhabitants.Gender {
	return b.Gender
}

func (b *Being) Unmarried() bool {
	return len(b.Spouses) == 0
}

func (b *Being) GetChildren() []inhabitants.Relatable {
	relatables := []inhabitants.Relatable{}
	for _, child := range b.Children {
		relatables = append(relatables, child)

	}
	return relatables
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

// Cousins returns the beings who are cousins of this being
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
func (b *Being) IsSiblingOf(with inhabitants.Relatable) bool {
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
func (b *Being) IsCloseRelativeOf(with inhabitants.Relatable) bool {
	close := false
	close = close || b.IsChildOf(with)
	close = close || b.IsParentOf(with)
	close = close || b.IsSiblingOf(with)
	return close
}

// Reproduce creates new Being objects from the 'parent' beings
func (b *Being) Reproduce(with *Being, c inhabitants.Cultured) ([]*Being, error) {
	if with == nil && b.Sex() != inhabitants.Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	child := &Being{Species: b.Species, Chronology: timeline.NewChronology()}

	child.Parents = Members{b, with}
	child.Randomize(c)
	b.Children = append(b.Children, child)
	with.Children = append(with.Children, child)
	b.Chronology.AddCurrentEvent(fmt.Sprintf("%s had a child %s with %s", b, child, with))
	with.Chronology.AddCurrentEvent(fmt.Sprintf("%s had a child %s with %s", with, child, b))
	return b.Children, nil
}

// Age returns the age of the being
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
	return strings.Trim(b.Name.GetDisplay(), " ")
}

// Alive returns whether this being is currently alive
func (b *Being) Alive() bool {
	return !b.Dead
}
