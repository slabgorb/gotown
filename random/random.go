package random

type Generator interface {
	Int() int
	Intn(int) int
	Float64() float64
}

type MockRandom struct{}

func NewMock() MockRandom {
	return MockRandom{}
}

func (m MockRandom) Int() int {
	return 4 // https://xkcd.com/221/
}

func (m MockRandom) Intn(max int) int {
	return max / 2
}

func (m MockRandom) Float64() float64 {
	return 0.40000
}
