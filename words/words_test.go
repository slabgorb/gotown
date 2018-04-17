package words_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
	. "github.com/slabgorb/gotown/words"
)

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	SetRandomizer(random.NewMock())
	defer persist.CloseTestDB()
	Seed()
	code := m.Run()
	os.Exit(code)
}

func TestWords(t *testing.T) {
	list, err := WordsList()
	if err != nil {
		panic(err)
	}
	found := false
	for _, v := range list {
		if v == "english town names" {
			found = true
		}
	}
	t.Log(list)
	if !found {
		t.Fatal("english town names not seeded")
	}

	w := &Words{Name: "english town names"}
	if err := w.Read(); err != nil {
		t.Fatal(err)
	}
	noun := w.Noun()
	if noun != "lantern" {
		t.Fail()
	}

}

type testRandomStringFunc func(w *Words) string

var testRandomStringTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"lantern", func(w *Words) string { return w.Noun() }},
	{"living", func(w *Words) string { return w.Adjective() }},
	{"lanterns", func(w *Words) string { return w.PluralNoun() }},
	{"lantern", func(w *Words) string { return w.StartNoun() }},
}

func TestStrings(t *testing.T) {
	w := &Words{Name: "english town names"}
	if err := w.Read(); err != nil {
		t.Fatal(err)
	}
	for _, ts := range testRandomStringTable {
		test := ts.f(w)
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}

func TestNameWords(t *testing.T) {
	w := &Words{Name: "viking male names"}
	if err := w.Read(); err != nil {
		t.Fail()
	}
	t.Log(w.Dictionary)
	pt := w.Patronymic()
	if pt != "son" {
		t.Errorf("Expected 'son' for patronymic, got %s", pt)
	}
}
