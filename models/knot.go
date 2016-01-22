package models

import "gopkg.in/mgo.v2/bson"

type Knot struct {
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Text    string `json:"text"`
	Title   string `json:"title"`
	Created int64 `json:"created"`
}

func (knot Knot) Validate() bool {
	return knot.Text != "" && knot.Title != ""
}
