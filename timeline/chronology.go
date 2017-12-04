package timeline

type Callback func(year int)

type Chronology struct {
	CurrentYear int
	Callbacks   []Callback
	Events      map[int][]*Event
	frozen      bool
}

// NewChronology returns a initialized Chronology
func NewChronology() *Chronology {
	return &Chronology{CurrentYear: 0, Events: make(map[int][]*Event)}
}

// Freeze stops the ticks on a Chronology
func (c *Chronology) Freeze() {
	c.frozen = true
}

// AddCurrentEvent adds an event to the Chronology in the current year
func (c *Chronology) AddCurrentEvent(description string) {
	c.AddEvent(&Event{Description: description, Year: c.CurrentYear})
}

// AddEvent adds an Event to the Chronology
func (c *Chronology) AddEvent(event *Event) {
	if _, ok := c.Events[event.Year]; !ok {
		c.Events[event.Year] = []*Event{}
	}
	c.Events[event.Year] = append(c.Events[event.Year], event)
}

// Register registers a timeline.Callback to the Chronology. Each Callback in
// the slice is called per year tick.
func (c *Chronology) Register(ca Callback) {
	c.Callbacks = append(c.Callbacks, ca)
}

// Tick increments the current year by one, if unfrozen, and calls each Callback
// in the registered callbacks.
func (c *Chronology) Tick() {
	if c.frozen {
		return
	}
	c.CurrentYear++
	for _, ca := range c.Callbacks {
		go ca(c.CurrentYear)
	}
}
