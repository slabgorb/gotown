package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/locations"
	"github.com/slabgorb/gotown/words"
)

func init() {
	rand.Seed(time.Now().UnixNano())

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

func renameHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{}{})
}

func townHandler(c echo.Context) error {
	// type townRequest struct {
	// 	size int `form:"size" query:"size"`
	// }
	// tr := townRequest{}
	// if err := c.Bind(tr); err != nil {
	// 	return err
	// }
	female := inhabitants.NewSpeciesGender(words.NorseFemaleNamer, inhabitants.Matronymic, 12, 48)
	male := inhabitants.NewSpeciesGender(words.NorseMaleNamer, inhabitants.Patronymic, 12, 65)
	s := inhabitants.NewSpecies("Northman", map[inhabitants.Gender]*inhabitants.SpeciesGender{
		inhabitants.Female: female,
		inhabitants.Male:   male,
	})
	area := locations.NewArea(locations.Town, nil, nil)
	for i := 0; i < 1000; i++ {
		being := inhabitants.Being{Species: s}
		being.Randomize()
		area.Add(&being)
	}

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
