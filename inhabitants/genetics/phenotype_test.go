package genetics_test

import (
	"testing"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

func init() {
	genetics.SetRandomizer(random.NewMock())

}

func TestExpression(t *testing.T) {
	exp := genetics.Expression{}

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

	trait := genetics.NewTrait("eye color", variants)

	exp.Add(trait)
	//runtime.Breakpoint()
	c := genetics.RandomChromosome(40)
	e := c.Express(exp)
	expected := "hazel"
	if e["eye color"] != expected {
		t.Errorf("Expected %s got %s", expected, e["eye color"])
	}
}
