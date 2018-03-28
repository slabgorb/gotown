package inhabitants

import (
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
)

// Expresser defines the behavior for anything that can express genetics
type Expresser interface {
	Express(e genetics.Expression) map[string]string
}

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}

func Seed() {
	if err := seedSpecies(); err != nil {
		panic(err)
	}
	if err := seedCultures(); err != nil {
		panic(err)
	}
}

func seedSpecies() error {
	var species = &Species{}
	return persist.SeedHelper("species", species)
}

func seedCultures() error {
	var culture = &Culture{}
	return persist.SeedHelper("cultures", culture)
}
