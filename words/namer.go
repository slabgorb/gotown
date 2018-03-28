package words

import (
	"bytes"
	"log"
	"regexp"
	"strings"
	"text/template"

	"github.com/slabgorb/gotown/persist"
)

type Namer struct {
	Words        *Words
	ID           int       `json:"id" storm:"increment"`
	Name         string    `json:"name" storm:"unique"`
	Patterns     []Pattern `json:"patterns"`
	WordsName    string    `json:"words"`
	NameStrategy string    `json:"name_strategy"`
}

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
	return persist.DB.One("Name", n.Name, n)
}

func (n *Namer) Reset() {
	n.Words = nil
	n.ID = 0
	n.Name = ""
	n.Patterns = []Pattern{}
	n.WordsName = ""
	n.NameStrategy = ""
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

func (n *Namer) CreateName() string {
	s, err := n.Execute(n.Words)
	if err != nil {
		log.Println(err)
	}
	return s
}

func NewNamer(patterns []string, words string, nameStrategy string) *Namer {
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
