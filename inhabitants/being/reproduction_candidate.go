package being

// ReproductionCandidate blah blah (maybe should be unexported)
type ReproductionCandidate struct {
	b     *Being
	score float64
}

func (rc *ReproductionCandidate) Score() float64 {
	return rc.score
}

func (rc *ReproductionCandidate) Being() *Being {
	return rc.b
}
