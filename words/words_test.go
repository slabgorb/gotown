package words_test

import (
	"os"
	"testing"

	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/random"
	. "github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
}

func TestMain(m *testing.M) {
	persist.OpenTestDB()
	defer persist.CloseTestDB()
	Seed()
	code := m.Run()
	os.Exit(code)
}

func TestWords(t *testing.T) {
	list, err := List()
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
	if noun != "lard" {
		t.Fail()
	}

}

type testRandomStringFunc func(w *Words) string

var testRandomStringTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"lard", func(w *Words) string { return w.Noun() }},
	{"living", func(w *Words) string { return w.Adjective() }},
	{"lards", func(w *Words) string { return w.PluralNoun() }},
	{"lard", func(w *Words) string { return w.StartNoun() }},
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

// func TestBackup(t *testing.T) {
// 	newWords := NewWords()
// 	newWords. = BaseWords
// 	newNamer := NewNamer([]string{"{{.Adjective}}{{.Noun}}"}, newWords, "")
// 	name := newNamer.Name()
// 	if name == "" {
// 		t.Errorf("Got empty string from newNamer.Name(), got nothing from backup")
// 	}
// 	if name != "Livinglard" {
// 		t.Errorf("Got wrong string, got %s expected %s", name, "Livinglard")
// 	}
// }

// var testTemplateTable = []struct {
// 	expected string
// 	f        testRandomStringFunc
// }{
// 	{"Never-Lard of the Lards", func(w *Words) string { return w.Name() }},
// }

// func TestTemplating(t *testing.T) {
// 	w := &Words{Name: "english town names"}
// 	if err := w.Read(); err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, ts := range testTemplateTable {
// 		test := ts.f(w)
// 		if test != ts.expected {
// 			t.Errorf("Got %s expected %s", test, ts.expected)
// 		}
// 	}
// }

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
