package inhabitants_test

import (
	"math/rand"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	words "github.com/slabgorb/gotown/words"
)

var nameTests = []struct {
	pattern  string
	expected string
}{
	{"{{.GivenName}} {{.FamilyName}}", "Finnbjorn Einarson"},
	{"{{.GivenName}} {{.FamilyName}} {{.OtherNames}}", "Gunnbjorn Kolfinnson"},
	{"{{.FamilyName}}", "Bendikson"},
}

func init() {
	rand.Seed(0)
}

func TestName(t *testing.T) {
	for _, nt := range nameTests {
		namer := words.NewNamer([]string{nt.pattern}, words.NorseMaleNameWords)
		speciesGender := NewSpeciesGender(namer, Patronymic, 12, 65)
		species := NewSpecies("Northman", map[Gender]*SpeciesGender{Male: speciesGender})
		being := &Being{Species: species}
		being.Randomize()
		if being.Sex != Male {
			t.Errorf("Expected Male got %s", being.Sex)
		}
		if being.String() != nt.expected {
			t.Errorf("Expected %s got %s", nt.expected, being.String())
		}
	}
}

func TestInheritedName(t *testing.T) {
	rand.Seed(6)
	male := NewSpeciesGender(words.NorseMaleNamer, Patronymic, 12, 65)
	female := NewSpeciesGender(words.NorseFemaleNamer, Matronymic, 12, 50)
	species := NewSpecies("Northman", map[Gender]*SpeciesGender{Male: male, Female: female})
	m := &Being{Species: species, Sex: Female}
	m.Name = female.NameStrategy(m)
	f := &Being{Species: species, Sex: Male}
	f.Name = male.NameStrategy(m)
	//runtime.Breakpoint()
	children, err := f.Reproduce(m)
	if err != nil {
		t.Errorf("%s", err)
	}

	child := children[0]

	if child.Name.FamilyName != m.Name.GivenName+"dottir" {
		t.Errorf("expected %s got %s", m.Name.GivenName+"dottir", child.Name.FamilyName)
	}
	//t.Errorf("%v", children)

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
