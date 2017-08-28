package genetics_test

import (
	"testing"

	"github.com/slabgorb/gotown/inhabitants/genetics"
)

func getVariants() []*genetics.Variant {
	vs := [][]string{
		[]string{"brown", "f(9|a|b|c|d)"},
		[]string{"hazel", "f(6|7|8)"},
		[]string{"blue", "f1|f2"},
		[]string{"gray", "f3"},
		[]string{"green", "f4|f5"},
	}

	variants := make([]*genetics.Variant, len(vs))
	for i, v := range vs {
		va, err := genetics.NewVariant(v[0], v[1])
		if err != nil {
			panic(err)
		}
		variants[i] = va
	}
	return variants
}

func TestTrait(t *testing.T) {
	trait := genetics.NewTrait("eye color", getVariants())
	exp := genetics.Expression{}
	exp.Add(trait)
	testCases := []struct {
		s        string
		expected string
	}{
		{
			"f1f4f4f9",
			"green",
		},
		{
			"f1f9f3fa",
			"brown",
		},
		{
			"f1f1f1",
			"blue",
		},
		{
			"f1f1f3",
			"blue",
		},
	}

	for _, tc := range testCases {
		e, _ := trait.Expression(tc.s)
		if e != tc.expected {
			t.Errorf("Expected %s got %s ", tc.expected, e)
		}
		c := genetics.Chromosome{}
		c.Add(tc.s)
		result := c.Express(exp)
		if result["eye color"] != tc.expected {
			t.Errorf("Expected %s got %s", tc.expected, result["eye_color"])
		}
	}

}
