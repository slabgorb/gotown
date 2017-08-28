package words

import (
	"bytes"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/slabgorb/gotown/random"
)

type Namer struct {
	Patterns []Pattern
	*Words
	randomizer random.Generator
}

func (n *Namer) Template() *template.Template {
	n.SetDefaultRandomizer()
	randomChoice := n.Patterns[n.randomizer.Intn(len(n.Patterns))]
	return randomChoice.Template()
}

func (n *Namer) SetRandomizer(g random.Generator) {
	n.randomizer = g
}

func (n *Namer) SetDefaultRandomizer() {
	if n.randomizer == nil {
		n.SetRandomizer(rand.New(rand.NewSource(time.Now().UTC().UnixNano())))
	}
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

func NewNamer(patterns []string, words *Words) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	return &Namer{Patterns: ps, Words: words}
}
