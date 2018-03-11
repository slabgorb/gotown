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

	bolt "github.com/coreos/bbolt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/timeline"
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

type contextWithSession struct {
	echo.Context
	session *bolt.DB
}

const (
	cultureBucket = "cultures"
	speciesBucket = "species"
)

type fetchable interface {
	fetch(key string, db *bolt.DB) error
}

func listBucketKeys(bucket string, db *bolt.DB) []string {
	names := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			names = append(names, string(k))
		}
		return nil
	})
	return names
}

func doFetch(j interface{}, bucket string, key string, db *bolt.DB) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		err := json.Unmarshal(v, j)
		if err != nil {
			return fmt.Errorf("Error unmarshaling %s %s", bucket, key)
		}
		return nil
	})
}

type fetchableCulture struct{ culture *inhabitants.Culture }

func newFetchableCulture() fetchableCulture {
	return fetchableCulture{culture: &inhabitants.Culture{}}
}

func (f fetchableCulture) fetch(key string, db *bolt.DB) error {
	return doFetch(f.culture, cultureBucket, key, db)
}

type fetchableSpecies struct{ species *inhabitants.Species }

func newFetchableSpecies() fetchableSpecies {
	return fetchableSpecies{species: &inhabitants.Species{}}
}

func (f fetchableSpecies) fetch(key string, db *bolt.DB) error {
	return doFetch(f.species, speciesBucket, key, db)
}

func main() {

	session, err := bolt.Open("gotown.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}

	defer session.Close()

	addSessionMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cws := &contextWithSession{
				Context: c,
				session: session,
			}
			return next(cws)
		}
	}

	e := echo.New()
	api := e.Group("/api")
	api.GET("/cultures", listCulturesHandler)
	api.GET("/cultures/:name", showCulturesHandler)
	api.GET("/species", listSpeciesHandler)
	api.GET("/species/:name", showSpeciesHandler)
	api.GET("/town/name", townNameHandler)
	api.POST("/towns/create", townHandler)
	api.GET("/being", beingHandler)
	//e.GET("/household", householdHandler)
	api.GET("/random/chromosome", randomChromosomeHandler)
	e.Static("/fonts", "web/fonts")
	e.Static("/styles", "web/styles")
	e.Static("/scripts", "web/scripts")
	e.Static("/data", "web/data")
	e.File("/manifest.json", "web/manifest.json")
	e.File("/", "web")
	e.File("/*", "web/index.html")
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
	cc := c.(*contextWithSession)
	return c.JSON(http.StatusOK, listBucketKeys(cultureBucket, cc.session))
}

func showCulturesHandler(c echo.Context) error {
	cc := c.(*contextWithSession)
	culture := &inhabitants.Culture{}
	err := cc.session.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cultureBucket))
		v := b.Get([]byte(c.Param("name")))
		return json.Unmarshal(v, culture)
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Echo().Logger.Debug(fmt.Sprintf("%#v", culture))
	return c.JSON(http.StatusOK, culture)
}

func listSpeciesHandler(c echo.Context) error {
	cc := c.(*contextWithSession)
	return c.JSON(http.StatusOK, listBucketKeys(speciesBucket, cc.session))
}

func showSpeciesHandler(c echo.Context) error {
	cc := c.(*contextWithSession)
	fs := newFetchableSpecies()
	err := fs.fetch(c.Param("name"), cc.session)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, fs.species)
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

type townHandlerRequest struct {
	Culture string `json:"culture" form:"culture" query:"culture"`
	Species string `json:"species" form:"species" query:"species"`
	Name    string `json:"name" form:"name" query:"name"`
}

func townHandler(c echo.Context) error {
	c.Logger().Debug("Ok in handler")
	cc := c.(*contextWithSession)
	req := new(townHandlerRequest)
	if err := cc.Bind(req); err != nil {
		c.Logger().Debug("Error in binding")
		c.Logger().Debug(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	c.Logger().Debug(req)
	fc := fetchableCulture{culture: &inhabitants.Culture{}}
	err := fc.fetch(req.Culture, cc.session)
	if err != nil || fc.culture == nil {
		c.Logger().Debug("error fetching culture")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	culture := fc.culture
	fs := fetchableSpecies{species: &inhabitants.Species{}}
	err = fs.fetch(req.Species, cc.session)
	if err != nil || fs.species == nil {
		c.Logger().Debug("error fetching species")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	species := fs.species
	c.Logger().Debug(culture)
	c.Logger().Debug(species)
	area := locations.NewArea(locations.Town, nil, nil)
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)
	c.Logger().Debug("got past prep, working on making")
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			being := inhabitants.Being{Species: species, Culture: culture, Chronology: timeline.NewChronology()}
			if err := being.Randomize(); err != nil {
				c.Logger().Debug("error in randomizing")
				c.Logger().Debug(err)
			}
			c.Logger().Debugf("%s", being.Name.Display)

			area.Add(&being)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	c.Logger().Debug(area)
	return c.JSON(http.StatusOK, area)
}

func townNameHandler(c echo.Context) error {
	area := locations.NewArea(locations.Town, nil, nil)
	return c.JSON(http.StatusOK, area.Name)
}
