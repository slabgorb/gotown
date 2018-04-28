package locations

import (
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

type Habitation struct {
	ID           int               `json:"id" storm:"id,increment"`
	Residents    *being.Population `json:"population"`
	Name         string            `json:"name"`
	NamerName    string            `json:"namer"`
	*words.Namer `json:"-"`
}

func NewHabitation(chronology *timeline.Chronology, culture inhabitants.Cultured) *Habitation {
	return &Habitation{Residents: being.NewPopulation([]int{})}
}

func (h *Habitation) SetNamer(namer *words.Namer) {
	h.Namer = namer
}

func (h *Habitation) Add(b *being.Being) (ok bool) {
	return h.Residents.Add(b)
}

func (h *Habitation) Remove(b *being.Being) (ok bool) {
	return h.Residents.Remove(b)
}

func (h *Habitation) Population() int {
	return h.Residents.Len()
}

func (h *Habitation) Resident(b *being.Being) (found bool) {
	return h.Residents.Exists(b)
}
