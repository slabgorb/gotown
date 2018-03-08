package inhabitants_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/random"
)

func init() {
	SetRandomizer(random.NewMock())
}

func TestToString(t *testing.T) {
	mockSpecies := helperMockSpecies(t)
	if mockSpecies.String() != "Human" {
		t.Error(mockSpecies.String())
	}
}

func TestGenders(t *testing.T) {
	mockSpecies := helperMockSpecies(t)
	if mockSpecies.Name != "Human" {
		t.Fail()
	}
	culture := helperMockCulture(t, "viking")
	b := &Being{Sex: Male, Species: mockSpecies, Culture: culture}
	name := culture.GetName(b)
	nameDisplay := name.Display
	expected := "Arnulf Arnulfson"
	if nameDisplay != expected {
		t.Errorf("expected %s got %s", expected, nameDisplay)
	}
}

func TestRandomAge(t *testing.T) {
	mockSpecies := helperMockSpecies(t)
	testCases := []struct {
		in, out int
	}{
		{0, 7},
		{30, 16},
		{40, 22},
		{99, 85},
	}
	for _, tc := range testCases {
		age := mockSpecies.RandomAge(tc.in)
		if age != tc.out {
			t.Errorf("expected %d got %d", tc.out, age)
		}
	}
}
