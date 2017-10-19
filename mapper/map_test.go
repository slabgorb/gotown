package mapper_test

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"
	vu "github.com/pzsz/voronoi/utils"
	. "github.com/slabgorb/gotown/mapper"
)

func TestDiagram(t *testing.T) {
	x := 300
	y := 300
	diagram := VoronoiDiagram(x, y, 50, 8)
	patterns := []gg.Pattern{}
	for i := 1; i < 13; i++ {
		im, err := gg.LoadPNG(fmt.Sprintf("draw/textures/seamless-watercolor-%d.png", i))
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
	for _, cell := range diagram.Cells {
		dc.SetLineJoinRound()
		dc.SetLineWidth(2.0)
		//		var v *voronoi.EdgeVertex
		for _, he := range cell.Halfedges {
			a := he.GetStartpoint()
			b := he.GetEndpoint()
			dc.LineTo(a.X, a.Y)
			dc.LineTo(b.X, b.Y)
			centroid := vu.CellCentroid(cell)
			dc.DrawPoint(centroid.X, centroid.Y, 5)
		}
		dc.SetFillStyle(patterns[currentPattern])
		dc.FillPreserve()
		dc.Stroke()
		currentPattern++
		if currentPattern == len(patterns) {
			currentPattern = 0
		}

	}
	// for _, edge := range diagram.Edges {
	// 	fmt.Printf("%#v", edge)
	// 	dc.DrawLine(edge.Va.X, edge.Va.Y, edge.Vb.X, edge.Vb.Y)
	// }
	// dc.SetFillStyle(patterns[currentPattern])
	// dc.Fill()

	dc.SavePNG("dg.png")
}
