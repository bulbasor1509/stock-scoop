package main

import (
	"github.com/bulbasor1509/stock-scoop/backend/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/ping", handlers.BulkHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start the server on port 8080")
	}
}
