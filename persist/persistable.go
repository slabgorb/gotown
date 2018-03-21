package persist

import (
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
