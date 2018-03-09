package words

import (
	"fmt"
	"text/template"
)

// Pattern is a template.Template denoting the pattern for a name, like
// '{{GivenName}} {{FamilyName}}'
type Pattern string

// Template compiles the pattern template and returns the compiled template.
// Will panic if the pattern cannot compile?
func (p Pattern) Template() *template.Template {
	return template.Must(template.New(fmt.Sprint(p)).Parse(fmt.Sprint(p)))
}
