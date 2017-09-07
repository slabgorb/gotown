package main

type genderNameSet struct {
	Gender     string   `json:"gender"`
	Patterns   []string `json:"patterns"`
	Givennames []string `json:"given_names"`
}

type nameSet struct {
	Name        string                   `json:"name"`
	Patronymics []string                 `json:"patronymics"`
	Matronymics []string                 `json:"matronymics"`
	GenderNames map[string]genderNameSet `json:"gender_names"`
}
