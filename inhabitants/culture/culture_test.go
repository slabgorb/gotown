package culture_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

var testSpecies = &species.Species{Name: "human"}

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	defer persist.CloseTestDB()
	words.Seed()
	species.Seed()
	Seed()
	if err := testSpecies.Read(); err != nil {
		panic(err)
	}
	code := m.Run()

	os.Exit(code)
}

type MarriageableStub struct {
	alive   bool
	married bool
	sex     inhabitants.Gender
	age     int
}

func (ms *MarriageableStub) Alive() bool {
	return ms.alive
}

func (ms *MarriageableStub) Unmarried() bool {
	return !ms.married
}

func (ms *MarriageableStub) Sex() inhabitants.Gender {
	return ms.sex
}

func (ms *MarriageableStub) IsCloseRelativeOf(Marriageable) bool {
	panic("not implemented")
}

func (ms *MarriageableStub) Age() int {
	return ms.age
}

func TestSeed(t *testing.T) {
	list, err := List()
	if err != nil {
		panic(err)
	}

	found := false
	for _, v := range list {
		if v == "italianate" {
			found = true
		}
	}
	t.Log(list)
	if !found {
		t.Fatal("italianate not seeded")
	}

	w := &Culture{Name: "italianate"}
	if err := w.Read(); err != nil {
		t.Fatal(err)
	}
	t.Log(w.NamerNames)
	namers := w.GetNamers()
	_, ok := namers[inhabitants.Male]
	if !ok {
		t.Error("did not load male names")
	}

}
