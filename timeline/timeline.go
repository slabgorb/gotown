package timeline

import "github.com/slabgorb/gotown/events"

type Callback func(year int)

type Chronology struct {
	CurrentYear int
	Callbacks   []Callback
	Events      []events.Event
}

func NewChronology() *Chronology {
	return &Chronology{CurrentYear: 0}
}

func (c *Chronology) AddEvent(description string) {
	c.Events = append(c.Events, events.Event{Description: description, Year: c.CurrentYear})
}

func (c *Chronology) Register(ca Callback) {
	c.Callbacks = append(c.Callbacks, ca)
}

func (c *Chronology) Tick() {
	c.CurrentYear++
	for _, ca := range c.Callbacks {
		go ca(c.CurrentYear)
	}
}
