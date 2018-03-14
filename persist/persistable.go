package persist

import (
	"github.com/asdine/storm"
)

var DB *storm.DB

func SetDB(newDb *storm.DB) {
	DB = newDb
}

// Persistable models database persistence
type Persistable interface {
	Save() error
	Read() error
	Delete() error
}
