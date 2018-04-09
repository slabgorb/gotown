package being_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

type mockCulture struct {
	name string
}

func (m *mockCulture) RandomName(sex inhabitants.Gender) *inhabitants.Name {
	if sex == inhabitants.Female {
		return &inhabitants.Name{
			GivenName:  "Arnulf",
			FamilyName: "Arnulfdottir",
			Display:    "Arnulf Arnulfdottir",
		}
	}
	return &inhabitants.Name{
		GivenName:  "Arnulf",
		FamilyName: "Arnulfson",
		Display:    "Arnulf Arnulfson",
	}
}

func (m *mockCulture) MaritalCandidate(a, b inhabitants.Marriageable) bool {
	return (a.Alive() && b.Alive())
	//return (a.Alive() && b.Alive()) && (a.Sex() != b.Sex()) && (a.Unmarried() && b.Unmarried())
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

var beingFixtures = make(map[string]*Being)

func TestMain(m *testing.M) {
	type beingFixture struct {
		label string
		name  string
		age   int
		sex   string
	}

	var beingFixtureRaw = []beingFixture{
		{
			label: "adam",
			name:  "Adam Man",
			age:   35,
			sex:   "male",
		},
		{
			label: "eve",
			name:  "Eve Woman",
			age:   35,
			sex:   "female",
		},
		{
			label: "steve",
			name:  "Steve Guy",
			age:   35,
			sex:   "male",
		},
		{
			label: "cain",
			name:  "Cain Man",
			age:   17,
			sex:   "male",
		},
		{
			label: "abel",
			name:  "Abel Man",
			age:   18,
			sex:   "male",
		},
		{
			label: "martha",
			name:  "Martha Man",
			age:   19,
			sex:   "female",
		},
		{
			label: "abigail",
			name:  "Abigail Man",
			age:   25,
			sex:   "female",
		},
	}
	for _, bf := range beingFixtureRaw {
		b := &Being{
			Gender:     inhabitants.Gender(bf.sex),
			Name:       inhabitants.NewName(bf.name),
			Chronology: timeline.NewChronology(),
			Species:    &mockSpecies{},
		}
		b.SetAge(bf.age)
		beingFixtures[bf.label] = b

	}
	code := m.Run()
	os.Exit(code)
}
func TestName(t *testing.T) {
	species := &mockSpecies{}
	culture := &mockCulture{}
	expected := "Arnulf Arnulfson"
	being := &Being{Species: species, Gender: inhabitants.Male}
	words.SetRandomizer(random.NewMock())
	being.RandomizeName(culture)
	if being.Sex() != inhabitants.Male {
		t.Errorf("Expected Male got %s", being.Sex())
	}
	if being.String() != expected {
		t.Errorf("Expected %s got %s", expected, being.String())
	}
}
func TestInheritedName(t *testing.T) {
	species := &mockSpecies{}
	culture := &mockCulture{}
	m := &Being{Species: species, Gender: inhabitants.Female, Chronology: timeline.NewChronology()}
	m.RandomizeName(culture)
	f := &Being{Species: species, Gender: inhabitants.Male, Chronology: timeline.NewChronology()}
	f.RandomizeName(culture)
	children, err := f.Reproduce(m, culture)
	if err != nil {
		t.Errorf("%s", err)
	}

	child := children[0]
	if child.Sex() == inhabitants.Male {
		if child.Name.GetFamilyName() != f.Name.GetGivenName()+"son" {
			t.Errorf("expected %s got %s", f.Name.GetGivenName()+"son", child.Name.GetFamilyName())
		}
	} else {

		if child.Name.GetFamilyName() != m.Name.GetGivenName()+"dottir" {
			t.Errorf("expected %s got %s", m.Name.GetGivenName()+"dottir", child.Name.GetFamilyName())
		}
	}
}

func TestSiblings(t *testing.T) {
	bf := beingFixtures
	t.Log(bf)
	bf["adam"].Marry(bf["eve"])
	bf["adam"].Children = []*Being{bf["cain"], bf["abel"]}
	bf["eve"].Children = []*Being{bf["cain"], bf["abel"]}
	bf["cain"].Parents = []*Being{bf["adam"], bf["eve"]}
	bf["abel"].Parents = []*Being{bf["adam"], bf["eve"]}
	if bf["cain"].Siblings()[0] != bf["abel"] {
		t.Errorf("expected cain to be abel's brother")
	}
	if bf["abel"].Siblings()[0] != bf["cain"] {
		t.Errorf("expected cain to be abel's brother")
	}
	if !bf["abel"].IsSiblingOf(bf["cain"]) {
		t.Errorf("expected cain to be abel's brother")
	}
}
func TestDeath(t *testing.T) {
	adam := &Being{Chronology: timeline.NewChronology()}
	if !adam.Alive() {
		t.Fail()
	}
	adam.Die()
	if adam.Alive() {
		t.Fail()
	}
	if len(adam.Chronology.Events) > 1 {
		t.Fail()
	}

}
