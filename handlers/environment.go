package handlers

import (
	"gopkg.in/mgo.v2"
)

type Env struct {
	db *mgo.Database
	session *mgo.Session
}

func InitEnv() *Env {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	db := session.DB("knots_db")
	return &Env{db, session}
}

func (env Env) Destroy()  {
	defer env.session.Close()
}