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

	// Knots handlers
	r.GET("/knots", env.AllKnot)
	r.POST("/knots", env.NewKnot)
	r.PUT("/knots/:id", env.UpdateKnot)
	r.DELETE("/knots/:id", env.DeleteKnot)

	r.Run(":8080")
}

