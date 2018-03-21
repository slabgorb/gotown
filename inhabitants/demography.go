package inhabitants

// DemographyBucket splits the ages of a being, i.e. Child, Elderly
type DemographyBucket int

// DemographyBucket enum
const (
	Child DemographyBucket = iota
	Teenager
	YoungAdult
	EarlyAdult
	Adult
	MiddleAge
	Senior
	Elderly
	Ancient
)

// Demo forms part of a demography
type Demo struct {
	MaxAge            int `json:"max_age"`
	CumulativePercent int `json:"cumulative_percent"`
}
