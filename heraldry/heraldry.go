package heraldry

import (
	"github.com/fogleman/gg"
	"image/color"
)

// Device models a heraldric device
type Device struct {
}

type Tincture map[string]color.RGBA

// Metals is a heraldric color set
var Metals = Tincture{
	"argent": color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	"or":     color.RGBA{R: 0xff, G: 0xd4, B: 0x00, A: 0xff},
}

// Colors is a heraldric color set
var Colors = Tincture{
	"sable":    color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	"gules":    color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
	"sanguine": color.RGBA{R: 0x66, G: 0x00, B: 0x00, A: 0xff},
	"azure":    color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
	"vert":     color.RGBA{R: 0x00, G: 0x66, B: 0x00, A: 0xff},
	"purpure":  color.RGBA{R: 0x80, G: 0x00, B: 0x80, A: 0xff},
	"murrey":   color.RGBA{R: 0x8c, G: 0x00, B: 0x4b, A: 0xff},
	"cendree":  color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff},
}

func Apply(dc *gg.Context) error {
	dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
	dc.SetColor(Colors["sable"])
	dc.Fill()
	heaterShield(dc)
	dc.SavePNG("test.png")
	return nil
}

func heaterShield(dc *gg.Context) {
	dc.SetColor(Metals["or"])
	w := float64(dc.Width())
	h := float64(dc.Height())
	margin := float64(w / 10)
	dc.MoveTo(margin, margin)
	dc.LineTo(w-margin, margin)
	dc.QuadraticTo(w/2, h-margin, margin, h-margin)
	dc.MoveTo(margin, margin)
	dc.QuadraticTo(w/2, h-margin, w-margin, h-margin)
	dc.ClosePath()
	dc.Fill()
}
