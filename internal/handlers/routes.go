package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "auth pong"})
	})
	auth := r.Group("/auth")
	RegisterAuthRoutes(auth)
}
