package townomatic

import "text/template"

const (
	Asexual = iota
	Male
	Female
)

type Gender int

type FertileAge struct {
	Gender
	Start int
	End   int
}

type Species struct {
	Name         string
	Genders      []Gender
	FertileAges  []FertileAge
	NameTemplate *template.Template
}

func NewSpecies(name string, genders []int, fertileAges []FertileAge, templateString string) (*Species, error) {
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
