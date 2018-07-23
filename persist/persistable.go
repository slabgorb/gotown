package persist

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/satori/go.uuid"
)

var DB *pool.Pool

func getUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}

func SetDB(p *pool.Pool) {
	DB = p
}

type IDPair struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Identifiable interface {
	GetID() string
	SetID(string)
}

type IdentifiableImpl struct {
	ID string
}

func (i *IdentifiableImpl) GetID() string {
	return i.ID
}

func (i *IdentifiableImpl) SetID(id string) {
	i.ID = id
}

// Persistable models database persistence
type Persistable interface {
	Save() error
	Read() error
	Delete() error
	Reset()
	Identifiable
}

func getConn(f func(conn *redis.Client) error) error {
	conn, err := DB.Get()
	if err != nil {
		return fmt.Errorf("cannot get connection: %s", err)
	}
	defer DB.Put(conn)
	return f(conn)
}

// Read reads in by id or name
func Read(i Identifiable) error {
	if i.GetID() == "" {
		return fmt.Errorf("cannot read without id")
	}
	return getConn(func(conn *redis.Client) error {
		j, err := conn.Cmd("GET", i.GetID()).Bytes()
		if err != nil {
			return fmt.Errorf("could not get %s from cache: %s", i.GetID(), err)
		}
		err = json.Unmarshal(j, &i)
		if err != nil {
			return fmt.Errorf("could not unmarshal json: %s", err)
		}
		return nil
	})
}

type operation func(client *redis.Client) error

// Open opens the connection to the database file
func Open() error {
	p, err := pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		return fmt.Errorf("could not open pool:%s", err)
	}
	if err != nil {
		return fmt.Errorf("could not open database:%s", err)
	}
	SetDB(p)
	return nil
}

// Save saves a Persistable
func Save(item Persistable) error {
	if item.GetID() == "" {
		item.SetID(getUUID())
	}
	return getConn(func(conn *redis.Client) error {
		j, err := json.Marshal(item)
		if err != nil {
			return fmt.Errorf("could not marshal item %s: %s", item.GetID(), err)
		}
		err = conn.Cmd("SET", item.GetID(), j).Err
		if err != nil {
			return fmt.Errorf("could not save item %s: %s", item.GetID(), err)
		}
		return nil
	})
}

func Update(item Persistable) error {
	return Save(item)
}

func Delete(item Persistable) error {
	if item.GetID() == "" {
		return fmt.Errorf("cannot read without id")
	}
	return getConn(func(conn *redis.Client) error {
		return conn.Cmd("DEL", item.GetID()).Err
	})
}

// SaveAll saves a slice of persistables
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

// Close closes the connection to the database file
func Close() error {
	return nil
}

// OpenTestDB sets up a connection to a test db instance
func OpenTestDB() {
	err := Open()
	if err != nil {
		panic(err)
	}
}

// CloseTestDB closes the file and deletes it
func CloseTestDB() {
}

func deleteAll() error {
	return getConn(func(conn *redis.Client) error {
		return conn.Cmd("FLUSHALL").Err
	})
}

func SeedHelper(pathname string, item Persistable) error {
	bundle := PersistBundle
	err := deleteAll()
	if err != nil {
		return err
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
		if err := Save(item); err != nil {
			return fmt.Errorf("could not save %s/%s: %s", pathname, name, err)
		}
		item.Reset()
	}
	return nil
}
