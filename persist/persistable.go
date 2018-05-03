package persist

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/asdine/storm"
)

// DB is the storm database
var DB *storm.DB

// SetDB sets the storm (bolt) database for the package
func SetDB(newDb *storm.DB) {
	DB = newDb
}

// Persistable models database persistence
type Persistable interface {
	Save() error
	Read() error
	Delete() error
	Reset()
}

func Open(path string) error {
	session, err := storm.Open(path, storm.Batch())
	if err != nil {
		return fmt.Errorf("could not open database at %s: %s", path, err)
	}
	SetDB(session)
	return nil
}

func SaveAll(items []Persistable) error {
	quit := make(chan struct{})
	errs := make(chan error)
	done := make(chan error)
	for _, i := range items {
		go func(i Persistable) {
			err := error(nil)
			ch := done
			if err = i.Save(); err != nil {
				ch = errs
			}
			select {
			case ch <- err:
				return
			case <-quit:
				return
			}
		}(i)
	}
	count := 0
	for {
		select {
		case err := <-errs:
			close(quit)
			return err
		case <-done:
			count++
			if count == len(items) {
				return nil
			}
		}
	}
}

func Close() error {
	return DB.Close()
}

func OpenTestDB() {
	err := Open("_gotown_test.db")
	if err != nil {
		panic(err)
	}
}

func CloseTestDB() {
	defer os.Remove("_gotown_test.db")
	DB.Close()
}

func SeedHelper(pathname string, item Persistable) error {
	bundle := PersistBundle
	if err := DB.Drop(item); err != nil {
		if err.Error() != "bucket not found" {
			return fmt.Errorf("could not delete bucket %s:%s", pathname, err)
		}
	}
	for _, name := range bundle.Files() {
		splits := strings.Split(name, "/")
		if splits[0] != pathname {
			continue
		}
		r, _ := bundle.Open(name)
		if err := json.NewDecoder(r).Decode(item); err != nil {
			return fmt.Errorf("could not decode %s/%s: %s", pathname, name, err)
		}
		if err := DB.Save(item); err != nil {
			return fmt.Errorf("could not save %s/%s: %s", pathname, name, err)
		}
		item.Reset()
	}

	return nil
}
