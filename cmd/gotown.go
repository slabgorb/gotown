package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/abiosoft/ishell"
	bolt "github.com/coreos/bbolt"
	"github.com/slabgorb/gotown/inhabitants"
	"github.com/slabgorb/gotown/locations"
)

var currentCulture *inhabitants.Culture
var currentSpecies *inhabitants.Species
var currentArea *locations.Area

func main() {
	ctx := context.Background()
	session, err := bolt.Open("gotown.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
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
		seedCommand(ctx, session),
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

func seedCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "seed",
		Help: "seed database with static json files",
		Func: func(c *ishell.Context) {
			if err := seedSpecies(session, c); err != nil {
				c.Err(err)
			}
			if err := seedCultures(session, c); err != nil {
				c.Err(err)
			}
		},
	}
}

func seedSpecies(session *bolt.DB, c *ishell.Context) error {
	speciesNames := []string{"human", "elf"}
	for _, name := range speciesNames {
		c.Printf("\tloading %s\n", name)
		r, err := os.Open(fmt.Sprintf("web/data/%s.json", name))
		if err != nil {
			return err
		}
		species := &inhabitants.Species{}
		if err = species.Load(r); err != nil {
			return err
		}
		if err = doSaveSpecies(species, session); err != nil {
			return err
		}
	}
	return nil
}

func seedCultures(session *bolt.DB, c *ishell.Context) error {
	cultureNames := []string{"italianate", "viking"}
	for _, name := range cultureNames {
		c.Printf("\tloading %s\n", name)
		r, err := os.Open(fmt.Sprintf("web/data/%s.json", name))
		if err != nil {
			return err
		}
		culture := &inhabitants.Culture{}
		if err = culture.Load(r); err != nil {
			return err
		}
		if err = doSaveCulture(culture, session); err != nil {
			return err
		}

	}
	return nil
}

func tickCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
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

func doSaveCulture(culture *inhabitants.Culture, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("cultures"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(culture)
		if err != nil {
			return err
		}
		return b.Put([]byte(culture.Name), encoded)
	})
}

func doLoadCulture(name string, db *bolt.DB) error {
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("cultures"))
		v = bucket.Get([]byte(name))
		return nil
	})
	culture := &inhabitants.Culture{}
	err := culture.Load(bytes.NewReader(v))
	if err != nil {
		return err
	}
	currentCulture = culture
	return nil
}

func loadCultureCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "load_culture",
		Help: "load a culture",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				c.Println("Please specify a name")
				return
			}
			name := c.Args[0]
			if err := doLoadCulture(name, session); err != nil {
				c.Err(err)
			}
			c.Println(currentCulture)
		},
	}

}

func loadSpeciesCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "load_species",
		Help: "load a species",
		Func: func(c *ishell.Context) {
			name := strings.ToLower(c.Args[0])
			if err := doLoadSpecies(name, session); err != nil {
				c.Err(err)
			}
			c.Println(currentSpecies)
		},
	}

}

func doLoadSpecies(name string, db *bolt.DB) error {
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("species"))
		v = bucket.Get([]byte(name))
		return nil
	})
	species := &inhabitants.Species{}
	err := species.Load(bytes.NewReader(v))
	if err != nil {
		return err
	}
	currentSpecies = species
	return nil
}

func doSaveSpecies(species *inhabitants.Species, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("species"))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(species)
		if err != nil {
			return err
		}
		return b.Put([]byte(species.Name), encoded)
	})
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

func showTownCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
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

func createTownCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
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

func statusCommand(ctx context.Context, session *bolt.DB) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "status",
		Help: "current status",
		Func: status,
	}
}
