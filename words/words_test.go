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
	{"fragrance", func() string { return BaseNamer.StartNoun() }},
}

func TestStrings(t *testing.T) {
	for _, ts := range testRandomStringTable {
		test := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}

var testTemplateTable = []struct {
	expected string
	f        testRandomStringFunc
}{
	{"Taperedinfamy", func() string { return BaseNamer.Name() }},
}

func TestTemplating(t *testing.T) {
	for _, ts := range testTemplateTable {
		test := ts.f()
		if test != ts.expected {
			t.Errorf("Got %s expected %s", test, ts.expected)
		}
	}
}
