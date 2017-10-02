package inhabitants

type DemographyBucket int

const (
	Child DemographyBucket = iota
	YoungAdult
	EarlyAdult
	Adult
	MiddleAge
	LateAdult
	Senior
	Elderly
	Ancient
)

type demo struct {
	max int
	pct int
}
type demography map[DemographyBucket]demo

// Demographies models the population curve, as well as providing information into fertility, age-related changes, etc.
var Demographies = map[string]demography{
	"medieval": demography{
		Child:      demo{14, 29},
		YoungAdult: demo{18, 36},
		EarlyAdult: demo{26, 50},
		Adult:      demo{31, 58},
		MiddleAge:  demo{41, 72},
		LateAdult:  demo{51, 84},
		Senior:     demo{61, 93},
		Elderly:    demo{71, 98},
		Ancient:    demo{100, 100},
	},
}
