package heraldry

import (
	"image/color"

	"github.com/fogleman/gg"
)

type div func(...color.Color) fill
type fill func(dc *gg.Context)

var divisions = map[string]div{
	"per fess":          perFess,
	"per cross":         perCross,
	"per chevron":       perChevron,
	"per bend":          perBend,
	"per saltire":       perSaltire,
	"per bend sinister": perBendSinister,
	"per pale":          perPale,
	"":                  solid,
}

// perFess divides the field horizontally
func perFess(c ...color.Color) fill {
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

func perCross(c ...color.Color) fill {
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

// perChevron divides the field 'after the manner of a chevron'
func perChevron(c ...color.Color) fill {
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

// perBend divides the field diagonally from upper left to lower right
func perBend(c ...color.Color) fill {
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

func perSaltire(c ...color.Color) fill {
	fill := perCross(c...)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(0.785398, w/2, h/2)
		dc.ScaleAbout(2.0, 2.0, w/2, h/2)
		fill(dc)
	}
}

// perBendSinister divides the field diagonally from upper right to lower left
func perBendSinister(c ...color.Color) fill {
	fill := perBend(c...)
	return func(dc *gg.Context) {
		h := float64(dc.Height())
		w := float64(dc.Width())
		dc.RotateAbout(3.14159, w/2, h/2)
		fill(dc)
	}
}

func perPale(c ...color.Color) fill {
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

func solid(c ...color.Color) fill {
	c = fillOutColors(c, 1)
	return func(dc *gg.Context) {
		dc.SetColor(c[0])
		dc.Fill()
	}
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
