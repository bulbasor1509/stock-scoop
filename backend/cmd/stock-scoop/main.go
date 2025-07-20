package main

import (
	"database/sql"
	"github.com/bulbasor1509/stock-scoop/backend/internal/config"
	"github.com/bulbasor1509/stock-scoop/backend/internal/database"
	"github.com/bulbasor1509/stock-scoop/backend/internal/handlers"
	"github.com/bulbasor1509/stock-scoop/backend/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {

	if os.Getenv("DOCKER_RUN") == "true" {
		log.Println("running in docker container, skipping .env load")
	} else if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("DB_URL environment variable not set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln("error while connecting to the database")
	}
	queries := database.New(conn)
	apiCfg := types.ApiConfig{
		DB: queries,
	}

	router := gin.Default()
	cfg := cors.DefaultConfig()
	config.CorsConfig(&cfg)
	router.Use(cors.New(cfg))

	//gin.SetMode(gin.ReleaseMode)

	router.GET("/ping", handlers.BulkHandler(&apiCfg))
	router.GET("/bulk", handlers.GetBulkHandler(&apiCfg))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start the server on port 8080")
	}
}
