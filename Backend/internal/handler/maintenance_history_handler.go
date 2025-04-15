package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMaintenanceHistories(c *gin.Context) {
	list, err := repository.GetAllMaintenanceHistories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch maintenance history"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetMaintenanceHistoryByMID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch maintenance_history"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	vid := h.VID
	list, err := repository.GetMaintenanceHistoryByMID(vid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch maintenance history"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

func CreateMaintenanceHistory(c *gin.Context) {
	var mh models.MaintenanceHistory
	if err := c.ShouldBindJSON(&mh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	if err := repository.CreateMaintenanceHistory(mh); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create maintenance history record"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record created"})
}

func DeleteMaintenanceHistory(c *gin.Context) {
	mid, err1 := strconv.Atoi(c.Param("m_id"))
	vid, err2 := strconv.Atoi(c.Param("v_id"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid M_ID or V_ID"})
		log.Println("Error(handler/maintenance_history): ", err1, " and ", err2)
		return
	}
	if err := repository.DeleteMaintenanceHistory(mid, vid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete maintenance history record"})
		log.Println("Error(handler/maintenance_history): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}
