package locations

import "github.com/slabgorb/gotown/random"

var randomizer random.Generator = random.Random

func SetRandomizer(g random.Generator) {
	randomizer = g
}
