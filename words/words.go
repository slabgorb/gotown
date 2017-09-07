package words

import (
	"math/rand"
	"time"

	"github.com/jinzhu/inflection"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func SetRandomizer(g random.Generator) {
	randomizer = g
}

const shortWord = 10

func chooseRandomString(s []string) string {
	return s[randomizer.Intn(len(s))]
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
	}
	return []string{""}
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

func (w *Words) GivenName() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("givenNames")) })
}

func (w *Words) FamilyName() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("familyNames")) })
}

func (w *Words) Matronymic() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("matronymics")) })
}
func (w *Words) Patronymic() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("patronymics")) })
}

func NewWords() *Words {
	return &Words{Dictionary: make(map[string][]string)}
}

func (w *Words) AddList(key string, list []string) {
	w.Dictionary[key] = list
}
