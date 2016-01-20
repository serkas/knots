package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", index)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}