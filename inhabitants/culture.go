package inhabitants

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/slabgorb/gotown/words"
)

// Culture represents the culture of a population, such as the naming schemes,
// marriage customs, etc.
type Culture struct {
	Name              string
	nameStrategies    map[Gender]NameStrategy
	maritalStrategies []maritalStrategy
	Namers            map[Gender]*words.Namer
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
			return (a.Age/2)+7 < b.Age && a.Age >= b.Age
		}
		return (b.Age/2)+7 < a.Age && b.Age >= a.Age
	},
	"close age female older": func(a, b *Being) bool {
		// divide by 2 add 7
		if a.Sex == Gender("female") {
			return (a.Age/2)+7 < b.Age && a.Age >= b.Age
		}
		return (b.Age/2)+7 < a.Age && b.Age >= a.Age
	},
	"close age": func(a, b *Being) bool {
		return (a.Age/2)+7 < b.Age && (b.Age/2)+7 < a.Age
	},
	"unrelated": func(a, b *Being) bool {
		return !a.IsCloseRelativeOf(b)
	},
}

func (c *Culture) Load(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&c)
	if err != nil {
		return fmt.Errorf("could not parse json file")
	}
	return nil
}

func (c *Culture) String() string {
	return fmt.Sprintf("%s", c.Name)
}

func (c *Culture) UnmarshalJSON(data []byte) error {
	type cultureLoader struct {
		Name              string   `json:"name"`
		MaritalStrategies []string `json:"marital_strategies"`
		Patronymics       []string `json:"patronymics"`
		Matronymics       []string `json:"matronymics"`
		FamilyNames       []string `json:"family_names"`
		GenderNames       []struct {
			Gender       Gender   `json:"gender"`
			Patterns     []string `json:"patterns"`
			GivenNames   []string `json:"given_names"`
			NameStrategy string   `json:"name_strategy"`
		} `json:"gender_names"`
	}
	cl := &cultureLoader{}
	err := json.Unmarshal(data, cl)
	if err != nil {
		return err
	}
	c.Name = cl.Name
	for _, ms := range cl.MaritalStrategies {
		c.maritalStrategies = append(c.maritalStrategies, maritalStrategies[ms])
	}
	c.nameStrategies = make(map[Gender]NameStrategy)
	c.Namers = make(map[Gender]*words.Namer)
	for _, gn := range cl.GenderNames {
		w := words.NewWords()
		w.AddList("patronymics", cl.Patronymics)
		w.AddList("matronymics", cl.Matronymics)
		w.AddList("givenNames", gn.GivenNames)
		w.AddList("familyNames", cl.FamilyNames)
		c.Namers[gn.Gender] = words.NewNamer(gn.Patterns, w, gn.NameStrategy)
		c.nameStrategies[gn.Gender] = NameStrategies[gn.NameStrategy]
	}

	return nil
}

func (c *Culture) MaritalCandidate(a, b *Being) bool {
	out := true
	for _, s := range c.maritalStrategies {
		out = out && s(a, b)
	}
	return out
}

func (c *Culture) GetName(b *Being) *Name {
	f := c.nameStrategies[b.Sex]
	return f(b)
}

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
