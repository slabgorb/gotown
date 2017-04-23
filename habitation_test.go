package gotown_test

import (
	"testing"
	"text/template"

	. "github.com/slabgorb/gotown"
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

var nameTemplate, _ = template.New("test").Parse("{{.GivenName}} {{.FamilyName}}")
var human = &Species{
	NameTemplate: nameTemplate,
}
var areaTest = []struct {
	area     *Area
	expected string
}{
	{NewArea(Castle, nil, nil), "Larval Field"},
	{NewArea(Castle, nil, nil), "Tepid Crossing"},
	{NewArea(Town, nil, nil), "Riskyton"},
	{NewArea(Castle, nil, nil), "Northfield"},
}

func TestAreaName(t *testing.T) {
	for _, ta := range areaTest {
		if ta.area.Name != ta.expected {
			t.Errorf("Expected %s got %s", ta.expected, ta.area.Name)
		}
	}
}
