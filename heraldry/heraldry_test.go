package heraldry_test

import (
	"fmt"
	"testing"

	"github.com/fogleman/gg"
	. "github.com/slabgorb/gotown/heraldry"
)

func compareImages(name string) bool {
	testResult, _ := gg.LoadImage(fmt.Sprintf("%s.png", name))
	reference, _ := gg.LoadImage(fmt.Sprintf("reference_images/%s.png", name))
	if testResult.Bounds() != reference.Bounds() {
		return false
	}
	bounds := reference.Bounds()
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			if testResult.At(i, j) != reference.At(i, j) {
				return false
			}
		}
	}
	return true
}

func TestRandom(t *testing.T) {
	e := RandomEscutcheon("square", true)
	e.Render()
	e.DC.SavePNG("random.png")
}

func TestEscutcheon(t *testing.T) {
	type testCase struct {
		draw     string
		division string
		charge   string
		name     string
	}
	testCases := []testCase{
		testCase{
			name:     "heater_per_chevron",
			draw:     "heater",
			division: "per chevron",
		},
		testCase{
			name:     "heater_per_pale",
			draw:     "heater",
			division: "per pale",
		},
		testCase{
			name:     "heater_per_fess",
			draw:     "heater",
			division: "per fess",
		},
		testCase{
			name:     "heater_per_bend",
			draw:     "heater",
			division: "per bend",
			charge:   "bears-head-couped",
		},
		testCase{
			name:     "heater_per_bend_sinister",
			draw:     "heater",
			division: "per bend sinister",
			charge:   "acorn",
		},
		testCase{
			name:     "heater_per_cross",
			draw:     "heater",
			division: "per cross",
		},
		testCase{
			name:     "heater_per_saltire",
			draw:     "heater",
			division: "per saltire",
		},
	}
	for _, tc := range testCases {
		dc := gg.NewContext(270, 270)
		e := Escutcheon{
			ShapeKey:    tc.draw,
			DivisionKey: tc.division,
			ChargeColor: "argent",
			FieldColors: []string{"tenne", "sable", "murrey", "purpure"},
			ChargeKey:   tc.charge,
			DC:          dc,
		}

		e.Render()
		fname := fmt.Sprintf("%s.png", tc.name)
		dc.SavePNG(fname)
		if !compareImages(tc.name) {
			t.Errorf("images are not equal for %s", tc.name)
		}
		//os.Remove(fname)
	}

}
