package inhabitants

import (
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

type Expresser interface {
	Express(e genetics.Expression) map[string]string
}

var randomizer random.Generator = random.Random

func SetRandomizer(g random.Generator) {
	randomizer = g
}
