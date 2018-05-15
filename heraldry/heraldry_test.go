package heraldry_test

import (
	"fmt"
	"github.com/fogleman/gg"
	. "github.com/slabgorb/gotown/heraldry"
	"image"
	"testing"
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

func TestEscutcheon(t *testing.T) {
	type testCase struct {
		draw   func(dc *gg.Context) *image.Alpha
		fill   func(dc *gg.Context)
		charge string
		name   string
	}
	dc := gg.NewContext(270, 270)
	testCases := []testCase{
		testCase{
			name: "heater_per_chevron",
			draw: HeaterShield,
			fill: PerChevron(Colors["sable"], Metals["or"]),
		},
		testCase{
			name:   "heater_per_pale",
			draw:   HeaterShield,
			fill:   PerPale(Colors["sable"], Metals["or"]),
			charge: "acorn.png",
		},
		testCase{
			name: "heater_per_fess",
			draw: HeaterShield,
			fill: PerFess(Colors["sable"], Metals["or"]),
		},
		testCase{
			name: "heater_per_bend",
			draw: HeaterShield,
			fill: PerBend(Colors["sable"], Metals["or"]),
		},
		testCase{
			name: "heater_per_bend_sinister",
			draw: HeaterShield,
			fill: PerBendSinister(Colors["sable"], Metals["or"]),
		},
	}
	for _, tc := range testCases {
		e := Escutcheon{
			Shape: tc.draw,
			Fill:  tc.fill,
		}
		if tc.charge != "" {
			e.Charge = &Charge{
				Color: Metals["argent"],
				Key:   tc.charge,
			}
		}
		e.Render(dc)
		dc.SavePNG(fmt.Sprintf("%s.png", tc.name))
		if !compareImages(tc.name) {
			t.Errorf("images are not equal for %s", tc.name)
		}
	}

}
