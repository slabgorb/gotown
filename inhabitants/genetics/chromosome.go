package genetics

import (
	"fmt"
	"math/rand"
	"strings"
)

// Chromosome represents a collection of genes
type Chromosome struct {
	Genes []Gene `json:"genes"`
}

// RandomChromosome returns a Chromosome with randomized genes.
func RandomChromosome(count int) *Chromosome {
	c := &Chromosome{}
	c.Randomize(count)
	return c
}

// Add adds a 'custom' gene to a Chromosome, used for testing, usually
func (c *Chromosome) Add(s string) {
	c.Genes = append(c.Genes, Gene(s))

}

// Index returns the gene at the passed in index number
func (c Chromosome) Index(i int) Gene {
	return c.Genes[i]
}

// Len returns the count of genes for this Chromosome
func (c Chromosome) Len() int {
	return len(c.Genes)
}

// String representation of the Chromosome, a concatenation of all the gene strings.
func (c Chromosome) String() string {
	s := make([]string, len(c.Genes))
	for i := 0; i < c.Len(); i++ {
		s[i] = c.Genes[i].String()
	}
	return strings.Join(s, "")
}

// Combine combines two Chromosomes, used to simulate sexual reproduction
func (c *Chromosome) Combine(other *Chromosome) (*Chromosome, error) {
	if c.Len() != other.Len() {
		return nil, fmt.Errorf("cannot combine chromosomes of differing gene counts, got %d and %d", c.Len(), other.Len())
	}
	combined := &Chromosome{}
	combined.Genes = make([]Gene, c.Len())
	for i := 0; i < c.Len(); i++ {
		if rand.Float64() > 0.5 {
			combined.Genes[i] = c.Genes[i]
		} else {
			combined.Genes[i] = other.Genes[i]
		}
	}
	return combined, nil
}

// Express applies an Expression to the Chromosome
func (c *Chromosome) Express(e Expression) map[string]string {
	set := c.String()
	express := make(map[string]string)
	for _, trait := range e.GetTraits() {
		k, _ := trait.Expression(set)
		express[trait.GetName()] = k
	}
	return express
}

// Randomize randomizes all the genes in this Chromosome.
func (c *Chromosome) Randomize(count int) {
	c.Genes = make([]Gene, count)
	for i := 0; i < count; i++ {
		c.Genes[i] = RandomGene()
	}
}
