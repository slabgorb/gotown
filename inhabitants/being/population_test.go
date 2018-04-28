package being_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/timeline"
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
	culture := &mockCulture{}
	chronology := timeline.NewChronology()
	beings := []inhabitants.Populatable{
		beingFixtures["adam"],
		beingFixtures["eve"],
		beingFixtures["martha"],
	}
	beingFixtures["eve"].Marry(beingFixtures["adam"])
	p := NewPopulation(beings, chronology, culture)
	candidates := p.ReproductionCandidates()
	if len(candidates) != 2 {
		t.Fail()
	}

}

func TestMaritalCandidates(t *testing.T) {
	culture := &mockCulture{name: "italian"}
	chronology := timeline.NewChronology()
	beings := []inhabitants.Populatable{
		beingFixtures["adam"],
		beingFixtures["eve"],
	}
	p := NewPopulation(beings, chronology, culture)
	candidates, err := p.MaritalCandidates(culture)
	if err != nil {
		t.Error(err)
	}
	if len(candidates) < 1 {
		t.Log(candidates)
		t.Errorf("expected a candidate pair")
	}
	a, b := candidates[0].Pair()
	if !(a == beings[0] || b == beings[0]) || !(a == beings[1] || b == beings[1]) {
		t.Errorf("expected adam and eve")
	}
	beings[1].Die()
	candidates, _ = p.MaritalCandidates(culture)
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and dead eve")
	}
	beings = []inhabitants.Populatable{
		beingFixtures["adam"],
		beingFixtures["steve"],
	}
	p = NewPopulation(beings, chronology, culture)
	candidates, _ = p.MaritalCandidates(culture)
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and steve")
	}
}

func TestAddAndRemove(t *testing.T) {
	b := &Being{Species: &mockSpecies{}}
	p := NewPopulation([]inhabitants.Populatable{b}, nil, nil)
	p.Remove(b)
	if p.Get(b) {
		t.Fail()
	}
	p.Add(b)
	if !p.Get(b) {
		t.Fail()
	}
}

func TestReproductionCandidatesScore(t *testing.T) {
	chronology := timeline.NewChronology()
	pop := []inhabitants.Populatable{}
	for _, v := range beingFixtures {
		pop = append(pop, v)
	}
	p := NewPopulation(pop, chronology, &mockCulture{})
	rcs := p.ReproductionCandidates()
	for _, rc := range rcs {
		fmt.Println(rc)
	}

}

func TestAdamEve(t *testing.T) {
	chronology := timeline.NewChronology()
	pop := []inhabitants.Populatable{}
	for _, v := range beingFixtures {
		pop = append(pop, v)
	}
	p := NewPopulation(pop, chronology, &mockCulture{})
	for i := 0; i < 100; i++ {
		chronology.Tick()
		fmt.Println(p.Chronology.EventsForYear(i))
	}
	for _, b := range p.Inhabitants() {
		fmt.Println(b)
	}

}
