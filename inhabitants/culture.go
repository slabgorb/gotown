package inhabitants

import "github.com/slabgorb/gotown/words"

// Culture represents the culture of a population, such as the naming schemes, marriange customs, etc.
type Culture struct {
	NameStrategy    string
	MaritalStrategy []string
	FamilyNames     []string
	Patterns        map[Gender]words.Pattern
	GivenNames      map[Gender][]string
}

// MaritalStrategy is a function which indicates whether the two beings are marriage candidates
type MaritalStrategy func(a, b *Being) bool

// NameStrategy is a function which describes how children are named
type NameStrategy func(b *Being) *Name

var MaritalStrategies = map[string]MaritalStrategy{
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
