package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerformsMaintenance(c *gin.Context) {
	list, err := repository.GetAllPerformsMaintenance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch performs_maintenance data"})
		log.Println("Error(handler/perform_maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetPerformsMaintenanceByMID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/route): ", err)
		return
	}
	list, err := repository.GetPerformsMaintenanceByMID(h.IDNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		log.Println("Error(handler/perform_maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

func CreatePerformsMaintenance(c *gin.Context) {
	var pm models.PerformsMaintenance
	if err := c.ShouldBindJSON(&pm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/perform_maintenance): ", err)
		return
	}
	if err := repository.CreatePerformsMaintenance(pm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert record"})
		log.Println("Error(handler/perform_maintenance): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record created"})
}

func DeletePerformsMaintenance(c *gin.Context) {
	mid, err1 := strconv.Atoi(c.Param("m_id"))
	staffID, err2 := strconv.Atoi(c.Param("staff_id"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid M_ID or STAFF_ID"})
		log.Println("Error(handler/perform_maintenance): ", err1, " and ", err2)
		return
	}
	if err := repository.DeletePerformsMaintenance(mid, staffID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		log.Println("Error(handler/perform_maintenance): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}
