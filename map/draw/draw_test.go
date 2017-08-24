package draw

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"
)

var testCases = []struct {
	name string
	s    *shape
	dc   *gg.Context
	f    shaper
}{
	{
		"roof",
		NewShape(50, 50, 25, 25, 30),
		gg.NewContext(200, 200),
		Roof,
	},
	{
		"road",
		NewShape(50, 50, 25, 125, 0),
		gg.NewContext(200, 200),
		Road,
	},
	{
		"road2",
		NewShape(50, 50, 125, 125, 0),
		gg.NewContext(200, 200),
		Road,
	},
	{
		"garden",
		NewShape(50, 50, 125, 125, 30),
		gg.NewContext(200, 200),
		Garden,
	},
	{
		"farm",
		NewShape(50, 50, 125, 125, 30),
		gg.NewContext(200, 200),
		Farm,
	},
	{
		"house_with_garden",
		NewShape(50, 50, 125, 125, 30),
		gg.NewContext(200, 200),
		HouseWithGarden,
	},
}

func TestDrawing(t *testing.T) {
	for _, tc := range testCases {
		tc.f(tc.dc, tc.s)
		tc.dc.SavePNG(fmt.Sprintf("%s.png", tc.name))
	}
}
