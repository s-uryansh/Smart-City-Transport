package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllAccidentHistory(c *gin.Context) {
	data, err := repository.GetAllAccidentHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accident history: "})
		log.Println("Error(handler/accident_history): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}
func GetAccidentHistoryByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accident_histroy"})
		log.Println("Error(handler/accident_history): ", err)
		return
	}
	vid := h.VID
	data, err := repository.GetAccidentHistoryByID(vid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}
func UpdateAccidentHistory(c *gin.Context) {
	var oldVID, oldIID int
	var newEntry models.AccidentHistory

	// Get old composite key from query params
	oldVID, err1 := strconv.Atoi(c.Query("old_vid"))
	oldIID, err2 := strconv.Atoi(c.Query("old_iid"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if err := c.ShouldBindJSON(&newEntry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := repository.UpdateAccidentHistory(oldVID, oldIID, newEntry.VID, newEntry.IID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update accident history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Accident history updated"})
}

func CreateAccidentHistory(c *gin.Context) {
	var ah models.AccidentHistory
	if err := c.ShouldBindJSON(&ah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("Error(handler/accident_history): ", err)
		return
	}
	if err := repository.CreateAccidentHistory(ah); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create accident history"})
		log.Println("Error(handler/accident_history): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Accident history created"})
}

func DeleteAccidentHistory(c *gin.Context) {
	vidStr := c.Param("vid")
	iidStr := c.Param("iid")

	vid, err1 := strconv.Atoi(vidStr)
	iid, err2 := strconv.Atoi(iidStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid V_ID or I_ID"})
		log.Println("Error(handler/accident_history): ", err1, "and", err2)
		return
	}

	if err := repository.DeleteAccidentHistory(vid, iid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete accident history"})
		log.Println("Error(handler/accident_history): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Accident history deleted"})
}
