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
