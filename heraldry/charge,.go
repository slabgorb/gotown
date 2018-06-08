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

type charge struct {
	key    string
	stroke color.Color
	fill   color.Color
}

func (c *charge) mask() (*image.Alpha, error) {
	r, err := resource.HeraldryBundle.Open(fmt.Sprintf("masks/mask-%s.png", c.key))
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

func (c *charge) decal() (*image.Alpha, error) {
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
		if strings.Contains(f, "mask-") {
			continue
		}
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
