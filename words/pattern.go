package words

import (
	"fmt"
	"text/template"
)

type Pattern string

func (p Pattern) Template() *template.Template {
	return template.Must(template.New(fmt.Sprint(p)).Parse(fmt.Sprint(p)))
}
