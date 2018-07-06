package being

import (
	//"encoding/json"
	"fmt"
	"strings"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}

// Relatable models relative
type Relatable interface {
	IsChildOf(int) bool
	IsParentOf(int) bool
	IsSiblingOf(int) bool
	IsCloseRelativeOf(int) bool
	GetChildren() ([]*Being, error)
	GetID() int
}

// Marriageable abstracts the ability to marry
type Marriageable interface {
	Unmarried() bool
	GetAge() int
	Alive() bool
}

// Cultured abstracts cultures
type Cultured interface {
	inhabitants.Readable
	GetMaritalStrategies() []culture.MaritalStrategy
	GetNamers() map[inhabitants.Gender]*words.Namer
}

// Being represents any being, like a human, a vampire, whatever.
type Being struct {
	ID          int                  `json:"id" storm:"id,increment"`
	Name        *Name                `json:"name"`
	SpeciesName string               `json:"species_name"`
	CultureName string               `json:"culture_name"`
	Parents     []int                `json:"parents"`
	Children    []int                `json:"children"`
	Spouses     []int                `json:"spouses"`
	Gender      inhabitants.Gender   `json:"gender"`
	Dead        bool                 `json:"dead"`
	Chromosome  *genetics.Chromosome `json:"chromosome"`
	Age         int                  `json:"age"`
	Species     *species.Species     `json:"-"`
	Culture     *culture.Culture     `json:"-"`
}

// New initializes a being
func New(s *species.Species, c *culture.Culture) *Being {
	return &Being{
		Name:        &Name{},
		SpeciesName: s.GetName(),
		CultureName: c.GetName(),
		Species:     s,
		Culture:     c,
		Chromosome:  genetics.RandomChromosome(30),
		Gender:      inhabitants.Asexual,
	}
}

func getBeingsFromIDS(IDS []int) ([]*Being, error) {
	beings := []*Being{}
	for _, id := range IDS {
		b := &Being{ID: id}
		if err := b.Read(); err != nil {
			return nil, err
		}
		beings = append(beings, b)
	}
	return beings, nil
}

func (b *Being) GetParents() ([]*Being, error) {
	return getBeingsFromIDS(b.Parents)
}

// GetChildren returns the children of the being
func (b *Being) GetChildren() ([]*Being, error) {
	return getBeingsFromIDS(b.Children)
}

func (b *Being) getSpouses() ([]*Being, error) {
	return getBeingsFromIDS(b.Spouses)
}

// GetID returns the id
func (b *Being) GetID() int {
	return b.ID
}

func (b *Being) GetNamer() *words.Namer {
	return b.Culture.GetNamers()[b.Gender]
}

func (b *Being) genderedParent(gender inhabitants.Gender) (*Being, error) {
	parents, err := b.GetParents()
	if err != nil {
		return nil, err
	}
	for _, b := range parents {
		if b.Sex() == gender {
			return b, nil
		}
	}
	return nil, nil
}

// Reset sets the culture back to zero
func (b *Being) Reset() {
	b.ID = 0
	b.Name = &Name{}
	b.SpeciesName = ""
	b.CultureName = ""
	b.Spouses = []int{}
	b.Children = []int{}
	b.Parents = []int{}
	b.Chromosome = genetics.RandomChromosome(30)
	b.Gender = inhabitants.Asexual
}

// Read implements persist.Persistable
func (b *Being) Read() error {
	if b.ID == 0 {
		return fmt.Errorf("cannot read being without id")
	}
	if err := persist.DB.One("ID", b.ID, b); err != nil {
		return fmt.Errorf("could not load being %d: %s", b.ID, err)
	}
	b.Species = &species.Species{}
	if err := persist.DB.One("Name", b.SpeciesName, b.Species); err != nil {
		return fmt.Errorf("could not load species %s for being %d: %s", b.SpeciesName, b.ID, err)
	}
	b.Culture = &culture.Culture{}
	if err := persist.DB.One("Name", b.CultureName, b.Culture); err != nil {
		return fmt.Errorf("could not load culture %s for being %d: %s", b.CultureName, b.ID, err)
	}
	return nil
}

// Save implements persist.Persistable
func (b *Being) Save() error {
	b.CultureName = b.Culture.GetName()
	b.SpeciesName = b.Species.GetName()
	return persist.DB.Save(b)
}

// Delete implements persist.Persistable
func (b *Being) Delete() error {
	return persist.DB.DeleteStruct(b)
}

// GetName returns the name object
func (b *Being) GetName() string {
	return b.Name.Display
}

// GetFullName returns the name object
func (b *Being) GetFullName() *Name {
	return b.Name
}

// Father returns a male parent of the Being
func (b *Being) Father() (*Being, error) {
	return b.genderedParent(inhabitants.Male)
}

// Mother returns a female parent of the Being
func (b *Being) Mother() (*Being, error) {
	return b.genderedParent(inhabitants.Female)
}

// Randomize scrambles a Being randomly
func (b *Being) Randomize() error {
	if b.SpeciesName == "" {
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
	b.Age = b.Species.RandomAge(slot)
}

// RandomizeGender randomizes the Being's gender based on the possible genders
// the species exposes.
func (b *Being) RandomizeGender() {
	b.Gender = b.Species.GetGenders()[randomizer.Intn(len(b.Species.GetGenders()))]
}

// RandomizeName creates a new random name based on the being's culture.
func (b *Being) RandomizeName() {
	namer := b.Culture.GetNamers()[b.Sex()]
	b.Name = NameStrategies[namer.NameStrategy](b)
}

// RandomizeChromosome randomizes the being's chromosome.
func (b *Being) RandomizeChromosome() {
	b.Chromosome = genetics.RandomChromosome(20)
}

// Expression returns the genetic expression of the being's chromosome in the
// context of the being's species.
func (b *Being) Expression() map[string]string {
	return b.Chromosome.Express(b.Species.Expression())
}

// Marry marries two beings together. Marry does not check whether the beings
// are compatible marriage partners based on cultural settings, it is up to the
// caller to make sure they should be candidates.
func (b *Being) Marry(with *Being) {
	b.Spouses = append(b.Spouses, with.ID)
	with.Spouses = append(with.Spouses, b.ID)
}

// IsParentOf returns true of the receiver is the parent of the passed in being
func (b *Being) IsParentOf(with int) bool {
	for _, id := range b.Children {
		if id == with {
			return true
		}
	}
	return false
}

// IsChildOf returns true if the receiver being is a child of the passed in
// being
func (b *Being) IsChildOf(with int) bool {
	parents, err := b.GetParents()
	if err != nil {
		return false
	}
	for _, p := range parents {
		if p.GetID() == with {
			return true
		}
	}
	return false
}

// Sex returns the gender
func (b *Being) Sex() inhabitants.Gender {
	return b.Gender
}

// Unmarried returns whether this being has no spouses
func (b *Being) Unmarried() bool {
	return len(b.Spouses) == 0
}

// Siblings gets all siblings (half and full) of the receiver
func (b *Being) Siblings() (*Population, error) {
	children := make(map[int]struct{})
	sibs := []int{}

	parents, err := b.GetParents()
	if err != nil {
		return nil, err
	}
	for _, p := range parents {
		for _, c := range p.Children {
			children[c] = struct{}{}
		}
	}
	for s := range children {
		if s != b.ID {
			sibs = append(sibs, s)
		}
	}
	return NewPopulation(sibs), nil
}

// Piblings returns aunts and uncles of the receiver
func (b *Being) Piblings() (*Population, error) {
	parentSiblings := NewPopulation([]int{})
	parents, err := b.GetParents()
	if err != nil {
		return nil, err
	}
	for _, p := range parents {
		siblings, err := p.Siblings()
		if err != nil {
			return nil, err
		}
		parentSiblings.appendIds(siblings.getIds()...)
	}
	return parentSiblings, nil
}

// Cousins returns the beings who are cousins of this being
func (b *Being) Cousins() (*Population, error) {
	piblings, err := b.Piblings()
	if err != nil {
		return nil, err
	}
	cousins := NewPopulation([]int{})
	pibBeings, err := piblings.Inhabitants()
	if err != nil {
		return nil, err
	}
	for _, p := range pibBeings {
		cousins.appendIds(p.Children...)
	}
	return cousins, nil

}

// Niblings returns nieces and nephews of the receiver
func (b *Being) Niblings() (*Population, error) {
	siblings, err := b.Siblings()
	if err != nil {
		return nil, err
	}
	niblings := NewPopulation([]int{})
	sibs, err := siblings.Inhabitants()
	if err != nil {
		return nil, err
	}
	for _, s := range sibs {
		niblings.appendIds(s.Children...)
	}
	return niblings, nil
}

// IsSiblingOf checks to see if the receiver is a sibling of the passed in being
func (b *Being) IsSiblingOf(with int) bool {
	siblings, err := b.Siblings()
	if err != nil {
		return false
	}
	for _, s := range siblings.getIds() {
		if s == with {
			return true
		}
	}
	return false
}

// IsCloseRelativeOf returns true if the receiver is a close relative of the
// passed in being
func (b *Being) IsCloseRelativeOf(with int) bool {
	close := false
	close = close || b.IsChildOf(with)
	close = close || b.IsParentOf(with)
	close = close || b.IsSiblingOf(with)
	return close
}

// Reproduce creates new Being objects from the 'parent' beings
func (b *Being) Reproduce(with *Being) (*Being, error) {
	if with == nil && b.Sex() != inhabitants.Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	child := New(b.Species, b.Culture)
	child.Parents = []int{b.ID, with.ID}
	child.Randomize()
	child.Age = 0
	chromosome, err := b.Chromosome.Combine(with.Chromosome)
	if err != nil {
		return nil, err
	}
	child.Chromosome = chromosome
	if err := child.Save(); err != nil {
		return nil, fmt.Errorf("could not save new child: %s", err)
	}
	b.Children = append(b.Children, child.ID)
	with.Children = append(with.Children, child.ID)
	return child, nil
}

// GetAge returns the age of the being
func (b *Being) GetAge() int {
	return b.Age
}

// SetAge sets the age of the being
func (b *Being) SetAge(age int) {
	b.Age = age
}

// Die makes the being dead.
func (b *Being) Die(explanation ...string) {
	if len(explanation) == 0 {
		explanation = append(explanation, "unknown causes")
	}
	b.Dead = true
}

// String returns the string representation of the being.
func (b *Being) String() string {
	return strings.Trim(b.Name.GetDisplay(), " ")
}

// Alive returns whether this being is currently alive
func (b *Being) Alive() bool {
	return !b.Dead
}

func (b *Being) AddChildren(ids ...int) {

}

func saveAll(beings []*Being) error {
	ps := []persist.Persistable{}
	for _, b := range beings {
		ps = append(ps, b)
	}
	return persist.SaveAll(ps)
}

type BeingAPI struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	SpeciesName string            `json:"species"`
	CultureName string            `json:"culture"`
	Parents     []string          `json:"parents"`
	Children    []string          `json:"children"`
	Spouses     []string          `json:"spouses"`
	Gender      string            `json:"gender"`
	Age         int               `json:"age"`
	Chromosome  string            `json:"chromosome"`
	Expression  map[string]string `json:"expression"`
}

func getStrings(beings []*Being) []string {
	display := []string{}
	for _, b := range beings {
		display = append(display, b.Name.GetDisplay())
	}
	return display
}

func (b *Being) API() (interface{}, error) {
	parents, err := b.GetParents()
	if err != nil {
		return nil, err
	}
	children, err := b.GetChildren()
	if err != nil {
		return nil, err
	}
	spouses, err := b.getSpouses()
	if err != nil {
		return nil, err
	}
	expression := b.Expression()
	return &BeingAPI{
		ID:          b.ID,
		Parents:     getStrings(parents),
		Children:    getStrings(children),
		Spouses:     getStrings(spouses),
		Name:        b.Name.GetDisplay(),
		SpeciesName: b.Species.Name,
		CultureName: b.Culture.Name,
		Age:         b.Age,
		Gender:      b.Gender.String(),
		Chromosome:  b.Chromosome.String(),
		Expression:  expression,
	}, nil
}

// List returns the names of the beings already in the database
func List() ([]persist.IDPair, error) {
	beings := []Being{}
	if err := persist.DB.All(&beings); err != nil {
		return nil, err
	}
	items := []persist.IDPair{}
	for _, c := range beings {
		items = append(items, persist.IDPair{Name: c.Name.Display, ID: c.ID})
	}
	return items, nil
}
