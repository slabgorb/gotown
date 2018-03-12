package inhabitants

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/persist"
)

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

// func (s *Species) UnmarshalJSON(data []byte) error {
// 	aux := &struct {
// 		Name       string               `json:"name"`
// 		Genders    []Gender             `json:"genders"`
// 		Expression *genetics.Expression `json:"expression"`
// 		DemoArray  []Demo               `json:"demography"`
// 	}{}
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}
// 	s.Name = aux.Name
// 	s.Genders = aux.Genders
// 	s.Expression = aux.Expression
// 	d := make(map[DemographyBucket]Demo)
// 	for i, a := range aux.DemoArray {
// 		d[DemographyBucket(i)] = a
// 	}
// 	s.Demography = d
// 	return nil
// }

// GetBucket implements persist.Persistable
func (s *Species) GetBucket() persist.Bucket {
	return persist.SpeciesBucket
}

// GetKey implements persist.Persistable
func (s *Species) GetKey() string {
	return s.Name
}

// Save implements persist.Persistable
func (s *Species) Save() error {
	return persist.DoSave(s)
}

// Delete implements persist.Persistable
func (s *Species) Delete() error {
	return persist.DoDelete(s)
}

// Fetch implements persist.Persistable
func (s *Species) Fetch() error {
	return persist.DoFetch(s)
}

func (s *Species) Load(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&s)
	if err != nil {
		return fmt.Errorf("could not load, %s", err)
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
