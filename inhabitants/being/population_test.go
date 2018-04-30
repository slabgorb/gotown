package being_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/being"
)

func TestSerialization(t *testing.T) {
	p := NewPopulation([]int{})
	j, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	q := &Population{}
	err = json.Unmarshal(j, q)
	if err != nil {
		t.Error(err)
	}
}
func TestAging(t *testing.T) {
	p := NewPopulation([]int{})
	count := 10
	beings := make([]*Being, count)
	for i := 0; i < count; i++ {
		beings[i] = New(testSpecies, testCulture)
		p.Add(beings[i])
	}
	p.Age()
	ages := []int{}
	beings, err := p.Inhabitants()
	if err != nil {
		t.Fatalf("fatal: %s", err)
	}
	for _, b := range beings {
		ages = append(ages, b.GetAge())
	}
	sort.Ints(ages)
	expectedAges := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range expectedAges {
		if v != ages[i] {
			t.Errorf("Expected %d got %d", v, ages[i])
		}
	}
}

func BenchmarkMaritalCandidates(b *testing.B) {
	ids := []int{}
	for i := 0; i < 100; i++ {
		bg := New(testSpecies, testCulture)
		bg.Randomize()
		if err := bg.Save(); err != nil {
			b.Fatalf("could not save being:%s", err)
		}
		ids = append(ids, bg.ID)
	}
	p := NewPopulation(ids)
	b.Run("mc benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = p.MaritalCandidates(testCulture)
		}
	})
}

func TestReproductionCandidates(t *testing.T) {
	ids := []int{
		beingFixtures["adam"].ID,
		beingFixtures["eve"].ID,
		beingFixtures["martha"].ID,
	}
	beingFixtures["eve"].Marry(beingFixtures["adam"])
	p := NewPopulation(ids)
	candidates := p.ReproductionCandidates()
	if len(candidates) != 2 {
		t.Fail()
	}

}

func TestMaritalCandidates(t *testing.T) {
	ids := []int{
		beingFixtures["adam"].ID,
		beingFixtures["eve"].ID,
	}
	p := NewPopulation(ids)
	candidates, err := p.MaritalCandidates(testCulture)
	if err != nil {
		t.Error(err)
	}
	if len(candidates) < 1 {
		t.Log(candidates)
		t.Errorf("expected a candidate pair")
	}
	beings, _ := p.Inhabitants()
	a, b := candidates[0].Pair()
	if !(a == beings[0] || b == beings[0]) || !(a == beings[1] || b == beings[1]) {
		t.Errorf("expected adam and eve")
	}
	beings[1].Die()
	candidates, _ = p.MaritalCandidates(testCulture)
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and dead eve")
	}
	ids = []int{
		beingFixtures["adam"].ID,
		beingFixtures["steve"].ID,
	}
	p = NewPopulation(ids)
	candidates, _ = p.MaritalCandidates(testCulture)
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and steve")
	}
}

func TestAddAndRemove(t *testing.T) {
	b := &Being{ID: 0, Species: testSpecies}
	p := NewPopulation([]int{})
	p.Remove(b)
	if p.Exists(b) {
		t.Fail()
	}
	p.Add(b)
	if !p.Exists(b) {
		t.Fail()
	}
}

func TestAdamEve(t *testing.T) {
	pop := []int{}
	for _, v := range beingFixtures {
		pop = append(pop, v.ID)
	}
	p := NewPopulation(pop)
	beings, _ := p.Inhabitants()
	for _, b := range beings {
		fmt.Println(b)
	}
}

func TestMaritalStrategy(t *testing.T) {
	testCases := []struct {
		name     string
		a        *Being
		b        *Being
		ages     []int
		expected bool
	}{
		{
			name:     "usual",
			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Female},
			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male},
			expected: true,
		},
		{
			name:     "hetero only for this culture (yes, sorry)",
			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male},
			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Male},
			expected: false,
		},
		{
			name:     "no bigamy",
			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male, Spouses: []int{0}},
			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Female},
			expected: false,
		},
	}
	i := 1
	for _, tc := range testCases {
		tc.a.ID = i
		tc.b.ID = i + 1
		i += 2
		p := NewPopulation([]int{tc.a.ID, tc.b.ID})
		candidates, _ := p.MaritalCandidates(testCulture)
		actual := len(candidates) > 0
		if tc.expected != actual {
			t.Errorf("%s expected %t got %t", tc.name, tc.expected, actual)
		}
	}
}
