package gotown_test

import (
	"testing"

	. "github.com/slabgorb/gotown"
	words "github.com/slabgorb/gotown/words"
)

var nameTests = []struct {
	pattern  string
	expected string
}{
	{"{{.GivenName}} {{.FamilyName}}", "Something"},
	{"{{.GivenName}} {{.FamilyName}} {{.OtherNames}}", "Something"},
	{"{{.FamilyName}}", "Something"},
}

func TestName(t *testing.T) {
	for _, nt := range nameTests {
		namer := words.NewNamer([]string{nt.pattern}, words.NorseMaleNameWords)
		speciesGender := NewSpeciesGender(namer, Patronymic, 12, 65)
		species := NewSpecies("Northman", map[Gender]*SpeciesGender{Male: speciesGender})
		being := &Being{Species: species}
		being.Randomize()
		if being.Gender != Male {
			t.Errorf("Expected Male got %s", being.Gender)
		}
		// if being.String() != nt.expected {
		// 	t.Errorf("Expected %s got %s", nt.expected, being.String())
		// }
	}
}

func TestDeath(t *testing.T) {
	adam := &Being{}
	if !adam.Alive() {
		t.Fail()
	}
	adam.Die()
	if adam.Alive() {
		t.Fail()
	}

}
