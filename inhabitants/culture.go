package inhabitants

import (
	"encoding/json"

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
	"monogamous": func(a, b *Being) bool {
		return len(a.Spouses) == 0 && len(b.Spouses) == 0
	},
	"heterosexual": func(a, b *Being) bool {
		return a.Sex != b.Sex
	},
	"homosexual": func(a, b *Being) bool {
		return a.Sex == b.Sex
	},
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

var NameStrategies = map[string]NameStrategy{
	"matrilineal": func(b *Being) *Name {
		namer := b.Species.Genders[b.Sex].Namer
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
		namer := b.Species.Genders[b.Sex].Namer
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
		namer := b.Species.Genders[b.Sex].Namer
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
		namer := b.Species.Genders[b.Sex].Namer
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
		namer := b.Species.Genders[b.Sex].Namer
		name := &Name{GivenName: namer.GivenName()}
		display, _ := namer.Execute(name)
		name.Display = display
		return name
	},
}