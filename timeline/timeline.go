package timeline

type Callback func(year int)

type Chronology struct {
	CurrentYear int
	Callbacks   []Callback
}

func NewChronology() *Chronology {
	return &Chronology{CurrentYear: 0}
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
