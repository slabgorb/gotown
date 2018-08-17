package being

import (
	"fmt"
	"sync"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/logger"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/timeline"
)

// Population is a set of Being
type Population struct {
	mux sync.Mutex
	persist.IdentifiableImpl
	IDS         map[string]struct{} `json:"ids"`
	logger      Logger
	inhabitants []*Being
	stale       bool
}

// ReproductionCandidate blah blah (maybe should be unexported)
type ReproductionCandidate struct {
	b     *Being
	score float64
}

// String implements fmt.Stringer
func (rc ReproductionCandidate) String() string {
	return fmt.Sprintf("%s (score %f)", rc.b, rc.score)
}

// NewPopulation initializes a Population
func NewPopulation(ids []string, logger Logger) *Population {
	p := &Population{IDS: make(map[string]struct{}), logger: logger}
	p.appendIds(ids...)
	return p
}

// Reset implements persist.Persistable
func (p *Population) Reset() {
	p.ID = ""
	p.IDS = make(map[string]struct{})
}

func (p *Population) String() string {
	return p.ID
}

// Delete implements persist.Persistable
func (p *Population) Delete() error {
	beings, err := p.Inhabitants()
	if err != nil {
		return err
	}
	for _, b := range beings {
		if err := b.Delete(); err != nil {
			return err
		}
	}
	return persist.Delete(p)
}

// Save implements persist.Persistable
func (p *Population) Save() error {
	return persist.Save(p)
}

type PopulationAPI struct {
	ID     string        `json:"id"`
	Beings []interface{} `json:"beings"`
}

// API returns the population as an API struct
func (p *Population) API() (interface{}, error) {
	apis := []interface{}{}
	beings, err := p.Inhabitants()
	if err != nil {
		return nil, err
	}
	for _, b := range beings {
		api, err := b.API()
		if err != nil {
			return nil, err
		}
		apis = append(apis, api)
	}
	return &PopulationAPI{
		ID:     p.ID,
		Beings: apis,
	}, nil
}

func (p *Population) GetName() string { return "" }

// Read implements persist.Persistable
func (p *Population) Read() error {
	if p.ID == "" {
		return fmt.Errorf("need id for population")
	}
	return persist.Read(p)
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

func reproduction(p *Population) timeline.Callback {
	return func(_ int) {
		rc := p.ReproductionCandidates()
		for _, r := range rc {
			if randomizer.Float64() < r.score {
				var with *Being
				// if b is married, choose spouse?
				if r.b.Spouses != nil && len(r.b.Spouses) > 0 {
					withID := r.b.Spouses[0]
					with := &Being{IdentifiableImpl: persist.IdentifiableImpl{ID: withID}}
					with.Read()
				} else {
					// choose random guy for now, will work on the choice later
					men, _ := p.ByGender(inhabitants.Male)
					with = men[randomizer.Intn(len(men))]
				}
				r.b.Reproduce(with)
			}
		}
	}
}

// Inhabitants returns the beings in the population
func (p *Population) Inhabitants() ([]*Being, error) {
	logger.TimeSet("loading inhabitants")
	defer func() { logger.TimeElapsed("loading inhabitants") }()
	if p.stale == false && len(p.inhabitants) > 0 {
		return p.inhabitants, nil
	}
	if p == nil {
		return nil, fmt.Errorf("nil population")
	}
	ids := []string{}
	for id := range p.IDS {
		ids = append(ids, id)
	}
	bs, err := readAll(ids)
	if err != nil {
		return nil, err
	}
	p.inhabitants = bs
	logger.Debug("population %d", len(bs))
	p.stale = false
	return bs, nil
}

// Age ages all the beings in this population and saves the beings
func (p *Population) Age() error {
	beings, err := p.Inhabitants()
	if err != nil {
		return err
	}
	for _, b := range beings {
		b.Age = b.Age + 1
	}
	p.stale = true
	return saveAll(beings)
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.IDS)
}

func (p *Population) addID(id string) bool {
	p.mux.Lock()
	defer p.mux.Unlock()
	_, found := p.IDS[id]
	if !found {
		p.IDS[id] = struct{}{}
	}
	p.stale = true
	return !found
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	return p.addID(b.ID)
}

func (p *Population) getIds() []string {
	out := []string{}
	for id := range p.IDS {
		out = append(out, id)
	}
	return out
}

func (p *Population) appendIds(ids ...string) {
	for _, i := range ids {
		p.addID(i)
	}
}

// Exists returns whether this being is in the Population
func (p *Population) Exists(b *Being) bool {
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

// ByGender returns the beings filtered by gender
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

type maritalStrategy func(a, b Marriageable) bool

func contains(mc []MaritalCandidate, blist ...*Being) bool {
	for _, k := range mc {
		for _, b := range blist {
			if k.Contains(b) {
				return true
			}
		}
	}
	return false
}

// MaritalCandidates scans the population for potential candidates for marrying
// one another.
func (p *Population) MaritalCandidates(c Cultured) ([]MaritalCandidate, error) {
	mc := make(map[MaritalCandidate]bool)
	males, _ := p.ByGender(inhabitants.Male)
	females, _ := p.ByGender(inhabitants.Female)
	// loop through the population, taking each member and looking for candidates
	maritalStrategies := c.GetMaritalStrategies()
	for _, a := range males {
		for _, b := range females {
			m := MaritalCandidate{male: a, female: b}
			mc[m] = true
			for _, f := range maritalStrategies {
				mc[m] = mc[m] && f(m.male, m.female)
			}
		}
	}
	result := []MaritalCandidate{}
	for k := range mc {
		if mc[k] && !contains(result, k.female) {
			result = append(result, k)
		}
	}
	return result, nil
}
