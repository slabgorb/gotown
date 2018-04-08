package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/inhabitants"
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

func main() {
	err := persist.Open("gotown.db")
	if err != nil {
		panic(err)
	}

	defer persist.Close()

	e := echo.New()
	api := e.Group("/api")
	api.GET("/cultures", listCulturesHandler)
	api.GET("/cultures/:name", showCulturesHandler)
	api.GET("/species", listSpeciesHandler)
	api.GET("/species/:name", showSpeciesHandler)
	api.GET("/town/name", townNameHandler)
	api.DELETE("/towns/:name", deleteAreaHandler)
	api.GET("/towns", listAreasHandler)
	api.GET("/towns/:name", showAreaHandler)
	api.POST("/towns/create", townHandler)
	api.GET("/being/:id", showBeingHandler)
	api.PUT("/seed", seedHandler)
	//e.GET("/household", householdHandler)
	api.GET("/random/chromosome", randomChromosomeHandler)
	e.Static("/fonts", "web/fonts")
	e.Static("/styles", "web/styles")
	e.Static("/scripts", "web/scripts")
	e.Static("/data", "web/data")
	e.File("/manifest.json", "web/manifest.json")
	e.File("/", "web")
	e.File("/*", "web/index.html")
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
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

func listCulturesHandler(c echo.Context) error {
	all := []*inhabitants.Culture{}
	if err := persist.DB.All(&all); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	names := []string{}
	for _, t := range all {
		names = append(names, t.String())
	}
	return c.JSON(http.StatusOK, names)
}

func showCulturesHandler(c echo.Context) error {
	var item inhabitants.Culture
	if err := persist.DB.One("Name", c.Param("name"), &item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

func listSpeciesHandler(c echo.Context) error {
	all := []*inhabitants.Species{}
	if err := persist.DB.All(&all); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	names := []string{}
	for _, t := range all {
		names = append(names, t.String())
	}
	return c.JSON(http.StatusOK, names)
}

func showSpeciesHandler(c echo.Context) error {
	var item inhabitants.Species
	if err := persist.DB.One("Name", c.Param("name"), &item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

func showBeingHandler(c echo.Context) error {
	var item inhabitants.Being
	if err := persist.DB.One("ID", c.Param("id"), &item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, item)
}

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
	var culture inhabitants.Culture
	if err := persist.DB.One("Name", req.Culture, &culture); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var species inhabitants.Species
	if err := persist.DB.One("Name", req.Species, &species); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	area, err := locations.NewArea(locations.Town, &culture, nil, nil)
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
			//being := inhabitants.Being{Species: &species, Culture: &culture, Chronology: cron}
			//being.Randomize()
			//area.Add(&being)
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
	area, err := locations.NewArea(locations.Town, nil, nil, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, area.Name)
}
