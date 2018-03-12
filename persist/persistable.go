package persist

import (
	"encoding/json"
	"fmt"

	bolt "github.com/coreos/bbolt"
)

type Bucket string

var db *bolt.DB

func SetDB(newDb *bolt.DB) {
	db = newDb
}

func (b Bucket) getBytes() []byte {
	return []byte(string(b))
}

const (
	CultureBucket Bucket = "cultures"
	SpeciesBucket Bucket = "species"
	AreaBucket    Bucket = "area"
)

// Persistable models database persistence
type Persistable interface {
	Save() error
	Fetch() error
	Delete() error
	GetBucket() Bucket
	GetKey() string
}

// DoSave saves the json to the bucket
func DoSave(j Persistable) error {
	bucket := j.GetBucket()
	key := j.GetKey()
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket.getBytes())
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(j)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), encoded)
	})
}

// DoFetch returns the marshaled data
func DoFetch(j Persistable) error {
	bucket := j.GetBucket()
	key := j.GetKey()
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket.getBytes())
		v := b.Get([]byte(key))
		err := json.Unmarshal(v, j)
		if err != nil {
			return fmt.Errorf("Error unmarshaling %s %s", bucket, key)
		}
		return nil
	})
}

// DoDelete deletes a key from a bucket
func DoDelete(j Persistable) error {
	bucket := j.GetBucket()
	key := j.GetKey()
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket.getBytes())
		if err != nil {
			return err
		}
		return b.Delete([]byte(key))
	})
}

// ListBucketKeys lists the index for a bucket
func ListBucketKeys(bucket Bucket) []string {
	names := []string{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket.getBytes())
		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			names = append(names, string(k))
		}
		return nil
	})
	return names
}
