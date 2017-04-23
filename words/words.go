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

const shortWord = 10

func chooseRandomString(s []string) string {
	return s[rand.Intn(len(s))]
}

type Words struct {
	Dictionary map[string][]string
	Backup     *Words
}

func (w *Words) Noun() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
}

func (w *Words) PluralNoun() string {
	if s := w.Noun(); s != "" {
		return inflection.Plural(s)
	}
	return ""
}

func (w *Words) Adjective() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("adjectives")) })
}

func (w *Words) Suffix() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("suffixes")) })
}

func (w *Words) ShortAdjective() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(shortFilter(w.listFromKey("adjectives"))) })
}

func (w *Words) ShortNoun() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(shortFilter(w.listFromKey("nouns"))) })
}

func shortFilter(list []string) []string {
	newList := []string{}
	for _, l := range list {
		if len(l) <= shortWord {
			newList = append(newList, l)
		}
	}
	return newList
}

func (w *Words) withBackup(f func(w *Words) string) string {
	if s := f(w); s != "" {
		return s
	}
	if w.Backup != nil {
		return w.Backup.withBackup(f)
	}
	return ""
}

func (w *Words) listFromKey(s string) []string {
	if list, ok := w.Dictionary[s]; ok {
		return list
	} else {
		return []string{""}
	}
}

func (w *Words) Prefix() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("prefixes")) })
}

func (w *Words) StartNoun() string {
	if r := w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("startNouns")) }); r != "" {
		return r
	}
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
}

func (w *Words) EndNoun() string {
	if r := w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("endNouns")) }); r != "" {
		return r
	}
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
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
	err := tmpl.Execute(buf, n.Words)
	if err != nil {
		log.Println(err)
	}
	return lowercaseJoiners(strings.Title(buf.String()))
}

func NewNamer(patterns []string, words *Words) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	return &Namer{ps, words}
}
