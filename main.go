package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

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
	api.GET("/cultures/:name", showCulturesHandler)
	api.GET("/species", listSpeciesHandler)
	api.GET("/species/:name", showSpeciesHandler)
	api.GET("/species/:name/expression", expressSpeciesHandler)
	api.GET("/namers", listNamersHandler)
	api.GET("/namers/:name", showNamersHandler)
	api.GET("/namers/:name/random", randomNameHandler)
	api.GET("/words", listWordsHandler)
	api.GET("/words/:name", showWordsHandler)
	api.GET("/town/name", townNameHandler)
	api.DELETE("/towns/:name", deleteAreaHandler)
	api.GET("/towns", listAreasHandler)
	api.GET("/towns/:name", showAreaHandler)
	api.POST("/towns/create", townHandler)
	//api.GET("/being/:id", showBeingHandler)
	api.PUT("/seed", seedHandler)
	//e.GET("/household", householdHandler)
	api.GET("/random/chromosome", randomChromosomeHandler)

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
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri}\t=>\t${status}\t${latency_human}\n${query} ${form}",
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

func list(c echo.Context, f func() ([]string, error)) error {
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
	return show(c, &culture.Culture{Name: c.Param("name")})
}

func listSpeciesHandler(c echo.Context) error { return list(c, species.List) }
func showSpeciesHandler(c echo.Context) error { return show(c, &species.Species{Name: c.Param("name")}) }

func expressSpeciesHandler(c echo.Context) error {
	item := &species.Species{Name: c.Param("name")}
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
func showNamersHandler(c echo.Context) error { return show(c, &words.Namer{Name: c.Param("name")}) }

func randomNameHandler(c echo.Context) error {
	n := &words.Namer{Name: c.Param("name")}
	if err := n.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, n.CreateName())
}

func listWordsHandler(c echo.Context) error { return list(c, words.WordsList) }
func showWordsHandler(c echo.Context) error { return show(c, &words.Words{Name: c.Param("name")}) }

//func showBeingHandler(c echo.Context) error { return show(c, &being.Being{ID: c.Param("id")}) }

type townHandlerRequest struct {
	Culture string `json:"culture" form:"culture" query:"culture"`
	Species string `json:"species" form:"species" query:"species"`
	Name    string `json:"name" form:"name" query:"name"`
}

func listAreasHandler(c echo.Context) error {
	all := []*locations.Area{}
	if err := persist.DB.All(&all); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	names := []string{}
	for _, t := range all {
		names = append(names, t.String())
	}
	return c.JSON(http.StatusOK, names)
}

func deleteAreaHandler(c echo.Context) error {
	var a locations.Area
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := persist.DB.One("Name", req.Name, &a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := persist.DB.DeleteStruct(&a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, struct{ success bool }{success: true})
}

func showAreaHandler(c echo.Context) error {
	var a locations.Area
	a.Name = c.Param("name")
	if err := persist.DB.One("Name", c.Param("name"), &a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, a)
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

func townHandler(c echo.Context) error {
	req := new(townHandlerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	culture := &culture.Culture{Name: req.Culture}
	if err := culture.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	species := &species.Species{Name: req.Species}
	if err := species.Read(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	area, err := locations.NewArea(locations.Town, culture, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if req.Name != "" {
		area.Name = req.Name
	}
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)
	//cron := timeline.NewChronology()
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			being := being.New(species)
			being.Randomize(culture)
			area.Add(being)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	a := *area
	if err := persist.DB.Save(&a); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, area)
}

func townNameHandler(c echo.Context) error {
	namer := words.Namer{Name: "english towns"}
	if err := namer.Read(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, namer.CreateName())
}
