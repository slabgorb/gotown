package inhabitants

import (
	"fmt"

	"github.com/slabgorb/gotown/persist"

	"github.com/slabgorb/gotown/words"
)

// Culture represents the culture of a population, such as the naming schemes,
// marriage customs, etc.
type Culture struct {
	ID                int `json:"id" storm:"id,increment"`
	Name              string                  `json:"name" storm:"unique"`
	NameStrategies    map[Gender]string       `json:"name_strategies"`
	MaritalStrategies []string                `json:"marital_strategies"`
	Namers            map[Gender]*words.Namer `json:"names"`
}

// maritalStrategy is a function which indicates whether the two beings are
// marriage candidates
type maritalStrategy func(a, b *Being) bool

// NameStrategy is a function which describes how children are named
type NameStrategy func(b *Being) *Name

var maritalStrategies = map[string]maritalStrategy{
	"living": func(a, b *Being) bool {
		return a.Alive() && b.Alive()
	},
	"monogamous": func(a, b *Being) bool {
		return len(a.Spouses) == 0 && len(b.Spouses) == 0
	},
	"heterosexual": func(a, b *Being) bool {
		return a.Sex != b.Sex
	},
	"homosexual": func(a, b *Being) bool {
		return a.Sex == b.Sex
	},
	"close age male older": func(a, b *Being) bool {
		// divide by 2 add 7
		if a.Sex == Gender("male") {
			return (a.Age()/2)+7 < b.Age() && a.Age() >= b.Age()
		}
		return (b.Age()/2)+7 < a.Age() && b.Age() >= a.Age()
	},
	"close age female older": func(a, b *Being) bool {
		// divide by 2 add 7
		if a.Sex == Gender("female") {
			return (a.Age()/2)+7 < b.Age() && a.Age() >= b.Age()
		}
		return (b.Age()/2)+7 < a.Age() && b.Age() >= a.Age()
	},
	"close age": func(a, b *Being) bool {
		return (a.Age()/2)+7 < b.Age() && (b.Age()/2)+7 < a.Age()
	},
	"unrelated": func(a, b *Being) bool {
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

// MaritalCandidate decides whether this pair of Beings is a valid candidate for
// marriage, based on the culture's marital rules.
func (c *Culture) MaritalCandidate(a, b *Being) bool {
	out := true
	for _, s := range c.MaritalStrategies {
		out = out && maritalStrategies[s](a, b)
	}
	return out
}

// GetName returns a name appropriate for the passed in Being
func (c *Culture) GetName(b *Being) *Name {
	f := c.NameStrategies[b.Sex]
	return NameStrategies[f](b)
}

// NameStrategies deliniates the various naming strategy functions
var NameStrategies = map[string]NameStrategy{
	"matrilineal": func(b *Being) *Name {
		namer := b.Culture.Namers[b.Sex]
		name := &Name{GivenName: namer.GivenName()}
		if b.Mother() != nil {
			name.FamilyName = b.Mother().FamilyName
			return name
		}
		name.FamilyName = namer.GivenName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patrilineal": func(b *Being) *Name {
		namer := b.Culture.Namers[b.Sex]
		name := &Name{GivenName: namer.GivenName()}
		if b.Father() != nil {
			name.FamilyName = b.Father().FamilyName
			return name
		}
		name.FamilyName = namer.GivenName()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"matronymic": func(b *Being) *Name {
		namer := b.Culture.Namers[b.Sex]
		name := &Name{GivenName: namer.GivenName()}
		if b.Mother() != nil {
			name.FamilyName = b.Mother().GivenName + namer.Matronymic()
			return name
		}
		name.FamilyName = namer.GivenName() + namer.Matronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"patronymic": func(b *Being) *Name {
		namer := b.Culture.Namers[b.Sex]
		name := &Name{GivenName: namer.GivenName()}
		if b.Father() != nil {
			name.FamilyName = b.Father().GivenName + namer.Patronymic()
			return name
		}
		name.FamilyName = namer.GivenName() + namer.Patronymic()
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
	"onename": func(b *Being) *Name {
		namer := b.Culture.Namers[b.Sex]
		name := &Name{GivenName: namer.GivenName()}
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
}
