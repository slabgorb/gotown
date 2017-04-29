package gotown

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type Name struct {
	GivenName  string
	FamilyName string
	Other      []string
	Display    string
}

func (n *Name) OtherNames() string {
	return strings.Join(n.Other, " ")
}

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

type Being struct {
	*Name
	*Species
	Parents  map[Gender]*Being
	Children []*Being
	Age      int
	Gender
	Dead bool
}

func (b *Being) genderedParent(gender Gender) *Being {
	if parent, ok := b.Parents[gender]; ok {
		return parent
	}
	return nil
}

func (b *Being) Father() *Being {
	return b.genderedParent(Male)
}

func (b *Being) Mother() *Being {
	return b.genderedParent(Female)
}

func (b *Being) Randomize() error {
	if b.Species == nil {
		return fmt.Errorf("Cannot randomize a being without a species")
	}
	possibleGenders := []Gender{}
	for g, _ := range b.Species.Genders {
		possibleGenders = append(possibleGenders, g)
	}

	b.Gender = possibleGenders[rand.Intn(len(possibleGenders))]
	b.Name = b.Species.Genders[b.Gender].NameStrategy(b)
	b.Age = rand.Intn(int(math.Floor(float64(b.Species.Genders[b.Gender].Fertility.End) * 1.3)))
	return nil
}

// Reproduce creates new Being objects from the 'parent' beings
func (b *Being) Reproduce(with *Being) ([]*Being, error) {
	if with == nil && b.Gender != Asexual {
		return nil, fmt.Errorf("Being %s cannot reproduce asexually", b)
	}
	//children := []*Being{}

	child := &Being{Species: b.Species}
	child.Parents = map[Gender]*Being{
		b.Gender:    b,
		with.Gender: with,
	}
	child.Randomize()

	b.Children = append(b.Children, child)

	return b.Children, nil
}

func (b *Being) Die() {
	b.Dead = true
}

func (b *Being) String() string {
	return strings.Trim(b.Name.Display, " ")
}

func (b *Being) Alive() bool {
	return !b.Dead
}
