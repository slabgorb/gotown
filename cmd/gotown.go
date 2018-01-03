package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/abiosoft/ishell"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/locations"
	mgo "gopkg.in/mgo.v2"
)

var currentCulture *inhabitants.Culture
var currentSpecies *inhabitants.Species
var currentArea *locations.Area

func main() {
	ctx := context.Background()
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	shell := ishell.New()
	commandList := []*ishell.Cmd{
		loadSpeciesCommand(ctx, session),
		loadCultureCommand(ctx, session),
		tickCommand(ctx, session),
		showTownCommand(ctx, session),
		createTownCommand(ctx, session),
		statusCommand(ctx, session),
	}
	for _, c := range commandList {
		shell.AddCmd(c)
	}

	shell.AddCmd(&ishell.Cmd{
		Name: "status",
		Help: "display current status",
		Func: status,
	})
	shell.SetHomeHistoryPath(".ishell_history")
	shell.Run()
}

func tickCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "tick",
		Help: "tick history",
		Func: func(c *ishell.Context) {
			if currentArea == nil {
				c.Err(fmt.Errorf("please create a town first"))
				return
			}
			years, err := strconv.Atoi(c.Args[0])
			if err != nil {
				years = 1
			}
			for i := 0; i < years; i++ {
				c.Printf("%d\r", i)
				currentArea.Residents.Chronology.Tick()
			}
			c.Println()
		},
	}
}

func doLoadCulture(name string) error {
	r, err := os.Open(fmt.Sprintf("web/data/%s.json", name))
	if err != nil {
		return err
	}
	culture := &inhabitants.Culture{}
	if err = culture.Load(r); err != nil {
		return err
	}
	currentCulture = culture
	return nil
}

func loadCultureCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "load_culture",
		Help: "load a culture",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				c.Println("Please specify a name")
				return
			}
			name := c.Args[0]
			if err := doLoadCulture(name); err != nil {
				c.Err(err)
			}
			c.Println(currentCulture)
		},
	}

}

func loadSpeciesCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "load_species",
		Help: "load a species",
		Func: func(c *ishell.Context) {
			name := strings.ToLower(c.Args[0])
			if err := doLoadSpecies(name); err != nil {
				c.Err(err)
			}
			c.Println(currentSpecies)
		},
	}

}

func doLoadSpecies(name string) error {
	r, err := os.Open(fmt.Sprintf("web/data/%s.json", name))
	if err != nil {
		return err
	}
	species, err := inhabitants.LoadSpecies(r)
	if err != nil {
		return err
	}
	currentSpecies = &species
	return nil
}

func showTown(c *ishell.Context) {
	output := bytes.Buffer{}
	w := tabwriter.NewWriter(&output, 1, 1, 1, ' ', tabwriter.Debug)
	c.Println(currentArea.Name)
	beings := currentArea.Residents.Beings()
	for i := 0; i < len(beings); i++ {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%d", beings[i].String(), beings[i].Sex, beings[i].Age()))
	}
	w.Flush()
	c.Print(output.String())
}

func showTownCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "show_town",
		Help: "display current working town",
		Func: showTown,
	}
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
			being := inhabitants.NewBeing(currentSpecies, currentCulture)
			being.Randomize()
			area.Add(being)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	currentArea = area
	showTown(c)
}

func createTownCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "create_town",
		Help: "create a new town",
		Func: createTown,
	}
}

func status(c *ishell.Context) {
	if currentCulture != nil {
		c.Println(currentCulture)
	}
	if currentSpecies != nil {
		c.Println(currentSpecies)
	}
	if currentArea != nil {
		showTown(c)
	}
}

func statusCommand(ctx context.Context, session *mgo.Session) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "status",
		Help: "current status",
		Func: status,
	}
}
