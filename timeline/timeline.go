package timeline

type Callback func(year int)

type Chronology struct {
	CurrentYear int
	Callbacks   []Callback
	Events      []Event
	frozen      bool
}

type Event struct {
	Description string
	Year        int
}

func NewChronology() *Chronology {
	return &Chronology{CurrentYear: 0}
}

func (c *Chronology) Freeze() {
	c.frozen = true
}

func (c *Chronology) AddEvent(description string) {
	c.Events = append(c.Events, Event{Description: description, Year: c.CurrentYear})
}

func (c *Chronology) Register(ca Callback) {
	c.Callbacks = append(c.Callbacks, ca)
}

func (c *Chronology) Tick() {
	if c.frozen {
		return
	}
	c.CurrentYear++
	for _, ca := range c.Callbacks {
		go ca(c.CurrentYear)
	}
}
