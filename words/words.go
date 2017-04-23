package words

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

func chooseRandomString(s []string) string {
	return s[rand.Intn(len(s))]
}

type Words struct {
	Dictionary map[string][]string
	Backup     *Words
}

func (w *Words) Noun() (string, bool) {
	return w.withBackup(func() (string, bool) { return w.listFromKey("nouns") })
}

func (w *Words) PluralNoun() (string, bool) {
	if s, ok := w.Noun(); ok {
		return inflection.Plural(s), true
	}
	return "", false
}

func (w *Words) Adjective() (string, bool) {
	return w.withBackup(func() (string, bool) { return w.listFromKey("adjectives") })
}

func (w *Words) withBackup(f func() (string, bool)) (string, bool) {
	if s, ok := f(); ok {
		log.Println(s)
		return s, true
	} else {
		log.Println("Not ok", s)
	}
	log.Println("checking backup")
	if w.Backup != nil {
		return w.Backup.withBackup(f)
	}
	return "", false
}

func (w *Words) listFromKey(s string) (string, bool) {
	if list, ok := w.Dictionary[s]; ok {
		return chooseRandomString(list), true
	} else {
		return "", false
	}
}

func (w *Words) Prefix() (string, bool) {
	return w.withBackup(func() (string, bool) { return w.listFromKey("prefixes") })
}

func (w *Words) StartNoun() (string, bool) {
	if r, ok := w.withBackup(func() (string, bool) { return w.listFromKey("startNouns") }); ok {
		return r, ok
	}
	return w.withBackup(func() (string, bool) { return w.listFromKey("nouns") })
}

func (w *Words) EndNoun() (string, bool) {
	if r, ok := w.withBackup(func() (string, bool) { return w.listFromKey("endNouns") }); ok {
		return r, ok
	}
	return w.withBackup(func() (string, bool) { return w.listFromKey("nouns") })
}

func NewWords() *Words {
	return &Words{Dictionary: make(map[string][]string)}
}

func (w *Words) AddList(key string, list []string) {
	w.Dictionary[key] = list
}

type Pattern string

func (p Pattern) Template() *template.Template {
	return template.Must(template.New(fmt.Sprint(p)).Parse(fmt.Sprint(p)))
}

type Namer struct {
	Patterns []Pattern
	*Words
}

func (n *Namer) template() *template.Template {
	randomChoice := n.Patterns[rand.Intn(len(n.Patterns))]
	return randomChoice.Template()
}

func lowercaseJoiners(s string) string {
	s = strings.Replace(s, " Of ", " of ", -1)
	s = strings.Replace(s, " The ", " the ", -1)
	s = strings.Replace(s, " And ", " and ", -1)
	return s
}

func (n *Namer) Name() string {
	tmpl := n.template()
	buf := bytes.NewBuffer([]byte(""))
	tmpl.Execute(buf, n.Words)
	return lowercaseJoiners(strings.Title(buf.String()))
}

func NewNamer(patterns []string, words *Words) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	return &Namer{ps, words}
}
