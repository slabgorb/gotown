package inhabitants_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/random"
	words "github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
}

var (
	female = NewSpeciesGender(words.NorseFemaleNamer, NameStrategies["matronymic"], 12, 48)
	male   = NewSpeciesGender(words.NorseMaleNamer, NameStrategies["patronymic"], 12, 65)
	s      = NewSpecies("Northman", map[Gender]*SpeciesGender{
		Female: female,
		Male:   male,
	}, nil)
)

func TestToString(t *testing.T) {
	if s.String() != "Northman" {
		t.Fail()
	}
}

func TestGenders(t *testing.T) {
	if s.Name != "Northman" {
		t.Fail()
	}
	b := &Being{Sex: Male, Species: s}
	name := male.NameStrategy(b)
	nameDisplay := name.Display
	expected := "Hauk Haukson"
	if nameDisplay != expected {
		t.Errorf("expected %s got %s", expected, nameDisplay)
	}
}

func TestRandomAge(t *testing.T) {
	testCases := []struct {
		in, out int
	}{
		{0, 7},
		{30, 16},
		{40, 22},
		{99, 85},
	}
	for _, tc := range testCases {
		age := male.RandomAge(tc.in)
		if age != tc.out {
			t.Errorf("expected %d got %d", tc.out, age)
		}
	}
}

func TestRandomBeing(t *testing.T) {
	b := s.RandomBeing()
	expected := "Hauk Haukson"
	if b.String() != expected {
		t.Errorf("expected %s got %s", expected, b.String())
	}
	if b.Sex != Male {
		t.Errorf("Wrong gender, got %s", b.Sex)
	}

}
