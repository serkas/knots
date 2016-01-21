package handlers

import (
	"log"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"knots/models"
)

func (env Env) NewKnot(c *gin.Context) {
	var new models.Knot
	c.BindJSON(&new)

	if new.Validate() {
		collection := env.db.C("knots")
		err := collection.Insert(&new)
		if err != nil {
			log.Fatal(err)
		} else {
			c.JSON(200, gin.H{
				"status": "posted",
				"title": new.Title,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"status": "entity_not_valid",
	})

}

func (env Env) AllKnot(c *gin.Context) {
	collection := env.db.C("knots")
	var result []models.Knot
	err := collection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		c.JSON(200, gin.H{
			"status": true,
			"knots": result,
		})
	}
}
