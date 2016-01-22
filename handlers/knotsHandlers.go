package handlers

import (
	"log"
	"github.com/gin-gonic/gin"
	"knots/models"
	"time"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func (env Env) NewKnot(c *gin.Context) {
	var new models.Knot
	err := c.BindJSON(&new)

	if err == nil && new.Validate() {
		new.Created = time.Now().Unix()
		collection := env.db.C("knots")
		err := collection.Insert(&new)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "posted",
				"title": new.Title,
			})
		}
		return
	}
	c.JSON(200, gin.H{
		"status": "invalid_data",
	})

}

func (env Env) DeleteKnot(c *gin.Context) {
	collection := env.db.C("knots")

	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.JSON(400, gin.H{
			"status": false,
			"error": "id not valid",
		})
	}
	mongoId := bson.ObjectIdHex(id)
	err := collection.RemoveId(mongoId)

	if err != nil {
		c.JSON(500, gin.H{
			"status": false,
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": true,
		})
	}
}

func (env Env) AllKnot(c *gin.Context) {
	collection := env.db.C("knots")
	var result []models.Knot
	err := collection.Find(nil).Sort("-created").All(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		c.JSON(200, gin.H{
			"status": true,
			"knots": result,
		})
	}
}
