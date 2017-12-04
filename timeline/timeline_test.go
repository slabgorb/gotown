package timeline_test

import (
	"strconv"
	"testing"

	. "github.com/slabgorb/gotown/timeline"
)

func TestEvents(t *testing.T) {
	c := NewChronology()
	c.Tick()
	e := c.EventsForYear(1)
	if len(e) > 0 {
		t.Error("Expeted 0 events")
	}
	c.AddCurrentEvent("Year One")
	e = c.EventsForYear(1)
	if len(e) < 1 {
		t.Error("Expected events ")
	}
	if e[0].Description != "Year One" {
		t.Errorf("Got bad description")
	}
}

func TestChronology(t *testing.T) {
	c := NewChronology()
	foo := "nope"
	c.Register(func(year int) {
		foo = strconv.Itoa(year)
	})
	c.Tick()
	if c.CurrentYear != 1 {
		t.Error("expected Tick() to increment year")
	}
	c.Tick()
	c.Tick()
	c.Tick()
	if foo != "4" {
		t.Errorf("Expected simple callback to succeed, expected '4' got '%s'", foo)
	}
	c.Freeze()
	c.Tick()
	if foo != "4" {
		t.Errorf("Expected frozen chronology callback to not be called, expected '4' got '%s'", foo)
	}
}
