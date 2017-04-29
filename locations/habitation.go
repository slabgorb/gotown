package locations

import (
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/words"
)

type Habitation struct {
	Residents []*inhabitants.Being
	Name      string
	*words.Namer
}

func (h *Habitation) SetNamer(namer *words.Namer) {
	h.Namer = namer
}

func (h *Habitation) Add(b *inhabitants.Being) (ok bool) {
	_, found := h.Resident(b)
	if found {
		return false
	}
	h.Residents = append(h.Residents, b)
	return true
}

func (h *Habitation) Remove(b *inhabitants.Being) (ok bool) {
	index, found := h.Resident(b)
	if !found {
		return false
	}
	h.Residents = append(h.Residents[:index], h.Residents[index+1:]...)
	return true
}

func (h *Habitation) Population() int {
	return len(h.Residents)
}

func (h *Habitation) Resident(b *inhabitants.Being) (index int, found bool) {
	for i, r := range h.Residents {
		if r == b {
			found = true
			index = i
			break
		}
	}
	return index, found
}
