package inhabitants_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
}

func TestName(t *testing.T) {
	species := helperMockSpecies(t)
	culture := helperMockCulture(t, "viking")
	if culture == nil {
		t.Error("culture not loaded")
	}
	expected := "Hauk Haukson"
	being := &Being{Species: species, Culture: culture, Sex: Male}
	words.SetRandomizer(random.NewMock())
	being.RandomizeName()
	if being.Sex != Male {
		t.Errorf("Expected Male got %s", being.Sex)
	}
	if being.String() != expected {
		t.Errorf("Expected %s got %s", expected, being.String())
	}
}

func TestInheritedName(t *testing.T) {
	species := helperMockSpecies(t)
	culture := helperMockCulture(t, "viking")
	m := &Being{Species: species, Sex: Female, Culture: culture, Chronology: timeline.NewChronology()}
	m.Name = m.Culture.GetName(m)
	f := &Being{Species: species, Sex: Male, Culture: culture, Chronology: timeline.NewChronology()}
	f.Name = f.Culture.GetName(f)
	//runtime.Breakpoint()
	children, err := f.Reproduce(m)
	if err != nil {
		t.Errorf("%s", err)
	}

	child := children[0]
	if child.Sex == Male {
		if child.Name.FamilyName != f.Name.GivenName+"son" {
			t.Errorf("expected %s got %s", f.Name.GivenName+"son", child.Name.FamilyName)
		}
	} else {

		if child.Name.FamilyName != m.Name.GivenName+"dottir" {
			t.Errorf("expected %s got %s", m.Name.GivenName+"dottir", child.Name.FamilyName)
		}
	}
	//t.Errorf("%v", children)

}

func TestSiblings(t *testing.T) {
	bf := beingFixtures(t, "italian")
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
