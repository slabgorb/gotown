package being

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/timeline"
)

// Population is a set of Being
type Population struct {
	mux sync.Mutex
	ID  int
	IDS map[int]struct{}
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
func NewPopulation(ids []int) *Population {
	p := &Population{}
	p.appendIds(ids...)
	return p
}

type populationSerializer struct {
	ID  int   `json:"id" storm:"id,increment"`
	IDS []int `json:"ids"`
}

func (p *Population) MarshalJSON() ([]byte, error) {
	ps := &populationSerializer{ID: p.ID, IDS: p.getIds()}
	return json.Marshal(ps)
}

func (p *Population) UnmarshalJSON(data []byte) error {
	ps := &populationSerializer{}
	if err := json.Unmarshal(data, ps); err != nil {
		return err
	}
	for id := range ps.IDS {
		b := &Being{ID: id}
		if err := b.Read(); err != nil {
			return err
		}
		p.Add(b)
	}
	return nil
}

func marry(p *Population, c Cultured) timeline.Callback {
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

func reproduction(p *Population, c inhabitants.Cultured) timeline.Callback {
	return func(_ int) {
		rc := p.ReproductionCandidates()
		for _, r := range rc {
			if randomizer.Float64() < r.score {
				var with *Being
				// if b is married, choose spouse?
				if r.b.Spouses != nil && len(r.b.Spouses) > 0 {
					withID := r.b.Spouses[0]
					with := &Being{ID: withID}
					with.Read()
				} else {
					// choose random guy for now, will work on the choice later
					men, _ := p.ByGender(inhabitants.Male)
					with = men[randomizer.Intn(len(men))]
				}
				r.b.Reproduce(with, c)
			}
		}
	}
}

// Inhabitants returns the beings in the population
func (p *Population) Inhabitants() ([]*Being, error) {
	bs := make([]*Being, p.Len())
	i := 0
	for id := range p.IDS {
		b := &Being{ID: id}
		if err := b.Read(); err != nil {
			return nil, err
		}
		bs[i] = b
		i++
	}
	return bs, nil
}

func saveAll(beings []*Being) error {
	ps := []persist.Persistable{}
	for _, b := range beings {
		ps = append(ps, b)
	}
	return persist.SaveAll(ps)
}

// Age ages all the beings in this population and saves the beings
func (p *Population) Age() error {
	beings, err := p.Inhabitants()
	if err != nil {
		return err
	}
	for _, b := range beings {
		b.Age++
	}
	return saveAll(beings)
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.IDS)
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	p.mux.Lock()
	defer p.mux.Unlock()
	_, found := p.IDS[b.ID]
	if !found {
		p.IDS[b.ID] = struct{}{}
	}
	return !found
}

func (p *Population) getIds() []int {
	out := []int{}
	for id := range p.IDS {
		out = append(out, id)
	}
	return out
}

func (p *Population) appendIds(ids ...int) {
	for _, i := range ids {
		b := &Being{ID: i}
		p.Add(b)
	}
}

// Get returns whether this being is in the Population
func (p *Population) Get(b *Being) bool {
	p.mux.Lock()
	_, found := p.IDS[b.ID]
	p.mux.Unlock()
	return found
}

// Remove removes a being from the population
func (p *Population) Remove(b *Being) bool {
	p.mux.Lock()
	_, found := p.IDS[b.ID]
	delete(p.IDS, b.ID)
	p.mux.Unlock()
	return found
}

func (p *Population) ByGender(g inhabitants.Gender) ([]*Being, error) {
	out := []*Being{}
	beings, err := p.Inhabitants()
	if err != nil {
		return nil, err
	}
	for _, b := range beings {
		if b.Sex() == g {
			out = append(out, b)
		}
	}
	return out, nil
}

// ReproductionCandidates scans the population for potential candidates for
// reproduction.
func (p *Population) ReproductionCandidates() []*ReproductionCandidate {
	candidates := []*ReproductionCandidate{}
	females, _ := p.ByGender(inhabitants.Female)
	for _, b := range females {
		maxAge := b.Species.MaxAge(inhabitants.Adult)
		minAge := b.Species.MaxAge(inhabitants.Child) + 1
		if b.GetAge() > maxAge || b.GetAge() < minAge {
			continue
		}
		var score float64
		r := float64(maxAge - minAge)
		adjustedAge := float64(b.GetAge() - minAge)
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
func (p *Population) MaritalCandidates(c Cultured) ([]*MaritalCandidate, error) {
	mc := make(map[MaritalCandidate]bool)
	males, _ := p.ByGender(inhabitants.Male)
	females, _ := p.ByGender(inhabitants.Female)
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
