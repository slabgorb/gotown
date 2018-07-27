package locations

import (
	"fmt"

	"github.com/slabgorb/gotown/heraldry"
	"github.com/slabgorb/gotown/logger"

	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

// Area represents a geographical area
type Area struct {
	persist.IdentifiableImpl
	Name         string             `json:"name" storm:"index"`
	PopulationID string             `json:"population_id"`
	Size         AreaSize           `json:"size"`
	GraveyardID  string             `json:"graveyard_id"`
	LocationID   string             `json:"location_id"`
	EnclosureIDS []string           `json:"enclosure_ids"`
	Residents    *being.Population  `json:"-"`
	Graveyard    *being.Population  `json:"-"`
	Location     *Area              `json:"-"`
	Enclosures   map[string]*Area   `json:"-"`
	Heraldry     *heraldry.Heraldry `json:"heraldry"`
}

type AreaAPI struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Residents interface{} `json:"residents"`
	Size      string      `json:"size"`
	Image     string      `json:"image"`
	Icon      string      `json:"icon"`
}

func (a *Area) API() (interface{}, error) {

	populationAPI, err := a.Residents.API()
	if err != nil {
		return nil, fmt.Errorf("could not load population api %d for area %d", a.PopulationID, a.ID)
	}
	api := &AreaAPI{
		ID:        a.GetID(),
		Name:      a.Name,
		Residents: populationAPI,
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
	// if err := a.Graveyard.Save(); err != nil {
	// 	return err
	// }
	// a.GraveyardID = a.Graveyard.ID
	return persist.Save(a)
}

// Read implements persist.Persistable
func (a *Area) Read() error {
	if a.GetID() == "" {
		return fmt.Errorf("need id for area")
	}
	if err := persist.Read(a); err != nil {
		return err
	}
	// if err := a.readGraveyard(); err != nil {
	// 	return fmt.Errorf("could not read graveyard %d for area %d: %s", a.GraveyardID, a.ID, err)
	// }
	if err := a.readResidents(); err != nil {
		return fmt.Errorf("could not read population %d for area %d: %s", a.PopulationID, a.ID, err)
	}

	if a.Residents == nil || a.Residents.Len() == 0 {
		return fmt.Errorf("did not read population id: %d", a.PopulationID)
	}
	return nil
}

func (a *Area) GetID() string   { return a.ID }
func (a *Area) GetName() string { return a.Name }

func (a *Area) Reset() {
	a.Name = ""
	a.ID = ""
	a.Residents = nil
	a.Graveyard = nil
	a.GraveyardID = ""
	a.PopulationID = ""
	a.EnclosureIDS = []string{}
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

	return persist.Delete(a)
}

// NewArea creates an area
func NewArea(size AreaSize, location *Area, namer *words.Namer) *Area {
	a := &Area{Size: size, Location: location}
	a.Graveyard = being.NewPopulation([]string{}, logger.Default)
	a.Residents = being.NewPopulation([]string{}, logger.Default)
	a.EnclosureIDS = []string{}
	a.Enclosures = make(map[string]*Area)
	a.Name = namer.CreateName()
	e := heraldry.RandomEscutcheon("square", true)
	a.Heraldry = heraldry.New(e)
	return a
}

// Population returns the total population of the enclosed area
func (a *Area) Population() (int, error) {
	if a.Residents == nil {
		err := a.readResidents()
		if err != nil {
			return 0, err
		}
	}
	pop := a.Residents.Len()
	for _, area := range a.Enclosures {
		p, err := area.Population()
		if err != nil {
			return 0, err
		}
		pop = pop + p
	}
	return pop, nil
}

// IsEnclosedBy returns whether the receiver is enclosed by the parameter
func (a *Area) IsEnclosedBy(area *Area) bool {
	var loc *Area

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
		area.Enclosures = make(map[string]*Area)
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
	p := &being.Population{IdentifiableImpl: persist.IdentifiableImpl{ID: a.PopulationID}}
	a.Residents = p
	return p.Read()
}
func (a *Area) readGraveyard() error {
	p := &being.Population{IdentifiableImpl: persist.IdentifiableImpl{ID: a.GraveyardID}}
	a.Graveyard = p
	return p.Read()
}

func (a *Area) readEnclosures() error {
	return nil
}
