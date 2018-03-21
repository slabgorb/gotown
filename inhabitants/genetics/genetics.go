package genetics

import "github.com/slabgorb/gotown/random"

var randomizer random.Generator = random.Random

// SetRandomizer sets a randomizer for the package. Generally used by tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}
