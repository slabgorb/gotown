package main

import (
	"bytes"
	"fmt"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/slabgorb/gotown/heraldry"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/inhabitants/being"
	"github.com/slabgorb/gotown/inhabitants/culture"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/inhabitants/species"
	"github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/persist"
	"github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func defineAPIHandlers(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/cultures", listCulturesHandler)
	api.GET("/cultures/:id", showCulturesHandler)
	api.GET("/species", listSpeciesHandler)
	api.GET("/species/:id", showSpeciesHandler)
	api.GET("/species/:name/expression", expressSpeciesHandler)
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
	//api.GET("/being/:id", showBeingHandler)
	api.PUT("/seed", seedHandler)
	//e.GET("/household", householdHandler)
	api.GET("/random/chromosome", randomChromosomeHandler)
	api.GET("/random/heraldry", randomHeraldryHandler)

}

func defineStaticHandlers(e *echo.Echo) {
	e.Static("/fonts", "web/fonts")
	e.Static("/styles", "web/styles")
	e.Static("/scripts", "web/scripts")
	e.Static("/data", "web/data")
	e.File("/manifest.json", "web/manifest.json")
	e.File("/", "web")
	e.File("/*", "web/index.html")
}

func main() {
	err := persist.Open("gotown.db")
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
	e.Start(":8003")
}

func seedHandler(c echo.Context) error {
	species.Seed()
	culture.Seed()
	words.Seed()
	return nil
}

func getID(c echo.Context) int {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0
	}
	return id
}

func list(c echo.Context, f func() ([]persist.IDPair, error)) error {
	names, err := f()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, names)
}

func show(c echo.Context, item persist.Persistable) error {
	if err := item.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

func listCulturesHandler(c echo.Context) error { return list(c, culture.List) }
func showCulturesHandler(c echo.Context) error {
	return show(c, &culture.Culture{ID: getID(c), Name: c.Param("id")})
}

func listSpeciesHandler(c echo.Context) error { return list(c, species.List) }
func showSpeciesHandler(c echo.Context) error {
	return show(c, &species.Species{ID: getID(c), Name: c.Param("id")})
}

func expressSpeciesHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not convert %s into a valid int: %s", c.Param("id"), err))
	}
	item := &species.Species{ID: id}
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
	return show(c, &words.Namer{ID: getID(c), Name: c.Param("id")})
}

func randomNameHandler(c echo.Context) error {
	n := &words.Namer{Name: c.Param("name")}
	if err := n.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, n.CreateName())
}

func listWordsHandler(c echo.Context) error { return list(c, words.WordsList) }
func showWordsHandler(c echo.Context) error {
	return show(c, &words.Words{ID: getID(c), Name: c.Param("id")})
}

//func showBeingHandler(c echo.Context) error { return show(c, &being.Being{ID: c.Param("id")}) }

type townHandlerRequest struct {
	Culture string `json:"culture" form:"culture" query:"culture"`
	Species string `json:"species" form:"species" query:"species"`
	Name    string `json:"name" form:"name" query:"name"`
	ID      int    `json:"id" form:"id" query:"id"`
	Count   int    `json:"count" form:"count" query:"count"`
}

type lister interface {
	GetId()
	String()
}

type listItem struct {
	S  string `json:"name"`
	ID int    `json:"id"`
}

func listAreasHandler(c echo.Context) error {
	all := []*locations.Area{}
	if err := persist.DB.All(&all); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	names := []listItem{}
	for _, t := range all {
		names = append(names, listItem{S: t.Name, ID: t.ID})
	}
	return c.JSON(http.StatusOK, names)
}

func deleteAreaHandler(c echo.Context) error {
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	a := &locations.Area{ID: req.ID}
	if err := a.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := a.Delete(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, struct{ success bool }{success: true})
}

func showAreaHandler(c echo.Context) error {
	a := &locations.Area{ID: getID(c), Name: c.Param("id")}
	if err := a.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	api, err := a.API()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not load api for area %d", a.ID))
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
// 	r, err := os.Open(fmt.Sprintf("./web/data/%s.json", "human"))
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

func createTownHandler(c echo.Context) error {
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	cl := &culture.Culture{Name: req.Culture}
	if err := cl.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not load culture %s: %s", req.Culture, err))
	}
	s := &species.Species{Name: req.Species}
	if err := s.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("could not load species %s: %s", req.Species, err))
	}

	namer := &words.Namer{Name: "english towns"}
	if err := namer.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not load namer %s: %s", namer.Name, err))
	}

	area := locations.NewArea(locations.Town, nil, namer)

	if req.Name != "" {
		area.Name = req.Name
	}
	var wg sync.WaitGroup
	for i := 0; i < req.Count; i++ {
		go func(wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()
			being := being.New(s, cl)
			being.Randomize()
			being.Save()
			area.Add(being)
		}(&wg)
	}
	wg.Wait()
	if err := area.Save(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not save created area: %s", err))
	}
	api, err := area.API()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("could not create area api for %d", area.ID))
	}
	return c.JSON(http.StatusOK, api)
}

func townNameHandler(c echo.Context) error {
	namer := words.Namer{Name: "english towns"}
	if err := namer.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, namer.CreateName())
}

func randomHeraldryHandler(c echo.Context) error {
	e := heraldry.RandomEscutcheon("square")
	m := e.Render()
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, m); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Blob(http.StatusOK, "image/png", buffer.Bytes())
}
