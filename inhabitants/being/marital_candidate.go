package being

import "fmt"

// MaritalCandidate is a pair of being
type MaritalCandidate struct {
	male, female *Being
}

// NewMaritalCandidate returns a new MaritalCandidate
func NewMaritalCandidate(male, female *Being) *MaritalCandidate {
	return &MaritalCandidate{male: male, female: female}
}

func (m *MaritalCandidate) String() string {
	return fmt.Sprintf("%s <-> %s", m.female.String(), m.male.String())
}

// Contains returns true if the passed in being is part of this pair.
func (m *MaritalCandidate) Contains(b *Being) bool {
	return m.male.ID == b.ID || m.female.ID == b.ID
}

// Pair returns the underlying beings of a marital candidate pair
func (m *MaritalCandidate) Pair() (*Being, *Being) {
	return m.male, m.female
}
