package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "auth pong"})
	})
	RegisterAuthRoutes(r.Group("/auth"))
	RegisterDiaryRoutes(r.Group("/diary"))
}
