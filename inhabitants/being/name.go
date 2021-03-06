package being

import (
	"strings"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/logger"
	"github.com/slabgorb/gotown/words"
)

// Name is the name of a being or other named thing, in theory
type Name struct {
	GivenName  string   `json:"given_name"`
	FamilyName string   `json:"family_name"`
	Nickname   string   `json:"nickname"`
	Other      []string `json:"other_name"`
	Display    string   `json:"display_name"`
}

// OtherNames returns any other names a being may have as a space-separated list
func (n *Name) OtherNames() string {
	return strings.Join(n.Other, " ")
}

// GetGivenName is a getter
func (n *Name) GetGivenName() string { return n.GivenName }

// GetFamilyName is a getter
func (n *Name) GetFamilyName() string { return n.FamilyName }

// GetOtherNames is a getter
func (n *Name) GetOtherNames() []string { return n.Other }

// GetDisplay is a getter
func (n *Name) GetDisplay() string { return n.Display }

// NewName tries valiantly to create a formal name from a string
func NewName(fullName string) *Name {
	name := &Name{Display: fullName}
	names := strings.Split(fullName, " ")
	if len(names) > 0 {
		name.GivenName = names[0]
	}
	if len(names) > 1 {
		name.FamilyName = names[1]
	}
	if len(names) > 2 {
		name.Other = names[2:]
	}
	return name
}

// Nameable abstracts...
type Nameable interface {
	Father() (*Being, error)
	Mother() (*Being, error)
	GetNamer() *words.Namer
	GetFullName() *Name
	Sex() inhabitants.Gender
}

// NameStrategy is a function which describes how children are named
type NameStrategy func(b Nameable) *Name

// NameStrategies deliniates the various naming strategy functions
var NameStrategies = map[string]NameStrategy{
	"matrilineal": func(b Nameable) *Name {
		namer := b.GetNamer()
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Mother()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetFullName().FamilyName
			return name
		}
		name.FamilyName = namer.Words.FamilyName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patrilineal": func(b Nameable) *Name {
		namer := b.GetNamer()
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Father()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetFullName().FamilyName
			return name
		}
		name.FamilyName = namer.Words.FamilyName()
		display, err := namer.Execute(name)
		if err != nil {
			logger.Error("Error executing template: %s", err.Error())
		}
		name.Display = display
		return name
	},
	"matronymic": func(b Nameable) *Name {
		namer := b.GetNamer()
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Mother()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetFullName().GivenName + namer.Words.Matronymic()
			return name
		}
		name.FamilyName = namer.Words.GivenName() + namer.Words.Matronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patronymic": func(b Nameable) *Name {
		namer := b.GetNamer()
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Father()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetFullName().GivenName + namer.Words.Patronymic()
			return name
		}
		name.FamilyName = namer.Words.GivenName() + namer.Words.Patronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"one name": func(b Nameable) *Name {
		namer := b.GetNamer()
		name := &Name{GivenName: namer.Words.GivenName(), Nickname: namer.Words.Nickname()}
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
}
