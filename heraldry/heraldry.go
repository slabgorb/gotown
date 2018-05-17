package heraldry

import (
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/fogleman/gg"
	"github.com/slabgorb/gotown/resource"
)

// Device models a heraldric device
type Device struct {
}

type Fill func(dc *gg.Context)
type Shape func(dc *gg.Context) *image.Alpha

type Charge struct {
	Key string
	color.Color
}

func (c *Charge) mask() (*image.Alpha, error) {
	r, err := resource.HeraldryBundle.Open(fmt.Sprintf("%s.png", c.Key))
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	dc := gg.NewContextForImage(img)
	return dc.AsMask(), nil
}

// Escutcheon represents the 'shield' part of a heraldric device
type Escutcheon struct {
	DivisionKey string   `json:"division"`
	FieldColors []string `json:"field_colors"`
	ShapeKey    string   `json:"shape"`
	ChargeKey   string   `json:"charge"`
	ChargeColor string   `json:"charge_color"`
}

type Div func(...color.Color) Fill

var Divisions = map[string]Div{
	"per fess":          PerFess,
	"per cross":         PerCross,
	"per chevron":       PerChevron,
	"per bend":          PerBend,
	"per saltire":       PerSaltire,
	"per bend sinister": PerBendSinister,
	"per pale":          PerPale,
	"":                  Solid,
}

var Shapes = map[string]Shape{
	"heater": HeaterShield,
}

func fillOutColors(c []color.Color, count int) []color.Color {
	ret := make([]color.Color, count)
	for i := 0; i < count; i++ {
		if len(c) >= i {
			ret[i] = c[i]
		} else {
			ret[i] = Tinctures["sable"]
		}
	}
	return ret
}

func (e Escutcheon) charge() *Charge {
	if e.ChargeKey == "" {
		return nil
	}
	return &Charge{
		Color: Tinctures[e.ChargeColor],
		Key:   e.ChargeKey,
	}
}

// Render draws the Escutcheon
func (e Escutcheon) Render(dc *gg.Context) {
	c := []color.Color{}
	for _, clr := range e.FieldColors {
		c = append(c, Tinctures[clr])
	}
	mask := Shapes[e.ShapeKey](dc)
	dc.SetMask(mask)
	Divisions[e.DivisionKey](c...)(dc)
	ch := e.charge()
	if ch != nil {
		mask, err := ch.mask()
		if err != nil {
			panic(err)
		}
		dc.SetMask(mask)
		dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
		dc.SetColor(ch.Color)
		dc.Fill()
	}
}

// PerFess divides the field horizontally
func PerFess(c ...color.Color) Fill {
	c = fillOutColors(c, 2)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		dc.DrawRectangle(0, 0, w, divide)
		dc.SetColor(c[0])
		dc.Fill()
		dc.DrawRectangle(0, divide, w, divide)
		dc.SetColor(c[1])
		dc.Fill()
	}
}

func PerCross(c ...color.Color) Fill {
	c = fillOutColors(c, 4)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		divideW := w / 2
		dc.DrawRectangle(0, 0, divideW, divide)
		dc.SetColor(c[0])
		dc.Fill()
		dc.DrawRectangle(divideW, 0, divideW, divide)
		dc.SetColor(c[1])
		dc.Fill()
		dc.DrawRectangle(0, divide, divideW, divide)
		dc.SetColor(c[2])
		dc.Fill()
		dc.DrawRectangle(divideW, divide, divideW, divide)
		dc.SetColor(c[3])
		dc.Fill()
	}
}

// PerChevron divides the field 'after the manner of a chevron'
func PerChevron(c ...color.Color) Fill {
	c = fillOutColors(c, 2)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		divideW := w / 2
		dc.SetColor(c[0])
		dc.DrawRectangle(0, 0, w, h)
		dc.Fill()
		dc.MoveTo(0, h)
		dc.LineTo(divideW, divide)
		dc.LineTo(w, h)
		dc.LineTo(0, h)
		dc.SetColor(c[1])
		dc.Fill()
	}
}

// PerBend divides the field diagonally from upper left to lower right
func PerBend(c ...color.Color) Fill {
	c = fillOutColors(c, 2)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.MoveTo(0, 0)
		dc.LineTo(0, h)
		dc.LineTo(w, h)
		dc.LineTo(0, 0)
		dc.SetColor(c[0])
		dc.Fill()
		dc.MoveTo(w, 0)
		dc.LineTo(w, h)
		dc.LineTo(0, 0)
		dc.LineTo(w, 0)
		dc.SetColor(c[1])
		dc.Fill()
	}
}

func PerSaltire(c ...color.Color) Fill {
	fill := PerCross(c...)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(0.785398, w/2, h/2)
		dc.ScaleAbout(2.0, 2.0, w/2, h/2)
		fill(dc)
	}
}

// PerBendSinister divides the field diagonally from upper right to lower left
func PerBendSinister(c ...color.Color) Fill {
	fill := PerBend(c...)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(3.14159, w/2, h/2)
		fill(dc)
	}
}

func PerPale(c ...color.Color) Fill {
	c = fillOutColors(c, 2)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := w / 2
		dc.DrawRectangle(0, 0, divide, h)
		dc.SetColor(c[1])
		dc.Fill()
		dc.DrawRectangle(divide, 0, divide, h)
		dc.SetColor(c[0])
		dc.Fill()
	}
}

func Solid(c ...color.Color) Fill {
	c = fillOutColors(c, 1)
	return func(dc *gg.Context) {
		dc.SetColor(c[0])
		dc.Fill()
	}
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
