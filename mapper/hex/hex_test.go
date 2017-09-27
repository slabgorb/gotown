package hex

import (
	"testing"

	"github.com/fogleman/gg"
)

func TestCube(t *testing.T) {
	c1 := Cube{10, 10, 10}
	c2 := Cube{11, 10, 10}
	if c2.distance(c1) != 1 {
		t.Errorf("Expected a distance of 1, got %f", c2.distance(c1))
	}
	a1 := c2.ToAxial()
	if a1.q != c2.x && a1.r != c2.z {
		t.Errorf("Unexpected axial conversion")
	}
}

func TestAxial(t *testing.T) {
	a1 := Axial{10, 10}
	a2 := Axial{11, 10}
	if a2.distance(a1) != 1 {
		t.Errorf("Expected a distance of 1, got %f", a2.distance(a1))
	}
	c1 := a2.ToCube()
	a3 := c1.ToAxial()
	if a2.q != a3.q && a2.r != a3.r {
		t.Errorf("conversion to cube and back to axial failed")
	}
	surrounding := a2.Surrounding()
	for _, hex := range surrounding {
		if hex.distance(a2) != 1 {
			t.Errorf("surrounding hex should have a distance of 1")
		}
	}
}
func TestHexPaper(t *testing.T) {
	ag := NewSquareGrid(12.0)
	for q := 0.0; q < 60.0; q++ {
		for r := 0.0; r < 60.0; r++ {
			ag.SetHex(q, r)
		}
	}
	dc := gg.NewContext(500, 500)
	for _, h := range ag.Hexes {
		poly := ag.HexPolygon(*h)
		dc.MoveTo(poly[0][0], poly[0][1])
		for _, point := range poly {
			dc.LineTo(point[0], point[1])

		}
		dc.Stroke()
	}
	dc.SavePNG("hexpaper.png")
}
