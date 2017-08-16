package voronoi

type Cell struct {
	Site
	HalfEdges []HalfEdge
}

type HalfEdge struct {
	Site
}

type Edge struct {
	Left, Right Site
	Va, Vb      *Vertex
}

type Site struct {
}

type Vertex struct {
	x int
	y int
}

func NewCell(site Site) *Cell {
	return &Cell{Site: site}
}

func (c *Cell) prepareHalfEdges() {

}
