package handlers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.File("./templates/index.html")
}
