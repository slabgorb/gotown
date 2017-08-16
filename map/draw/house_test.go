package draw

import (
	"testing"

	"github.com/fogleman/gg"
)

func TestRoof(t *testing.T) {
	s := NewShape(50, 50, 25, 25, 30)
	dc := gg.NewContext(200, 200)
	Roof(dc, s)
	s = NewShape(70, 50, 25, 25, -30)
	Roof(dc, s)
	dc.SavePNG("roof.png")
}

func TestGarden(t *testing.T) {
	s := NewShape(150, 150, 125, 125, 30)
	dc := gg.NewContext(300, 300)
	Garden(dc, s)
	s = NewShape(170, 120, 125, 125, -30)
	Garden(dc, s)
	dc.SavePNG("garden.png")
}

func TestHouseWithGarden(t *testing.T) {
	s := NewShape(150, 150, 125, 125, 30)
	dc := gg.NewContext(300, 300)
	HouseWithGarden(dc, s)
	s = NewShape(170, 120, 125, 125, -30)
	HouseWithGarden(dc, s)
	dc.SavePNG("house_with_garden.png")
}
