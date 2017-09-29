package inhabitants_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/genetics"
)

var (
	mockSpecies = NewSpecies("Human", []Gender{Male, Female}, nil)
)

type beingFixture struct {
	Label   string `json:"label"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Culture string `json:"culture"`
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func helperMockCulture(t *testing.T, name string) *Culture {
	data := helperLoadBytes(t, fmt.Sprintf("mock_culture_%s.json", name))
	c := &Culture{}
	err := json.Unmarshal(data, c)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func beingFixtures(t *testing.T, cultureName string) map[string]*Being {
	var v []beingFixture
	culture := helperMockCulture(t, cultureName)
	beings := make(map[string]*Being)
	data := helperLoadBytes(t, "being_fixtures.json")
	err := json.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range v {
		name := NewName(f.Name)
		beings[f.Label] = &Being{Name: name, Age: f.Age, Sex: Gender(f.Sex), Culture: culture, Chromosome: genetics.RandomChromosome(30)}
	}
	return beings
}
