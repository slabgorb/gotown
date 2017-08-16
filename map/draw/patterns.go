package draw

import (
	"log"
	"math"

	"github.com/fogleman/gg"
)

func Hatch(dc *gg.Context, s *shape) {
	dc.SetLineWidth(1.0)
	for i := 0.0; i < s.width; i += 2 {
		dc.DrawLine(s.x+i, s.y, s.x+i, s.y+s.height)
		dc.Stroke()
	}
}

func Stipple(dc *gg.Context, s *shape) {
	const density = 7
	const radius = 2
	for j := 0.0; j < s.height; j += density {
		for i := 0.0; i < s.width; i += density {
			x := i + s.x
			d := float64(int(i) % (density + 1))
			f := 0.0
			if d != 0.0 {
				f = math.Pi * 1 / d
			}
			y := j + (math.Sin(f) * density) + s.y
			log.Printf("x %f y %f ", x, y)
			if s.inBounds(x, y) {
				dc.DrawCircle(x, y, radius)
			}

		}
	}

	dc.Fill()

}
