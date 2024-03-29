package locations_test

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/slabgorb/gotown/inhabitants/being"
	. "github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/logger"
)

func makePop(t *testing.T, count int) (*being.Population, []*being.Being) {
	ids := []string{}
	beings := []*being.Being{}
	var wg sync.WaitGroup
	beingQueue := make(chan *being.Being, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			bg := being.New(testSpecies, testCulture, logger.Default)
			bg.Randomize()
			if err := bg.Save(); err != nil {
				t.Fatalf("could not save being:%s", err)
			}
			beingQueue <- bg
		}()
	}
	go func(wg *sync.WaitGroup) {
		for b := range beingQueue {
			ids = append(ids, b.ID)
			beings = append(beings, b)
			wg.Done()
		}
	}(&wg)
	wg.Wait()
	return being.NewPopulation(ids, logger.Default), beings
}

func TestJsonEncode(t *testing.T) {
	area := NewArea(Town, nil, testNamer)
	_, err := json.Marshal(area)
	if err != nil {
		t.Error(err)
	}
}

func persistRoundtripArea(area *Area) error {
	if err := area.Save(); err != nil {
		return err
	}
	id := area.ID
	area.Reset()
	area.ID = id
	if err := area.Read(); err != nil {
		return err
	}
	return nil
}

func TestPersist(t *testing.T) {
	area := NewArea(Town, nil, testNamer)
	p, _ := makePop(t, 10)
	area.Residents = p
	if err := persistRoundtripArea(area); err != nil {
		t.Fatal(err)
	}
	pop, err := area.Population()
	if err != nil {
		t.Fatal(err)
	}
	if pop != 10 {
		t.Errorf("Expected 10 got %d", pop)
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

func TestAPI(t *testing.T) {
	area := NewArea(Town, nil, testNamer)
	p, _ := makePop(t, 10)
	area.Residents = p
	if err := persistRoundtripArea(area); err != nil {
		t.Fatal(err)
	}
	_, err := area.API()
	if err != nil {
		t.Fatal(err)
	}
}
