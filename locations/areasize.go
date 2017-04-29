package locations

import "strings"

type AreaSize int

const (
	Hut AreaSize = iota
	Cottage
	House
	Tower
	Castle
	Hamlet
	Palace
	Village
	Town
	City
	Region
	NationState
	Empire
)

func (s *AreaSize) article() string {
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
