package words_test

import (
	"math/rand"
	"testing"

	. "github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(0)
}

type testRandomStringFunc func() string

var testRandomStringTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	// Note: order is important here to match the rand.Seed of 0
	{"trough", func() string { return BaseNamer.Noun() }},
	{"colorless", func() string { return BaseNamer.Adjective() }},
	{"reputations", func() string { return BaseNamer.PluralNoun() }},
	{"stick", func() string { return BaseNamer.StartNoun() }},
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
	newNamer := NewNamer([]string{"{{.Adjective}}{{.Noun}}"}, newWords)
	name := newNamer.Name()
	if name == "" {
		t.Errorf("Got empty string from newNamer.Name(), got nothing from backup")
	}
	if name != "Brilliantonion" {
		t.Errorf("Got wrong string, got %s expected %s", name, "Brilliantonion")
	}
}

var testTemplateTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"Worry of the Blots", func() string { return BaseNamer.Name() }},
	{"Bait of the Games", func() string { return BaseNamer.Name() }},
	{"The Leaders", func() string { return BaseNamer.Name() }},
	{"Everthief", func() string { return BaseNamer.Name() }},
}

func TestTemplating(t *testing.T) {
	for _, ts := range testTemplateTable {
		test := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}
