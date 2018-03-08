package words_test

import (
	"encoding/json"
	"testing"

	"github.com/slabgorb/gotown/words"
)

func TestMarshal(t *testing.T) {
	w := words.NewWords()
	w.AddList("firstName", []string{"Bob", "Joe"})
	n := words.NewNamer([]string{"firstName"}, w, "patronymic")
	bytes, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"patterns":["firstName"],"words":{"dictionary":{"firstName":["Bob","Joe"]}},"name_strategy":"patronymic"}`
	if string(bytes) != expected {
		t.Errorf("expected %s got %s", expected, string(bytes))
	}
	n = &words.Namer{}
	err = json.Unmarshal(bytes, n)
	if err != nil {
		t.Fatal(err)
	}
	if n.Patterns[0] != "firstName" {
		t.Fail()
	}

}
