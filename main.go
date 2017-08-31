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
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func main() {
	e := echo.New()
	e.GET("/town_names", townNamesHandler)
	e.GET("/town", townHandler)
	e.GET("/being", beingHandler)
	e.File("/", "web")
	e.Static("/fonts", "web/fonts")
	e.Static("/styles", "web/styles")
	e.Static("/scripts", "web/scripts")
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

func townHandler(c echo.Context) error {
	type townRequest struct {
		size int64 `json:"size" form:"size" query:"size"`
	}
	area := locations.NewArea(locations.Town, nil, nil)

	for i := 0; i < 1000; i++ {
		being := inhabitants.Being{}
		_ = being.Randomize()
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
