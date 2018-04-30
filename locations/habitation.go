package locations

import (
	"fmt"

	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

// Habitation represents somewhere people live
type Habitation struct {
	ID                int `json:"id" storm:"id,increment"`
	*being.Population `json:"population"`
	Name              string `json:"name"`
	NamerName         string `json:"namer"`
	*words.Namer      `json:"-"`
}

// NewHabitation initializes a habitation
func NewHabitation() *Habitation {
	return &Habitation{Population: being.NewPopulation([]int{})}
}

// SetNamer sets the namer for this habiation
func (h *Habitation) SetNamer(namer *words.Namer) {
	h.Namer = namer
}

// Read implements persist.Persistable
func (h *Habitation) Read() error {
	if h.ID == 0 {
		return fmt.Errorf("cannot read habitation without id")
	}
	if err := persist.DB.One("ID", h.ID, h); err != nil {
		return fmt.Errorf("could not read habitation %d: %s", h.ID, err)
	}
	h.Namer = &words.Namer{Name: h.NamerName}
	return h.Namer.Read()
}

// Save implements persist.Persistable
func (h *Habitation) Save() error {
	return persist.DB.Save(h)
}

// Delete implements persist.Persistable
func (h *Habitation) Delete() error {
	return persist.DB.DeleteStruct(h)
}

// Reset implements persist.Persistable
func (h *Habitation) Reset() {
	h.ID = 0
	h.Population = being.NewPopulation([]int{})

}
