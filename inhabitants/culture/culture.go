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

type Marriageable interface {
	Unmarried() bool
	GetAge() int
	Alive() bool
	Sex() inhabitants.Gender
	IsCloseRelativeOf(with Marriageable) bool
}

// maritalStrategy is a function which indicates whether the two beings are
// marriage candidates
type maritalStrategy func(a, b Marriageable) bool

var maritalStrategies = map[string]maritalStrategy{
	"living": func(a, b Marriageable) bool {
		return a.Alive() && b.Alive()
	},
	"monogamous": func(a, b Marriageable) bool {
		return a.Unmarried() && b.Unmarried()
	},
	"heterosexual": func(a, b Marriageable) bool {
		return a.Sex() != b.Sex()
	},
	"homosexual": func(a, b Marriageable) bool {
		return a.Sex() == b.Sex()
	},
	"close age male older": func(a, b Marriageable) bool {
		// divide by 2 add 7
		if a.Sex() == inhabitants.Gender("male") {
			return (a.GetAge()/2)+7 < b.GetAge() && a.GetAge() >= b.GetAge()
		}
		return (b.GetAge()/2)+7 < a.GetAge() && b.GetAge() >= a.GetAge()
	},
	"close age female older": func(a, b Marriageable) bool {
		// divide by 2 add 7
		if a.Sex() == inhabitants.Gender("female") {
			return (a.GetAge()/2)+7 < b.GetAge() && a.GetAge() >= b.GetAge()
		}
		return (b.GetAge()/2)+7 < a.GetAge() && b.GetAge() >= a.GetAge()
	},
	"close age": func(a, b Marriageable) bool {
		return (a.GetAge()/2)+7 < b.GetAge() && (b.GetAge()/2)+7 < a.GetAge()
	},
	"unrelated": func(a, b Marriageable) bool {
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
	if c.Name == "" {
		return fmt.Errorf("cannot read culture without name")
	}
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

// Reset sets the culture back to zero
func (c *Culture) Reset() {
	c.ID = 0
	c.Name = ""
	c.MaritalStrategies = []string{}
	c.Namers = make(map[inhabitants.Gender]*words.Namer)
	c.NamerNames = make(map[inhabitants.Gender]string)
}

// MaritalCandidate decides whether this pair of Beings is a valid candidate for
// marriage, based on the culture's marital rules.
func (c *Culture) MaritalCandidate(a, b Marriageable) bool {
	out := true
	for _, s := range c.MaritalStrategies {
		out = out && maritalStrategies[s](a, b)
	}
	return out
}

// RandomName returns a name appropriate for the passed in Being
func (c *Culture) RandomName(b inhabitants.Nameable) *inhabitants.Name {
	namer := c.Namers[b.Sex()]
	return inhabitants.NameStrategies[namer.NameStrategy](b, c)
}

// GetNamers returns the namer objects for the culture
func (c *Culture) GetNamers() map[inhabitants.Gender]*words.Namer {
	return c.Namers
}

func (c *Culture) GetName() string { return c.Name }

// Seed seeds the database with initial cultures.
func Seed() error {
	var culture = &Culture{}
	return persist.SeedHelper("cultures", culture)
}

// List returns the names of the cultures already in tha database
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
