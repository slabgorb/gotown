package townomatic

import "testing"

func testToString(t *testing.T) {
	species := Species{Name: "Vampire"}
	if species.String() != "Vampire" {
		t.Fail()
	}
}
