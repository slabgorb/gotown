package inhabitants

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/slabgorb/gotown/inhabitants/genetics"
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

func (g *Gender) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	*g = Gender(s)
	return nil

}

type Fertility struct {
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name       string                    `json:"name"`
	Genders    []Gender                  `json:"genders"`
	Expression *genetics.Expression      `json:"expression"`
	Demography map[DemographyBucket]Demo `json:"demography"`
}

func (s *Species) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Name       string               `json:"name"`
		Genders    []Gender             `json:"genders"`
		Expression *genetics.Expression `json:"expression"`
		DemoArray  []Demo               `json:"demography"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.Name = aux.Name
	s.Genders = aux.Genders
	s.Expression = aux.Expression
	d := make(map[DemographyBucket]Demo)
	for i, a := range aux.DemoArray {
		d[DemographyBucket(i)] = a
	}
	s.Demography = d
	return nil
}

func (s *Species) Load(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		return fmt.Errorf("could not load")
	}
	return nil
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
		if dmo.CumulativePercent >= slot {
			return randomizer.Intn(dmo.MaxAge-min) + min
		}
		min = dmo.MaxAge
	}
	return 0
}
