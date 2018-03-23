package locations_test

import (
	"encoding/json"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/random"
	"github.com/slabgorb/gotown/words"
)

func init() {
	SetRandomizer(random.NewMock())
	words.SetRandomizer(random.NewMock())
}

func TestJsonEncode(t *testing.T) {
	area, _ := NewArea(Town, nil, nil, nil)
	_, err := json.Marshal(area)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTo(t *testing.T) {
	a1, _ := NewArea(Town, nil, nil, nil)
	a2, _ := NewArea(Castle, nil, nil, nil)
	a3, _ := NewArea(Town, nil, nil, nil)
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

func BenchmarkPop(b *testing.B) {
	culture := &inhabitants.Culture{}
	if err := culture.Read(); err != nil {
		panic(err)
	}
	s := inhabitants.NewSpecies("Human", []inhabitants.Gender{inhabitants.Male, inhabitants.Female}, nil, nil)
	area, _ := NewArea(Town, nil, nil, nil)
	for i := 0; i < b.N; i++ {
		being := inhabitants.Being{Species: s, Culture: culture}
		being.Randomize()
		area.Add(&being)
	}

}
