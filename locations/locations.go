package locations

import "github.com/slabgorb/gotown/random"

var randomizer random.Generator = random.Random

// SetRandomizer initializes the package-level randomizer
func SetRandomizer(g random.Generator) {
	randomizer = g
}
