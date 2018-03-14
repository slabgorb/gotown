package locations

import (
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
)

type Area struct {
	ID   int    `json:"id" storm:"id,increment"`
	Name string `storm:"index"`
	*Habitation
	Size  AreaSize           `json:"size"`
	Ruler *inhabitants.Being `json:"ruler"`
	Graveyard
	Location   *Area            `json:"location"`
	Enclosures map[string]*Area `json:"enclosures"`
}

func NewArea(size AreaSize, culture *inhabitants.Culture, ruler *inhabitants.Being, location *Area) *Area {
	var n *words.Namer
	if location != nil {
		n = location.Namer
	} else {
		n = words.TownNamer
	}
	a := &Area{Size: size, Ruler: ruler, Location: location}
	a.Habitation = NewHabitation(timeline.NewChronology(), culture)
	a.Enclosures = make(map[string]*Area)
	a.SetNamer(n)
	a.Name = a.Namer.Name()
	return a
}

func (a *Area) Population() int {
	pop := a.Habitation.Population()
	for _, area := range a.Enclosures {
		pop += area.Population()
	}
	return pop
}

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

func (a *Area) Encloses(area *Area) bool {
	for _, c := range a.Enclosures {
		if c == area {
			return true
		}
		return c.Encloses(area)
	}
	return false
}

func (a *Area) Detach() {
	if a.Location == nil {
		return
	}
	a.DetachFrom(a.Location)
}

func (a *Area) DetachFrom(area *Area) {
	delete(area.Enclosures, a.Name)
	a.Location = nil
}

func (a *Area) AttachTo(area *Area) bool {
	// make sure we don't have circular relationship
	loc := area
	for {
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

func (a *Area) String() string {
	return a.Name
}

// func (a *Area) String() string {
// 	loc := ""
// 	if a.Location != nil {
// 		loc += fmt.Sprintf("%s, %s %s within %s", a.Name, a.Size.article(), a.Size, a.Location.Name)
// 	} else {
// 		loc += fmt.Sprintf("%s, %s %s", a.Name, a.Size.article(), a.Size)
// 	}
// 	if a.Ruler != nil {
// 		loc += fmt.Sprintf(" ruled by %s", a.Ruler)
// 	}
// 	loc += "."
// 	return loc
// }
