package inhabitants

type DemographyBucket int

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

type Demo struct {
	MaxAge            int `json:"max_age"`
	CumulativePercent int `json:"cumulative_percent"`
}
