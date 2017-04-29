package words

import (
	"bytes"
	"log"
	"math/rand"
	"strings"
	"text/template"
)

type Namer struct {
	Patterns []Pattern
	*Words
}

func (n *Namer) Template() *template.Template {
	randomChoice := n.Patterns[rand.Intn(len(n.Patterns))]
	return randomChoice.Template()
}

func lowercaseJoiners(s string) string {
	s = strings.Replace(s, " Of ", " of ", -1)
	s = strings.Replace(s, " The ", " the ", -1)
	s = strings.Replace(s, " And ", " and ", -1)
	return s
}

func (n *Namer) Execute(with interface{}) (string, error) {
	tmpl := n.Template()
	buf := bytes.NewBuffer([]byte(""))
	err := tmpl.Execute(buf, with)
	return lowercaseJoiners(strings.Title(buf.String())), err
}

func (n *Namer) Name() string {
	s, err := n.Execute(n.Words)
	if err != nil {
		log.Println(err)
	}
	return s
}

func NewNamer(patterns []string, words *Words) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	return &Namer{ps, words}
}
