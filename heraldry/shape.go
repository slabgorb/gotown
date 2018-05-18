package heraldry

import (
	"image"

	"github.com/fogleman/gg"
)

type Shape func(dc *gg.Context) *image.Alpha

var Shapes = map[string]Shape{
	"heater": HeaterShield,
}

func HeaterShield(dc *gg.Context) *image.Alpha {
	w := float64(dc.Width())
	h := float64(dc.Height())
	margin := float64(w / 10)
	halfHeater(dc, w, h, margin)
	dc.ScaleAbout(-1, 1, w/2, h/2)
	halfHeater(dc, w, h, margin)
	dc.Fill()
	return dc.AsMask()
}

func halfHeater(dc *gg.Context, w, h, margin float64) {
	dc.MoveTo(w/2, margin)
	dc.LineTo(w-margin, margin)
	dc.QuadraticTo(w-margin, h-margin, w/2, h)
}
