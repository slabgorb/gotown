package being_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/persist"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

var beingFixtures = make(map[string]*Being)

var testSpecies = &species.Species{Name: "human"}
var testCulture = &culture.Culture{Name: "viking"}

type beingFixture struct {
	label string
	name  string
	age   int
	sex   string
}

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	defer persist.CloseTestDB()
	words.Seed()
	species.Seed()
	culture.Seed()
	if err := testCulture.Read(); err != nil {
		panic(err)
	}
	if err := testSpecies.Read(); err != nil {
		panic(err)
	}
	beingFixtureRaw := []beingFixture{
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
	id := 1
	for _, bf := range beingFixtureRaw {
		b := &Being{
			Gender:  inhabitants.Gender(bf.sex),
			Name:    inhabitants.NewName(bf.name),
			Species: testSpecies,
			Culture: testCulture,
			ID:      id,
		}
		id++
		b.SetAge(bf.age)
		beingFixtures[bf.label] = b

	}

	code := m.Run()
	os.Exit(code)
}
func TestName(t *testing.T) {
	expected := "Arnulf Arnulfson"
	being := &Being{Species: testSpecies, Gender: inhabitants.Male}
	words.SetRandomizer(random.NewMock())
	being.RandomizeName()
	if being.Sex() != inhabitants.Male {
		t.Errorf("Expected Male got %s", being.Sex())
	}
	if being.String() != expected {
		t.Errorf("Expected %s got %s", expected, being.String())
	}
}
func TestInheritedName(t *testing.T) {
	m := New(testSpecies, testCulture)
	m.Gender = inhabitants.Male
	m.RandomizeName()
	f := New(testSpecies, testCulture)
	f.Gender = inhabitants.Female
	f.RandomizeName()
	children, err := f.Reproduce(m)
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
	bf["adam"].Children = []int{bf["cain"].ID, bf["abel"].ID}
	bf["eve"].Children = []int{bf["cain"].ID, bf["abel"].ID}
	bf["cain"].Parents = []int{bf["adam"].ID, bf["eve"].ID}
	bf["abel"].Parents = []int{bf["adam"].ID, bf["eve"].ID}
	sibs, _ := bf["cain"].Siblings()
	if !sibs.Exists(bf["abel"]) {
		t.Errorf("expected cain to be abel's brother")
	}
	sibs, _ = bf["cain"].Siblings()
	if !sibs.Exists(bf["cain"]) {
		t.Errorf("expected cain to be abel's brother")
	}
	if !bf["abel"].IsSiblingOf(bf["cain"]) {
		t.Errorf("expected cain to be abel's brother")
	}
}
func TestDeath(t *testing.T) {
	adam := New(testSpecies, testCulture)
	if !adam.Alive() {
		t.Fail()
	}
	adam.Die()
	if adam.Alive() {
		t.Fail()
	}

}
