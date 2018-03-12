package locations

import (
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

type Habitation struct {
	Residents    *inhabitants.Population `json:"population"`
	Name         string                  `json:"name"`
	*words.Namer `json:"-"`
}

func NewHabitation(chronology *timeline.Chronology, culture *inhabitants.Culture) *Habitation {
	return &Habitation{Residents: inhabitants.NewPopulation([]*inhabitants.Being{}, chronology, culture)}
}

func (h *Habitation) SetNamer(namer *words.Namer) {
	h.Namer = namer
}

func (h *Habitation) Add(b *inhabitants.Being) (ok bool) {
	return h.Residents.Add(b)
}

func (h *Habitation) Remove(b *inhabitants.Being) (ok bool) {
	return h.Residents.Remove(b)
}

func (h *Habitation) Population() int {
	return h.Residents.Len()
}

func (h *Habitation) Resident(b *inhabitants.Being) (found bool) {
	return h.Residents.Get(b)
}
