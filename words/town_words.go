package words

var (
	TownWords *Words
	TownNamer *Namer
)

func init() {
	TownNamer = NewNamer(
		[]string{"{{.Prefix}}{{.Suffix}}",
			"{{.Prefix}} {{.Noun}}{{.Suffix}}",
			"{{.ShortAdjective}}{{.ShortNoun}}{{.Suffix}}",
			"{{.Prefix}} {{.ShortAdjective}}{{.ShortNoun}}",
			"{{.Noun}}{{.Suffix}}",
			"{{.Noun}}{{.Suffix}}",
			"{{.Adjective}}{{.Suffix}}",
			"{{.Adjective}}{{.Suffix}}",
			"{{.ShortAdjective}}{{.Suffix}}",
			"{{.ShortAdjective}}{{.Suffix}}",
			"{{.Adjective}} {{.EndNoun}}",
			"{{.Noun}} {{.EndNoun}}",
			"{{.Adjective}} {{.EndNoun}}"},
		TownWords,
	)
	TownWords = NewWords()
	TownWords.Backup = BaseWords
	TownWords.AddList(
		"prefixes",
		[]string{
			"south",
			"north",
			"east",
			"west",
			"south",
			"north",
			"east",
			"west",
			"upper",
			"lower",
			"old",
			"new",
			"northeast",
			"northwest",
			"southeast",
			"southwest"},
	)
	TownWords.AddList(
		"suffixes",
		[]string{
			"borough",
			"bridge",
			"burg",
			"burn",
			"cross",
			"dale",
			"end",
			"ey",
			"field",
			"ford",
			"gate",
			"green",
			"ham",
			"harbor",
			"hill",
			"hold",
			"ing",
			"ingley",
			"ington",
			"land",
			"lea",
			"leagh",
			"lin",
			"moor",
			"more",
			"port",
			"river",
			"stone",
			"sty",
			"thorpe",
			"ton",
			"ton",
			"town",
			"town",
			"ville",
			"ville",
			"wick",
			"wood",
			"worth",
			"yard",
		},
	)
	TownWords.AddList(
		"endNoun",
		[]string{
			"crossing",
			"field",
			"bend",
			"road",
			"town",
			"town",
			"city",
			"city",
			"green",
			"yard",
			"head",
			"harbor",
			"port",
		},
	)
}
