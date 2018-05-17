package heraldry

import "image/color"

type Tincture map[string]color.RGBA

// Metals is a heraldric color set
var Metals = Tincture{
	"argent": color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	"or":     color.RGBA{R: 0xff, G: 0xd4, B: 0x00, A: 0xff},
}

var Stains = Tincture{
	"murrey":   color.RGBA{R: 0x8c, G: 0x00, B: 0x4b, A: 0xff},
	"sanguine": color.RGBA{R: 0x66, G: 0x00, B: 0x00, A: 0xff},
	"tenne":    color.RGBA{},
}

// Colors is a heraldric color set
var Colors = Tincture{
	"sable":   color.RGBA{R: 0x00, G: 0x00, B: 0x0, A: 0xff},
	"gules":   color.RGBA{R: 0xff, G: 0x00, B: 0x0, A: 0xff},
	"azure":   color.RGBA{R: 0x00, G: 0x00, B: 0xf, A: 0xff},
	"vert":    color.RGBA{R: 0x00, G: 0x66, B: 0x0, A: 0xff},
	"purpure": color.RGBA{R: 0x80, G: 0x00, B: 0x8, A: 0xff},
	"cendree": color.RGBA{R: 0x80, G: 0x80, B: 0x8, A: 0xff},
}
