package gotown_test

import (
	"testing"

	. "github.com/slabgorb/gotown"
	words "github.com/slabgorb/gotown/words"
)

func TestToString(t *testing.T) {
	species := Species{Name: "Vampire"}

	if species.String() != "Vampire" {
		t.Fail()
	}
}

func TestGenders(t *testing.T) {
	female := NewSpeciesGender(words.NorseFemaleNamer, Matronymic, 12, 48)
	male := NewSpeciesGender(words.NorseMaleNamer, Patronymic, 12, 65)
	s := NewSpecies("Northman", map[Gender]*SpeciesGender{
		Female: female,
		Male:   male,
	})
	if s.Name != "Northman" {
		t.Fail()
	}
	b := &Being{Gender: Male, Species: s}
	name := male.NameStrategy(b).Display
	if name != "Folkvald Gunnerson" {
		t.Errorf("expected Folkvald Gunnerson got %s", name)
	}
}
