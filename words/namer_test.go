package words_test

// func TestMarshal(t *testing.T) {
// 	w := words.NewWords("base")
// 	w.AddList("firstName", []string{"Bob", "Joe"})
// 	w.Save()
// 	n := words.NewNamer([]string{"firstName"}, w.Name, "patronymic")
// 	bytes, err := json.Marshal(n)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	n = &words.Namer{}
// 	err = json.Unmarshal(bytes, n)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if n.Patterns[0] != "firstName" {
// 		t.Fail()
// 	}

// }
