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
		AllowOrigins:     []string{"https://smart-city-transport.vercel.app"},
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

	go func() {
		log.Println("Connecting to database in background...")
		if err := db.Connect(); err != nil {
			log.Printf("Database connection error: %v", err)
			// optionally: trigger shutdown if DB is critical
		} else {
			log.Println("Database connected successfully (background init)")
		}
	}()

	r.Use(gin.Logger(), gin.Recovery())
	routes.InitRoutes(r)
	r.Run(":" + port)
	db.Close()
}
