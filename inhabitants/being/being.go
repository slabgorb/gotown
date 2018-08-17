package being

import (
	//"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/logger"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

// Logger is a logging interface (TODO: am I really using this?)
type Logger interface {
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	SetOutput(out io.Writer)
	TimeSet(key string)
	TimeElapsed(key string)
}

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
	persist.IdentifiableImpl
	PopulationID string               `json:"population_id"`
	Name         *Name                `json:"name"`
	SpeciesID    string               `json:"species_id"`
	CultureID    string               `json:"culture_id"`
	Parents      []string             `json:"parents"`
	Children     []string             `json:"children"`
	Spouses      []string             `json:"spouses"`
	Gender       inhabitants.Gender   `json:"gender"`
	Dead         bool                 `json:"dead"`
	Chromosome   *genetics.Chromosome `json:"chromosome"`
	Age          int                  `json:"age"`
	Species      *species.Species     `json:"-"`
	Culture      *culture.Culture     `json:"-"`
	logger       Logger
}

// New initializes a being
func New(s *species.Species, c *culture.Culture, logger Logger) *Being {
	return &Being{
		Name:       &Name{},
		SpeciesID:  s.GetID(),
		CultureID:  c.GetID(),
		Species:    s,
		Culture:    c,
		Chromosome: genetics.RandomChromosome(30),
		Gender:     inhabitants.Asexual,
		logger:     logger,
	}
}

func getBeingsFromIDS(IDS []string) ([]*Being, error) {
	beings := []*Being{}
	for _, id := range IDS {
		b := &Being{IdentifiableImpl: persist.IdentifiableImpl{ID: id}}
		if err := b.Read(); err != nil {
			return nil, err
		}
		beings = append(beings, b)
	}
	return beings, nil
}

// GetParents returns the parents for this being
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

// GetNamer returns the namer for this being from the associated culture
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
	b.ID = ""
	b.Name = &Name{}
	b.SpeciesID = ""
	b.Species = nil
	b.CultureID = ""
	b.Culture = nil
	b.Spouses = []string{}
	b.Children = []string{}
	b.Parents = []string{}
	b.Chromosome = genetics.RandomChromosome(30)
	b.Gender = inhabitants.Asexual
}

// Read implements persist.Persistable
func (b *Being) Read() error {
	if b.ID == "" {
		return fmt.Errorf("cannot read being without id")
	}
	if err := persist.Read(b); err != nil {
		return fmt.Errorf("could not load being %s: %s", b.ID, err)
	}
	b.Species = &species.Species{}
	b.Species.SetID(b.SpeciesID)
	if err := persist.Read(b.Species); err != nil {
		return fmt.Errorf("could not load species %s for being %s: %s", b.SpeciesID, b.ID, err)
	}
	b.Culture = &culture.Culture{}
	b.Culture.SetID(b.CultureID)
	if err := persist.Read(b.Culture); err != nil {
		return fmt.Errorf("could not load culture %s for being %s: %s", b.CultureID, b.ID, err)
	}
	return nil
}

// Save implements persist.Persistable
func (b *Being) Save() error {
	b.CultureID = b.Culture.GetID()
	b.SpeciesID = b.Species.GetID()
	return persist.Save(b)
}

func (b *Being) Update() error {
	return persist.Update(b)
}

// Delete implements persist.Persistable
func (b *Being) Delete() error {
	return persist.Delete(b)
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
	if b.SpeciesID == "" {
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
func (b *Being) IsParentOf(with string) bool {
	for _, id := range b.Children {
		if id == with {
			return true
		}
	}
	return false
}

// IsChildOf returns true if the receiver being is a child of the passed in
// being
func (b *Being) IsChildOf(with string) bool {
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
	children := make(map[string]struct{})
	sibs := []string{}

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
	return NewPopulation(sibs, b.logger), nil
}

// Piblings returns aunts and uncles of the receiver
func (b *Being) Piblings() (*Population, error) {
	parentSiblings := NewPopulation([]string{}, b.logger)
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
	cousins := NewPopulation([]string{}, b.logger)
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
	niblings := NewPopulation([]string{}, b.logger)
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
func (b *Being) IsSiblingOf(with string) bool {
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
func (b *Being) IsCloseRelativeOf(with string) bool {
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
	child := New(b.Species, b.Culture, b.logger)
	child.Parents = []string{b.ID, with.ID}
	child.Randomize()
	child.Age = 0
	child.PopulationID = b.PopulationID
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
	k := fmt.Sprintf("saving %d beings", len(beings))
	logger.TimeSet(k)
	ps := []persist.Persistable{}
	for _, b := range beings {
		ps = append(ps, b)
	}
	defer func() { logger.TimeElapsed(k) }()
	return persist.SaveAll(ps)
}

func readAll(ids []string) ([]*Being, error) {
	beings := []*Being{}
	m := make(map[string]*Being)
	for _, s := range ids {
		m[s] = &Being{}
		m[s].SetID(s)
	}
	ps := make(map[string]persist.Persistable)
	for k, v := range m {
		ps[k] = v
	}
	logger.TimeSet("mread")
	if err := persist.Mread(ids, ps); err != nil {
		return nil, err
	}
	logger.TimeElapsed("mread")

	for _, v := range ps {
		b := v.(*Being)
		beings = append(beings, b)
	}

	speciesCache := make(map[string]*species.Species)
	cultureCache := make(map[string]*culture.Culture)

	getSpecies := func(id string) (*species.Species, error) {
		if s, ok := speciesCache[id]; ok {
			return s, nil
		}
		item := &species.Species{}
		item.SetID(id)
		if err := item.Read(); err != nil {
			return nil, err
		}
		speciesCache[id] = item
		return item, nil
	}

	getCulture := func(id string) (*culture.Culture, error) {
		if s, ok := cultureCache[id]; ok {
			return s, nil
		}
		item := &culture.Culture{}
		item.SetID(id)
		if err := item.Read(); err != nil {
			return nil, err
		}
		cultureCache[id] = item
		return item, nil
	}
	logger.TimeSet("getting species/culture")
	for _, b := range beings {
		s, err := getSpecies(b.SpeciesID)
		if err != nil {
			return nil, err
		}
		c, err := getCulture(b.CultureID)
		if err != nil {
			return nil, err
		}
		b.Species = s
		b.Culture = c
	}
	logger.TimeElapsed("getting species/culture")
	return beings, nil
}

type BeingAPI struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	SpeciesName string            `json:"species"`
	SpeciesID   string            `json:"species_id"`
	CultureName string            `json:"culture"`
	CultureID   string            `json:"culture_id"`
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
		SpeciesID:   b.Species.ID,
		CultureName: b.Culture.Name,
		CultureID:   b.Culture.ID,
		Age:         b.Age,
		Gender:      b.Gender.String(),
		Chromosome:  b.Chromosome.String(),
		Expression:  expression,
	}, nil
}

// List returns the names of the beings already in the database
func List() (map[string]string, error) {
	list, err := persist.List("Being")
	if err != nil {
		return nil, err
	}
	return list, nil
}
