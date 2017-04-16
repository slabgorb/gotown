package townomatic_test

import (
	"fmt"
	"testing"

	. "github.com/slabgorb/townomatic"
)

var (
	fa       = []FertileAge{FertileAge{Gender(Male), 12, 65}, FertileAge{Gender(Female), 12, 55}}
	human, _ = NewSpecies("Human", []int{Male, Female}, fa, "{{.GivenName}} {{.FamilyName}}")
	adam     = &Being{Name: Name{GivenName: "Adam", FamilyName: "Man"}, Species: human}
	eve      = &Being{Name: Name{GivenName: "Eve", FamilyName: "Man"}, Species: human}
)

func TestName(t *testing.T) {
	fmt.Println(adam.Species.NameTemplate)
	if adam.String() != "Adam Man" {
		t.Errorf("Expected 'Adam Man' got %v", adam.String())
	}
}

func TestReproduction(t *testing.T) {
	adam.Reproduce(eve)

}
