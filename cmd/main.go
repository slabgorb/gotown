package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"
	"text/tabwriter"

	"github.com/abiosoft/ishell"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/inhabitants/genetics"
	"github.com/slabgorb/gotown/locations"
)

var currentCulture *inhabitants.Culture
var currentSpecies *inhabitants.Species
var currentArea *locations.Area

func main() {
	shell := ishell.New()
	shell.AddCmd(&ishell.Cmd{
		Name: "load_culture",
		Help: "load a culture",
		Func: loadCulture,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "load_species",
		Help: "load a species",
		Func: loadSpecies,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "status",
		Help: "display current status",
		Func: status,
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "create_town",
		Help: "make a town",
		Func: createTown,
	})
	shell.SetHomeHistoryPath(".ishell_history")
	shell.Run()
}

func status(c *ishell.Context) {
	c.Println(currentArea)
	c.Println(currentCulture)
}

func loadSpecies(c *ishell.Context) {
	name := c.Args[0]
	r, err := os.Open("../web/data/human.json")
	if err != nil {
		c.Println(err)
	}
	expression, err := genetics.LoadExpression(r)
	if err != nil {
		c.Println(err)
	}
	currentSpecies = &inhabitants.Species{
		Name:       name,
		Expression: &expression,
		Genders:    []inhabitants.Gender{inhabitants.Male, inhabitants.Female},
		Demography: inhabitants.Demographies[name],
	}
	c.Println(currentSpecies)
}

func loadCulture(c *ishell.Context) {
	name := c.Args[0]
	r, err := os.Open(fmt.Sprintf("../web/data/%s.json", name))
	if err != nil {
		c.Println(err)
		return
	}
	culture := &inhabitants.Culture{}
	if err = culture.Load(r); err != nil {
		c.Println(err)
	}
	currentCulture = culture
	c.Println(currentCulture)
}

func createTown(c *ishell.Context) {
	if currentCulture == nil || currentSpecies == nil {
		c.Println("Please load a species and a culture before creating a town")
		return
	}
	if len(c.Args) == 0 {
		c.Println("Please specify a population")
		return
	}
	pop, err := strconv.Atoi(c.Args[0])
	if err != nil {
		c.Println(err)
		return
	}
	area := locations.NewArea(locations.Town, nil, nil)
	var wg sync.WaitGroup
	wg.Add(pop)
	for i := 0; i < pop; i++ {
		go func(wg *sync.WaitGroup) {
			being := inhabitants.Being{Species: currentSpecies, Culture: currentCulture}
			being.Randomize()
			area.Add(&being)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	currentArea = area
	output := bytes.Buffer{}
	w := tabwriter.NewWriter(&output, 1, 1, 1, ' ', tabwriter.Debug)
	c.Println(currentArea.Name)
	beings := currentArea.Residents.Beings()
	for i := 0; i < len(beings); i++ {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%d", beings[i].String(), beings[i].Age))
	}
	w.Flush()
	c.Print(output.String())
}
