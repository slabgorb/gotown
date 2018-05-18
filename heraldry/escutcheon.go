package heraldry

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/fogleman/gg"
	"github.com/slabgorb/gotown/resource"
)

// Escutcheon represents the 'shield' part of a heraldric device
type Escutcheon struct {
	DivisionKey string      `json:"division"`
	FieldColors []string    `json:"field_colors"`
	ShapeKey    string      `json:"shape"`
	ChargeKey   string      `json:"charge"`
	ChargeColor string      `json:"charge_color"`
	DC          *gg.Context `json:"-"`
}

const (
	width  = 270
	height = 270
)

func RandomEscutcheon() Escutcheon {
	e := Escutcheon{
		DC:          gg.NewContext(width, height),
		DivisionKey: RandomDivisionKey(),
		ShapeKey:    RandomShapeKey(),
		ChargeKey:   RandomChargeKey(),
		ChargeColor: RandomMetalKey(),
	}
	return e
}

// Render draws the Escutcheon
func (e Escutcheon) Render() image.Image {
	c := []color.Color{}
	for _, clr := range e.FieldColors {
		c = append(c, Tinctures[clr])
	}
	mask := Shapes[e.ShapeKey](e.DC)
	e.DC.SetMask(mask)
	divisions[e.DivisionKey](c...)(e.DC)
	ch := e.charge()
	if ch != nil {
		mask, err := ch.mask()
		if err != nil {
			panic(err)
		}
		e.DC.SetMask(mask)
		e.DC.DrawRectangle(0, 0, float64(e.DC.Width()), float64(e.DC.Height()))
		e.DC.SetColor(ch.c)
		e.DC.Fill()
	}
	return e.DC.Image()
}

func (e Escutcheon) charge() *charge {
	if e.ChargeKey == "" {
		return nil
	}
	return &charge{
		c:   Tinctures[e.ChargeColor],
		key: e.ChargeKey,
	}
}

type charge struct {
	key string
	c   color.Color
}

func (c *charge) mask() (*image.Alpha, error) {
	r, err := resource.HeraldryBundle.Open(fmt.Sprintf("%s.png", c.key))
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

// ListCharges returns a list of potential names of charges.
func ListCharges() []string {
	ret := []string{}
	for _, f := range resource.HeraldryBundle.Files() {
		ret = append(ret, strings.TrimSuffix(f, ".png"))
	}
	return ret
}

func RandomChargeKey() string {
	keys := []string{}
	for _, k := range ListCharges() {
		keys = append(keys, k)
	}
	return keys[randomizer.Intn(len(keys))]
}
