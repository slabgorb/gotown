package locations_test

import (
	"testing"

	. "github.com/slabgorb/gotown/locations"
)

var (
	areaTest = []struct {
		area     *Area
		expected string
	}{
		{NewArea(Castle, nil, nil), "Larval Field"},
		{NewArea(Castle, nil, nil), "Tepid Crossing"},
		{NewArea(Town, nil, nil), "Riskyton"},
		{NewArea(Castle, nil, nil), "Northfield"},
	}
)

func TestAreaName(t *testing.T) {
	for _, ta := range areaTest {
		if ta.area.Name != ta.expected {
			t.Errorf("Expected %s got %s", ta.expected, ta.area.Name)
		}
	}
}

func TestAddTo(t *testing.T) {
	a1 := NewArea(Town, nil, nil)
	a2 := NewArea(Castle, nil, nil)
	a2.AttachTo(a1)
	if !a1.Encloses(a2) {
		t.Errorf("Encloses not registering added area")
	}
	a2.DetachFrom(a1)
	if a1.Encloses(a2) {
		t.Errorf("Detach from not detaching")
	}

}
