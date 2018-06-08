package heraldry

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
)

// Escutcheon represents the 'shield' part of a heraldric device
type Escutcheon struct {
	DivisionKey string         `json:"division"`
	FieldColors []string       `json:"field_colors"`
	ShapeKey    string         `json:"shape"`
	ChargeKey   string         `json:"charge"`
	ChargeColor string         `json:"charge_color"`
	Quarters    [4]*Escutcheon `json:"quarters,omitempty"`
	DC          *gg.Context    `json:"-"`
}

const (
	width  = 270
	height = 270
)

func RandomEscutcheon(shape string, allowQuartering bool) Escutcheon {
	dKey := RandomDivisionKey()
	cColor, cField := randomColorStrategy()()
	var qs quarterStrategy = func() [4]*Escutcheon {
		return [4]*Escutcheon{}
	}
	if allowQuartering {
		qs = randomQuarterStrategy()
	}
	ck := RandomChargeKey()
	quarters := qs()
	any := false
	for _, q := range quarters {
		if !(q == nil) {
			any = true
		}
	}
	if any {
		ck = ""
	}
	e := Escutcheon{
		DC:          gg.NewContext(width, height),
		DivisionKey: dKey,
		ShapeKey:    shape,
		ChargeKey:   ck,
		ChargeColor: cColor,
		FieldColors: cField,
		Quarters:    quarters,
	}
	return e
}

func (e Escutcheon) RenderAtPercent(size float64) image.Image {
	img := e.image()
	if size == 1.0 {
		return img
	}
	img = resize.Resize(uint(width*size), 0, img, resize.NearestNeighbor)
	return img
}

func (e *Escutcheon) image() image.Image {
	if e.DC == nil {
		e.DC = gg.NewContext(width, height)
	}
	c := []color.Color{}
	for _, clr := range e.FieldColors {
		c = append(c, Tinctures[clr])
	}
	mask := Shapes[e.ShapeKey](e.DC)
	e.DC.SetMask(mask)
	divisions[e.DivisionKey](c...)(e.DC)
	e.DC.Identity()
	e.DC.ResetClip()
	e.DC.ClearPath()
	ch := e.charge()
	if ch != nil {
		mask, err := ch.mask()
		if err != nil {
			panic(err)
		}
		e.DC.SetMask(mask)
		e.DC.DrawRectangle(0, 0, float64(e.DC.Width()), float64(e.DC.Height()))
		e.DC.SetColor(ch.fill)
		e.DC.Fill()
		decal, err := ch.decal()
		if err != nil {
			panic(err)
		}
		e.DC.SetMask(decal)
		e.DC.DrawRectangle(0, 0, float64(e.DC.Width()), float64(e.DC.Height()))
		e.DC.SetColor(ch.stroke)
		e.DC.Fill()
	}
	uniform := &image.Uniform{color.RGBA{A: 255}}
	dc := gg.NewContext(width, height)
	dc.DrawImage(uniform, 0, 0)
	e.DC.SetMask(dc.AsMask())
	e.drawQuarters()
	return e.DC.Image()

}

func (e Escutcheon) drawQuarter(i int, x int, y int) {
	q := e.Quarters[i]
	if !(q == nil) {
		img := (*q).RenderAtPercent(0.5)
		e.DC.DrawImage(img, x, y)
	}
}

func (e Escutcheon) drawQuarters() {
	centerW := e.DC.Width() / 2
	centerY := e.DC.Height() / 2

	// top left
	e.drawQuarter(0, 0, 0)
	// top right
	e.drawQuarter(1, centerW, 0)
	// lower right
	e.drawQuarter(2, centerW, centerY)
	// lower left
	e.drawQuarter(3, 0, centerY)
}

// Render draws the Escutcheon
func (e Escutcheon) Render() image.Image {
	return e.RenderAtPercent(1.0)
}

func (e Escutcheon) charge() *charge {
	if e.ChargeKey == "" {
		return nil
	}
	return &charge{
		fill:   Tinctures[e.ChargeColor],
		stroke: Tinctures["sable"],
		key:    e.ChargeKey,
	}
}

type quarterStrategy func() [4]*Escutcheon

func randomQuarterStrategy() quarterStrategy {
	randomQuarterStrategies := map[string]quarterStrategy{
		"none1": func() [4]*Escutcheon {
			return [4]*Escutcheon{}
		},
		"none2": func() [4]*Escutcheon {
			return [4]*Escutcheon{}
		},
		"none3": func() [4]*Escutcheon {
			return [4]*Escutcheon{}
		},
		"none4": func() [4]*Escutcheon {
			return [4]*Escutcheon{}
		},
		"charge dexter": func() [4]*Escutcheon {
			e := RandomEscutcheon("square", false)
			return [4]*Escutcheon{&e, nil, nil, nil}

		},
		"charge sinister": func() [4]*Escutcheon {
			e := RandomEscutcheon("square", false)
			return [4]*Escutcheon{nil, &e, nil, nil}
		},
		"alternate quarters": func() [4]*Escutcheon {
			a := RandomEscutcheon("square", false)
			b := RandomEscutcheon("square", false)
			return [4]*Escutcheon{&a, &b, &a, &b}
		},
	}
	keys := []string{}
	for k := range randomQuarterStrategies {
		keys = append(keys, k)
	}
	return randomQuarterStrategies[keys[randomizer.Intn(len(keys))]]
}

type colorStrategy func() (string, []string)

func randomColorStrategy() colorStrategy {

	randomColorStrategies := map[string]colorStrategy{
		"argent on": func() (string, []string) {
			a := RandomColorKey()
			b := RandomColorKey()
			return argent, []string{a, b, a, b, a, b}
		},
		"sable on": func() (string, []string) {
			a := RandomSableOn()
			b := RandomSableOn()
			return sable, []string{a, b, a, b, a, b}
		},
		"sable on stain": func() (string, []string) {
			a := RandomSableOn()
			b := RandomStainKey()
			return sable, []string{a, b, a, b, a, b}
		},
		"argent on stain": func() (string, []string) {
			a := RandomStainKey()
			b := RandomStainKey()
			return argent, []string{a, b, a, b, a, b}
		},
		"alternate_1": func() (string, []string) {
			a := RandomColorKey()
			b := RandomColorKey()
			c := RandomMetalKey()
			return c, []string{a, b, a, b, a, b}
		},
		"alternate_2": func() (string, []string) {
			a := argent
			b := or
			c := RandomColorKey()
			return c, []string{a, b, a, b, a, b}
		},
		"alternate_3": func() (string, []string) {
			b := argent
			a := or
			c := RandomColorKey()
			return c, []string{a, b, a, b, a, b}
		},
	}
	keys := []string{}
	for k := range randomColorStrategies {
		keys = append(keys, k)
	}
	return randomColorStrategies[keys[randomizer.Intn(len(keys))]]
}
