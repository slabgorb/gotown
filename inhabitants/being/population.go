package being

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/timeline"
)

// Population is a set of Being
type Population struct {
	mux        sync.Mutex
	Beings     map[*Being]struct{}  `json:"inhabitants"`
	Chronology *timeline.Chronology `json:"history"`
}

type MaritalCandidate struct {
	male, female *Being
}

type ReproductionCandidate struct {
	b     *Being
	score float64
}

func (rc ReproductionCandidate) String() string {
	return fmt.Sprintf("%s (score %f)", rc.b, rc.score)
}

func (mc *MaritalCandidate) Pair() (*Being, *Being) {
	return mc.male, mc.female
}

// NewPopulation initializes a Population
func NewPopulation(beings []inhabitants.Populatable, chronology *timeline.Chronology, culture inhabitants.Cultured) *Population {
	p := &Population{Chronology: chronology}
	if chronology == nil {
		p.Chronology = timeline.NewChronology()
	}
	p.Chronology.Register(reproduction(p, culture))
	p.Chronology.Register(marry(p, culture))
	p.Beings = make(map[*Being]struct{})
	for _, b := range beings {
		if being, ok := b.(*Being); ok {
			p.Add(being)
		}
	}
	return p
}

type populationSerializer struct {
	Inhabitants []*Being             `json:"inhabitants"`
	History     *timeline.Chronology `json:"history"`
}

func (p *Population) MarshalJSON() ([]byte, error) {
	ps := &populationSerializer{Inhabitants: p.Inhabitants(), History: p.Chronology}
	return json.Marshal(ps)
}

func (p *Population) UnmarshalJSON(data []byte) error {
	ps := &populationSerializer{}
	if err := json.Unmarshal(data, ps); err != nil {
		return err
	}
	for _, b := range ps.Inhabitants {
		p.Add(b)
	}
	p.Chronology = ps.History
	return nil
}

func marry(p *Population, c inhabitants.Cultured) timeline.Callback {
	return func(_ int) {
		mc, _ := p.MaritalCandidates(c)
		for _, m := range mc {
			r := randomizer.Float64()
			if r < 0.10 {
				m.female.Marry(m.male)
			}
		}
	}
}

func (p *Population) History() *timeline.Chronology {
	return p.Chronology
}

func reproduction(p *Population, c inhabitants.Cultured) timeline.Callback {
	return func(_ int) {
		rc := p.ReproductionCandidates()
		for _, r := range rc {
			if randomizer.Float64() < r.score {
				var with *Being
				// if b is married, choose spouse?
				if r.b.Spouses != nil && len(r.b.Spouses) > 0 {
					with = r.b.Spouses[0]
				} else {
					// choose random guy for now, will work on the choice later
					men := p.ByGender(inhabitants.Male)
					with = men[randomizer.Intn(len(men))]
				}
				r.b.Reproduce(with, c)
			}
		}
	}
}

// Inhabitants returns the beings in the population
func (p *Population) Inhabitants() []*Being {
	bs := make([]*Being, p.Len())
	i := 0
	for b := range p.Beings {
		bs[i] = b
		i++
	}
	return bs
}

// Age ages all the beings in this population
func (p *Population) Age() {
	for b := range p.Beings {
		b.History().Tick()
	}
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.Beings)
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	p.mux.Lock()
	_, found := p.Beings[b]
	p.Beings[b] = struct{}{}
	p.History().Register(func(year int) {
		b.History().Tick()
	})
	p.mux.Unlock()
	return !found
}

// Get returns whether this being is in the Population
func (p *Population) Get(b *Being) bool {
	p.mux.Lock()
	_, found := p.Beings[b]
	p.mux.Unlock()
	return found
}

// Remove removes a being from the population
func (p *Population) Remove(b *Being) bool {
	p.mux.Lock()
	_, found := p.Beings[b]
	delete(p.Beings, b)
	p.mux.Unlock()
	return found
}

func (p *Population) ByGender(g inhabitants.Gender) []*Being {
	out := []*Being{}
	for b := range p.Beings {
		if b.Sex() == g {
			out = append(out, b)
		}
	}
	return out
}

// ReproductionCandidates scans the population for potential candidates for
// reproduction.
func (p *Population) ReproductionCandidates() []*ReproductionCandidate {
	candidates := []*ReproductionCandidate{}

	for _, b := range p.ByGender(inhabitants.Female) {
		maxAge := b.Species.MaxAge(inhabitants.Adult)
		minAge := b.Species.MaxAge(inhabitants.Child) + 1
		if b.Age() > maxAge || b.Age() < minAge {
			continue
		}
		var score float64
		r := float64(maxAge - minAge)
		adjustedAge := float64(b.Age() - minAge)
		splits := []float64{r * 0.3, r * 0.6, r}
		switch {
		case adjustedAge < splits[1]:
			score = 0.2
		case adjustedAge < splits[2]:
			score = 0.15
		}

		if b.Spouses == nil || len(b.Spouses) == 0 {
			score -= 0.1
		}
		candidates = append(candidates, &ReproductionCandidate{b: b, score: score})
	}
	return candidates
}

// MaritalCandidates scans the population for potential candidates for marrying
// one another.
func (p *Population) MaritalCandidates(c inhabitants.Cultured) ([]*MaritalCandidate, error) {
	mc := make(map[MaritalCandidate]bool)
	males := p.ByGender(inhabitants.Male)
	females := p.ByGender(inhabitants.Female)
	// loop through the population, taking each member and looking for candidates
	for _, a := range males {
		for _, b := range females {
			m := MaritalCandidate{male: a, female: b}
			if _, ok := mc[m]; ok {
				continue
			}
			mc[m] = c.MaritalCandidate(a, b)
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
