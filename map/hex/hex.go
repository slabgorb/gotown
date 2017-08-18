package hex

import "math"

var directions = [][]float64{{0, -1}, {1, -1}, {1, 0}, {-1, 1}, {-1, 0}}

type Cube struct {
	x, y, z float64
}

func (c Cube) distance(c2 Cube) float64 {
	return math.Max(math.Max(math.Abs(c.x-c2.x), math.Abs(c.y-c2.y)), math.Abs(c.z-c2.z))
}

func (c Cube) ToAxial() Axial {
	return Axial{q: c.x, r: c.z}
}

type Axial struct {
	q, r float64
}

func (a Axial) ToCube() Cube {
	return Cube{a.q, -a.q - a.r, a.r}
}

func (a Axial) distance(a2 Axial) float64 {
	return a.ToCube().distance(a2.ToCube())
}

func (a Axial) Coordinates() (float64, float64) {
	return a.q, a.r
}

func (a Axial) Surrounding() []Axial {
	axials := []Axial{}
	for _, d := range directions {
		axials = append(axials, Axial{q: a.q + d[0], r: a.r + d[1]})
	}
	return axials
}

type coord struct {
	q, r float64
}

type AxialGrid struct {
	hexRay float64
	Hexes  map[coord]*Axial
}

func NewAxialGrid(hr float64) AxialGrid {
	ag := AxialGrid{hexRay: hr}
	ag.Hexes = make(map[coord]*Axial)
	return ag
}

func (ag AxialGrid) SetHex(q, r float64) {
	ag.Hexes[coord{q, r}] = &Axial{q, r}
}

func (ag AxialGrid) halfHeight() float64 {
	return ag.hexRay
}

func (ag AxialGrid) hexHeight() float64 {
	return ag.hexRay * 2.0
}

func (ag AxialGrid) hexWidth() float64 {
	return math.Sqrt(3) / 2.0 * ag.hexHeight()
}

func (ag AxialGrid) halfWidth() float64 {
	return ag.hexWidth() / 2.0
}

func (ag AxialGrid) quarterHeight() float64 {
	return ag.hexHeight() / 4
}

func (ag AxialGrid) GetHex(q, r float64) (*Axial, bool) {
	h, ok := ag.Hexes[coord{q, r}]
	return h, ok
}

func (ag AxialGrid) toXY(a Axial) (float64, float64) {
	x := (ag.hexRay * math.Sqrt(3) * (a.q + a.r/2.0))
	y := (ag.hexRay * 3 / 2 * a.r)
	x += ag.halfWidth()
	y += ag.halfHeight()
	return x, y
}

func (ag AxialGrid) HexPolygon(a Axial) [][]float64 {
	x, y := ag.toXY(a)
	return [][]float64{
		{x - ag.halfWidth(), y + ag.quarterHeight()},
		{x, y + ag.halfHeight()},
		{x + ag.halfWidth(), y + ag.quarterHeight()},
		{x + ag.halfWidth(), y - ag.quarterHeight()},
		{x, y - ag.halfHeight()},
		{x - ag.halfWidth(), y - ag.quarterHeight()},
	}
}

type SquareGrid struct {
	AxialGrid
}

func NewSquareGrid(hr float64) SquareGrid {
	return SquareGrid{NewAxialGrid(hr)}
}

func (sq SquareGrid) evenQToAxialHex(col, row int) Axial {
	x := float64(col - (row-(row&1))/2)
	z := float64(row)
	y := float64(-x - z)
	c := Cube{x: x, y: y, z: z}
	return c.ToAxial()
}

func (sq SquareGrid) SetHex(q, r float64) {
	a := sq.evenQToAxialHex(int(q), int(r))
	sq.Hexes[coord{q, r}] = &a
}
