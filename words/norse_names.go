package words

var (
	NorseFemaleNameWords *Words = NewWords()
	NorseMaleNameWords   *Words = NewWords()
	NorseMaleNamer       *Namer = NewNamer(
		[]string{
			"{{.GivenName}} {{.GivenName}}{{.Patronymic}}",
		}, NorseMaleNameWords,
	)
	NorseFemaleNamer *Namer = NewNamer(
		[]string{
			"{{.GivenName}} {{.GivenName}}{{.Matronymic}}",
		}, NorseFemaleNameWords,
	)
)

func init() {
	NorseFemaleNameWords.AddList("givenNames",
		[]string{
			"Aadny", "Aafrid", "Aagunn", "Aalaug", "Aasa", "Aasbjorg", "Aase",
			"Aasfrid", "Aasgerd", "Aashild", "Aaslaug", "Aasne", "Aasrunn",
			"Aasveig ", "Agnhild", "Alfhild", "Alfrid", "Annbjorg", "Annfrid",
			"Anngjerd", "Annhild", "Annlaug", "Annveig", "Arnbjorg", "Arnfrid",
			"Arnhild", "Arnhill", "Arnlaug", "Arnveig", "Asbjorg", "Asfrid",
			"Asgjerd", "Aslaug", "Asrunn", "Asta ", "Astrid ", "Asveig", "Aud",
			"Audbjorg", "Audfrid", "Audgerd", "Audgunn", "Audhild", "Audrun",
			"Audveig", "Bergfrid", "Berghild", "Bergljot", "Bergthora", "Bergunn",
			"Bjarnhild", "Bjorg", "Bjorgfrid", "Bjorghild", "Bjorgunn",
			"Bjorgveig", "Bjornhild", "Bodhild", "Bodil", "Borghild", "Borgny",
			"Borgunn", "Bryngerd", "Brynhild", "Dagbjorg", "Dageid", "Dagfrid",
			"Daghild", "Dagmoy", "Dagny", "Dagrun", "Dagveig", "Eidbjorg",
			"Eidfrid", "Eidunn", "Eirunn", "Eivor", "Eldbjorg", "Eldfrid",
			"Eldrid", "Embla", "Erika", "Eydis", "Eyrun", "Eyvor", "Finngerd",
			"Finnlaug", "Fredbjorg", "Fredgunn", "Fredhild", "Fredlaug",
			"Fredmoy", "Freidun", "Freydis", "Frida", "Fridbjorg", "Fridgunn",
			"Fridhild", "Geirhild", "Geirunn", "Gerd", "Gjartrud", "Gro",
			"Gudbjorg", "Gudfrid", "Gudlaug", "Gudrun ", "Gudveig", "Gudvor",
			"Gunn", "Gunbjorg", "Gunda", "Gunnbjorg", "Gunnfrid", "Gunnhild",
			"Gunn ", "Gunnlaug", "Gunnveig", "Gunnvor", "Gyda", "Hallbjorg",
			"Hallfrid", "Hallgerd", "Hallgunn", "Hallrid", "Hallveig", "Heidrun",
			"Helga", "Helgunn", "Hella", "Herbjorg", "Herborg", "Herfrid",
			"Hergunn", "Hervor", "Hild ", "Hilda", "Hilde", "Hildebjorg",
			"Hildeborg", "Hildegunn", "Hill", "Hjartrud", "Hjordis", "Holmfrid",
			"Hladgerd", "Idunn", "Inga", "Ingebjorg", "Ingegerd", "Inger",
			"Inghild", "Inglaug", "Ingrid", "Ingunn", "Ingvild", "Jarlaug",
			"Jarlfrid", "Jarnhild", "Joreid", "Jorgunn", "Jorhild", "Jorfrid",
			"Jorlaug", "Jorunn", "Kjellaug", "Kjellbjorg", "Kjellfrid",
			"Kjellrun", "Kjellvor", "Kjerlaug", "Knuthild", "Kollbjorg",
			"Lagertha", "Lidveig", "Lidvor", "Liv", "Liveig", "Livunn", "Lydveig",
			"Lydvor", "Magnhild", "Magnvor", "Malfrid", "Malmfrid", "Mildfrid",
			"Mildrun", "Modgunn", "Modhild", "Moyfrid", "Norbjorg", "Norfrid",
			"Norgunn", "Norhild", "Norlaug", "Norveig", "Oda", "Oddbjorg",
			"Oddfrid", "Oddgunn", "Oddlaug", "Oddrun", "Oddrunn", "Oddveig",
			"Odfrid", "Odlaug", "Odveig", "Olbjorg", "Oleiv", "Olrun", "Olveig",
			"Oslaug", "Oyfrid", "Oygunn", "Oylaug", "Oyonn", "Oyvor", "Ragna",
			"Ragnfrid", "Ragnhild", "Rambjorg", "Ramfrid", "Randi", "Ranfrid",
			"Rannveig", "Ranveig", "Reidhild", "Reidun", "Reinhild", "Runa",
			"Runhild", "Sabjorg", "Salmoy", "Sebjorg", "Sif", "Sigbjorg",
			"Sigfrid", "Signe ", "Signhild", "Signy", "Sigrid", "Sigrunn",
			"Sigurda", "Sigvalda", "Sigvarda", "Sigveig", "Sigvor", "Siv",
			"Skjaldvor", "Snefrid", "Solaug", "Solbjorg", "Solfrid", "Solgerd",
			"Solgunn", "Solhild", "Sollaug", "Solmoy", "Solrunn", "Solunn",
			"Solveig", "Solvor", "Steinunn", "Steinvor", "Svanbjorg", "Svanfrid",
			"Svanhild", "Svanlaug", "Thora", "Thorbjorg", "Thorborg", "Thordis",
			"Thordny", "Thordun", "Thorfrid", "Torgerd", "Thorgunn", "Thorhild",
			"Torgerd", "Thorunn", "Thorvalda", "Thurid", "Tjodhild", "Tjodunn",
			"Tjodvor", "Tova", "Tove", "Trude", "Trudi", "Tryghild", "Turid",
			"Tuva", "Udbjorg", "Ulvhild", "Unn", "Unnfrid", "Unnhild", "Unni",
			"Unnlaug", "Unnveig", "Urda", "Valbjorg", "Valfrid", "Valgerd",
			"Varunn", "Vebjorg", "Velaug", "Venhild", "Vidrun", "Vigdis",
			"Vighild", "Vilbjorg", "Vilborg", "Vilfrid", "Vilgerd", "Vilgunn",
			"Ymbjorg", "Yngvild",
		},
	)
	NorseFemaleNameWords.AddList("matronymics",
		[]string{"dottir"},
	)
	NorseMaleNameWords.AddList("givenNames",
		[]string{
			"Aasbjorn", "Aasgeir", "Aasgrim", "Aaskell", "Aaskjell",
			"Aasleik", "Aasleif", "Aasleiv", "Aasvald", "Aasvard", "Ag",
			"Aggrim", "Agmund", "Agnar", "Aake", "Almar", "Alvald",
			"Alvar", "Alvbjørn", "Alvfinn", "Alvgaut", "Alvgeir", "Arn",
			"Arnbjorn", "Arnfinn", "Arnfred", "Arngeir", "Arngrim", "Arni",
			"Arnkjell", "Arnleif", "Arnleiv", "Arnljot", "Arnmod",
			"Arnstein", "Arnthor", "Arnulf", "Arnvald", "Arnvid", "Atli",
			"Audar", "Audbjorn", "Audfinn", "Audgrim", "Audkjell",
			"Audmund", "Audstein", "Audulv", "Audun", "Audvald", "Bard",
			"Bendik", "Berg", "Bergsvein", "Berulf", "Berulv", "Besse",
			"Birgir", "Bjarni", "Bjermund", "Bjorgulv", "Bjorn", "Bjornar",
			"Bjornulv", "Bodolf", "Bogi", "Boldolv", "Bodvar", "Borgi–Borgir",
			"Borgulv", "Borgvald", "Borri", "Botolv", "Bredi",
			"Bragi", "Brusi", "Brynjulf", "Dagfinn", "Dreng", "Dyri",
			"Egil", "Eigil", "Eilif", "Eiliv", "Eimund", "Einar",
			"Einvald", "Eldar", "Eric", "Erik", "Erland", "Erlend",
			"Erling", "Eystein", "Eyyolf", "Eyyolv", "Fartein", "Faste",
			"Finn", "Finnbjorn", "Finngard", "Folki", "Folkvald", "Freyr",
			"Fridleiv", "Fridthjof", "Frodi", "Gauti", "Geir",
			"Geirbrand", "Geirmund", "Geirolf", "Geirolv", "Geirstein",
			"Geirulf", "Geirulv", "Gisli", "Gissur", "Glumir", "Gorm",
			"Grim", "Grimkjell", "Gudbjorn", "Gudleik", "Gudleiv",
			"Gudmund", "Gudolf", "Gudolv", "Gudstein", "Gullbjorn",
			"Gullbrand", "Gullik", "Gunbjorn", "Gunnar", "Gunnbjorn",
			"Gunner", "Gunnolf", "Gunnolv", "Gunnsten", "Gunnstein",
			"Gunnvald", "Gunnvar", "Gutorm", "Guttorm", "Haabjorn",
			"Hagbart", "Haagen", "Haaken", "Haakon", "Haarek", "Haarik",
			"Haavald", "Haavard", "Haavid", "Halfdan", "Hallbjorn",
			"Hallfred", "Hallgeir", "Hallgrim", "Halli", "Hallkjell",
			"Hallstein", "Hallthor", "Hallvard", "Hallvor", "Harald", "Hauk",
			"Heming", "Hemming", "Herbjorn", "Herbrad", "Herleif",
			"Herleik", "Herleiv", "Hermod", "Hermund", "Hervard",
			"Hildebrand", "Hildemar", "Hjalmar", "Hogni", "Holgir",
			"Ingar", "Ingebjorn", "Ingemar", "Ingolf", "Ingvald", "Ingvar",
			"Ivar", "Iver", "Joar", "Jogeir", "Jor", "Jorik", "Jorulv",
			"Jorund", "Jostein", "Kaalv", "Kaarbjorn", "Kaare", "Ketil",
			"Kjetil", "Kleng", "Knut", "Kol", "Kolbein", "Kolbjorn",
			"Kolfinn", "Kolgrim", "Koll", "Kolstein", "Leid", "Leidolv",
			"Leidulf", "Leidulv", "Leidvor", "Leif", "Leiv", "Magnar",
			"Magni", "Narfi", "Narvi", "Nottolv", "Olaf", "Olav", "Odd",
			"Oddbjorn", "Oddfred", "Oddgeir", "Oddleif", "Oddleiv",
			"Oddmar", "Oddmund", "Oddvar", "Odin", "Olbjorn", "Olgir",
			"Olgeir", "Olvar", "Ottar", "Raadgeir", "Raadmund", "Raadolf",
			"Raadolv", "Ragnar", "Ragnvald", "Randmod", "Randolf",
			"Randolv", "Reidalv", "Reidar", "Reidmar", "Reidolf",
			"Reidolv", "Reidulf", "Reidulv", "Roald", "Roar", "Rolf",
			"Rolleif", "Rolleiv", "Runolv", "Rorek", "Rorik", "Runi",
			"Sigbjorn", "Sigfred", "Sigstein", "Sigtrygg", "Sigurd",
			"Sigvald", "Sigvat", "Sjaundi", "Skjalg", "Skuli", "Snorri",
			"Solmund", "Stein", "Steinar", "Steinbjorn", "Steingrim",
			"Stig", "Styrbjorn", "Svein", "Sveinar", "Sveinbjorn",
			"Sveinung", "Sverre", "Sverri", "Saebjorn", "Saemund",
			"Saevald", "Tali", "Tarjei", "Tjodgeir", "Tjodolf", "Tjoldolv",
			"Tjodrek", "Tjodrik", "Toki", "Thor", "Thoralf", "Thoralv",
			"Thorbjorn", "Thord", "Thorfinn", "Thorgeir", "Thorgrim",
			"Thorkel", "Thorkell", "Thorleik", "Thorleiv", "Thormod",
			"Thorodd", "Thorolf", "Thorolv", "Thorstein", "Thorsten",
			"Thorvald", "Trond", "Tryggvi", "Ulrik", "Ulf", "Ulv",
			"Vebjorn", "Vegard", "Vegeir", "Vemund", "Vidar", "Vidkunn",
			"Vigbjorn", "Volund", "Yngvi",
		})

	NorseMaleNameWords.AddList("patronymics",
		[]string{"son"},
	)
}