package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": false,
		"error": err.Error(),
	})
}

func sendBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": false,
		"error": err.Error(),
	})
}
