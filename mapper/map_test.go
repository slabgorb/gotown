package mapper_test

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"
	. "github.com/slabgorb/gotown/mapper"
)

func TestDiagram(t *testing.T) {
	diagram := VoronoiDiagram(300, 300, 150)
	dc := gg.NewContext(300, 300)
	dc.SetLineWidth(0.5)
	for _, edge := range diagram.Edges {
		fmt.Printf("%#v", edge)
		dc.DrawLine(edge.Va.X, edge.Va.Y, edge.Vb.X, edge.Vb.Y)
	}
	dc.Stroke()
	dc.SavePNG("dg.png")
}
