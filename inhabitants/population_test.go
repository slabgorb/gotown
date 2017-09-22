package inhabitants_test

import (
	"sort"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants"
)

func TestAging(t *testing.T) {
	p := NewPopulation([]*Being{}, nil, nil)
	count := 10
	beings := make([]*Being, count)
	for i := 0; i < count; i++ {
		beings[i] = &Being{Species: mockSpecies, Age: i}
		p.Add(beings[i])
	}
	p.Age()
	ages := []int{}
	for _, b := range p.Beings() {
		ages = append(ages, b.Age)

	}
	sort.Ints(ages)
	expectedAges := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range expectedAges {
		if v != ages[i] {
			t.Errorf("Expected %d got %d", v, ages[i])
		}
	}
}

func TestGender(t *testing.T) {

}

func TestAddAndRemove(t *testing.T) {
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
