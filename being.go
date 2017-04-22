package gotown

import (
	"bytes"
	"fmt"
	"text/template"
)

type Name struct {
	GivenName  string
	FamilyName string
	OtherNames []string
}

func NewName(names ...string) Name {
	name := Name{}
	if len(names) > 0 {
		name.GivenName = names[0]
	}
	if len(names) > 1 {
		name.FamilyName = names[1]
	}
	if len(names) > 2 {
		name.OtherNames = names[2:]
	}
	return name
}

func (n *Name) Patterned(t *template.Template) string {
	buf := bytes.NewBuffer([]byte(""))
	t.Execute(buf, n)
	return buf.String()
}

type Being struct {
	Name
	*Species
	Parents  []Being
	Children []Being
	Age      int
	Gender
	Dead bool
}

func (b *Being) Reproduce(with *Being) ([]*Being, error) {
	if with == nil && b.Gender != Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	children := []*Being{}
	return children, nil
}

func (b *Being) Die() {
	b.Dead = true
}

func (b *Being) String() string {
	return b.Name.Patterned(b.Species.NameTemplate)
}

func (b *Being) Alive() bool {
	return !b.Dead
}
