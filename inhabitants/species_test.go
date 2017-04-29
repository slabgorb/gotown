package inhabitants_test

import (
	"math/rand"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	words "github.com/slabgorb/gotown/words"
)

var (
	female = NewSpeciesGender(words.NorseFemaleNamer, Matronymic, 12, 48)
	male   = NewSpeciesGender(words.NorseMaleNamer, Patronymic, 12, 65)
	s      = NewSpecies("Northman", map[Gender]*SpeciesGender{
		Female: female,
		Male:   male,
	})
)

func init() {
	rand.Seed(0)
}

func TestToString(t *testing.T) {
	if s.String() != "Northman" {
		t.Fail()
	}
}

func TestGenders(t *testing.T) {
	if s.Name != "Northman" {
		t.Fail()
	}
	b := &Being{Gender: Male, Species: s}
	name := male.NameStrategy(b)
	nameDisplay := name.Display
	expected := "Eimund Borrison"
	if nameDisplay != expected {
		t.Errorf("expected %s got %s", expected, nameDisplay)
	}
}

func TestRandomBeing(t *testing.T) {
	b := s.RandomBeing()
	if b.String() != "Bjornar Leidulfson" {
		t.Fail()
	}
	if b.Gender != Male {
		t.Fail()
	}
	b = s.RandomBeing()
	if b.String() != "Gjartrud Odveigdottir" {
		t.Fail()
	}
	if b.Gender != Female {
		t.Fail()
	}
}
