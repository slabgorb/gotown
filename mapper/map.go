package mapper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pzsz/voronoi"
	vu "github.com/pzsz/voronoi/utils"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

type HeightMap struct {
	diagram voronoi.Diagram
	heights map[int]float64
}

func SetRandomizer(g random.Generator) {
	randomizer = g
}

func RandomHeight() float64 {
	return randomizer.Float64()
}

func RandomCell(diagram voronoi.Diagram) *voronoi.Cell {
	r := randomizer.Intn(len(diagram.Cells))
	return diagram.Cells[r]
}

func VoronoiDiagram(x, y int, pointCount int, relaxation int) *voronoi.Diagram {
	t := time.Now()
	sites := []voronoi.Vertex{}
	for i := 0; i < pointCount; i++ {
		vx := randomizer.Float64() * float64(x)
		vy := randomizer.Float64() * float64(y)
		sites = append(sites, voronoi.Vertex{X: vx, Y: vy})
	}
	fmt.Println(time.Since(t))
	t = time.Now()
	bbox := voronoi.NewBBox(0, float64(x), 0, float64(y))
	d := voronoi.ComputeDiagram(sites, bbox, true)
	fmt.Println(time.Since(t))
	for i := 0; i < relaxation; i++ {
		t = time.Now()
		v := vu.LloydRelaxation(d.Cells)
		d = voronoi.ComputeDiagram(v, bbox, true)
		fmt.Println(time.Since(t))
	}
	return d
}

func FindNeighborCells(diagram *voronoi.Diagram, cell *voronoi.Cell) []*voronoi.Cell {
	cells := []*voronoi.Cell{}
	for _, h := range cell.Halfedges {
		cells = append(cells, h.Edge.GetOtherCell(cell))
	}
	return cells

}
