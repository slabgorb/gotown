package inhabitants_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/random"
	words "github.com/slabgorb/gotown/words"
)

var nameTests = []struct {
	pattern  string
	expected string
}{
	{"{{.GivenName}} {{.FamilyName}}", "Leidolv Thorbjornson"},
	{"{{.GivenName}} {{.FamilyName}} {{.OtherNames}}", "Oddleif Borgulvson"},
	{"{{.FamilyName}}", "Bendikson"},
}

func TestName(t *testing.T) {
	for _, nt := range nameTests {
		namer := words.NewNamer([]string{nt.pattern}, words.NorseMaleNameWords)
		speciesGender := NewSpeciesGender(namer, Patronymic, 12, 65)
		species := NewSpecies("Northman", map[Gender]*SpeciesGender{Male: speciesGender})
		species.SetRandomizer(random.NewMock())
		being := &Being{Species: species}
		being.SetRandomizer(random.NewMock())
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
	male := NewSpeciesGender(words.NorseMaleNamer, Patronymic, 12, 65)
	female := NewSpeciesGender(words.NorseFemaleNamer, Matronymic, 12, 50)
	species := NewSpecies("Northman", map[Gender]*SpeciesGender{Male: male, Female: female})
	species.SetRandomizer(random.NewMock())
	m := &Being{Species: species, Sex: Female}
	m.SetRandomizer(random.NewMock())
	m.Name = female.NameStrategy(m)
	f := &Being{Species: species, Sex: Male}
	f.SetRandomizer(random.NewMock())
	f.Name = male.NameStrategy(m)
	//runtime.Breakpoint()
	children, err := f.Reproduce(m)
	if err != nil {
		t.Errorf("%s", err)
	}

	child := children[0]

	if child.Name.FamilyName != f.Name.GivenName+"son" {
		t.Errorf("expected %s got %s", f.Name.GivenName+"son", child.Name.FamilyName)
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
