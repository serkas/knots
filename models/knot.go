package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type Knot struct {
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Raw      string `json:"raw"`
	Title    string `json:"title"`
	Created  int64 `json:"created"`
	Modified int64 `json:"modified"`
}

func (knot Knot) Validate() bool {
	return knot.Raw != ""
}

func (knot Knot) Insert(db *mgo.Database) error {
	knot.Created = time.Now().Unix()
	collection := db.C("knots")
	return collection.Insert(&knot)
}

func (knot Knot) Update(db *mgo.Database, mongoId bson.ObjectId) error {
	knot.Modified = time.Now().Unix()
	collection := db.C("knots")
	return collection.UpdateId(mongoId, knot)
}
