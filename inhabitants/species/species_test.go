package species_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
)

var demog = map[int]Demo{
	inhabitants.Child:      {MaxAge: 14, CumulativePercent: 29},
	inhabitants.Teenager:   {MaxAge: 18, CumulativePercent: 36},
	inhabitants.YoungAdult: {MaxAge: 26, CumulativePercent: 50},
	inhabitants.EarlyAdult: {MaxAge: 31, CumulativePercent: 58},
	inhabitants.Adult:      {MaxAge: 41, CumulativePercent: 72},
	inhabitants.MiddleAge:  {MaxAge: 51, CumulativePercent: 74},
	inhabitants.Senior:     {MaxAge: 61, CumulativePercent: 93},
	inhabitants.Elderly:    {MaxAge: 71, CumulativePercent: 98},
	inhabitants.Ancient:    {MaxAge: 100, CumulativePercent: 100},
}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapped(m))
}

func testMainWrapped(m *testing.M) int {
	persist.OpenTestDB()
	defer persist.CloseTestDB()
	Seed()
	SetRandomizer(random.NewMock())
	return m.Run()

}
func TestSeed(t *testing.T) {
	list, err := List()
	if err != nil {
		panic(err)
	}

	found := false
	for _, v := range list {
		if v == "human" {
			found = true
		}
	}
	t.Log(list)
	if !found {
		t.Fatal("human not seeded")
	}

	w := &Species{Name: "elf"}
	if err := persist.ReadByName(w.Name, "Species", w); err != nil {
		t.Fatal(err)
	}

}

// func TestGenders(t *testing.T) {
// 	mockSpecies := helperMockSpecies(t)
// 	if mockSpecies.Name != "Human" {
// 		t.Fail()
// 	}
// 	culture := helperMockCulture(t, "viking")
// 	b := &Being{Sex: Male, Species: mockSpecies, Culture: culture}
// 	name := culture.GetName(b)
// 	nameDisplay := name.Display
// 	expected := "Arnulf Arnulfson"
// 	if nameDisplay != expected {
// 		t.Errorf("expected %s got %s", expected, nameDisplay)
// 	}
// }

// func TestRandomAge(t *testing.T) {
// 	mockSpecies := helperMockSpecies(t)
// 	testCases := []struct {
// 		in, out int
// 	}{
// 		{0, 7},
// 		{30, 16},
// 		{40, 22},
// 		{99, 85},
// 	}
// 	for _, tc := range testCases {
// 		age := mockSpecies.RandomAge(tc.in)
// 		if age != tc.out {
// 			t.Errorf("expected %d got %d", tc.out, age)
// 		}
// 	}
// }

func TestSpecies_RandomAge(t *testing.T) {

	tests := []struct {
		name string
		args int
		max  int
		min  int
	}{
		{
			name: "teenager",
			args: 1,
			max:  18,
			min:  15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Species{
				Demography: demog,
			}
			if got := s.RandomAge(tt.args); got >= tt.max || got <= tt.min {
				t.Errorf("Species.RandomAge() = %v, max %v min %v", got, tt.max, tt.min)
			}
		})
	}
}
