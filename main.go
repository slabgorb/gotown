package main

import (
	"bytes"
	"fmt"
	"image/png"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/heraldry"
	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/logger"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	file, _ := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	logger.SetOutput(file)
}

type response interface {
	persist.Persistable
	API() (interface{}, error)
}

func definePprofHandlers(e *echo.Echo) {
	e.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))
}

func defineAPIHandlers(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/populations", listPopulationHandler)
	api.GET("/populations/:id", showPopulationHandler)
	api.GET("/cultures", listCulturesHandler)
	api.GET("/cultures/:id", showCulturesHandler)
	api.GET("/species", listSpeciesHandler)
	api.GET("/species/:id/expression", expressSpeciesHandler)
	api.GET("/species/:id", showSpeciesHandler)
	api.GET("/namers", listNamersHandler)
	api.GET("/namers/:id", showNamersHandler)
	api.GET("/namers/:id/random", randomNameHandler)
	api.GET("/words", listWordsHandler)
	api.GET("/words/:id", showWordsHandler)
	api.GET("/town/name", townNameHandler)
	api.DELETE("/towns/:id", deleteAreaHandler)
	api.GET("/towns", listAreasHandler)
	api.GET("/towns/:id", showAreaHandler)
	api.POST("/towns/create", createTownHandler)
	api.GET("/beings/:id", showBeingHandler)
	api.GET("/beings", listBeingsHandler)
	api.PUT("/seed", seedHandler)
	//e.GET("/household", householdHandler)
	api.GET("/random/chromosome", randomChromosomeHandler)
	api.GET("/random/heraldry", randomHeraldryHandler)

}

func defineStaticHandlers(e *echo.Echo) {
	e.Static("/fonts", "/docroot/fonts")
	e.Static("/styles", "/docroot/styles")
	e.Static("/scripts", "/docroot/scripts")
	e.Static("/data", "/docroot/data")
	e.File("/manifest.json", "/docroot/manifest.json")
	e.File("/", "/docroot")
}

func main() {
	err := persist.Open()
	if err != nil {
		panic(err)
	}
	defer persist.Close()

	e := echo.New()
	defineAPIHandlers(e)
	defineStaticHandlers(e)

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 6 << 10,
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${remote_ip} ${method} ${uri}\t=>\t${status}\t${latency_human}\n${query} ${form} ",
	}))
	e.Logger.SetLevel(log.DEBUG)
	e.HTTPErrorHandler = customHTTPErrorHandler(e.DefaultHTTPErrorHandler)
	definePprofHandlers(e)
	e.Start(":8003")
}

func customHTTPErrorHandler(f echo.HTTPErrorHandler) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		c.Logger().Error(err)
		f(err, c)
	}
}

func seedHandler(c echo.Context) error {
	persist.DeleteAll()
	species.Seed()
	culture.Seed()
	words.Seed()
	return nil
}

func getID(c echo.Context) string {
	return c.Param("id")
}

func list(c echo.Context, f func() (map[string]string, error)) error {
	names, err := f()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	type pair struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	pairs := []pair{}

	for k, v := range names {
		pairs = append(pairs, pair{ID: k, Name: v})
	}
	return c.JSON(http.StatusOK, pairs)
}

func show(c echo.Context, item response) error {
	if err := item.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	api, err := item.API()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, api)
}

func listCulturesHandler(c echo.Context) error {
	return list(c, culture.List)
}

func showCulturesHandler(c echo.Context) error {
	cu := &culture.Culture{}
	cu.SetID(getID(c))
	return show(c, cu)
}

func listSpeciesHandler(c echo.Context) error {
	return list(c, species.List)
}

func showSpeciesHandler(c echo.Context) error {
	s := &species.Species{}
	s.SetID(getID(c))
	return show(c, s)
}

func listPopulationHandler(c echo.Context) error {
	pops := []*being.Population{}
	list, err := persist.List("Population")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	for _, v := range list {
		p := &being.Population{}
		p.SetID(v)
		err := p.Read()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		pops = append(pops, p)
	}
	apis := []interface{}{}
	for _, p := range pops {
		api, err := p.API()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		apis = append(apis, api)
	}
	return c.JSON(http.StatusOK, apis)
}

func showPopulationHandler(c echo.Context) error {
	p := &being.Population{}
	p.SetID(getID(c))
	return show(c, p)
}

func expressSpeciesHandler(c echo.Context) error {
	item := &species.Species{}
	item.SetID(getID(c))
	if err := item.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	geneString := strings.Split(c.Param("genes"), "|")
	genes := []genetics.Gene{}
	for _, s := range geneString {
		genes = append(genes, genetics.Gene(s))
	}

	chromosome := genetics.Chromosome{Genes: genes}
	e := chromosome.Express(item.GeneticExpression)
	return c.JSON(http.StatusOK, e)
}

func listNamersHandler(c echo.Context) error { return list(c, words.NamerList) }
func showNamersHandler(c echo.Context) error {
	n := &words.Namer{}
	n.SetID(getID(c))
	return show(c, n)
}

func randomNameHandler(c echo.Context) error {
	n := &words.Namer{}
	n.SetID(getID(c))
	if err := n.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, n.CreateName())
}

func listWordsHandler(c echo.Context) error { return list(c, words.WordsList) }
func showWordsHandler(c echo.Context) error {
	w := &words.Words{}
	w.SetID(getID(c))
	return show(c, w)
}

func showBeingHandler(c echo.Context) error {
	b := &being.Being{}
	b.SetID(getID(c))
	return show(c, b)
}
func listBeingsHandler(c echo.Context) error { return list(c, being.List) }

type lister interface {
	GetId()
	String()
}

type listItem struct {
	S     string `json:"name"`
	ID    string `json:"id"`
	Icon  string `json:"icon"`
	Image string `json:"image"`
}

func listAreasHandler(c echo.Context) error {
	logger.TimeSet()
	list, err := persist.List("Area")
	logger.TimeElapsed("loading area list")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	names := []interface{}{}
	for k, v := range list {
		logger.TimeSet()
		a := &locations.Area{}
		a.SetID(k)
		logger.TimeElapsed(fmt.Sprintf("getting details for list %s", v))
		if err := a.Read(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		logger.TimeSet()
		t, err := a.ListItemAPI()
		logger.TimeElapsed(fmt.Sprintf("api for %s", v))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		names = append(names, t)
	}
	return c.JSON(http.StatusOK, names)
}

func deleteAreaHandler(c echo.Context) error {
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	a := &locations.Area{}
	a.SetID(req.ID)

	if err := a.Delete(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, struct{ success bool }{success: true})
}

func showAreaHandler(c echo.Context) error {
	a := &locations.Area{}
	a.SetID(getID(c))
	if err := a.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	api, err := a.API()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not load api for area %d: %s", a.ID, err))
	}
	return c.JSON(http.StatusOK, api)
}

func randomChromosomeHandler(c echo.Context) error {
	count, err := strconv.Atoi(c.QueryParam("count"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	chromosome := genetics.RandomChromosome(count)
	return c.JSON(http.StatusOK, chromosome)
}

func renameHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{}{})
}

// func householdHandler(c echo.Context) error {
// 	filename := c.QueryParam("culture")
// 	culture, err := loadCulture(filename)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err)
// 	}
// 	r, err := os.Open(fmt.Sprintf("./docroot/data/%s.json", "human"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
// 	}
// 	s, err := inhabitants.LoadSpecies(r)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
// 	}
// 	mom := &inhabitants.Being{Species: &s, Sex: inhabitants.Female, Culture: culture, Chronology: timeline.NewChronology()}
// 	mom.RandomizeName()
// 	mom.RandomizeChromosome()
// 	//mom.RandomizeAge(2)
// 	dad := &inhabitants.Being{Species: &s, Sex: inhabitants.Male, Culture: culture, Chronology: timeline.NewChronology()}
// 	dad.RandomizeName()
// 	dad.RandomizeChromosome()
// 	mom.Name.FamilyName = dad.Name.FamilyName
// 	mom.Reproduce(dad)
// 	return c.JSON(http.StatusOK, []*inhabitants.Being{mom, dad})
// }

type townHandlerRequest struct {
	Culture string `json:"culture" form:"culture" query:"culture"`
	Species string `json:"species" form:"species" query:"species"`
	Name    string `json:"name" form:"name" query:"name"`
	ID      string `json:"id" form:"id" query:"id"`
	Size    int    `json:"size" form:"size" query:"size"`
}

func createTownHandler(c echo.Context) error {
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	cl := &culture.Culture{Name: req.Culture}
	if err := persist.ReadByName(cl.Name, "Culture", cl); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not load culture %s: %s", req.Culture, err))
	}
	s := &species.Species{Name: req.Species}
	if err := persist.ReadByName(s.Name, "Species", s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not load species %s: %s", req.Species, err))
	}

	namer, err := getDefaultNamer()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := namer.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not load namer %s: %s", namer.Name, err))
	}

	area := locations.NewArea(locations.Town, nil, namer)
	logger.TimeSet()
	if req.Name != "" {
		area.Name = req.Name
	}
	h := heraldry.RandomEscutcheon("square", true)
	area.Heraldry = heraldry.New(h)
	var wg sync.WaitGroup
	size := locations.AreaSize(req.Size)
	count := size.Population()
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()
			being := being.New(s, cl, logger.Default)
			being.Randomize()
			being.Save()
			area.Add(being)
		}(&wg)
	}
	wg.Wait()
	logger.TimeElapsed("creation")
	logger.TimeSet()
	for i := 0; i < 10; i++ {
		if err := area.Residents.Age(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed aging population in year  %d: %s", i+1, err))
		}
		mcs, err := area.Residents.MaritalCandidates(cl)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed getting marital candidates year  %d: %s", i+1, err))
		}
		for _, mc := range mcs {
			a, b := mc.Pair()
			a.Marry(b)
		}
		// rcs := area.Residents.ReproductionCandidates()
		// for _, rc := range rcs {
		// 	rc
		// }
	}
	logger.TimeElapsed("aging")
	if err := area.Save(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not save created area: %s", err))
	}
	id := area.ID
	area.Reset()
	area.ID = id
	if err := area.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not read area %d: %s", area.ID, err))
	}
	api, err := area.API()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not create area api for %d: %s", area.ID, err))
	}
	return c.JSON(http.StatusOK, api)
}

func getDefaultNamer() (*words.Namer, error) {
	list, err := persist.List("Namer")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	namerID := ""
	for k, v := range list {
		if v == "english towns" {
			namerID = k
			break
		}
	}

	namer := &words.Namer{}
	namer.SetID(namerID)
	return namer, nil
}

func townNameHandler(c echo.Context) error {

	namer, err := getDefaultNamer()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := namer.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, namer.CreateName())
}

func randomHeraldryHandler(c echo.Context) error {
	e := heraldry.RandomEscutcheon("square", true)
	m := e.Render()
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, m); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Blob(http.StatusOK, "image/png", buffer.Bytes())
}
