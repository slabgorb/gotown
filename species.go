package gotown

import "text/template"

type Gender int

const (
	Asexual Gender = iota
	Male
	Female
)

type FertileAge struct {
	Gender
	Start int
	End   int
}

// Species represents a species or a race.
type Species struct {
	Name         string
	Genders      []Gender
	FertileAges  []FertileAge
	NameTemplate *template.Template
}

func NewSpecies(name string, genders []Gender, fertileAges []FertileAge, templateString string) (*Species, error) {
	nameTemplate, err := template.New(name + "_nameTemplate").Parse(templateString)
	if err != nil {
		return nil, err
	}
	g := []Gender{}
	for _, i := range genders {
		g = append(g, Gender(i))
	}
	s := &Species{
		Name:         name,
		Genders:      g,
		FertileAges:  fertileAges,
		NameTemplate: nameTemplate,
	}
	return s, nil
}

func (s *Species) String() string {
	return s.Name
}
