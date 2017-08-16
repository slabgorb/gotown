package voronoi

import "math"

const (
	ε    = math.E
	invε = 1 / 1e-9
)

// Diagram is the voronoi diagram
type Diagram struct {
	Height float64
	Width  float64
}

func equalε(a, b float64) bool {
	return math.Abs(a-b) < ε
}

func gtε(a, b float64) bool {
	return a-b > ε
}

func ltε(a, b float64) bool {
	return b-a > ε
}
