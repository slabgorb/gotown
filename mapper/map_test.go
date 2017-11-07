package mapper_test

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"
	. "github.com/slabgorb/gotown/mapper"
)

func TestDiagram(t *testing.T) {
	x := 600
	y := 400
	diagram := VoronoiDiagram(x, y, 8000, 1)
	patterns := []gg.Pattern{}
	for i := 0; i < 12; i++ {
		im, err := gg.LoadPNG(fmt.Sprintf("draw/textures/%d.png", i))
		if err != nil {
			t.Fatal("Could not load pattern")
		}
		pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
		patterns = append(patterns, pattern)
	}
	dc := gg.NewContext(x, y)
	dc.SetLineWidth(0.5)
	dc.SetHexColor("#000000")
	currentPattern := 0
	dc.SetLineJoinRound()
	for _, cell := range diagram.Cells {
		dc.SetLineWidth(1.0)
		//		var v *voronoi.EdgeVertex
		for _, he := range cell.Halfedges {
			a := he.GetStartpoint()
			b := he.GetEndpoint()
			dc.LineTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
			//centroid := vu.CellCentroid(cell)
			//dc.DrawPoint(centroid.X, centroid.Y, 3)
		}
		dc.SetFillStyle(patterns[currentPattern])
		dc.FillPreserve()
		dc.Stroke()
		currentPattern++
		if currentPattern == len(patterns) {
			currentPattern = 0
		}

	}

	dc.SavePNG("dg.png")
}
