package gotown_test

import (
	"testing"

	. "github.com/slabgorb/gotown"
)

func testToString(t *testing.T) {
	species := Species{Name: "Vampire"}
	if species.String() != "Vampire" {
		t.Fail()
	}
}
