package inhabitants

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

// go:generate stringer -type=Gender
type Gender int

const (
	Asexual Gender = iota
	Male
	Female
)

type NameStrategy func(b *Being) *Name

var (
	Matrilineal NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		if b.Mother() != nil {
			name.FamilyName = b.Mother().FamilyName
			return name
		}
		name.FamilyName = namer.GivenName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}

	Patrilineal NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		if b.Father() != nil {
			name.FamilyName = b.Father().FamilyName
			return name
		}
		name.FamilyName = namer.GivenName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}

	Matronymic NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		if b.Mother() != nil {
			name.FamilyName = b.Mother().GivenName + namer.Matronymic()
			return name
		}
		name.FamilyName = namer.GivenName() + namer.Matronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}

	Patronymic NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		if b.Father() != nil {
			fmt.Println(b.Father())
			fmt.Println(namer.Patronymic())
			name.FamilyName = b.Father().GivenName + namer.Patronymic()
			return name
		}
		name.FamilyName = namer.GivenName() + namer.Patronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}

	OneName NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}
)

type SpeciesGender struct {
	Fertility
	*words.Namer
	NameStrategy
	randomizer random.Generator
}

func NewSpeciesGender(namer *words.Namer, ns NameStrategy, start, end int) *SpeciesGender {
	return &SpeciesGender{Fertility: Fertility{start, end}, Namer: namer, NameStrategy: ns}
}

func (s *SpeciesGender) RandomName() {
	s.Name()
}

func (s *SpeciesGender) SetRandomizer(g random.Generator) {
	s.randomizer = g
}

func (s *SpeciesGender) SetDefaultRandomizer() {
	if s.randomizer == nil {
		s.SetRandomizer(rand.New(rand.NewSource(time.Now().UTC().UnixNano())))
	}
}

func (s *SpeciesGender) RandomAge() int {
	s.SetDefaultRandomizer()
	return s.randomizer.Intn(s.Fertility.End * 3)
}

type Fertility struct {
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name          string                       `json:"name"`
	Genders       map[Gender]*SpeciesGender    `json:"genders"`
	MultipleBirth func(g random.Generator) int `json:"-"`
	Expression    genetics.Expression          `json:"-"`
	randomizer    random.Generator
}

// NewSpecies creates and initializes a *Species
func NewSpecies(name string, genders map[Gender]*SpeciesGender) *Species {
	return &Species{
		Name:    name,
		Genders: genders,
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

func (s *Species) GetGenders() map[Gender]*SpeciesGender {
	return s.Genders
}

func (s *Species) SetRandomizer(g random.Generator) {
	s.randomizer = g
}

func (s *Species) SetDefaultRandomizer() {
	if s.randomizer == nil {
		s.SetRandomizer(rand.New(rand.NewSource(time.Now().UTC().UnixNano())))
	}
}

func (s *Species) RandomBeing() *Being {
	b := &Being{Species: s}
	b.SetRandomizer(s.randomizer)
	b.Randomize()
	return b
}
