package gotown

import (
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
		namer := b.Species.Genders[b.Gender].Namer
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
		namer := b.Species.Genders[b.Gender].Namer
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
		namer := b.Species.Genders[b.Gender].Namer
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
		namer := b.Species.Genders[b.Gender].Namer
		name := &Name{GivenName: namer.GivenName()}
		if b.Father() != nil {
			name.FamilyName = b.Father().GivenName + namer.Patronymic()
			return name
		}
		name.FamilyName = namer.GivenName() + namer.Patronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	}
	OneName NameStrategy = func(b *Being) *Name {
		namer := b.Species.Genders[b.Gender].Namer
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
}

func NewSpeciesGender(namer *words.Namer, ns NameStrategy, start, end int) *SpeciesGender {
	return &SpeciesGender{Fertility{start, end}, namer, ns}
}

func (s *SpeciesGender) RandomName() {
	s.Name()
}

type Fertility struct {
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name    string
	Genders map[Gender]*SpeciesGender
}

// NewSpecies creates and initializes a *Species
func NewSpecies(name string, genders map[Gender]*SpeciesGender) *Species {
	return &Species{
		Name:    name,
		Genders: genders,
	}
}

func (s *Species) String() string {
	return s.Name
}
