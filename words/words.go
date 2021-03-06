package words

import (
	"fmt"

	"github.com/jinzhu/inflection"
	"github.com/slabgorb/gotown/persist"
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
	persist.IdentifiableImpl
	Name       string              `json:"name" storm:"unique"`
	Dictionary map[string][]string `json:"dictionary"`
	backup     *Words
	BackupName string `json:"backup"`
}

func (w *Words) GetName() string           { return w.Name }
func (w *Words) String() string            { return w.Name }
func (w *Words) API() (interface{}, error) { return w, nil }

// Save implements persist.Persistable
func (w *Words) Save() error {
	return persist.Save(w)
}

// Delete implements persist.Persistable
func (w *Words) Delete() error {
	return persist.Delete(w)
}

// Fetch implements persist.Persistable
func (w *Words) Read() error {
	if err := persist.Read(w); err != nil {
		return fmt.Errorf("could not load words: %s", err)
	}
	return w.loadBackup()
}

func (w *Words) Reset() {
	w.ID = ""
	w.Name = ""
	w.Dictionary = make(map[string][]string)
	w.backup = nil
	w.BackupName = ""
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

func (w *Words) loadBackup() error {
	backupWords := &Words{}
	if w.BackupName == "" {
		return nil
	}
	if err := persist.ReadByName(w.BackupName, "Words", backupWords); err != nil {
		return fmt.Errorf("could not load backup words %s: %s", w.BackupName, err)
	}
	if backupWords == nil {
		return fmt.Errorf("could not load backup words %s", w.BackupName)
	}
	w.backup = backupWords
	return nil
}

func (w *Words) GetBackup() *Words {
	if w.backup == nil && w.BackupName != "" {
		w.loadBackup()
	}
	return w.backup
}

func (w *Words) SetBackup(b *Words) {
	w.backup = b
}

func (w *Words) withBackup(f func(w *Words) string) string {
	if s := f(w); s != "" {
		return s
	}
	if w.backup != nil && w.backup.ID != "" {
		return w.backup.withBackup(f)
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
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("given_names")) })
}

// FamilyName returns a family name
func (w *Words) FamilyName() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("family_names")) })
}

func (w *Words) Nickname() string {
	return w.withBackup(func(w *Words) string { return chooseRandomString(w.listFromKey("nicknames")) })
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
func NewWords(name string) *Words {
	return &Words{Name: name, Dictionary: make(map[string][]string)}
}

// AddList adds a list of words to a particular key, e.g. 'noun'
func (w *Words) AddList(key string, list []string) {
	w.Dictionary[key] = list
}

func Seed() {
	if err := seedWords(); err != nil {
		panic(err)
	}
	if err := seedNamers(); err != nil {
		panic(err)
	}
}

func WordsList() (map[string]string, error) {
	list, err := persist.List("Words")
	if err != nil {
		return nil, err
	}
	return list, err
}

func seedWords() error {
	var words = &Words{}
	return persist.SeedHelper("words", words)
}

func seedNamers() error {
	var namer = &Namer{}
	return persist.SeedHelper("namers", namer)
}
