package inhabitants

import (
	"sort"

	"github.com/slabgorb/gotown/inhabitants/genetics"
)

// Species represents a species or a race.
type Species struct {
	ID         int                       `json:"id" storm:"id,increment"`
	Name       string                    `json:"name"`
	Genders    []Gender                  `json:"genders"`
	Expression *genetics.Expression      `json:"expression"`
	Demography map[DemographyBucket]Demo `json:"demography"`
}

// NewSpecies creates and initializes a *Species
func NewSpecies(name string, genders []Gender, e *genetics.Expression, d map[DemographyBucket]Demo) *Species {
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

// GetGenders returns the genders appropriate for this species
func (s *Species) GetGenders() []Gender {
	return s.Genders
}

// RandomAge provides a random age within a 'slot', or a demography bucket.
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
		if dmo.CumulativePercent >= slot {
			return randomizer.Intn(dmo.MaxAge-min) + min
		}
		min = dmo.MaxAge
	}
	return 0
}
