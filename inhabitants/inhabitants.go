package inhabitants

import (
	"math/rand"
	"time"

	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func SetRandomizer(g random.Generator) {
	randomizer = g
}
