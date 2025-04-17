package main

import (
	"SmartCityTransportSystem/config"
	"SmartCityTransportSystem/pkg/db"
	"SmartCityTransportSystem/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}
	r := gin.New() // Gin router
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.RedirectTrailingSlash = false
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	r.Run(":" + port)
	err = db.Connect()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	r.Use(gin.Logger(), gin.Recovery())

	routes.InitRoutes(r)
	// r.Run(port) // Start the server on port 8080
	db.Close()
}
