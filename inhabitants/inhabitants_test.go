package inhabitants_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/slabgorb/gotown/words"

	. "github.com/slabgorb/gotown/inhabitants"
)

var (
	female      = NewSpeciesGender(words.NorseFemaleNamer, NameStrategies["matronymic"], 12, 48)
	male        = NewSpeciesGender(words.NorseMaleNamer, NameStrategies["patronymic"], 12, 65)
	mockSpecies = NewSpecies("Northman", map[Gender]*SpeciesGender{
		Female: female,
		Male:   male,
	}, nil)
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

func beingFixtures(t *testing.T) map[string]*Being {
	var v []beingFixture
	beings := make(map[string]*Being)
	data := helperLoadBytes(t, "being_fixtures.json")
	err := json.Unmarshal(data, v)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range v {
		name := NewName(f.Name)
		beings[f.Label] = &Being{Name: name, Age: f.Age}
	}
	return beings
}
