package gotown_test

import (
	"fmt"
	"testing"
	"text/template"

	. "github.com/slabgorb/gotown"
)

var nameTests = []struct {
	pattern  string
	name     Name
	expected string
}{
	{"{{.GivenName}} {{.FamilyName}}", NewName("Adam", "Man"), "Adam Man"},
	{"{{.GivenName}} {{.FamilyName}} {{.OtherNames}}", NewName("Adam", "Man", "The"), "Adam Man [The]"},
	{"{{.FamilyName}}", NewName("Adam", "Man"), "Man"},
}

func TestName(t *testing.T) {
	for i, nt := range nameTests {
		nameTemplate, _ := template.New(fmt.Sprintf("%v_nameTemplate", i)).Parse(nt.pattern)
		actual := nt.name.Patterned(nameTemplate)
		if actual != nt.expected {
			t.Errorf("Name.Patterned() expected %s got %s", nt.expected, actual)
		}
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
