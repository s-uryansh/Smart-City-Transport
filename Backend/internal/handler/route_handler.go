package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all routes
func GetAllRoutes(c *gin.Context) {
	routes, err := repository.GetAllRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve routes"})
		log.Println("Error(handler/route): ", err)
		return
	}
	c.JSON(http.StatusOK, routes)
}

// Get one route by ID
func GetRouteByID(c *gin.Context) {
	routeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route): ", err)
		return
	}
	route, err := repository.GetRouteByID(routeID)
	// log.Println(route.JourneyTime)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		log.Println("Error(handler/route): ", err)
		return
	}

	c.JSON(http.StatusOK, route)
}

// Create a new route
func CreateRoute(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		log.Println("Error reading body: ", err)
		return
	}

	var r models.Route
	// Unmarshal the JSON directly into the Route struct
	if err := json.Unmarshal(bodyBytes, &r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error unmarshaling into struct: ", err)
		return
	}

	// Ensure that journey_time is a string and is set correctly
	// You can add validation here if needed
	if r.JourneyTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Journey time is required"})
		log.Println("Error: Journey time is required")
		return
	}

	// Call repository to create the route'
	r.RID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route): ", err)
		return
	}
	log.Println(r.RID)
	if err := repository.CreateRoute(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create route"})
		log.Println("Error(handler/route): ", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Route created successfully"})
}

// Update an existing route
func UpdateRoute(c *gin.Context) {
	routeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route): ", err)
		return
	}
	var r models.Route
	r.RID = routeID
	log.Println("Route ID: ", routeID)
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		log.Println("Error(handler/route): ", err)
		return
	}

	if err := repository.UpdateRoute(routeID, r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update route"})
		log.Println("Error(handler/route): ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Route updated successfully"})
}

// Delete a route
func DeleteRoute(c *gin.Context) {
	routeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route): ", err)
		return
	}
	log.Println(routeID)
	if err := repository.DeleteRoute(routeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete route"})
		log.Println("Error(handler/route): ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Route deleted successfully"})
}
