package being_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/words"
)

type stubLogger struct {
}

func (s stubLogger) Info(format string, v ...interface{})  {}
func (s stubLogger) Debug(format string, v ...interface{}) {}
func (s stubLogger) Error(format string, v ...interface{}) {}
func (s stubLogger) SetOutput(out io.Writer)               {}

type beingFixture map[string]*Being

func (bf beingFixture) iterate() map[string]*Being {
	return bf
}

func (bf beingFixture) beings() []*Being {
	bgs := []*Being{}
	for _, b := range bf.iterate() {
		bgs = append(bgs, b)
	}
	return bgs
}

var beingFixtures = make(beingFixture)

var testSpecies = &species.Species{Name: "human"}
var testCulture = &culture.Culture{Name: "italianate"}

type rawBeingFixture struct {
	label string
	name  string
	age   int
	sex   string
}

func testMainWrapped(m *testing.M) int {
	persist.OpenTestDB()

	words.Seed()
	species.Seed()
	culture.Seed()
	if err := persist.ReadByName(testCulture.Name, "Culture", testCulture); err != nil {
		list, _ := culture.List()
		panic(fmt.Sprintf("could not load test culture %s, have %#v: %s", testCulture.Name, list, err))
	}
	if err := persist.ReadByName(testSpecies.Name, "Species", testSpecies); err != nil {
		panic(fmt.Sprintf("could not load test species: %s", err))
	}
	beingFixtureRaw := []rawBeingFixture{
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
			Gender:  inhabitants.Gender(bf.sex),
			Name:    NewName(bf.name),
			Species: testSpecies,
			Culture: testCulture,
		}
		b.Chromosome = genetics.RandomChromosome(30)
		b.SetAge(bf.age)
		if err := b.Save(); err != nil {
			panic(err)
		}
		beingFixtures[bf.label] = b

	}

	defer persist.CloseTestDB()
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapped(m))
}

func TestName(t *testing.T) {
	expected := "Leone Giovanelli"
	being := &Being{Species: testSpecies, Culture: testCulture, Gender: inhabitants.Male}
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
	m := New(testSpecies, testCulture, stubLogger{})
	m.Gender = inhabitants.Male
	m.RandomizeName()
	if err := m.Save(); err != nil {
		t.Fatal(err)
	}
	f := New(testSpecies, testCulture, stubLogger{})
	f.Gender = inhabitants.Female
	f.RandomizeName()
	if err := f.Save(); err != nil {
		t.Fatal(err)
	}
	child, err := f.Reproduce(m)
	if err != nil {
		t.Errorf("%s", err)
	}
	if child.Name.GetFamilyName() != f.Name.GetFamilyName() {
		t.Errorf("expected %s got %s", f.Name.GetFamilyName(), child.Name.GetFamilyName())
	}
}

func TestSiblings(t *testing.T) {
	bf := beingFixtures
	bf["adam"].Marry(bf["eve"])
	bf["adam"].Children = []string{bf["cain"].ID, bf["abel"].ID}
	bf["eve"].Children = []string{bf["cain"].ID, bf["abel"].ID}
	bf["cain"].Parents = []string{bf["adam"].ID, bf["eve"].ID}
	bf["abel"].Parents = []string{bf["adam"].ID, bf["eve"].ID}
	for _, b := range bf {
		b.Save()
	}
	sibs, _ := bf["cain"].Siblings()
	if !sibs.Exists(bf["abel"]) {
		t.Errorf("expected cain to be abel's brother")
	}
	sibs, _ = bf["abel"].Siblings()
	if !sibs.Exists(bf["cain"]) {
		t.Errorf("expected abel to be cain's brother")
	}
	if !bf["abel"].IsSiblingOf(bf["cain"].ID) {
		t.Errorf("expected cain to be abel's brother")
	}
	if !bf["cain"].IsSiblingOf(bf["abel"].ID) {
		t.Errorf("expected abel to be cain's brother")
	}
}

func TestParents(t *testing.T) {
	bf := beingFixtures
	bf["adam"].Marry(bf["eve"])
	bf["adam"].Children = []string{bf["cain"].ID, bf["abel"].ID}
	bf["eve"].Children = []string{bf["cain"].ID, bf["abel"].ID}
	bf["cain"].Parents = []string{bf["adam"].ID, bf["eve"].ID}
	bf["abel"].Parents = []string{bf["adam"].ID, bf["eve"].ID}
	beings, err := bf["abel"].GetParents()
	if err != nil {
		t.Fail()
	}
	if beings[0].ID != bf["adam"].ID && beings[1].ID != bf["adam"].ID {
		t.Errorf("expected parent to be adam")
	}
	if beings[0].ID != bf["eve"].ID && beings[1].ID != bf["eve"].ID {
		t.Errorf("expected parent to be eve")
	}
	mother, err := bf["abel"].Mother()
	if err != nil {
		t.Fatal(err)
	}
	if mother.ID != bf["eve"].ID {
		t.Errorf("expected eve to be abel's mom")
	}
	father, err := bf["abel"].Father()
	if err != nil {
		t.Fatal(err)
	}
	if father.ID != bf["adam"].ID {
		t.Errorf("expected eve to be abel's mom")
	}
}
func TestDeath(t *testing.T) {
	adam := New(testSpecies, testCulture, stubLogger{})
	if !adam.Alive() {
		t.Fail()
	}
	adam.Die()
	if adam.Alive() {
		t.Fail()
	}

}

func TestRepro(t *testing.T) {
	bf := beingFixtures
	adam := bf["adam"]
	ac := adam.Chromosome
	eve := bf["eve"]
	ec := eve.Chromosome
	fmt.Println(ac, ec)
	adam.Marry(eve)
	child, err := eve.Reproduce(adam)
	if err != nil {
		t.Errorf("failed repro: %s", err)
	}
	if child.Age != 0 {
		t.Errorf("expected 0 for age, got %v", child.Age)
	}

	if child.ID == "" {
		t.Errorf("expected child to have a non-empty id")
	}

	for _, g := range child.Chromosome.Genes {
		fa := adam.Chromosome.Find(g)
		fe := eve.Chromosome.Find(g)
		if fa < 0 && fe < 0 {
			t.Errorf("expected all genes of the child to be members of the genes of the parents")
		}
	}

}
