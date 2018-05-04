package locations_test

import (
	"encoding/json"
	"testing"

	. "github.com/slabgorb/gotown/locations"
)

func TestJsonEncode(t *testing.T) {
	area := NewArea(Town, nil, testNamer)
	_, err := json.Marshal(area)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTo(t *testing.T) {
	a1 := NewArea(Town, nil, testNamer)
	a2 := NewArea(Castle, nil, testNamer)
	a3 := NewArea(Town, nil, testNamer)
	ok := a2.AttachTo(a1)
	if !ok {
		t.Fail()
	}
	if !a1.Encloses(a2) {
		t.Errorf("Encloses not registering added area")
	}
	if !a2.IsEnclosedBy(a1) {
		t.Error("Enclosed by not registering")
	}
	a2.DetachFrom(a1)
	if a1.Encloses(a2) {
		t.Errorf("Detach from not detaching")
	}
	if ok := a3.AttachTo(a1); !ok {
		t.Fail()
	}
	if ok := a2.AttachTo(a3); !ok {
		t.Fail()
	}
	if ok := a1.AttachTo(a2); ok {
		t.Error("Should not allow adding in a circular relationship")
	}

}
