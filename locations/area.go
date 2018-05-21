package locations

import (
	"fmt"
	"github.com/slabgorb/gotown/heraldry"

	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

// Area represents a geographical area
type Area struct {
	ID           int                `json:"id" storm:"id,increment"`
	Name         string             `json:"name" storm:"index"`
	PopulationID int                `json:"population_id"`
	Size         AreaSize           `json:"size"`
	GraveyardID  int                `json:"graveyard_id"`
	LocationID   int                `json:"location_id"`
	EnclosureIDS []int              `json:"enclosure_ids"`
	Residents    *being.Population  `json:"-"`
	Graveyard    *being.Population  `json:"-"`
	Location     *Area              `json:"-"`
	Enclosures   map[int]*Area      `json:"-"`
	Heraldry     *heraldry.Heraldry `json:"heraldry"`
}

type AreaAPI struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Residents []*being.API `json:"residents"`
	Size      string       `json:"size"`
	Image     string       `json:"image"`
	Icon      string       `json:"icon"`
}

func (a *Area) API() (*AreaAPI, error) {
	beingsApi := []*being.API{}
	if a.Residents != nil {
		beings, err := a.Residents.Inhabitants()
		if err != nil {
			return nil, fmt.Errorf("could not load residents for area %d: %s", a.ID, err)
		}
		beingsApi, err = being.APIList(beings)
		if err != nil {
			return nil, fmt.Errorf("could not load being api for area %d: %s", a.ID, err)
		}
	}

	api := &AreaAPI{
		ID:        a.ID,
		Name:      a.Name,
		Residents: beingsApi,
		Size:      a.Size.String(),
		Icon:      a.Heraldry.Icon,
		Image:     a.Heraldry.Image,
	}
	return api, nil
}

// Add adds a being to the area
func (a *Area) Add(b *being.Being) {
	a.Residents.Add(b)
}

// Save implements persist.Persistable
func (a *Area) Save() error {
	if err := a.Residents.Save(); err != nil {
		return err
	}
	a.PopulationID = a.Residents.ID
	if err := a.Graveyard.Save(); err != nil {
		return err
	}
	a.GraveyardID = a.Graveyard.ID
	return persist.DB.Save(a)
}

// Read implements persist.Persistable
func (a *Area) Read() error {
	if a.ID == 0 {
		return fmt.Errorf("need id for area")
	}
	if err := persist.DB.One("ID", a.ID, a); err != nil {
		return err
	}
	// if err := a.readGraveyard(); err != nil {
	// 	return fmt.Errorf("could not read graveyard %d for area %d: %s", a.GraveyardID, a.ID, err)
	// }
	// if err := a.readResidents(); err != nil {
	// 	return fmt.Errorf("could not read population %d for area %d: %s", a.PopulationID, a.ID, err)
	// }
	return nil
}

func (a *Area) GetID() int      { return a.ID }
func (a *Area) GetName() string { return a.Name }

func (a *Area) Reset() {
	a.Name = ""
	a.ID = 0
	a.Residents = nil
	a.Graveyard = nil
	a.GraveyardID = 0
	a.PopulationID = 0
	a.EnclosureIDS = []int{}
}

// Delete implements persist.Persistable
func (a *Area) Delete() error {
	if a.Residents != nil {
		if err := a.Residents.Delete(); err != nil {
			return err
		}
	}

	if a.Graveyard != nil {
		if err := a.Graveyard.Delete(); err != nil {
			return err
		}
	}

	return persist.DB.DeleteStruct(a)
}

// NewArea creates an area
func NewArea(size AreaSize, location *Area, namer *words.Namer) *Area {
	a := &Area{Size: size, Location: location}
	a.Graveyard = being.NewPopulation([]int{})
	a.Residents = being.NewPopulation([]int{})
	a.EnclosureIDS = []int{}
	a.Enclosures = make(map[int]*Area)
	a.Name = namer.CreateName()
	e := heraldry.RandomEscutcheon("square", true)
	a.Heraldry = heraldry.New(e)
	return a
}

// Population returns the total population of the enclosed area
func (a *Area) Population() int {
	if a.Residents == nil {
		a.readResidents()
	}
	pop := a.Residents.Len()
	for _, area := range a.Enclosures {
		pop = pop + area.Population()
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
	delete(area.Enclosures, a.ID)
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
		area.Enclosures = make(map[int]*Area)
	}
	area.EnclosureIDS = append(area.EnclosureIDS, a.ID)
	area.Enclosures[a.ID] = a
	a.Location = area
	a.LocationID = area.ID
	return true
}

// String implements fmt.Stringer
func (a *Area) String() string {
	return a.Name
}

func (a *Area) readResidents() error {
	p := &being.Population{ID: a.PopulationID}
	return p.Read()
}
func (a *Area) readGraveyard() error {
	p := &being.Population{ID: a.GraveyardID}
	return p.Read()
}

func (a *Area) readEnclosures() error {
	return nil
}
