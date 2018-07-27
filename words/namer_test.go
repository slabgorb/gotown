package words_test

import (
	"testing"

	"github.com/slabgorb/gotown/persist"

	. "github.com/slabgorb/gotown/words"
)

func TestSeed(t *testing.T) {
	list, err := NamerList()
	if err != nil {
		t.Fail()
	}
	if len(list) == 0 {
		t.Error("no namers seeded")
	}
}

func TestPatterns(t *testing.T) {
	n := Namer{Name: "english towns"}
	if err := persist.ReadByName(n.Name, "Namer", &n); err != nil {
		t.Fail()
	}
	if len(n.Patterns) == 0 {
		t.Fail()
	}
	expected := "{{.Noun}}{{.Suffix}}"
	template := n.Template()
	actual := template.Name()
	if expected != actual {
		t.Errorf("expected %s actual %s", expected, actual)
	}
}

func TestBackup(t *testing.T) {
	n := Namer{Name: "english towns"}
	if err := persist.ReadByName(n.Name, "Namer", &n); err != nil {
		t.Fail()
	}
	name := n.CreateName()
	if name == "" {
		t.Errorf("Got empty string from n.Name(), got nothing from backup")
	}
	if name != "Lanternkirk" {
		t.Errorf("Got wrong string, got %s expected %s", name, "Lanternkirk")
	}
}
