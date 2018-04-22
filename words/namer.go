package words

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"strings"
	"text/template"

	"github.com/slabgorb/gotown/persist"
)

// Namer 'names' things using an underlying Words struct and a set of template patterns.
type Namer struct {
	Words        *Words
	ID           int       `json:"id" storm:"increment"`
	Name         string    `json:"name" storm:"unique"`
	Patterns     []Pattern `json:"patterns"`
	WordsName    string    `json:"words"`
	NameStrategy string    `json:"name_strategy"`
}

// PatternList returns the set of patterns as a slice of string
func (n *Namer) PatternList() []string {
	pl := []string{}
	for _, p := range n.Patterns {
		pl = append(pl, string(p))
	}
	return pl
}

// Save implements persist.Persistable
func (n *Namer) Save() error {
	return persist.DB.Save(n)
}

// Delete implements persist.Persistable
func (n *Namer) Delete() error {
	return persist.DB.DeleteStruct(n)
}

// Fetch implements persist.Persistable
func (n *Namer) Read() error {
	if err := persist.DB.One("Name", n.Name, n); err != nil {
		return err
	}
	w := Words{Name: n.WordsName}
	if err := w.Read(); err != nil {
		return err
	}
	n.Words = &w
	return nil
}

// Reset implements persist.Persistable
func (n *Namer) Reset() {
	n.Words = nil
	n.ID = 0
	n.Name = ""
	n.Patterns = []Pattern{}
	n.WordsName = ""
	n.NameStrategy = ""
}

// Template chooses a random template pattern from the list of patterns
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

// Execute performs interpolations on a random template pattern using the
// underlying Words
func (n *Namer) Execute(with interface{}) (string, error) {
	return n.ExecuteWithTemplate(with, n.Template())
}

// Executable abstracts templates
type Exeutable interface {
	Execute(io.Writer, interface{}) error
}

// ExecuteWithTemplate performs Execute with a selected template pattern
func (n *Namer) ExecuteWithTemplate(with interface{}, tmpl Exeutable) (string, error) {
	buf := bytes.NewBuffer([]byte(""))
	err := tmpl.Execute(buf, with)
	return edgeCases(lowercaseJoiners(strings.Title(buf.String()))), err
}

// CreateName makes a random name
func (n *Namer) CreateName() string {
	s, err := n.Execute(n.Words)
	if err != nil {
		log.Println(err)
	}
	return s
}

// CreateNameWithPattern makes a random name with the specified pattern
func (n *Namer) CreateNameWithPattern(tmpl Exeutable) string {
	s, err := n.ExecuteWithTemplate(n.Words, tmpl)
	if err != nil {
		log.Println(err)
	}
	return s
}

// New returns an initialized Namer
func New(patterns []string, words string, nameStrategy string) *Namer {
	ps := []Pattern{}
	for _, p := range patterns {
		ps = append(ps, Pattern(p))
	}
	w := &Words{Name: words}
	if err := w.Read(); err != nil {
		panic(err)
	}
	return &Namer{Patterns: ps, WordsName: words, Words: w, NameStrategy: nameStrategy}
}

// NamerList returns a list of namers (as []string)
func NamerList() ([]string, error) {
	ns := []Namer{}
	if err := persist.DB.All(&ns); err != nil {
		return nil, err
	}
	names := []string{}
	for _, n := range ns {
		names = append(names, n.Name)
	}
	return names, nil
}
