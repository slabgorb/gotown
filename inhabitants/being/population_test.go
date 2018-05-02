package being_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants/being"
)

func makePop(t *testing.T, count int) (*Population, []*Being) {
	ids := []int{}
	beings := []*Being{}
	var wg sync.WaitGroup
	beingQueue := make(chan *Being, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			bg := New(testSpecies, testCulture)
			bg.Randomize()
			if err := bg.Save(); err != nil {
				t.Fatalf("could not save being:%s", err)
			}
			beingQueue <- bg
		}()
	}
	go func(wg *sync.WaitGroup) {
		for b := range beingQueue {
			ids = append(ids, b.ID)
			beings = append(beings, b)
			wg.Done()
		}
	}(&wg)
	wg.Wait()
	return NewPopulation(ids), beings
}

func TestNew(t *testing.T) {
	p, _ := makePop(t, 0)
	if p.Len() != 0 {
		t.Errorf("expected 0 count for zero pop")
	}
	count := 100
	p, beings := makePop(t, count)
	if p.Len() != count {
		t.Errorf("expected %d count for generated pop got %d", count, p.Len())
	}
	j, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	p.Reset()
	if p.Len() != 0 {
		t.Errorf("expected 0 count for reset pop")
	}
	err = json.Unmarshal(j, p)
	if err != nil {
		t.Error(err)
	}
	if p.Len() != count {
		t.Errorf("expected %d count for unmarshaled pop", count)
	}
	for _, b := range beings {
		if !p.Exists(b) {
			t.Errorf("expected being %d to exist in population", b.ID)
		}
	}
}

func TestAging(t *testing.T) {
	fmt.Println("aging")
	_, beings := makePop(t, 10)
	expectedAges := []int{}
	for _, b := range beings {
		expectedAges = append(expectedAges, b.GetAge())
	}
	//p.Age()

	sort.Ints(ages)
	expectedAges := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range expectedAges {
		if v != ages[i] {
			t.Errorf("Expected %d got %d", v, ages[i])
		}
	}
}

// func BenchmarkMaritalCandidates(b *testing.B) {
// 	ids := []int{}
// 	for i := 0; i < 100; i++ {
// 		bg := New(testSpecies, testCulture)
// 		bg.Randomize()
// 		if err := bg.Save(); err != nil {
// 			b.Fatalf("could not save being:%s", err)
// 		}
// 		ids = append(ids, bg.ID)
// 	}
// 	p := NewPopulation(ids)
// 	b.Run("mc benchmark", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			_, _ = p.MaritalCandidates(testCulture)
// 		}
// 	})
// }

// func TestReproductionCandidates(t *testing.T) {
// 	ids := []int{
// 		beingFixtures["adam"].ID,
// 		beingFixtures["eve"].ID,
// 		beingFixtures["martha"].ID,
// 	}
// 	beingFixtures["eve"].Marry(beingFixtures["adam"])
// 	p := NewPopulation(ids)
// 	candidates := p.ReproductionCandidates()
// 	if len(candidates) != 2 {
// 		t.Fail()
// 	}

// }

// func TestMaritalCandidates(t *testing.T) {
// 	ids := []int{
// 		beingFixtures["adam"].ID,
// 		beingFixtures["eve"].ID,
// 	}
// 	p := NewPopulation(ids)
// 	candidates, err := p.MaritalCandidates(testCulture)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(candidates) < 1 {
// 		t.Log(candidates)
// 		t.Errorf("expected a candidate pair")
// 	}
// 	beings, _ := p.Inhabitants()
// 	a, b := candidates[0].Pair()
// 	if !(a == beings[0] || b == beings[0]) || !(a == beings[1] || b == beings[1]) {
// 		t.Errorf("expected adam and eve")
// 	}
// 	beings[1].Die()
// 	candidates, _ = p.MaritalCandidates(testCulture)
// 	if len(candidates) > 0 {
// 		t.Errorf("Did not expect adam and dead eve")
// 	}
// 	ids = []int{
// 		beingFixtures["adam"].ID,
// 		beingFixtures["steve"].ID,
// 	}
// 	p = NewPopulation(ids)
// 	candidates, _ = p.MaritalCandidates(testCulture)
// 	if len(candidates) > 0 {
// 		t.Errorf("Did not expect adam and steve")
// 	}
// }

// func TestAddAndRemove(t *testing.T) {
// 	b := New(testSpecies, testCulture)

// 	p := NewPopulation([]int{})
// 	p.Remove(b)
// 	if p.Exists(b) {
// 		t.Fail()
// 	}
// 	p.Add(b)
// 	if !p.Exists(b) {
// 		t.Fail()
// 	}
// }

// func TestAdamEve(t *testing.T) {
// 	pop := []int{}
// 	for _, v := range beingFixtures {
// 		pop = append(pop, v.ID)
// 	}
// 	p := NewPopulation(pop)
// 	beings, _ := p.Inhabitants()
// 	for _, b := range beings {
// 		fmt.Println(b)
// 	}
// }

// func TestMaritalStrategy(t *testing.T) {
// 	testCases := []struct {
// 		name     string
// 		a        *Being
// 		b        *Being
// 		ages     []int
// 		expected bool
// 	}{
// 		{
// 			name:     "usual",
// 			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Female},
// 			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male},
// 			expected: true,
// 		},
// 		{
// 			name:     "hetero only for this culture (yes, sorry)",
// 			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male},
// 			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Male},
// 			expected: false,
// 		},
// 		{
// 			name:     "no bigamy",
// 			a:        &Being{Culture: testCulture, Species: testSpecies, Age: 20, Gender: inhabitants.Male, Spouses: []int{0}},
// 			b:        &Being{Culture: testCulture, Species: testSpecies, Age: 19, Gender: inhabitants.Female},
// 			expected: false,
// 		},
// 	}
// 	i := 1
// 	for _, tc := range testCases {
// 		tc.a.ID = i
// 		tc.b.ID = i + 1
// 		i += 2
// 		p := NewPopulation([]int{tc.a.ID, tc.b.ID})
// 		candidates, _ := p.MaritalCandidates(testCulture)
// 		actual := len(candidates) > 0
// 		if tc.expected != actual {
// 			t.Errorf("%s expected %t got %t", tc.name, tc.expected, actual)
// 		}
// 	}
// }
