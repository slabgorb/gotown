package inhabitants

import (
	"math/rand"
	"time"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

type Expresser interface {
	Express(e genetics.Expression) map[string]string
}

var randomizer random.Generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func SetRandomizer(g random.Generator) {
	randomizer = g
}
