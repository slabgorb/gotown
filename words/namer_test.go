package words_test

import (
	"testing"

	. "github.com/slabgorb/gotown/words"
)

func TestSeed(t *testing.T) {
	list, err := NamerList()
	if err != nil {
		t.Fail()
	}
	t.Log(list)
	t.Fail()
}

// func TestPatterns(t *testing.T) {
// 	n := Namer{Name: "english towns"}
// 	if err := n.Read(); err != nil {
// 		t.Fail()
// 	}
// 	if len(n.Patterns) == 0 {
// 		t.Fail()
// 	}
// 	t.Log(n.Words)
// 	t.Log(n.Patterns)
// 	t.Fail()
// 	//n.Template()
// }

// func TestBackup(t *testing.T) {
// 	n := Namer{Name: "english towns"}
// 	if err := n.Read(); err != nil {
// 		t.Fail()
// 	}
// 	name := n.CreateName()
// 	if name == "" {
// 		t.Errorf("Got empty string from n.Name(), got nothing from backup")
// 	}
// 	if name != "Livinglard" {
// 		t.Errorf("Got wrong string, got %s expected %s", name, "Livinglard")
// 	}
// }

var testTemplateTable = []struct {
	expected string
	f        func(w *Namer) string
}{
	{"Never-Lard of the Lards", func(w *Namer) string { return w.CreateName() }},
}

// func TestTemplating(t *testing.T) {
// 	n := Namer{Name: "english towns"}
// 	if err := n.Read(); err != nil {
// 		t.Fail()
// 	}
// 	for _, ts := range testTemplateTable {
// 		test := ts.f(&n)
// 		if test != ts.expected {
// 			t.Errorf("Got %s expected %s", test, ts.expected)
// 		}
// 	}
// }
