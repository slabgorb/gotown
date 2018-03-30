package species

// DemographyBucket enum
const (
	Child int = iota
	Teenager
	YoungAdult
	EarlyAdult
	Adult
	MiddleAge
	Senior
	Elderly
	Ancient
	MaxDemographyBucket
)

const Random int = -1

// Demo forms part of a demography
type Demo struct {
	MaxAge            int `json:"max_age"`
	CumulativePercent int `json:"cumulative_percent"`
}
