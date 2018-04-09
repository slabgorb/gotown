package culture

import (
	"fmt"

	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

// Culture represents the culture of a population, such as the naming schemes,
// marriage customs, etc.
type Culture struct {
	ID                int                                 `json:"id" storm:"id,increment"`
	Name              string                              `json:"name" storm:"unique"`
	NameStrategies    map[inhabitants.Gender]string       `json:"name_strategies"`
	MaritalStrategies []string                            `json:"marital_strategies"`
	Namers            map[inhabitants.Gender]*words.Namer `json:"names"`
}

// maritalStrategy is a function which indicates whether the two beings are
// marriage candidates
type maritalStrategy func(a, b inhabitants.Marriageable) bool

var maritalStrategies = map[string]maritalStrategy{
	"living": func(a, b inhabitants.Marriageable) bool {
		return a.Alive() && b.Alive()
	},
	"monogamous": func(a, b inhabitants.Marriageable) bool {
		return a.Unmarried() && b.Unmarried()
	},
	"heterosexual": func(a, b inhabitants.Marriageable) bool {
		return a.Sex() != b.Sex()
	},
	"homosexual": func(a, b inhabitants.Marriageable) bool {
		return a.Sex() == b.Sex()
	},
	"close age male older": func(a, b inhabitants.Marriageable) bool {
		// divide by 2 add 7
		if a.Sex() == inhabitants.Gender("male") {
			return (a.Age()/2)+7 < b.Age() && a.Age() >= b.Age()
		}
		return (b.Age()/2)+7 < a.Age() && b.Age() >= a.Age()
	},
	"close age female older": func(a, b inhabitants.Marriageable) bool {
		// divide by 2 add 7
		if a.Sex() == inhabitants.Gender("female") {
			return (a.Age()/2)+7 < b.Age() && a.Age() >= b.Age()
		}
		return (b.Age()/2)+7 < a.Age() && b.Age() >= a.Age()
	},
	"close age": func(a, b inhabitants.Marriageable) bool {
		return (a.Age()/2)+7 < b.Age() && (b.Age()/2)+7 < a.Age()
	},
	"unrelated": func(a, b inhabitants.Marriageable) bool {
		return !a.IsCloseRelativeOf(b)
	},
}

// String implements fmt.Stringer
func (c *Culture) String() string {
	return fmt.Sprintf("%s", c.Name)
}

// UnmarshalJSON implements json.Unmarshaler
// func (c *Culture) UnmarshalJSON(data []byte) error {
// 	cl := &cultureSerializer{}
// 	err := json.Unmarshal(data, cl)
// 	if err != nil {
// 		return err
// 	}
// 	c.Name = cl.Name
// 	c.MaritalStrategies = cl.MaritalStrategies
// 	c.NameStrategies = make(map[Gender]string)
// 	c.Namers = make(map[Gender]*words.Namer)
// 	for _, gn := range cl.GenderNames {
// 		w := words.NewWords()
// 		w.AddList("patronymics", cl.Patronymics)
// 		w.AddList("matronymics", cl.Matronymics)
// 		w.AddList("givenNames", gn.GivenNames)
// 		w.AddList("familyNames", cl.FamilyNames)
// 		c.Namers[gn.Gender] = words.NewNamer(gn.Patterns, w, gn.NameStrategy)
// 		c.NameStrategies[gn.Gender] = gn.NameStrategy
// 	}

// 	return nil
// }

// // MarshalJSON implements json.marshaler
// func (c *Culture) MarshalJSON() ([]byte, error) {
// 	cl := &cultureSerializer{}
// 	cl.Name = c.Name
// 	cl.MaritalStrategies = c.MaritalStrategies
// 	cl.GenderNames = []genderNamesSerializer{}
// 	for gender, gn := range c.Namers {
// 		cl.Patronymics = gn.Dictionary["patronymics"]
// 		cl.Matronymics = gn.Dictionary["matronymics"]
// 		cl.FamilyNames = gn.Dictionary["familyNames"]
// 		cl.GenderNames = append(cl.GenderNames, genderNamesSerializer{
// 			Gender:       gender,
// 			Patterns:     gn.PatternList(),
// 			NameStrategy: c.NameStrategies[gender],
// 			GivenNames:   gn.Dictionary["givenNames"],
// 		})

// 	}
// 	return json.Marshal(cl)
// }

// Save implements persist.Persistable
func (c *Culture) Save() error {
	return persist.DB.Save(c)
}

// Delete implements persist.Persistable
func (c *Culture) Delete() error {
	return persist.DB.DeleteStruct(c)
}

// Fetch implements persist.Persistable
func (c *Culture) Read() error {
	return persist.DB.One("Name", c.Name, c)
}

func (c *Culture) Reset() {
	c.ID = 0
	c.Name = ""
	c.NameStrategies = make(map[inhabitants.Gender]string)
	c.MaritalStrategies = []string{}
	c.Namers = make(map[inhabitants.Gender]*words.Namer)
}

// MaritalCandidate decides whether this pair of Beings is a valid candidate for
// marriage, based on the culture's marital rules.
func (c *Culture) MaritalCandidate(a, b inhabitants.Marriageable) bool {
	out := true
	for _, s := range c.MaritalStrategies {
		out = out && maritalStrategies[s](a, b)
	}
	return out
}

// GetName returns a name appropriate for the passed in Being
func (c *Culture) GetName(b inhabitants.Nameable) *inhabitants.Name {
	f := c.NameStrategies[b.Sex()]
	return inhabitants.NameStrategies[f](b, c)
}

func (c *Culture) GetNamers() map[inhabitants.Gender]*words.Namer {
	return c.Namers
}

func (c *Culture) RandomName(sex inhabitants.Gender, b inhabitants.Nameable) *inhabitants.Name {
	f := c.NameStrategies[sex]
	return inhabitants.NameStrategies[f](b, c)
}

func Seed() error {
	var culture = &Culture{}
	return persist.SeedHelper("cultures", culture)
}
