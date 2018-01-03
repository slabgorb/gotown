package culture

import (
	"github.com/slabgorb/gotown/inhabitants"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "cultures"
const database = "gotown"

// All gets all the cultures from the mongodb collection
func All(session *mgo.Session) ([]*inhabitants.Culture, error) {
	sessionCopy := session.Copy()
	c := sessionCopy.DB(database).C(collection)
	cultures := []*inhabitants.Culture{}
	err := c.Find(bson.M{}).All(&cultures)
	return cultures, err
}
