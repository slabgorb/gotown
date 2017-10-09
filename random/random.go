package random

import (
	"math/rand"
	"sync"
	"time"
)

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

var Random *rand.Rand

func init() {
	Random = rand.New(
		&lockedRandSource{
			src: rand.NewSource(time.Now().UnixNano()),
		},
	)
}

// locked to prevent concurrent use of the underlying source
type lockedRandSource struct {
	lock sync.Mutex // protects src
	src  rand.Source
}

// to satisfy rand.Source interface
func (r *lockedRandSource) Int63() int64 {
	r.lock.Lock()
	ret := r.src.Int63()
	r.lock.Unlock()
	return ret
}

// to satisfy rand.Source interface
func (r *lockedRandSource) Seed(seed int64) {
	r.lock.Lock()
	r.src.Seed(seed)
	r.lock.Unlock()
}
