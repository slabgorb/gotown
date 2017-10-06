package inhabitants

import (
	"encoding/json"
	"sort"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

type Gender string

const (
	Asexual Gender = "neuter"
	Male    Gender = "male"
	Female  Gender = "female"
)

func (g Gender) String() string {
	return string(g)
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

type Fertility struct {
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name          string                       `json:"name"`
	Genders       []Gender                     `json:"-"`
	MultipleBirth func(g random.Generator) int `json:"-"`
	Expression    *genetics.Expression         `json:"-"`
	randomizer    random.Generator
	Demography    `json:"demography"`
}

// NewSpecies creates and initializes a *Species
func NewSpecies(name string, genders []Gender, e *genetics.Expression) *Species {
	return &Species{
		Name:       name,
		Genders:    genders,
		Expression: e,
		Demography: Demographies["human"],
		MultipleBirth: func(g random.Generator) int {
			if g.Float64() < 0.05 {
				return 4
			}
			if g.Float64() < 0.1 {
				return 3
			}
			if g.Float64() < 0.3 {
				return 2
			}
			return 1
		},
	}
}

func (s *Species) String() string {
	return s.Name
}

func (s *Species) GetGenders() []Gender {
	return s.Genders
}

func (s *Species) RandomAge(slot int) int {
	if slot == -1 {
		slot = randomizer.Intn(101)
	}
	min := 0
	keys := make([]int, len(s.Demography))
	i := 0
	for k := range s.Demography {
		keys[i] = int(k)
		i++
	}
	sort.Ints(keys)
	for _, k := range keys {
		dmo := s.Demography[DemographyBucket(k)]
		if dmo.pct >= slot {
			return randomizer.Intn(dmo.max-min) + min
		}
		min = dmo.max
	}
	return 0
}
