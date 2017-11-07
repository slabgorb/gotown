package inhabitants

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/slabgorb/gotown/timeline"
)

// Population is a set of Being
type Population struct {
	mux    sync.Mutex
	beings map[*Being]bool
	*timeline.Chronology
	*Culture
}

func (p *Population) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Beings     []*Being             `json:"residents"`
		Chronology *timeline.Chronology `json:"chronology"`
	}{
		Beings:     p.Beings(),
		Chronology: p.Chronology,
	})
}

type MaritalCandidate struct {
	a, b *Being
}

func (mc *MaritalCandidate) Pair() (*Being, *Being) {
	return mc.a, mc.b
}

// NewPopulation initializes a Population
func NewPopulation(beings []*Being, chronology *timeline.Chronology, culture *Culture) *Population {
	p := &Population{Chronology: chronology, Culture: culture}
	if chronology == nil {
		p.Chronology = timeline.NewChronology()
	}
	p.beings = make(map[*Being]bool)
	for _, b := range beings {
		p.Add(b)
	}
	return p
}

// Beings returns the beings in the population
func (p *Population) Beings() []*Being {
	bs := make([]*Being, p.Len())
	i := 0
	for b := range p.beings {
		bs[i] = b
		i++
	}
	return bs
}

// Age ages all the beings in this population
func (p *Population) Age() {
	for b := range p.beings {
		b.Chronology.Tick()
	}
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.beings)
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	p.mux.Lock()
	_, found := p.beings[b]
	p.beings[b] = true
	p.mux.Unlock()
	return !found
}

// Get returns whether this being is in the Population
func (p *Population) Get(b *Being) bool {
	p.mux.Lock()
	_, found := p.beings[b]
	p.mux.Unlock()
	return found
}

// Remove removes a being from the population
func (p *Population) Remove(b *Being) bool {
	p.mux.Lock()
	_, found := p.beings[b]
	delete(p.beings, b)
	p.mux.Unlock()
	return found
}

func (p Population) ByGender(g Gender) []*Being {
	out := make([]*Being, p.Len())
	i := 0
	for b := range p.beings {
		if b.Sex == g {
			out[i] = b
			i++
		}
	}
	return out
}

func (p *Population) MaritalCandidates() ([]*MaritalCandidate, error) {
	mc := []*MaritalCandidate{}
	if p.Culture == nil {
		return nil, fmt.Errorf("no culture for population, cannot assess marital candidates")
	}
	beings := p.Beings()
	// loop through the population, taking each member and looking for candidates
	for _, a := range beings {
		for _, b := range beings {
			if a == b {
				continue
			}
			if p.Culture.MaritalCandidate(a, b) {
				mc = append(mc, &MaritalCandidate{a: a, b: b})
			}
		}
	}
	return mc, nil
}
