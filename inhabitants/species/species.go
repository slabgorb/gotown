package species

import (
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
	ID         int                  `json:"id" storm:"id,increment"`
	Name       string               `json:"name" storm:"index,unique"`
	Genders    []inhabitants.Gender `json:"genders"`
	Expression genetics.Expression  `json:"expression"`
	Demography map[int]Demo         `json:"demography"`
}

// New creates and initializes a *Species
func New(name string, genders []inhabitants.Gender, e genetics.Expression, d map[int]Demo) *Species {
	return &Species{
		Name:       name,
		Genders:    genders,
		Expression: e,
		Demography: d,
	}
}

// String implements fmt.Stringer
func (s *Species) String() string {
	return s.Name
}

// Save implements persist.Persistable
func (s *Species) Save() error {
	return persist.DB.Save(s)
}

// Reset implements persist.Persistable
func (s *Species) Reset() {
	s.Name = ""
	s.ID = 0
	s.Genders = []inhabitants.Gender{}
	s.Expression = genetics.Expression{}
	s.Demography = make(map[int]Demo)
}

// Delete implements persist.Persistable
func (s *Species) Delete() error {
	return persist.DB.DeleteStruct(s)
}

// Fetch implements persist.Persistable
func (s *Species) Read() error {
	return persist.DB.One("Name", s.Name, s)
}

// GetGenders returns the genders appropriate for this species
func (s *Species) GetGenders() []inhabitants.Gender {
	return s.Genders
}

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

func Seed() error {
	var s = &Species{}
	return persist.SeedHelper("species", s)
}

func List() ([]string, error) {
	species := []Species{}
	if err := persist.DB.All(&species); err != nil {
		return nil, err
	}
	names := []string{}
	for _, c := range species {
		names = append(names, c.Name)
	}
	return names, nil
}
