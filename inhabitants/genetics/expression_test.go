package genetics_test

import (
	"encoding/json"
	"testing"

	. "github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/random"
)

var expressionJSON = `
{
   "traits":[
      {
         "name":"eye color",
         "variants":[
            {
               "name":"brown",
               "match":"f(9|a|b|c|d)"
            },
            {
               "name":"hazel",
               "match":"f(6|7|8)"
            },
            {
               "name":"blue",
               "match":"f1|f2"
            },
            {
               "name":"gray",
               "match":"f3"
            },
            {
               "name":"green",
               "match":"f4|f5"
            }
         ]
      },
      {
         "name":"eye shape",
         "variants":[
            {
               "name":"almond",
               "match":"e(1|2|3)"
            },
            {
               "name":"close set",
               "match":"e(4|5)"
            },
            {
               "name":"deep set",
               "match":"e(6|7)"
            },
            {
               "name":"protruding",
               "match":"e(8,9)"
            },
            {
               "name":"deep set",
               "match":"e(a|b)"
            }
         ]
      },
      {
         "name":"lip shape",
         "variants":[
            {
               "name":"bee sting",
               "match":"d1"
            },
            {
               "name":"thin",
               "match":"d(2|3)"
            },
            {
               "name":"cupid bow",
               "match":"d4"
            },
            {
               "name":"natural",
               "match":"d(5,6)"
            }
         ]
      }
   ]
}
`

func init() {
	SetRandomizer(random.NewMock())

}

func TestExpression(t *testing.T) {
	exp := Expression{}
	reader := []byte(expressionJSON)
	err := json.Unmarshal(reader, &exp)
	if err != nil {
		t.Error(err)
	}

	c := RandomChromosome(40)
	e := c.Express(exp)
	expected := "hazel"
	if e["eye color"] != expected {
		t.Errorf("Expected %s got %s", expected, e["eye color"])
	}
}
