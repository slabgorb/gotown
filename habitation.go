package townomatic

import (
	"fmt"
	"strings"
)

type Size int

const (
	Hut Size = iota
	Cottage
	House
	Tower
	Castle
	Palace
	Hamlet
	Village
	Town
	City
	NationState
	Empire
)

func (s *Size) article() string {
	art := "a"
	vowels := []string{"A", "E", "I", "O", "U"}
	for _, vowel := range vowels {
		if strings.HasPrefix(s.String(), vowel) {
			art = "an"
			break
		}
	}
	return art
}

type Habitation struct {
	Residents []*Being
}

type Area struct {
	Habitation
	Name string
	Size
	Ruler    *Being
	Location *Area
}

func NewArea(name string, size Size, ruler *Being, location *Area) *Area {
	return &Area{Name: name, Size: size, Ruler: ruler, Location: location}
}

func (a *Area) String() string {
	loc := ""
	if a.Location != nil {
		loc += fmt.Sprintf("%s, %s %s within %s", a.Name, a.Size.article(), a.Size, a.Location.Name)
	} else {
		loc += fmt.Sprintf("%s, %s %s", a.Name, a.Size.article(), a.Size)
	}
	if a.Ruler != nil {
		loc += fmt.Sprintf(" ruled by %s", a.Ruler)
	}
	loc += "."
	return loc
}

func (h *Habitation) Resident(b *Being) (int, bool) {
	found := false
	index := 0
	for i, r := range h.Residents {
		if r == b {
			found = true
			index = i
			break
		}
	}
	return index, found
}

func (h *Habitation) Add(b *Being) (ok bool) {
	_, found := h.Resident(b)
	if found {
		return false
	}
	h.Residents = append(h.Residents, b)
	return true
}

func (h *Habitation) Remove(b *Being) (ok bool) {
	index, found := h.Resident(b)
	if !found {
		return false
	}
	h.Residents = append(h.Residents[:index], h.Residents[index+1:]...)
	return true
}

func (h *Habitation) Population() int {
	return len(h.Residents)
}
