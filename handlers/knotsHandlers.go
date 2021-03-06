package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"knots/models"
	"fmt"
)

func (env Env) NewKnot(c *gin.Context) {
	var new models.Knot
	err := c.BindJSON(&new)

	if err == nil && new.Validate() {
		err := new.Insert(env.db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "posted",
				"title": new.Title,
			})
		}
	} else {
		sendBadRequest(c, fmt.Errorf("invalid_data"))
	}
}

func (env Env) UpdateKnot(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		sendBadRequest(c, fmt.Errorf("id_not_valid"))
	}
	mongoId := bson.ObjectIdHex(id)

	var new models.Knot
	var err error
	err = c.BindJSON(&new)

	if err != nil || !new.Validate() {
		sendBadRequest(c, fmt.Errorf("invalid_data"))
	} else {
		err = new.Update(env.db, mongoId)
		if err != nil {
			sendError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"status": true})
		}
	}
}

func (env Env) DeleteKnot(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		sendBadRequest(c, fmt.Errorf("id_not_valid"))
	}else{
		mongoId := bson.ObjectIdHex(id)

		collection := env.db.C("knots")
		err := collection.RemoveId(mongoId)

		if err != nil {
			sendError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": true,
			})
		}
	}

}

func (env Env) AllKnot(c *gin.Context) {
	collection := env.db.C("knots")
	var result []models.Knot
	err := collection.Find(nil).Sort("-created").All(&result)
	if err != nil {
		sendError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"knots": result,
		})
	}
}

func (env Env) OneKnot(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		sendBadRequest(c, fmt.Errorf("id_not_valid"))
	}else {
		mongoId := bson.ObjectIdHex(id)

		collection := env.db.C("knots")
		var result models.Knot
		err := collection.FindId(mongoId).One(&result)
		if err != nil {
			sendError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": true,
				"knot": result,
			})
		}
	}

}