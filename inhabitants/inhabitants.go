package inhabitants

import (
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}

// Expresser defines the behavior for anything that can express genetics
type Expresser interface {
	Express(e genetics.Expression) map[string]string
}
