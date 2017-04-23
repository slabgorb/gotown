package words_test

import (
	"math/rand"
	"testing"

	. "github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(0)
}

type testRandomStringFunc func() (string, bool)

var testRandomStringTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	// Note: order is important here to match the rand.Seed of 0
	{"trough", func() (string, bool) { return BaseNamer.Noun() }},
	{"colorless", func() (string, bool) { return BaseNamer.Adjective() }},
	{"reputations", func() (string, bool) { return BaseNamer.PluralNoun() }},
	{"fragrance", func() (string, bool) { return BaseNamer.StartNoun() }},
}

func TestStrings(t *testing.T) {
	for _, ts := range testRandomStringTable {
		test, _ := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}

var testTemplateTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"Taperedinfamy", func() (string, bool) { return BaseNamer.Name(), true }},
	{"Patternedonion", func() (string, bool) { return BaseNamer.Name(), true }},
	{"Odor of the Worries", func() (string, bool) { return BaseNamer.Name(), true }},
	{"Renown of the Spawn", func() (string, bool) { return BaseNamer.Name(), true }},
}

func TestTemplating(t *testing.T) {
	for _, ts := range testTemplateTable {
		test, _ := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}
