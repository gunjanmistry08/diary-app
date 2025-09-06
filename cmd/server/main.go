package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gunjanmistry08/diary-app/internal/database"
	"github.com/gunjanmistry08/diary-app/internal/handlers"
	"github.com/gunjanmistry08/diary-app/internal/server"
)

func main() {
	database.Connect()

	r := gin.Default()
	handlers.RegisterRoutes(r)

	addr := fmt.Sprintf(":%s", database.Config().Port)
	server.Run(r, addr)
}
