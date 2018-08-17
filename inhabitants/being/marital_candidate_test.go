package being_test

import (
	"testing"

	. "github.com/slabgorb/gotown/inhabitants/being"
)

func TestContains(t *testing.T) {
	bf := beingFixtures
	adam := bf["adam"]
	eve := bf["eve"]
	mc := NewMaritalCandidate(adam, eve)
	t.Log(mc)
	cain := bf["cain"]
	if mc.Contains(cain) {
		t.Error("failed, should not contain Cain")
	}
	if !mc.Contains(adam) {
		t.Error("failed, should contain Adam")
	}
	if !mc.Contains(eve) {
		t.Error("failed, should contain Eve ")
	}
}
