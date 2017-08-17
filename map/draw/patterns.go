package draw

import (
	"math"

	"github.com/fogleman/gg"
)

func clipPattern(dc *gg.Context, s *shape, f shaper) {
	dc.DrawRectangle(s.x, s.y, s.width, s.height)
	dc.Clip()
	f(dc, s)
	dc.ResetClip()

}

func Hatch(dc *gg.Context, s *shape) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		dc.SetLineWidth(1.0)
		for i := 0.0; i < s.width; i += 2 {
			dc.DrawLine(s.x+i, s.y, s.x+i, s.y+s.height)
			dc.Stroke()
		}
	})

}

func Stipple(dc *gg.Context, s *shape) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
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
				dc.DrawCircle(x, y, radius)

			}
		}

		dc.Fill()
	})

}
