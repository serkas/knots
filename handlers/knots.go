package handlers

import (
	"github.com/gin-gonic/gin"
	"knots/models"
)

func NewKnot(c *gin.Context) {
	var new models.Knot
	c.BindJSON(&new)

	if new.Validate() {
		c.JSON(200, gin.H{
			"status": "posted",
			"title": new.Title,
		})
	}else{
		c.JSON(200, gin.H{
			"status": "entity_not_valid",
		})
	}

}
