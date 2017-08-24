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

func TestToString(t *testing.T) {
	if s.String() != "Northman" {
		t.Fail()
	}
}

func TestGenders(t *testing.T) {
	rand.Seed(0)
	if s.Name != "Northman" {
		t.Fail()
	}
	b := &Being{Sex: Male, Species: s}
	name := male.NameStrategy(b)
	nameDisplay := name.Display
	expected := "Finnbjorn Finnbjornson"
	if nameDisplay != expected {
		t.Errorf("expected %s got %s", expected, nameDisplay)
	}
}

func TestRandomBeing(t *testing.T) {
	rand.Seed(0)
	b := s.RandomBeing()
	expected := "Annfrid Solunndottir"
	if b.String() != expected {
		t.Errorf("expected %s got %s", expected, b.String())
	}
	if b.Sex != Female {
		t.Errorf("Wrong gender, got %s", b.Sex)
	}
	b = s.RandomBeing()
	expected = "Gunnbjorn Kolfinnson"
	if b.String() != expected {

		t.Errorf("expected %s got %s", expected, b.String())
	}

	if b.Sex != Male {
		t.Errorf("Wrong gender, got %s", b.Sex)
	}
}
