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
	DivisionKey string   `json:"division"`
	FieldColors []string `json:"field_colors"`
	ShapeKey    string   `json:"shape"`
	ChargeKey   string   `json:"charge"`
	ChargeColor string   `json:"charge_color"`
}

// Render draws the Escutcheon
func (e Escutcheon) Render(dc *gg.Context) {
	c := []color.Color{}
	for _, clr := range e.FieldColors {
		c = append(c, Tinctures[clr])
	}
	mask := Shapes[e.ShapeKey](dc)
	dc.SetMask(mask)
	divisions[e.DivisionKey](c...)(dc)
	ch := e.charge()
	if ch != nil {
		mask, err := ch.mask()
		if err != nil {
			panic(err)
		}
		dc.SetMask(mask)
		dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
		dc.SetColor(ch.c)
		dc.Fill()
	}
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
