package inhabitants_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
}

func TestName(t *testing.T) {
	species := NewSpecies("Northman", []Gender{Male, Female}, nil)
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
	species := NewSpecies("Northman", []Gender{Male, Female}, nil)
	culture := helperMockCulture(t, "viking")
	m := &Being{Species: species, Sex: Female, Culture: culture}
	m.Name = m.Culture.GetName(m)
	f := &Being{Species: species, Sex: Male, Culture: culture}
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
