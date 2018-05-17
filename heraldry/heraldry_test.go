package heraldry_test

import (
	"fmt"
	"image"
	"image/color"
	"os"
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

func TestEscutcheon(t *testing.T) {
	type testCase struct {
		draw     func(dc *gg.Context) *image.Alpha
		division string
		colors   []color.Color
		charge   string
		name     string
	}
	testCases := []testCase{
		testCase{
			name:     "heater_per_chevron",
			draw:     HeaterShield,
			division: "per chevron",
			colors:   []color.Color{Colors["sable"], Metals["or"]},
		},
		testCase{
			name:     "heater_per_pale",
			draw:     HeaterShield,
			division: "per pale",
			colors:   []color.Color{Colors["sable"], Metals["or"]},
		},
		testCase{
			name:     "heater_per_fess",
			draw:     HeaterShield,
			division: "per fess",
			colors:   []color.Color{Colors["sable"], Metals["or"]},
		},
		testCase{
			name:     "heater_per_bend",
			draw:     HeaterShield,
			division: "per bend",
			colors:   []color.Color{Colors["sable"], Metals["or"]},
			charge:   "bears-head-couped",
		},
		testCase{
			name:     "heater_per_bend_sinister",
			draw:     HeaterShield,
			division: "per bend sinister",
			colors:   []color.Color{Colors["sable"], Metals["or"]},
			charge:   "acorn",
		},
		testCase{
			name:     "heater_per_cross",
			draw:     HeaterShield,
			division: "per cross",
			colors:   []color.Color{Colors["sable"], Metals["or"], Stains["murrey"], Metals["argent"]},
		},
		testCase{
			name:     "heater_per_saltire",
			draw:     HeaterShield,
			division: "per saltire",
			colors:   []color.Color{Colors["sable"], Metals["or"], Stains["murrey"], Metals["argent"]},
		},
	}
	for _, tc := range testCases {
		dc := gg.NewContext(270, 270)
		e := Escutcheon{
			Shape: tc.draw,
			Fill:  Divisions[tc.division](tc.colors...),
		}
		if tc.charge != "" {
			e.Charge = &Charge{
				Color: Metals["argent"],
				Key:   tc.charge,
			}
		}
		e.Render(dc)
		fname := fmt.Sprintf("%s.png", tc.name)
		dc.SavePNG(fname)
		if !compareImages(tc.name) {
			t.Errorf("images are not equal for %s", tc.name)
		}
		os.Remove(fname)
	}

}
