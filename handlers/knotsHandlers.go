package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"knots/models"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "invalid_data",
		})
	}


}

func (env Env) DeleteKnot(c *gin.Context) {
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "id not valid",
		})
	}
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

func sendError(c *gin.Context, err error)  {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": false,
		"error": err.Error(),
	})
}