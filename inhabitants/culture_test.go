package inhabitants_test

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
)

func helperMockCulture(t *testing.T, name string) *Culture {
	data := helperLoadBytes(t, fmt.Sprintf("mock_culture_%s.json", name))
	c := &Culture{}
	err := json.Unmarshal(data, c)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestUnmarshal(t *testing.T) {
	c := helperMockCulture(t, "italian")
	if c.Name != "Italianate" {
		t.Error("did not get name")
	}
}

func TestMaritalStrategy(t *testing.T) {
	c := helperMockCulture(t, "italian")
	testCases := []struct {
		name     string
		a        *Being
		b        *Being
		expected bool
	}{
		{
			name:     "usual",
			a:        &Being{Species: mockSpecies, Age: 20, Sex: Male},
			b:        &Being{Species: mockSpecies, Age: 19, Sex: Female},
			expected: true,
		},
		{
			name:     "hetero only for this culture (yes, sorry)",
			a:        &Being{Species: mockSpecies, Age: 20, Sex: Male},
			b:        &Being{Species: mockSpecies, Age: 19, Sex: Male},
			expected: false,
		},
		{
			name:     "no bigamy",
			a:        &Being{Species: mockSpecies, Age: 20, Sex: Male, Spouses: []*Being{&Being{}}},
			b:        &Being{Species: mockSpecies, Age: 19, Sex: Female},
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
