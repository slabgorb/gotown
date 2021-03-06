package persist

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/alicebob/miniredis"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/satori/go.uuid"
)

var DB *pool.Pool
var Test *miniredis.Miniredis

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

func getType(v interface{}) string {
	rawType := reflect.TypeOf(v).String()
	splits := strings.Split(rawType, ".")
	return splits[len(splits)-1]
}

type Identifiable interface {
	GetID() string
	SetID(string)
}

type IdentifiableImpl struct {
	ID string `json:"id"`
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
	fmt.Stringer
	Reset()
	Identifiable
}

type PersistableImpl struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

func getConn(f func(conn *redis.Client) error) error {
	conn, err := DB.Get()
	if err != nil {
		return fmt.Errorf("cannot get connection: %s", err)
	}
	defer DB.Put(conn)
	return f(conn)
}

// Read reads in by id
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

func Mread(ids []string, items map[string]Persistable) error {
	return getConn(func(conn *redis.Client) error {
		js, err := conn.Cmd("MGET", ids).Array()
		if err != nil {
			return fmt.Errorf("could not mget: %s", err)
		}
		for _, r := range js {
			j, err := r.Bytes()
			if err != nil {
				return err
			}

			type ider struct {
				ID string `json:"id"`
			}

			iding := &ider{}
			err = json.Unmarshal(j, iding)
			if err != nil {
				return fmt.Errorf("could not unmarshal json: %s", err)
			}
			item := items[iding.ID]
			err = json.Unmarshal(j, item)
			if err != nil {
				return fmt.Errorf("could not unmarshal json: %s", err)
			}
		}
		return nil
	})
}

type operation func(client *redis.Client) error

// Open opens the connection to the database file
func Open() error {
	p, err := pool.New("tcp", "redis:6379", 10)
	if err != nil {
		return fmt.Errorf("could not open pool:%s", err)
	}
	if err != nil {
		return fmt.Errorf("could not open database:%s", err)
	}
	SetDB(p)
	return nil
}

func List(setKey string) (map[string]string, error) {
	pairs := make(map[string]string)
	keys := []interface{}{}
	iter := 0
	err := getConn(func(conn *redis.Client) error {
		for {
			s, err := conn.Cmd("SSCAN", setKey, iter).Array()
			if err != nil {
				return fmt.Errorf("could not scan %s: %s", setKey, err)
			}
			iter, err = s[0].Int()
			if err != nil {
				return fmt.Errorf("could not get cursor: %s", err)
			}
			ary, err := s[1].Array()
			if err != nil {
				return fmt.Errorf("could not get array: %s", err)
			}
			for _, sc := range ary {
				str, err := sc.Str()
				if err != nil {
					return fmt.Errorf("could not get string: %s", err)
				}
				keys = append(keys, str)

			}
			if iter == 0 {
				break
			}

		}
		if len(keys) == 0 {
			return nil
		}
		items, err := conn.Cmd("MGET", keys...).Array()
		if err != nil {
			return fmt.Errorf("could not get array: %s", err)
		}
		type nameAndId struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		}
		for _, i := range items {
			j, err := i.Bytes()
			if err != nil {
				return err
			}
			pair := nameAndId{}
			err = json.Unmarshal(j, &pair)
			if err != nil {
				return fmt.Errorf("could not unmarshal json: %s", err)
			}
			pairs[pair.ID] = pair.Name
		}
		return nil
	})
	return pairs, err
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
		return conn.Cmd("SADD", getType(item), item.GetID()).Err
	})
}

func ReadByName(name string, set string, item Persistable) error {
	list, err := List(set)
	if err != nil {
		return err
	}
	id := ""
	for k, v := range list {
		if v == name {
			id = k
			break
		}
	}
	item.SetID(id)
	return item.Read()
}

func Update(item Persistable) error {
	return Save(item)
}

func Delete(item Persistable) error {
	if item.GetID() == "" {
		return fmt.Errorf("cannot read without id")
	}
	return getConn(func(conn *redis.Client) error {
		err := conn.Cmd("DEL", item.GetID()).Err
		if err != nil {
			return fmt.Errorf("could not del item %s: %s", item.GetID(), err)
		}
		return conn.Cmd("SREM", getType(item), item.GetID()).Err
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
	var err error
	Test, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	p, err := pool.New("tcp", Test.Addr(), 10)
	if err != nil {
		panic(fmt.Errorf("could not open pool:%s", err))
	}
	if err != nil {
		panic(fmt.Errorf("could not open database:%s", err))
	}
	SetDB(p)
}

// CloseTestDB closes the file and deletes it
func CloseTestDB() {
	Test.Close()
}

func DeleteAll() error {
	return getConn(func(conn *redis.Client) error {
		return conn.Cmd("FLUSHALL").Err
	})
}

func SeedHelper(pathname string, item Persistable) error {
	bundle := PersistBundle
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
