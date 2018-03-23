package persist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

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
}

func Open(path string) error {
	session, err := storm.Open(path)
	if err != nil {
		return err
	}
	SetDB(session)
	return nil
}

func Close() error {
	return DB.Close()
}

func OpenTestDB() {
	err := Open("gotown_test.db")
	if err != nil {
		panic(err)
	}
}

func CloseTestDB() {
	defer os.Remove("gotown_test.db")
	DB.Close()
}

func SeedHelper(pathname string, item Persistable) error {
	files, err := ioutil.ReadDir(pathname)
	if err != nil {
		return err
	}
	if err := DB.Drop(item); err != nil {
		if err.Error() != "bucket not found" {
			return fmt.Errorf("could not delete bucket %s:%s", pathname, err)
		}
	}
	for _, file := range files {
		filepath := path.Join(pathname, file.Name())
		r, err := os.Open(filepath)
		if err != nil {
			return err
		}
		if err := json.NewDecoder(r).Decode(item); err != nil {
			return fmt.Errorf("could not decode %s from file %s: %s", pathname, file.Name(), err)
		}
		if err := DB.Save(item); err != nil {
			return fmt.Errorf("could not save %s: %s", pathname, err)
		}
	}
	return nil
}
