package inhabitants

import (
	"fmt"

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
	demog string
}

type demo struct {
	max int
	pct int
}
type demography []demo

var demographies = map[string]demography{
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

func NewSpeciesGender(namer *words.Namer, ns NameStrategy, start, end int) *SpeciesGender {
	words.SetRandomizer(randomizer)
	return &SpeciesGender{Fertility: Fertility{start, end}, Namer: namer, NameStrategy: ns, demog: "medieval"}
}

func (s *SpeciesGender) RandomName() {
	s.Name()
}

func (s *SpeciesGender) RandomAge(slot int) int {
	d, ok := demographies[s.demog]
	if !ok {
		return 0
	}
	if slot == -1 {
		slot = randomizer.Intn(101)
	}
	min := 0
	for _, dmo := range d {
		if dmo.pct >= slot {
			return randomizer.Intn(dmo.max-min) + min
		}
		min = dmo.max
	}
	return 0
}

type Fertility struct {
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name          string                       `json:"name"`
	Genders       map[Gender]*SpeciesGender    `json:"-"`
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

func (s *Species) RandomBeing() *Being {
	b := &Being{Species: s}
	b.Randomize()
	return b
}
