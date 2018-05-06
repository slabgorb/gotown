package inhabitants

import (
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/persist"
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
	Express(e Expresser) map[string]string
	Expression(string) (string, int)
	GetTraits() []Expresser
	GetName() string
}

// demographic constants
const (
	Child int = iota
	Teenager
	YoungAdult
	EarlyAdult
	Adult
	MiddleAge
	Senior
	Elderly
	Ancient
	MaxDemographyBucket
)

// Populatable abstracts...
type Populatable interface {
	Alive() bool
	Sex() Gender
	GetAge() int
	Die(...string)
}

// Readable abstracts the ability to read from a database
type Readable interface {
	persist.Persistable
}

// Specieser abstracts...
type Specieser interface {
	Readable
	RandomAge(slot int) int
	MaxAge(slot int) int
	GetGenders() []Gender
	Expression() genetics.Expression
}

// Namer abstracts ...
type Namer interface {
	GetDisplay() string
	GetGivenName() string
	GetFamilyName() string
}
