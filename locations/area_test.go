package locations_test

import (
	"testing"

	. "github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
	words.SetRandomizer(random.NewMock())
}

func TestAddTo(t *testing.T) {
	a1 := NewArea(Town, nil, nil)
	a2 := NewArea(Castle, nil, nil)
	a3 := NewArea(Town, nil, nil)
	a2.AttachTo(a1)
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
