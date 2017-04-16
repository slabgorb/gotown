package townomatic_test

import (
	"testing"

	. "github.com/slabgorb/townomatic"
)

func TestPopulation(t *testing.T) {
	h := Area{}
	if h.Population() != 0 {
		t.Errorf("Expected 0 got %d", h.Population())
	}
	b := &Being{Name: NewName("Adam")}
	h.Add(b)
	if h.Population() != 1 {
		t.Errorf("Expected 1 got %d", h.Population())
	}
	_, found := h.Resident(b)
	if !found {
		t.Errorf("Being not resident")
	}
	h.Remove(b)
	if h.Population() != 0 {
		t.Errorf("Expected 0 got %d", h.Population())
	}
}
