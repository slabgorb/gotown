package main

import (
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
	"github.com/slabgorb/gotown/timeline"
	"github.com/slabgorb/gotown/words"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

const (
	mongoDBHosts = "localhost"
	mongoDBName  = "gotown"
)

type ContextWithSession struct {
	echo.Context
	session *mgo.Session
}

func main() {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  60 * time.Second,
		Database: mongoDBName,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	addSessionMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cws := &ContextWithSession{
				Context: c,
				session: mongoSession,
			}
			return next(cws)
		}
	}

	e := echo.New()
	e.GET("/cultures", listCulturesHandler)
	e.GET("/species", listSpeciesHandler)
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
	e.Use(addSessionMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Logger.SetLevel(log.DEBUG)
	e.Logger.Fatal(e.Start(":8003"))
}

func listCulturesHandler(c echo.Context) error {
	cc := c.(*ContextWithSession)
	sessionCopy := cc.session.Copy()
	collection := sessionCopy.DB(mongoDBName).C("cultures")
	cultures := []*inhabitants.Culture{}
	if err := collection.Find(bson.M{}).All(&cultures); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cultures)
}

func listSpeciesHandler(c echo.Context) error {
	cc := c.(*ContextWithSession)
	sessionCopy := cc.session.Copy()
	collection := sessionCopy.DB(mongoDBName).C("species")
	species := []*inhabitants.Species{}
	if err := collection.Find(bson.M{}).All(&species); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, species)
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
	if err = culture.Load(r); err != nil {
		return nil, err
	}
	return culture, nil
}

func householdHandler(c echo.Context) error {
	filename := c.QueryParam("culture")
	culture, err := loadCulture(filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	r, err := os.Open(fmt.Sprintf("./web/data/%s.json", "human"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
	}
	s, err := inhabitants.LoadSpecies(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
	}
	mom := &inhabitants.Being{Species: &s, Sex: inhabitants.Female, Culture: culture, Chronology: timeline.NewChronology()}
	mom.RandomizeName()
	mom.RandomizeChromosome()
	//mom.RandomizeAge(2)
	dad := &inhabitants.Being{Species: &s, Sex: inhabitants.Male, Culture: culture, Chronology: timeline.NewChronology()}
	dad.RandomizeName()
	dad.RandomizeChromosome()
	mom.Name.FamilyName = dad.Name.FamilyName
	mom.Reproduce(dad)
	return c.JSON(http.StatusOK, []*inhabitants.Being{mom, dad})
}

func townHandler(c echo.Context) error {
	culture, err := loadCulture(c.QueryParam("culture"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	r, err := os.Open(fmt.Sprintf("./web/data/%s.json", "human"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
	}
	s, err := inhabitants.LoadSpecies(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not parse json file")
	}
	area := locations.NewArea(locations.Town, nil, nil)
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			being := inhabitants.Being{Species: &s, Culture: culture, Chronology: timeline.NewChronology()}
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
