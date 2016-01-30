package main

import (
	"github.com/gin-gonic/gin"
	"knots/handlers"
)


func main() {
	r := gin.Default()
	authorized := r.Group("", gin.BasicAuth(gin.Accounts{
		"admin":    "admin",
	}))

	env := handlers.InitEnv()
	defer env.Destroy()

	r.Static("/assets", "./assets")
	r.GET("/", handlers.Index)

	// Knots handlers
	r.GET("/knots", env.AllKnot)
	r.GET("/knots/:id", env.OneKnot)
	authorized.POST("/knots", env.NewKnot)
	authorized.PUT("/knots/:id", env.UpdateKnot)
	authorized.DELETE("/knots/:id", env.DeleteKnot)

	r.Run(":8001")
}

