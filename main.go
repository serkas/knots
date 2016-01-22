package main

import (
	"github.com/gin-gonic/gin"
	"knots/handlers"
)


func main() {
	r := gin.Default()

	env := handlers.InitEnv()
	defer env.Destroy()

	r.Static("/assets", "./assets")
	r.GET("/", handlers.Index)
	r.GET("/knots", env.AllKnot)
	r.POST("/knots", env.NewKnot)
	r.DELETE("/knots/:id", env.DeleteKnot)

	r.Run(":8080")
}

