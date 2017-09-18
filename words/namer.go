package words

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"text/template"
)

type Namer struct {
	Patterns []Pattern
	*Words
	NameStrategy string
}

func (n *Namer) Template() *template.Template {
	randomChoice := n.Patterns[randomizer.Intn(len(n.Patterns))]
	return randomChoice.Template()
}

func lowercaseJoiners(s string) string {
	s = strings.Replace(s, " Of ", " of ", -1)
	s = strings.Replace(s, " The ", " the ", -1)
	s = strings.Replace(s, " And ", " and ", -1)
	return s
}

func edgeCases(s string) string {
	re := regexp.MustCompile(`yey$`)
	return re.ReplaceAllString(s, "y")
}

func (n *Namer) Execute(with interface{}) (string, error) {
	tmpl := n.Template()
	buf := bytes.NewBuffer([]byte(""))
	err := tmpl.Execute(buf, with)
	return edgeCases(lowercaseJoiners(strings.Title(buf.String()))), err
}

func (n *Namer) Name() string {
	s, err := n.Execute(n.Words)
	if err != nil {
		log.Println(err)
	}
	return s
}

func NewNamer(patterns []string, words *Words, nameStrategy string) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	return &Namer{Patterns: ps, Words: words, NameStrategy: nameStrategy}
}
