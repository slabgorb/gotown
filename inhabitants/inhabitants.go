package inhabitants

import (
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
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

type Populatable interface {
	Alive() bool
	Sex() Gender
	Age() int
	History() *timeline.Chronology
	GetName() *Name
	Die(...string)
}

type Relatable interface {
	IsChildOf(Relatable) bool
	IsParentOf(Relatable) bool
	IsSiblingOf(Relatable) bool
	IsCloseRelativeOf(Relatable) bool
	GetChildren() []Relatable
}

type Marriageable interface {
	Relatable
	Populatable
	Unmarried() bool
}

type Specieser interface {
	RandomAge(slot int) int
	MaxAge(slot int) int
	GetGenders() []Gender
	Expression() genetics.Expression
}

type Cultured interface {
	GetName(Nameable) *Name
	MaritalCandidate(Marriageable, Marriageable) bool
	GetNamers() map[Gender]*words.Namer
}

type Nameable interface {
	Father() Nameable
	Mother() Nameable
	//Culture() Cultured
	GetName() *Name
	Sex() Gender
}
type Namer interface {
	GetDisplay() string
	GetGivenName() string
	GetFamilyName() string
}
