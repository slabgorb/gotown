package words

import (
	"github.com/jinzhu/inflection"
	"github.com/slabgorb/gotown/random"
)

var randomizer random.Generator = random.Random

// SetRandomizer sets the package randomizer. Used in tests
func SetRandomizer(g random.Generator) {
	randomizer = g
}

const shortWord = 10

func chooseRandomString(s []string) string {
	return s[randomizer.Intn(len(s))]
}

type Words struct {
	Dictionary map[string][]string `json:"dictionary"`
	Backup     *Words              `json:"-"`
}

// Noun returns a noun
func (w *Words) Noun() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
}

// PluralNoun returns a pluralized noun
func (w *Words) PluralNoun() string {
	if s := w.Noun(); s != "" {
		return inflection.Plural(s)
	}
	return ""
}

// Adjective returns an adjective
func (w *Words) Adjective() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("adjectives")) })
}

// Suffix returns a suffix
func (w *Words) Suffix() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("suffixes")) })
}

// ShortAdjective returns an adjective shorter than the constant shortWord
func (w *Words) ShortAdjective() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(shortFilter(w.listFromKey("adjectives"))) })
}

// ShortNoun returns a noun shorter than the constant shortWord
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

// Prefix returns a prefix
func (w *Words) Prefix() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("prefixes")) })
}

// StartNoun returns a noun appropriate for starting a name, or a random noun if
// no start nouns are defined.
func (w *Words) StartNoun() string {
	if r := w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("startNouns")) }); r != "" {
		return r
	}
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
}

// EndNoun returns a noun appropriate for ending a name, or a random noun if
// no start nouns are defined.
func (w *Words) EndNoun() string {
	if r := w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("endNouns")) }); r != "" {
		return r
	}
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nouns")) })
}

// GivenName returns a given name
func (w *Words) GivenName() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("givenNames")) })
}

// FamilyName returns a family name
func (w *Words) FamilyName() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("familyNames")) })
}

// Matronymic returns a matronymic name
func (w *Words) Matronymic() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("matronymics")) })
}

// Patronymic returns a patronymic name
func (w *Words) Patronymic() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("patronymics")) })
}

// NewWords initializes a Words struct
func NewWords() *Words {
	return &Words{Dictionary: make(map[string][]string)}
}

// AddList adds a list of words to a particular key, e.g. 'noun'
func (w *Words) AddList(key string, list []string) {
	w.Dictionary[key] = list
}
