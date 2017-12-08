package inhabitants_test

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/timeline"
)

func TestAging(t *testing.T) {
	mockSpecies := helperMockSpecies(t)
	p := NewPopulation([]*Being{}, nil, nil)
	count := 10
	beings := make([]*Being, count)
	for i := 0; i < count; i++ {
		beings[i] = &Being{Species: mockSpecies, Chronology: &timeline.Chronology{CurrentYear: i}}
		p.Add(beings[i])
	}
	p.Age()
	ages := []int{}
	for _, b := range p.Beings() {
		ages = append(ages, b.Age())

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
	t := &testing.T{}
	culture := helperMockCulture(t, "italian")
	beings := []*Being{}
	for i := 0; i < 100; i++ {
		b := &Being{}
		b.Randomize()
		beings = append(beings, b)
	}
	var bc []*MaritalCandidate
	chronology := timeline.NewChronology()
	p := NewPopulation(beings, chronology, culture)
	b.Run("mc benchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bc, _ = p.MaritalCandidates()
		}
	})
}

func TestReproductionCandidates(t *testing.T) {
	culture := helperMockCulture(t, "italian")
	beingFixtures := beingFixtures(t, "italian")
	chronology := timeline.NewChronology()
	beings := []*Being{
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
	culture := helperMockCulture(t, "italian")
	beingFixtures := beingFixtures(t, "italian")
	chronology := timeline.NewChronology()
	beings := []*Being{
		beingFixtures["adam"],
		beingFixtures["eve"],
	}
	p := NewPopulation(beings, chronology, culture)
	candidates, _ := p.MaritalCandidates()
	a, b := candidates[0].Pair()
	if !(a == beings[0] || b == beings[0]) || !(a == beings[1] || b == beings[1]) {
		t.Errorf("expected adam and eve")
	}
	beings[1].Die()
	candidates, _ = p.MaritalCandidates()
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and dead eve")
	}
	beings = []*Being{
		beingFixtures["adam"],
		beingFixtures["steve"],
	}
	p = NewPopulation(beings, chronology, culture)
	candidates, _ = p.MaritalCandidates()
	if len(candidates) > 0 {
		t.Errorf("Did not expect adam and steve")
	}
}

func TestAddAndRemove(t *testing.T) {
	mockSpecies := helperMockSpecies(t)
	b := &Being{Species: mockSpecies}
	p := NewPopulation([]*Being{b}, nil, nil)
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
	culture := helperMockCulture(t, "italian")
	beingFixtures := beingFixtures(t, "italian")
	chronology := timeline.NewChronology()
	pop := []*Being{}
	for _, v := range beingFixtures {
		pop = append(pop, v)
	}
	p := NewPopulation(pop, chronology, culture)
	rcs := p.ReproductionCandidates()
	for _, rc := range rcs {
		fmt.Println(rc)
	}

}

func TestAdamEve(t *testing.T) {
	culture := helperMockCulture(t, "italian")
	beingFixtures := beingFixtures(t, "italian")
	chronology := timeline.NewChronology()
	pop := []*Being{}
	for _, v := range beingFixtures {
		pop = append(pop, v)
	}
	p := NewPopulation(pop, chronology, culture)
	for i := 0; i < 100; i++ {
		chronology.Tick()
		fmt.Println(p.Chronology.EventsForYear(i))
	}
	for _, b := range p.Beings() {
		fmt.Println(b)
	}

}
