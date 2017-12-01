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
	beings map[*Being]struct{}
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
	male, female *Being
}

type ReproductionCandidate struct {
	b     *Being
	score float64
}

func (mc *MaritalCandidate) Pair() (*Being, *Being) {
	return mc.male, mc.female
}

// NewPopulation initializes a Population
func NewPopulation(beings []*Being, chronology *timeline.Chronology, culture *Culture) *Population {
	p := &Population{Chronology: chronology, Culture: culture}
	if chronology == nil {
		p.Chronology = timeline.NewChronology()
	}
	p.beings = make(map[*Being]struct{})
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
	p.beings[b] = struct{}{}
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
	out := []*Being{}
	for b := range p.beings {
		if b.Sex == g {
			out = append(out, b)
		}
	}
	return out
}

// ReproductionCandidates scans the population for potential candidates for
// reproduction.
func (p *Population) ReproductionCandidates() []*ReproductionCandidate {
	candidates := []*ReproductionCandidate{}

	for _, b := range p.ByGender(Gender("female")) {
		maxAge := b.Species.Demography[Adult].MaxAge
		minAge := b.Species.Demography[Child].MaxAge + 1
		if b.Age() > maxAge || b.Age() < minAge {
			continue
		}
		score := 0.05
		if b.Spouses != nil && len(b.Spouses) > 0 {
			score += 0.05
		}
		candidates = append(candidates, &ReproductionCandidate{b: b, score: score})
	}
	return candidates
}

// MaritalCandidates scans the population for potential candidates for marrying
// one another.
func (p *Population) MaritalCandidates() ([]*MaritalCandidate, error) {
	mc := make(map[MaritalCandidate]bool)
	if p.Culture == nil {
		return nil, fmt.Errorf("no culture for population, cannot assess marital candidates")
	}
	males := p.ByGender(Male)
	females := p.ByGender(Female)
	// loop through the population, taking each member and looking for candidates
	for _, a := range males {
		for _, b := range females {
			m := MaritalCandidate{male: a, female: b}
			if _, ok := mc[m]; ok {
				continue
			}
			mc[m] = p.Culture.MaritalCandidate(a, b)
		}
	}
	result := []*MaritalCandidate{}
	for k, v := range mc {
		if v {
			result = append(result, &k)

		}
	}
	return result, nil
}
