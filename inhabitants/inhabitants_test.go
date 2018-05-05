package inhabitants_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

func TestMain(m *testing.M) {
	os.Exit(testMainWrapped(m))
}

func testMainWrapped(m *testing.M) int {
	persist.OpenTestDB()
	defer persist.CloseTestDB()
	words.Seed()
	culture.Seed()
	species.Seed()
	return m.Run()
}
