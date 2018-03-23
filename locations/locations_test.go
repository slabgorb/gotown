package locations_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

type tester interface {
	Fail()
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	inhabitants.Seed()
	words.Seed()
	code := m.Run()
	os.Exit(code)
}
