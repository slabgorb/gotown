package inhabitants

import "github.com/slabgorb/gotown/timeline"

// Population is a set of Being
type Population struct {
	beings     map[*Being]bool
	chronology timeline.Chronology
}

type MaritalCandidate struct {
	female, male *Being
}

// NewPopulation initializes a Population
func NewPopulation(beings []*Being, chronology *timeline.Chronology) *Population {
	p := &Population{}
	if chronology == nil {

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
		b.Age++
	}
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.beings)
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	_, found := p.beings[b]
	p.beings[b] = true
	return !found
}

// Get returns whether this being is in the Population
func (p *Population) Get(b *Being) bool {
	_, found := p.beings[b]
	return found
}

// Remove removes a being from the population
func (p *Population) Remove(b *Being) bool {
	_, found := p.beings[b]
	delete(p.beings, b)
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

func (p *Population) MaritalCandidates(strategy MaritalStrategy) []*MaritalCandidate {
	mc := []*MaritalCandidate{}
	// loop through the population, taking each member and looking for candidates
	for _, m := range p.ByGender(Male) {
		for _, f := range p.ByGender(Female) {
			if strategy(m, f) {
				mc = append(mc, &MaritalCandidate{male: m, female: f})
			}
		}
	}
	return mc
}
