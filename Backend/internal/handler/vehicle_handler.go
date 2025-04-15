package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllVehicles(c *gin.Context) {
	vehicles, err := repository.GetAllVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch vehicles"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func GetVehicleByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	// log.Println("ieubriawundruipajsd")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	id := h.VID
	// log.Println("Entering GetVehicleByID ")
	vehicle, err := repository.GetVehicleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func CreateVehicle(c *gin.Context) {
	var v models.Vehicle
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	v.LastUpdate = time.Now()
	err := repository.CreateVehicle(v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create vehicle"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Vehicle created"})
}

func UpdateVehicle(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	id := h.VID

	var v models.Vehicle
	v.VehicleID = id
	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	log.Println("Vehicle Updating")
	err = repository.UpdateVehicle(id, v)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update vehicle"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle updated"})
}

func DeleteVehicle(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	id := h.VID
	err = repository.DeleteVehicle(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete vehicle"})
		log.Println("Error(handler/vehicle): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vehicle deleted"})
}
