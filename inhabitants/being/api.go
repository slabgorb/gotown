package being

type API struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	SpeciesName string            `json:"species"`
	CultureName string            `json:"culture"`
	Parents     []string          `json:"parents"`
	Children    []string          `json:"children"`
	Spouses     []string          `json:"spouses"`
	Gender      string            `json:"gender"`
	Age         int               `json:"age"`
	Chromosome  string            `json:"chromosome"`
	Expression  map[string]string `json:"expression"`
}

func APIList(beings []*Being) ([]*API, error) {
	apis := []*API{}
	for _, b := range beings {
		api, err := b.GetAPI()
		if err != nil {
			return nil, err
		}
		apis = append(apis, api)
	}
	return apis, nil
}

func getStrings(beings []*Being) []string {
	display := []string{}
	for _, b := range beings {
		display = append(display, b.Name.GetDisplay())
	}
	return display
}

func (b *Being) GetAPI() (*API, error) {
	parents, err := b.GetParents()
	if err != nil {
		return nil, err
	}
	children, err := b.GetChildren()
	if err != nil {
		return nil, err
	}
	spouses, err := b.getSpouses()
	if err != nil {
		return nil, err
	}
	expression := b.Expression()
	return &API{
		ID:          b.ID,
		Parents:     getStrings(parents),
		Children:    getStrings(children),
		Spouses:     getStrings(spouses),
		Name:        b.Name.GetDisplay(),
		SpeciesName: b.Species.Name,
		CultureName: b.Culture.Name,
		Age:         b.Age,
		Gender:      b.Gender.String(),
		Chromosome:  b.Chromosome.String(),
		Expression:  expression,
	}, nil
}
