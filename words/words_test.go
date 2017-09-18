package words_test

import (
	"testing"

	"github.com/slabgorb/gotown/random"
	. "github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
}

type testRandomStringFunc func() string

var testRandomStringTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"lard", func() string { return BaseNamer.Noun() }},
	{"living", func() string { return BaseNamer.Adjective() }},
	{"lards", func() string { return BaseNamer.PluralNoun() }},
	{"lard", func() string { return BaseNamer.StartNoun() }},
}

func TestStrings(t *testing.T) {
	for _, ts := range testRandomStringTable {
		test := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}

func TestBackup(t *testing.T) {
	newWords := NewWords()
	newWords.Backup = BaseWords
	newNamer := NewNamer([]string{"{{.Adjective}}{{.Noun}}"}, newWords, "")
	name := newNamer.Name()
	if name == "" {
		t.Errorf("Got empty string from newNamer.Name(), got nothing from backup")
	}
	if name != "Livinglard" {
		t.Errorf("Got wrong string, got %s expected %s", name, "Livinglard")
	}
}

var testTemplateTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"Never-Lard of the Lards", func() string { return BaseNamer.Name() }},
}

func TestTemplating(t *testing.T) {
	for _, ts := range testTemplateTable {
		test := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}

func TestNameWords(t *testing.T) {
	w := NorseMaleNameWords
	pt := w.Patronymic()
	if pt != "son" {
		t.Errorf("Expected 'son' for patronymic, got %s", pt)
	}
}
