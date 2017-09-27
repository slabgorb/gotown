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

// TightHatch draws a tight pattern of lines
func TightHatch(dc *gg.Context, s *shape) {
	Hatch(dc, s, 2)
}

// LooseHatch draws a looser pattern of lines
func LooseHatch(dc *gg.Context, s *shape) {
	Hatch(dc, s, 4)
}

// Hatch draws a one-directional series of lines.
func Hatch(dc *gg.Context, s *shape, spacing float64) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		dc.SetLineWidth(0.5)
		for i := 0.0; i < s.width; i += spacing {
			dc.DrawLine(s.x+i, s.y, s.x+i, s.y+s.height)
		}
		dc.Stroke()
		dc.SetLineWidth(0.5)
	})

}

// Thatch draws a thatching pattern, suitable for huts
func Thatch(dc *gg.Context, s *shape) {
	var spacing = 2.5
	dc.SetLineWidth(0.25)
	for y := 0.0; y < s.height; y += spacing {
		dc.DrawLine(s.x, s.y+y, s.width+s.x, s.y+y+(0.5-random.Float64()))
	}
}

// CrossHatch draws a crosshatch pattern
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

// Cobblestone draws a cobblestone pattern, suitable for your fancier roads
func Cobblestone(dc *gg.Context, s *shape) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		const density = 7.0
		const radius = 1.0
		dc.SetLineWidth(0.25)
		for column := 0.0; column < s.width+density; column += density {
			for row := 0.0; row < s.height+density/2; row += (density / 2) {
				r := radius
				if random.Float64() > 0.75 {
					r += 0.5
				}
				y := row + s.y
				x := column + s.x + float64(((int(y)&1)-1))*(density/2)
				dc.DrawRoundedRectangle(x, y, density, density/2, r)
			}
		}
		dc.Stroke()
		dc.SetLineWidth(1.0)
	})
}

// Stipple draws a slightly-wavy pattern of dots.
func Stipple(dc *gg.Context, s *shape) {
	clipPattern(dc, s, func(dc *gg.Context, s *shape) {
		const density = 7
		const radius = 1.0

		for j := 0.0; j < s.height; j += density {
			for i := 0.0; i < s.width; i += density {
				r := radius
				if random.Float64() > 0.75 {
					r += 0.25
				}
				x := i + s.x
				d := i * 2 * math.Pi / s.height * 8
				y := s.y + j + (math.Sin(d)+1)/2*density
				dc.DrawCircle(x, y, r)
			}
		}
		dc.Fill()
	})

}
