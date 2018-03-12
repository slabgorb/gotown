package locations_test

import (
	"encoding/json"
	"testing"

	"github.com/slabgorb/gotown/inhabitants"
	. "github.com/slabgorb/gotown/locations"
)

func TestPopulation(t *testing.T) {
	culture := helperMockCulture(t, "italian")
	h := NewArea(Village, culture, nil, nil)
	if h.Population() != 0 {
		t.Errorf("Expected 0 got %d", h.Population())
	}
	b := &inhabitants.Being{Name: inhabitants.NewName("Adam")}
	h.Add(b)
	if h.Population() != 1 {
		t.Errorf("Expected 1 got %d", h.Population())
	}
	found := h.Resident(b)
	if !found {
		t.Errorf("Being not resident")
	}
	h.Remove(b)
	if h.Population() != 0 {
		t.Errorf("Expected 0 got %d", h.Population())
	}
	j, err := json.Marshal(h)
	if err != nil {
		t.Error(err)
	}
	h2 := &Area{}
	err = json.Unmarshal(j, h2)
	if err != nil {
		t.Error(err)
	}
}
