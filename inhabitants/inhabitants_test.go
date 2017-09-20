package inhabitants_test

import (
	"github.com/slabgorb/gotown/words"

	. "github.com/slabgorb/gotown/inhabitants"
)

var (
	female      = NewSpeciesGender(words.NorseFemaleNamer, NameStrategies["matronymic"], 12, 48)
	male        = NewSpeciesGender(words.NorseMaleNamer, NameStrategies["patronymic"], 12, 65)
	mockSpecies = NewSpecies("Northman", map[Gender]*SpeciesGender{
		Female: female,
		Male:   male,
	}, nil)
)
