package townomatic

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
	Alive bool
}

func (b *Being) Reproduce(with *Being) ([]*Being, error) {
	if with == nil && b.Gender != Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	//children := []*Being{}

	return nil, nil
}

func (b *Being) Die() {
	b.Alive = false
}

func (b *Being) String() string {
	return b.Name.Patterned(b.Species.NameTemplate)
}
