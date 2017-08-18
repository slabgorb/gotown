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

func TightHatch(dc *gg.Context, s *shape) {
	Hatch(dc, s, 2)
}

func LooseHatch(dc *gg.Context, s *shape) {
	Hatch(dc, s, 4)
}

func Hatch(dc *gg.Context, s *shape, spacing float64) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		dc.SetLineWidth(1.0)
		for i := 0.0; i < s.width; i += spacing {
			dc.DrawLine(s.x+i, s.y, s.x+i, s.y+s.height)
			dc.Stroke()
		}
	})

}

func CrossHatch(dc *gg.Context, s *shape) {
	LooseHatch(dc, s)
	rotated := s.grow(s.width * 1.5)
	rotated.angle = s.angle + 90
	if rotated.angle > 360 {
		rotated.angle = math.Mod(rotated.angle, 360)
	}
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		withRotation(dc, rotated, LooseHatch)
	})

}

func Stipple(dc *gg.Context, s *shape) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		const density = 7
		const radius = 1
		for j := 0.0; j < s.height; j += density {
			for i := 0.0; i < s.width; i += density {
				x := i + s.x
				d := i * 2 * math.Pi / s.height * 8
				y := s.y + j + (math.Sin(d)+1)/2*density
				dc.DrawCircle(x, y, radius)
			}
		}
		dc.Fill()
	})

}
