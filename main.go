package main

import (
	"github.com/gin-gonic/gin"
	"knots/handlers"
)


func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	r.GET("/", handlers.Index)
	r.POST("/new", handlers.NewKnot)

	r.Run(":8080")
}