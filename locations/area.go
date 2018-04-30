package locations

import (
	"encoding/json"
	"fmt"

	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

// Area represents a geographical area
type Area struct {
	ID   int
	Name string
	*Habitation
	Size       AreaSize
	Graveyard  *being.Population
	Location   *Area
	Enclosures map[string]*Area
}

type serializeArea struct {
	ID           int      `json:"id" storm:"id,increment"`
	Name         string   `storm:"index"`
	HabitationID int      `json:"habitation_id"`
	Size         AreaSize `json:"size"`
	LocationID   int      `json:"location_id"`
	GraveyardID  int      `json:"graveyard_id"`
	EnclosureIDS []int    `json:"enclosure_ids"`
}

// MarshalJSON implements json.Marshaler
func (a *Area) MarshalJSON() ([]byte, error) {
	eids := []int{}
	for _, e := range a.Enclosures {
		eids = append(eids, e.ID)
	}
	s := &serializeArea{
		Name:         a.Name,
		HabitationID: a.Habitation.ID,
		Size:         a.Size,
		LocationID:   a.Location.ID,
		GraveyardID:  a.Graveyard.ID,
		EnclosureIDS: eids,
	}
	return json.Marshal(s)
}

// UnmarshalJSON implements json.Unmarshaler
func (a *Area) UnmarshalJSON(data []byte) error {
	s := &serializeArea{}
	if err := json.Unmarshal(data, s); err != nil {
		return fmt.Errorf("could not unmarshal area: %s", err)
	}
	a.ID = s.ID
	a.Habitation = &Habitation{ID: s.HabitationID}
	if err := a.Habitation.Read(); err != nil {
		return fmt.Errorf("could not load habitation %d: %s", a.Habitation.ID, err)
	}
	a.Size = s.Size
	a.Location = &Area{ID: s.LocationID}
	if err := a.Location.Read(); err != nil {
		return fmt.Errorf("could not load location %d: %s", a.Location.ID, err)
	}
	a.Graveyard = &being.Population{ID: s.GraveyardID}
	if err := a.Graveyard.Read(); err != nil {
		return fmt.Errorf("could not read graveyard(population) %d: %s", a.Graveyard.ID, err)
	}
	return nil
}

// Save implements persist.Persistable
func (a *Area) Save() error {
	return persist.DB.Save(a)
}

// Read implements persist.Persistable
func (a *Area) Read() error {
	if a.ID == 0 {
		return fmt.Errorf("need id for area")
	}
	return persist.DB.One("ID", a.ID, a)
}

// Delete implements persist.Persistable
func (a *Area) Delete() error {
	return persist.DB.DeleteStruct(a)
}

// NewArea creates an area
func NewArea(size AreaSize, location *Area) (*Area, error) {
	var n *words.Namer
	if location != nil {
		n = location.Namer
	} else {
		n = &words.Namer{Name: "english towns"}
		if err := n.Read(); err != nil {
			return nil, fmt.Errorf("error loading default words: %s", err)
		}
	}
	a := &Area{Size: size, Location: location}
	a.Habitation = NewHabitation()
	a.Enclosures = make(map[string]*Area)
	a.SetNamer(n)
	a.Name = a.Namer.CreateName()
	return a, nil
}

// Population returns the total population of the enclosed area
func (a *Area) Population() int {
	pop := a.Habitation.Len()
	for _, area := range a.Enclosures {
		pop += area.Population()
	}
	return pop
}

// IsEnclosedBy returns whether the receiver is enclosed by the parameter
func (a *Area) IsEnclosedBy(area *Area) bool {
	loc := a

	for {
		loc = a.Location
		if loc == nil {
			break
		}
		if loc == area {
			return true
		}
	}
	return false
}

// Encloses returns whether the reciever includes the parameter
func (a *Area) Encloses(area *Area) bool {
	for _, c := range a.Enclosures {
		if c == area {
			return true
		}
		return c.Encloses(area)
	}
	return false
}

// Detach removes the area from the Location it is currently in
func (a *Area) Detach() {
	if a.Location == nil {
		return
	}
	a.DetachFrom(a.Location)
}

// DetachFrom removes the passed in area from the receiver
func (a *Area) DetachFrom(area *Area) {
	delete(area.Enclosures, a.Name)
	a.Location = nil
}

// AttachTo puts the receiver into the passed in area
func (a *Area) AttachTo(area *Area) bool {
	// make sure we don't have circular relationship
	loc := area
	for {
		if loc == nil {
			break
		}
		loc = loc.Location
		if loc == nil {
			break
		}
		if loc == a {
			return false
		}
	}
	if a.Location != nil {
		a.DetachFrom(a.Location)
	}
	if area.Enclosures == nil {
		area.Enclosures = make(map[string]*Area)
	}
	area.Enclosures[a.Name] = a
	a.Location = area
	return true
}

// String implements fmt.Stringer
func (a *Area) String() string {
	return a.Name
}
