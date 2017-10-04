package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

type namerLoad struct {
	Name        string   `json:"name"`
	Species     string   `json:"species"`
	Patronymics []string `json:"patronymics"`
	Matronymics []string `json:"matronymics"`
	FamilyNames []string `json:"family_names"`
	GenderNames []struct {
		Gender       string   `json:"gender"`
		Patterns     []string `json:"patterns"`
		GivenNames   []string `json:"given_names"`
		NameStrategy string   `json:"name_strategy"`
	} `json:"gender_names"`
}

type namerKey struct {
	name   string
	gender inhabitants.Gender
}

var namers = make(map[namerKey]*words.Namer)

func main() {
	e := echo.New()
	e.GET("/town_names", townNamesHandler)
	e.GET("/town", townHandler)
	e.GET("/being", beingHandler)
	e.GET("/household", householdHandler)
	e.GET("/random/chromosome", randomChromosomeHandler)
	e.File("/", "web")
	e.Static("/fonts", "web/fonts")
	e.Static("/styles", "web/styles")
	e.Static("/scripts", "web/scripts")
	e.Static("/data", "web/data")
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":8003"))
}

func beingHandler(c echo.Context) error {
	species := &inhabitants.Species{Name: "human"}
	being := &inhabitants.Being{
		Species: species,
	}
	return c.JSON(http.StatusOK, being)
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

func loadCulture(name string) (*inhabitants.Culture, error) {
	r, err := os.Open(fmt.Sprintf("./web/data/%s.json", name))
	if err != nil {
		return nil, fmt.Errorf("could not load internal data file")
	}
	culture := &inhabitants.Culture{}
	err = json.NewDecoder(r).Decode(&culture)
	if err != nil {
		return nil, fmt.Errorf("could not parse json file")
	}
	return culture, nil
}

func householdHandler(c echo.Context) error {
	filename := c.QueryParam("culture")
	culture, err := loadCulture(filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	expression, err := loadHuman()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
	}
	s := inhabitants.NewSpecies("human", []inhabitants.Gender{inhabitants.Male, inhabitants.Female}, expression)
	mom := &inhabitants.Being{Species: s, Sex: inhabitants.Female, Culture: culture}
	mom.RandomizeName()
	mom.RandomizeChromosome()
	mom.RandomizeAge(2)
	dad := &inhabitants.Being{Species: s, Sex: inhabitants.Male, Culture: culture}
	dad.RandomizeName()
	dad.RandomizeChromosome()
	mom.Name.FamilyName = dad.Name.FamilyName
	mom.Reproduce(dad)
	return c.JSON(http.StatusOK, []*inhabitants.Being{mom, dad})
}

func loadHuman() (*genetics.Expression, error) {
	r, err := os.Open("./web/data/human.json")
	if err != nil {
		return nil, err
	}
	expression, err := genetics.LoadExpression(r)
	if err != nil {
		return nil, err
	}
	return &expression, nil
}

func townHandler(c echo.Context) error {
	expression, err := loadHuman()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	culture, err := loadCulture(c.QueryParam("culture"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	s := inhabitants.NewSpecies("Northman", []inhabitants.Gender{inhabitants.Male, inhabitants.Female}, expression)
	area := locations.NewArea(locations.Town, nil, nil)
	count := 10000
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			being := inhabitants.Being{Species: s, Culture: culture}
			being.Randomize()
			area.Add(&being)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	return c.JSON(http.StatusOK, area)

}

func townNamesHandler(c echo.Context) error {
	count := 1000
	var wg sync.WaitGroup
	names := []string{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			area := locations.NewArea(locations.Town, nil, nil)
			names = append(names, area.Name)
		}()

	}
	wg.Wait()
	return c.JSON(http.StatusOK, names)
}
