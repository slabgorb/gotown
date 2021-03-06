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

var testNamer = &words.Namer{Name: "english towns"}
var testSpecies = &species.Species{Name: "human"}
var testCulture = &culture.Culture{Name: "italianate"}

func testMainWrapped(m *testing.M) int {
	SetRandomizer(random.NewMock())
	words.SetRandomizer(random.NewMock())
	persist.OpenTestDB()
	species.Seed()
	culture.Seed()
	words.Seed()
	if err := persist.ReadByName(testNamer.Name, "Namer", testNamer); err != nil {
		panic(err)
	}
	if err := persist.ReadByName(testCulture.Name, "Culture", testCulture); err != nil {
		panic(err)
	}
	if err := persist.ReadByName(testSpecies.Name, "Species", testSpecies); err != nil {
		panic(err)
	}

	defer persist.CloseTestDB()
	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapped(m))
}
