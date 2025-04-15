package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMaintenance(c *gin.Context) {
	data, err := repository.GetAllMaintenance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch maintenance records"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetMaintenanceByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch operates_on"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	vid := h.VID
	data, err := repository.GetMaintenanceByID(vid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Maintenance record not found"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateMaintenance(c *gin.Context) {
	var m models.Maintenance
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	if err := repository.CreateMaintenance(m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create maintenance record"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Maintenance record created"})
}

func UpdateMaintenance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	var m models.Maintenance
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	if err := repository.UpdateMaintenance(id, m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update maintenance record"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Maintenance record updated"})
}

func DeleteMaintenance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	if err := repository.DeleteMaintenance(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete maintenance record"})
		log.Println("Error(handler/maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Maintenance record deleted"})
}
