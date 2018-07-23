package species

import (
	"fmt"
	"sort"

	"github.com/slabgorb/gotown/inhabitants/genetics"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}

// Species represents a species or a race.
type Species struct {
	persist.IdentifiableImpl
	Name              string               `json:"name" storm:"index,unique"`
	Genders           []inhabitants.Gender `json:"genders"`
	GeneticExpression genetics.Expression  `json:"expression"`
	Demography        map[int]Demo         `json:"demography"`
}

// New creates and initializes a *Species
func New(name string, genders []inhabitants.Gender, e genetics.Expression, d map[int]Demo) *Species {
	return &Species{
		Name:              name,
		Genders:           genders,
		GeneticExpression: e,
		Demography:        d,
	}
}

// Expression returns the species' genetic expression
func (s *Species) Expression() genetics.Expression {
	return s.GeneticExpression
}

// String implements fmt.Stringer
func (s *Species) String() string {
	return s.Name
}

func (s *Species) API() (interface{}, error) {
	return s, nil
}

// Save implements persist.Persistable
func (s *Species) Save() error {
	return persist.Save(s)
}

// Reset implements persist.Persistable
func (s *Species) Reset() {
	s.Name = ""
	s.IdentifiableImpl = persist.IdentifiableImpl{ID: ""}
	s.Genders = []inhabitants.Gender{}
	s.GeneticExpression = genetics.Expression{}
	s.Demography = make(map[int]Demo)
}

// Delete implements persist.Persistable
func (s *Species) Delete() error {
	return persist.Delete(s)
}

// Fetch implements persist.Persistable
func (s *Species) Read() error {
	if err := persist.Read(s); err != nil {
		return fmt.Errorf("cannot read species: %s", err)
	}
	return nil
}

// GetGenders returns the genders appropriate for this species
func (s *Species) GetGenders() []inhabitants.Gender {
	return s.Genders
}

// GetName returns the name of the specis
func (s *Species) GetName() string { return s.Name }

// RandomAge provides a random age within a 'slot', or a demography bucket. If th
func (s *Species) RandomAge(slot int) int {
	if slot == -1 || slot > int(inhabitants.MaxDemographyBucket) {
		slot = randomizer.Intn(int(inhabitants.MaxDemographyBucket))
	}
	min := 0
	var keys []int
	for k := range s.Demography {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	for _, k := range keys {
		dmo := s.Demography[int(k)]
		if k == slot {
			return randomizer.Intn(dmo.MaxAge-min) + min
		}
		min = dmo.MaxAge
	}
	return 0
}

// MaxAge returns the maximum age for the passed in demography slot
func (s *Species) MaxAge(slot int) int {
	dmo := s.Demography[slot]
	return dmo.MaxAge
}

// Seed populates the database with samples
func Seed() error {
	var s = &Species{}
	return persist.SeedHelper("species", s)
}

// List returns species names from the database
func List() ([]persist.IDPair, error) {
	species := []Species{}
	if err := persist.DB.All(&species); err != nil {
		return nil, err
	}
	names := []persist.IDPair{}
	for _, c := range species {
		names = append(names, persist.IDPair{Name: c.Name, ID: c.ID})
	}
	return names, nil
}
