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
	Shape
	Fill
	*Charge
}

// Render draws the Escutcheon
func (e Escutcheon) Render(dc *gg.Context) {
	mask := e.Shape(dc)
	dc.SetMask(mask)
	e.Fill(dc)
	if e.Charge != nil {
		mask, err := e.Charge.mask()
		if err != nil {
			panic(err)
		}
		dc.SetMask(mask)
		dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
		dc.SetColor(e.Charge.Color)
		dc.Fill()
	}
}

// PerFess divides the field horizontally
func PerFess(top color.Color, bottom color.Color) Fill {
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		dc.DrawRectangle(0, 0, w, divide)
		dc.SetColor(top)
		dc.Fill()
		dc.DrawRectangle(0, divide, w, divide)
		dc.SetColor(bottom)
		dc.Fill()
	}
}

func PerCross(ul color.Color, ur color.Color, ll color.Color, lr color.Color) Fill {
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		divideW := w / 2
		dc.DrawRectangle(0, 0, divideW, divide)
		dc.SetColor(ul)
		dc.Fill()
		dc.DrawRectangle(divideW, 0, divideW, divide)
		dc.SetColor(ur)
		dc.Fill()
		dc.DrawRectangle(0, divide, divideW, divide)
		dc.SetColor(ll)
		dc.Fill()
		dc.DrawRectangle(divideW, divide, divideW, divide)
		dc.SetColor(lr)
		dc.Fill()
	}
}

// PerChevron divides the field 'after the manner of a chevron'
func PerChevron(top color.Color, bottom color.Color) Fill {
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := h / 2
		divideW := w / 2
		dc.SetColor(top)
		dc.DrawRectangle(0, 0, w, h)
		dc.Fill()
		dc.MoveTo(0, h)
		dc.LineTo(divideW, divide)
		dc.LineTo(w, h)
		dc.LineTo(0, h)
		dc.SetColor(bottom)
		dc.Fill()
	}
}

// PerBend divides the field diagonally from upper left to lower right
func PerBend(left color.Color, right color.Color) Fill {
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.MoveTo(0, 0)
		dc.LineTo(0, h)
		dc.LineTo(w, h)
		dc.LineTo(0, 0)
		dc.SetColor(left)
		dc.Fill()
		dc.MoveTo(w, 0)
		dc.LineTo(w, h)
		dc.LineTo(0, 0)
		dc.LineTo(w, 0)
		dc.SetColor(right)
		dc.Fill()
	}
}

func PerSaltire(ul color.Color, ur color.Color, ll color.Color, lr color.Color) Fill {
	fill := PerCross(ul, ur, ll, lr)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(0.785398, w/2, h/2)
		dc.ScaleAbout(2.0, 2.0, w/2, h/2)
		fill(dc)
	}
}

// PerBendSinister divides the field diagonally from upper right to lower left
func PerBendSinister(left color.Color, right color.Color) Fill {
	fill := PerBend(left, right)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(3.14159, w/2, h/2)
		fill(dc)
	}
}

func PerPale(left color.Color, right color.Color) Fill {
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		divide := w / 2
		dc.DrawRectangle(0, 0, divide, h)
		dc.SetColor(left)
		dc.Fill()
		dc.DrawRectangle(divide, 0, divide, h)
		dc.SetColor(right)
		dc.Fill()
	}
}

func Solid(fillColor color.Color) Fill {
	return func(dc *gg.Context) {
		dc.SetColor(fillColor)
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
