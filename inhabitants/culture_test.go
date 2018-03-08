package inhabitants_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/timeline"
)

func TestUnmarshal(t *testing.T) {
	c := helperMockCulture(t, "italian")
	if c.Name != "Italianate" {
		t.Error("did not get name")
	}
}

func TestMarshal(t *testing.T) {
	c := helperMockCulture(t, "viking")
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}
	culture := &inhabitants.Culture{}
	err = json.Unmarshal(bytes, culture)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", culture)
	t.Logf(string(bytes))
	if culture.Name != "viking" {
		t.Fail()
	}
	expectedMS := []string{"heterosexual", "monogamous"}
	if !reflect.DeepEqual(expectedMS, culture.MaritalStrategies) {
		t.Fail()
	}
	expectedNS := map[inhabitants.Gender]string{"male": "patronymic", "female": "matronymic"}
	if !reflect.DeepEqual(expectedNS, culture.NameStrategies) {
		t.Errorf("expected %#v got %#v", expectedNS, culture.NameStrategies)
	}
	if culture.Namers[inhabitants.Female] == nil {
		t.Errorf("did not get female namer")

	}
	if culture.Namers[inhabitants.Female].Matronymic() != "dottir" {
		t.Errorf("expected dottir got %s", culture.Namers[inhabitants.Female].Matronymic())
	}
}

func TestMaritalStrategy(t *testing.T) {
	c := helperMockCulture(t, "italian")
	mockSpecies := helperMockSpecies(t)
	testCases := []struct {
		name     string
		a        *Being
		b        *Being
		expected bool
	}{
		{
			name:     "usual",
			a:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 20}, Sex: Male},
			b:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 19}, Sex: Female},
			expected: true,
		},
		{
			name:     "hetero only for this culture (yes, sorry)",
			a:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 20}, Sex: Male},
			b:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 19}, Sex: Male},
			expected: false,
		},
		{
			name:     "no bigamy",
			a:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 20}, Sex: Male, Spouses: []*Being{&Being{}}},
			b:        &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: 29}, Sex: Female},
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
