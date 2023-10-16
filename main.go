package main

import (
	"log"

	"github.com/Lalu-Mahato/go-dhan-das/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load the .env file")
	}
	// config.DatabaseConnection()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	routes.SetupRouter(r)
	r.Run()
}
