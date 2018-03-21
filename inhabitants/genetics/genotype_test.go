package genetics_test

import (
	"testing"

	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

func init() {
	genetics.SetRandomizer(random.NewMock())

}
func TestGene(t *testing.T) {
	g := genetics.RandomGene()
	if len(g) != 6 {
		t.Errorf("Expected a length of 6, got %d", len(g))
	}
	var expected = int64(8388607)
	if g.Int64() != expected {
		t.Errorf("Expected %d, got %d", expected, g.Int64())
	}

}

func TestRandomize(t *testing.T) {
	c := genetics.RandomChromosome(10)
	if c.Len() != 10 {
		t.Errorf("Expected a length of 10, got %d", c.Len())
	}
	expected := "7fffff"
	if string(c.Index(0)) != expected {
		t.Errorf("Got %s expected %s", c.Index(0), expected)
	}
}

func TestCombined(t *testing.T) {
	c1 := genetics.RandomChromosome(10)
	c2 := genetics.RandomChromosome(10)
	c3 := genetics.RandomChromosome(11)
	c, err := c2.Combine(c3)
	if err == nil && c != nil {
		t.Errorf("expected error when combining two chromosomes with differing counts")
	}
	c, err = c1.Combine(c2)
	if err != nil || c == nil {
		t.Errorf("expected to be able to combine two chromosomes with same counts ")
	}
	for i := 0; i < c.Len(); i++ {
		if !(c.Index(i) == c1.Index(i) || c.Index(i) == c2.Index(i)) {
			t.Errorf("Index %d should have been from one of the two parent chromosomes", i)
		}
	}

}
