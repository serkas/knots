package handlers

import (
	"log"
	"github.com/gin-gonic/gin"
	"knots/models"
	"time"
)

func (env Env) NewKnot(c *gin.Context) {
	var new models.Knot
	c.BindJSON(&new)

	if new.Validate() {
		new.Created = time.Now().Unix()
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
