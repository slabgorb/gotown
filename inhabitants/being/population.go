package being

import (
	"fmt"
	"sync"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/timeline"
)

// Population is a set of Being
type Population struct {
	mux sync.Mutex
	ID  int              `json:"id" storm:"id,increment"`
	IDS map[int]struct{} `json:"ids"`
}

// MaritalCandidate is a pair of being
type MaritalCandidate struct {
	male, female *Being
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

// Pair returns the underlying beings of a marital candidate pair
func (mc *MaritalCandidate) Pair() (*Being, *Being) {
	return mc.male, mc.female
}

// NewPopulation initializes a Population
func NewPopulation(ids []int) *Population {
	p := &Population{IDS: make(map[int]struct{})}
	p.appendIds(ids...)
	return p
}

// Reset implements persist.Persistable
func (p *Population) Reset() {
	p.ID = 0
	p.IDS = make(map[int]struct{})
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
	return persist.DB.DeleteStruct(p)
}

// Save implements persist.Persistable
func (p *Population) Save() error {
	return persist.DB.Save(p)
}

type PopulationAPI struct {
	ID     int           `json:"id"`
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

func (p *Population) GetID() int      { return p.ID }
func (p *Population) GetName() string { return "" }

// Read implements persist.Persistable
func (p *Population) Read() error {
	if p.ID == 0 {
		return fmt.Errorf("need id for population")
	}
	return persist.DB.One("ID", p.ID, p)
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
					with := &Being{ID: withID}
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
	if p == nil {
		return nil, fmt.Errorf("nil population")
	}
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

// Age ages all the beings in this population and saves the beings
func (p *Population) Age() error {
	beings, err := p.Inhabitants()
	if err != nil {
		return err
	}
	for _, b := range beings {
		b.Age = b.Age + 1
	}
	return saveAll(beings)
}

// Len returns the number of beings in the population
func (p *Population) Len() int {
	return len(p.IDS)
}

func (p *Population) addID(id int) bool {
	p.mux.Lock()
	defer p.mux.Unlock()
	_, found := p.IDS[id]
	if !found {
		p.IDS[id] = struct{}{}
	}
	return !found
}

// Add adds a being to the population and returns whether it was actually added.
func (p *Population) Add(b *Being) bool {
	return p.addID(b.ID)
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

// MaritalCandidates scans the population for potential candidates for marrying
// one another.
func (p *Population) MaritalCandidates(c Cultured) ([]*MaritalCandidate, error) {
	mc := make(map[MaritalCandidate]bool)
	males, _ := p.ByGender(inhabitants.Male)
	females, _ := p.ByGender(inhabitants.Female)
	// loop through the population, taking each member and looking for candidates
	maritalStrategies := c.GetMaritalStrategies()
	for _, a := range males {
		for _, b := range females {
			m := MaritalCandidate{male: a, female: b}
			if _, ok := mc[m]; ok {
				continue
			}
			for _, f := range maritalStrategies {
				mc[m] = mc[m] && f(a, b)
			}
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
