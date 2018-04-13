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
	MaritalStrategies []string                            `json:"marital_strategies"`
	Namers            map[inhabitants.Gender]*words.Namer `json:"namers"`
	NamerNames        map[inhabitants.Gender]string       `json:"namer_names"`
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
	if err := persist.DB.One("Name", c.Name, c); err != nil {
		return err
	}
	c.Namers = make(map[inhabitants.Gender]*words.Namer)
	for gender, namerName := range c.NamerNames {
		fmt.Println(gender, namerName)
		n := &words.Namer{Name: namerName}
		if err := n.Read(); err != nil {
			return err
		}
		c.Namers[gender] = n
	}
	return nil
}

func (c *Culture) Reset() {
	c.ID = 0
	c.Name = ""
	c.MaritalStrategies = []string{}
	c.Namers = make(map[inhabitants.Gender]*words.Namer)
	c.NamerNames = make(map[inhabitants.Gender]string)
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
	namer := c.Namers[b.Sex()]
	return inhabitants.NameStrategies[namer.NameStrategy](b, c)
}

func (c *Culture) GetNamers() map[inhabitants.Gender]*words.Namer {
	return c.Namers
}

func (c *Culture) RandomName(sex inhabitants.Gender, b inhabitants.Nameable) *inhabitants.Name {
	namer := c.Namers[sex]
	return inhabitants.NameStrategies[namer.NameStrategy](b, c)
}

func Seed() error {
	var culture = &Culture{}
	return persist.SeedHelper("cultures", culture)
}

func List() ([]string, error) {
	cultures := []Culture{}
	if err := persist.DB.All(&cultures); err != nil {
		return nil, err
	}
	names := []string{}
	for _, c := range cultures {
		names = append(names, c.Name)
	}
	return names, nil
}
