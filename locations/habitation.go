package locations

import (
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

type Habitation struct {
	Residents    *being.Population `json:"population"`
	Name         string            `json:"name"`
	*words.Namer `json:"-"`
}

func NewHabitation(chronology *timeline.Chronology, culture inhabitants.Cultured) *Habitation {
	return &Habitation{Residents: being.NewPopulation([]inhabitants.Populatable{}, chronology, culture)}
}

func (h *Habitation) SetNamer(namer *words.Namer) {
	h.Namer = namer
}

func (h *Habitation) Add(b inhabitants.Populatable) (ok bool) {
	return h.Residents.Add(b)
}

func (h *Habitation) Remove(b inhabitants.Populatable) (ok bool) {
	return h.Residents.Remove(b)
}

func (h *Habitation) Population() int {
	return h.Residents.Len()
}

func (h *Habitation) Resident(b inhabitants.Populatable) (found bool) {
	return h.Residents.Get(b)
}
