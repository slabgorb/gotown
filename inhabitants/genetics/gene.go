package genetics

import (
	"fmt"
	"strconv"
)

// Gene is part of a Chromosome, represented by a hexstring
type Gene string

// Int64 converts the hexstring of a Gene into an integer
func (g Gene) Int64() int64 {
	v, _ := strconv.ParseInt(string(g), 16, 64)
	return v
}

// String implements fmt.Stringer
func (g Gene) String() string {
	return string(g)
}

// RandomGene creates a randomized gene
func RandomGene() Gene {
	r := randomizer.Intn(16777215)
	return Gene(fmt.Sprintf("%06x", r))
}
