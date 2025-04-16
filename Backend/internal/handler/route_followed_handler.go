package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all route_followed entries
func GetAllRouteFollowed(c *gin.Context) {
	data, err := repository.GetAllRouteFollowed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch entries"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}
func GetRouteByVehicleID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/route): ", err)
		return
	}
	vid := h.VID
	route_id, err := repository.GetRoutesByVehicleID(vid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/route): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"route_id": route_id})
}

// Create a new route_followed entry
func CreateRouteFollowed(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/route): ", err)
		return
	}
	vid := h.VID
	var rf models.RouteFollowed
	rf.VehicleID = vid
	if err := c.ShouldBindJSON(&rf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	if err := repository.CreateRouteFollowed(rf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entry"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Route-followed entry created"})
}
func CreateRouteFollowedAdmin(c *gin.Context) {
	routeID, err := strconv.Atoi(c.Param("route_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	vid, err := strconv.Atoi(c.Param("vid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	var rf models.RouteFollowed
	rf.RouteID = routeID
	rf.VehicleID = vid
	if err := repository.CreateRouteFollowed(rf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create entry"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Route-followed entry created"})
}

// Delete a route_followed entry
func DeleteRouteFollowed(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/route): ", err)
		return
	}
	vid := h.VID
	routeID, err := strconv.Atoi(c.Param("route_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	if err := repository.DeleteRouteFollowed(vid, routeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entry"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
func DeleteRouteFollowedAdmin(c *gin.Context) {
	routeID, err := strconv.Atoi(c.Param("route_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	vid, err := strconv.Atoi(c.Param("vid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	if err := repository.DeleteRouteFollowed(vid, routeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete entry"})
		log.Println("Error(handler/route_followed): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}
