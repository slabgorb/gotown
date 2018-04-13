package locations_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/species"
	. "github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

type tester interface {
	Fail()
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func TestMain(m *testing.M) {
	SetRandomizer(random.NewMock())
	words.SetRandomizer(random.NewMock())
	persist.OpenTestDB()
	species.Seed()
	culture.Seed()
	words.Seed()
	code := m.Run()
	os.Exit(code)
}
