package heraldry_test

import (
	"testing"

	"github.com/fogleman/gg"
	. "github.com/slabgorb/gotown/heraldry"
)

func TestApply(t *testing.T) {
	dc := gg.NewContext(300, 300)
	Apply(dc)
}
