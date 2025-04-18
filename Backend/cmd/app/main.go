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
		AllowOrigins:     []string{"https://smart-city-transport.vercel.app/"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Run(":" + port)
	log.Println("Server started on port " + port)

	r.RedirectTrailingSlash = false
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	log.Println("Config loaded successfully")
	err = db.Connect()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	log.Println("Connected to database successfully")

	r.Use(gin.Logger(), gin.Recovery())
	routes.InitRoutes(r)

	log.Println("Routes initialized successfully")
	db.Close()
}
