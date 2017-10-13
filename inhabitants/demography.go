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

type Demo struct {
	MaxAge            int `json:"max_age"`
	CumulativePercent int `json:"cumulative_percent"`
}

//
// // Demographies models the population curve, as well as providing information into fertility, age-related changes, etc.
// var Demographies = map[string]Demography{
// 	"human": Demography{
// 		Child:      Demo{14, 29},
// 		YoungAdult: Demo{18, 36},
// 		EarlyAdult: Demo{26, 50},
// 		Adult:      Demo{31, 58},
// 		MiddleAge:  Demo{41, 72},
// 		LateAdult:  Demo{51, 84},
// 		Senior:     Demo{61, 93},
// 		Elderly:    Demo{71, 98},
// 		Ancient:    Demo{100, 100},
// 	},
// }
