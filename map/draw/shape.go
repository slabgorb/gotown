package draw

import "github.com/fogleman/gg"

type shape struct {
	x, y          float64
	height, width float64
	angle         float64
}

type color struct {
	r, g, b float64
}

var (
	Black = color{0, 0, 0}
	Gray  = color{0.5, 0.5, 0.5}
)

func (s *shape) midWidth() float64 {
	return s.x + (s.width / 2)
}
func (s *shape) midHeight() float64 {
	return s.y + (s.height / 2)
}
func (s *shape) inBounds(x, y float64) bool {
	return (x > s.x && x < s.x+s.width) && (y > s.y && y < s.y+s.height)
}

func (s *shape) shrink(amount float64) *shape {
	return NewShape(s.x+amount, s.y+amount, s.height-(amount*2), s.width-(amount*2), s.angle)
}

func (s *shape) grow(amount float64) *shape {
	return NewShape(s.x-amount, s.y-amount, s.height+(amount*2), s.width+(amount*2), s.angle)
}

func (s *shape) offset(amount float64) *shape {
	return NewShape(s.x+amount, s.y+amount, s.height, s.width, s.angle)
}

func withRotation(dc *gg.Context, s *shape, f shaper) {
	dc.Push()
	dc.RotateAbout(gg.Radians(s.angle), s.midWidth(), s.midHeight())
	f(dc, s)
	dc.Pop()
}

func NewShape(x, y, height, width, angle float64) *shape {
	return &shape{x: x, y: y, height: height, width: width, angle: angle}
}

type shaper func(dc *gg.Context, s *shape)
