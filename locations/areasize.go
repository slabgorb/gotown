package locations

import "strings"

// AreaSize is an enum
type AreaSize int

// enum for AreaSize
const (
	Hut AreaSize = iota
	Cottage
	House
	Tower
	Hamlet
	Palace
	Village
	Castle
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

func (s AreaSize) Population() int {
	switch s {
	case Hut:
		return randomizer.Intn(4)
	case Cottage:
		return randomizer.Intn(5)
	case House:
		return randomizer.Intn(10)
	case Tower:
		return randomizer.Intn(25)
	case Hamlet:
		return randomizer.Intn(60) + 10
	case Palace:
		return randomizer.Intn(100) + 25
	case Village:
		return randomizer.Intn(200) + 50
	case Castle:
		return randomizer.Intn(250)
	case Town:
		return randomizer.Intn(1000) + 100
	default:
		return randomizer.Intn(1000) + 100 // let's not go higher than this until we figure out scaling
	}
}
