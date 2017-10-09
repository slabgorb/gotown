package locations_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/slabgorb/gotown/inhabitants"
)

type tester interface {
	Fail()
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func helperLoadBytes(t tester, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func helperMockCulture(t tester, name string) *inhabitants.Culture {
	data := helperLoadBytes(t, fmt.Sprintf("mock_culture_%s.json", name))
	c := &inhabitants.Culture{}
	err := json.Unmarshal(data, c)
	if err != nil {
		t.Fatal(err)
	}
	return c
}
