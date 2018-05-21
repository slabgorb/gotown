package heraldry

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the random generator for the package. Generally used by
// tests.
func SetRandomizer(g random.Generator) {
	randomizer = g
}

type Heraldry struct {
	*Escutcheon `json:"escutcheon`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
}

func New(e Escutcheon) *Heraldry {
	return &Heraldry{
		Escutcheon: &e,
		Icon:       renderImageToBase64(e, 0.25),
		Image:      renderImageToBase64(e, 1.0),
	}
}

func renderImageToBase64(e Escutcheon, pct float64) string {
	buf := bytes.NewBuffer([]byte{})
	png.Encode(buf, e.RenderAtPercent(pct))
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	return fmt.Sprintf("data:image/png;base64,%s", b64)
}
