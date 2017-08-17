package draw

import "github.com/fogleman/gg"

func HouseWithGarden(dc *gg.Context, s *shape) {

	amount := s.height / 4
	roofShape := s.shrink(amount).offset(amount)

	Roof(dc, roofShape)

	dc.DrawRectangle(roofShape.x, roofShape.y, roofShape.width, roofShape.height)
	dc.Clip()
	Garden(dc, s)
}

func Roof(dc *gg.Context, s *shape) {
	withRotation(dc, s, func(dc *gg.Context, s *shape) {
		dc.SetRGB(0, 0, 0)
		dc.DrawRectangle(s.x, s.y, s.width, s.height)
		dc.Stroke()
		dc.SetRGB(0, 0, 0)
		Hatch(dc, s)
		dc.SetLineWidth(2.0)
		dc.DrawLine(s.midWidth(), s.y, s.midWidth(), s.y+s.height)
		dc.Stroke()
	})
}

func Garden(dc *gg.Context, s *shape) {
	withRotation(dc, s, func(dc *gg.Context, s *shape) {

		dc.SetRGB(0, 0, 0)

		//dc.Stroke()
		Stipple(dc, s)

	})
}
