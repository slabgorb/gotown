package inhabitants

import "strings"

// Name is the name of a being or other named thing, in theory
type Name struct {
	GivenName  string   `json:"given_name"`
	FamilyName string   `json:"family_name"`
	Other      []string `json:"other_name"`
	Display    string   `json:"display_name"`
}

// OtherNames returns any other names a being may have as a space-separated list
func (n *Name) OtherNames() string {
	return strings.Join(n.Other, " ")
}

func (n *Name) GetGivenName() string    { return n.GivenName }
func (n *Name) GetFamilyName() string   { return n.FamilyName }
func (n *Name) GetOtherNames() []string { return n.Other }
func (n *Name) GetDisplay() string      { return n.Display }

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

// NameStrategy is a function which describes how children are named
type NameStrategy func(b Nameable, c Cultured) *Name

// NameStrategies deliniates the various naming strategy functions
var NameStrategies = map[string]NameStrategy{
	"matrilineal": func(b Nameable, c Cultured) *Name {
		namer := c.GetNamers()[b.Sex()]
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Mother()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetName().FamilyName
			return name
		}
		name.FamilyName = namer.Words.FamilyName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patrilineal": func(b Nameable, c Cultured) *Name {
		namer := c.GetNamers()[b.Sex()]
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Father()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetName().FamilyName
			return name
		}
		name.FamilyName = namer.Words.FamilyName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"matronymic": func(b Nameable, c Cultured) *Name {
		namer := c.GetNamers()[b.Sex()]
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Mother()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetName().GivenName + namer.Words.Matronymic()
			return name
		}
		name.FamilyName = namer.Words.GivenName() + namer.Words.Matronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patronymic": func(b Nameable, c Cultured) *Name {
		namer := c.GetNamers()[b.Sex()]
		name := &Name{GivenName: namer.Words.GivenName()}
		parent, err := b.Father()
		if parent != nil && err == nil {
			name.FamilyName = parent.GetName().GivenName + namer.Words.Patronymic()
			return name
		}
		name.FamilyName = namer.Words.GivenName() + namer.Words.Patronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"onename": func(b Nameable, c Cultured) *Name {
		namer := c.GetNamers()[b.Sex()]
		name := &Name{GivenName: namer.Words.GivenName()}
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
}
