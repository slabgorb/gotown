package draw

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func HouseWithGarden(dc *gg.Context, s *shape) {
	amount := s.height / 4
	roofShape := s.shrink(amount).offset(amount)
	ShingledRoof(dc, roofShape)
	Garden(dc, s)
}

func ShingledRoof(dc *gg.Context, s *shape) {
	roof(dc, s, TightHatch)
}

func ThatchedRoof(dc *gg.Context, s *shape) {
	roof(dc, s, Thatch)
}

func roof(dc *gg.Context, s *shape, pattern shaper) {
	withRotation(dc, s, func(dc *gg.Context, s *shape) {
		dc.DrawRectangle(s.x, s.y, s.width, s.height)
		dc.Stroke()
		pattern(dc, s)
		dc.SetLineWidth(2.0)
		dc.DrawLine(s.midWidth(), s.y, s.midWidth(), s.y+s.height)
		dc.Stroke()
	})
}

func Road(dc *gg.Context, s *shape) {
	withRotation(dc, s, Cobblestone)
}

func Farm(dc *gg.Context, s *shape) {
	withRotation(dc, s, CrossHatch)
}

func Garden(dc *gg.Context, s *shape) {
	dc.SetRGB(0, 0, 0)
	withRotation(dc, s, Stipple)
}

// Trees draws some happy little trees
func Trees(dc *gg.Context, s *shape) {
	// basically, we want an irregular 'leafy' outline, which hopefully isn't too squared off.
	// pick some points around a rounded rectangle and stroke some quadratics?

}
