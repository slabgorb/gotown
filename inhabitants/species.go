package inhabitants

import (
	"encoding/json"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

// go:generate stringer -type=Gender
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

type demo struct {
	max int
	pct int
}
type demography []demo

var Demographies = map[string]demography{
	"medieval": demography{
		demo{14, 29},
		demo{18, 36},
		demo{26, 50},
		demo{31, 58},
		demo{41, 72},
		demo{51, 84},
		demo{61, 93},
		demo{71, 98},
		demo{100, 100},
	},
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
	demog         demography
}

// NewSpecies creates and initializes a *Species
func NewSpecies(name string, genders []Gender, e *genetics.Expression) *Species {
	return &Species{
		Name:       name,
		Genders:    genders,
		Expression: e,
		demog:      Demographies["medieval"],
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
	for _, dmo := range s.demog {
		if dmo.pct >= slot {
			return randomizer.Intn(dmo.max-min) + min
		}
		min = dmo.max
	}
	return 0
}
