package culture_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/being"
	. "github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	words.Seed()
	Seed()
	code := m.Run()
	persist.CloseTestDB()
	os.Exit(code)
}

type mockSpecies struct {
	name string
}

func (m *mockSpecies) RandomAge(slot int) int {
	return slot * 10
}

func (m *mockSpecies) MaxAge(slot int) int {
	return slot * 12
}

func (m *mockSpecies) GetGenders() []inhabitants.Gender {
	return []inhabitants.Gender{inhabitants.Male, inhabitants.Female}
}

func (m *mockSpecies) Expression() inhabitants.Expresser {
	panic("not implemented")
}

type MarriageableStub struct {
	alive   bool
	married bool
	sex     inhabitants.Gender
	age     int
}

func (ms *MarriageableStub) Alive() bool {
	return ms.alive
}

func (ms *MarriageableStub) Unmarried() bool {
	return !ms.married
}

func (ms *MarriageableStub) Sex() inhabitants.Gender {
	return ms.sex
}

func (ms *MarriageableStub) IsCloseRelativeOf(inhabitants.Marriageable) bool {
	panic("not implemented")
}

func (ms *MarriageableStub) Age() int {
	return ms.age
}

func TestSeed(t *testing.T) {
	list, err := List()
	if err != nil {
		panic(err)
	}

	found := false
	for _, v := range list {
		if v == "italianate" {
			found = true
		}
	}
	t.Log(list)
	if !found {
		t.Fatal("italianate not seeded")
	}

	w := &Culture{Name: "italianate"}
	if err := w.Read(); err != nil {
		t.Fatal(err)
	}
	t.Log(w.NamerNames)
	namers := w.GetNamers()
	_, ok := namers[inhabitants.Male]
	if !ok {
		t.Error("did not load male names")
	}

}

func TestMaritalStrategy(t *testing.T) {
	c := Culture{Name: "italianate"}
	if err := c.Read(); err != nil {
		t.Fail()
	}
	t.Log(c.MaritalStrategies)
	t.Log(c)
	ms := &mockSpecies{}
	testCases := []struct {
		name     string
		a        *being.Being
		b        *being.Being
		expected bool
	}{
		{
			name:     "usual",
			a:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 20}, Gender: inhabitants.Male},
			b:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 19}, Gender: inhabitants.Female},
			expected: true,
		},
		{
			name:     "hetero only for this culture (yes, sorry)",
			a:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 20}, Gender: inhabitants.Male},
			b:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 19}, Gender: inhabitants.Male},
			expected: false,
		},
		{
			name:     "no bigamy",
			a:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 20}, Gender: inhabitants.Male, Spouses: []*being.Being{&being.Being{}}},
			b:        &being.Being{Species: ms, Chronology: &timeline.Chronology{CurrentYear: 29}, Gender: inhabitants.Female},
			expected: false,
		},
	}
	for _, tc := range testCases {
		actual := c.MaritalCandidate(tc.a, tc.b)
		if tc.expected != actual {
			t.Errorf("%s expected %t got %t", tc.name, tc.expected, actual)
		}
	}
}
