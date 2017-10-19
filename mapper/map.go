package mapper

import (
	"math/rand"
	"time"

	"github.com/pzsz/voronoi"
	vu "github.com/pzsz/voronoi/utils"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func SetRandomizer(g random.Generator) {
	randomizer = g
}

func VoronoiDiagram(x, y int, pointCount int, relaxation int) *voronoi.Diagram {
	sites := []voronoi.Vertex{}
	for i := 0; i < pointCount; i++ {
		vx := randomizer.Float64() * float64(x)
		vy := randomizer.Float64() * float64(y)
		sites = append(sites, voronoi.Vertex{X: vx, Y: vy})
	}
	bbox := voronoi.NewBBox(0, float64(x), 0, float64(y))
	d := voronoi.ComputeDiagram(sites, bbox, true)
	for i := 0; i < relaxation; i++ {
		v := vu.LloydRelaxation(d.Cells)
		d = voronoi.ComputeDiagram(v, bbox, true)
	}
	return d

}
