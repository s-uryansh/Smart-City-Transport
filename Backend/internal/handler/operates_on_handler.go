package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /operates_on
func GetAllOperatesOn(c *gin.Context) {
	list, err := repository.GetAllOperatesOn()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch operates_on records"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

// POST /operates_on
func CreateOperatesOn(c *gin.Context) {
	var op models.OperatesOn
	if err := c.ShouldBindJSON(&op); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	err := repository.CreateOperatesOn(op)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create operates_on record"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Operates_On record created successfully"})
}

// GET /operates_on/:vid/:sid
func GetOperatesOnByIDs(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch operates_on"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	vid := h.VID
	op, err := repository.GetOperatesOnByIDs(vid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	c.JSON(http.StatusOK, op)
}

// DELETE /operates_on/:vid/:sid
func DeleteOperatesOn(c *gin.Context) {
	vID, err1 := strconv.Atoi(c.Param("vid"))
	sID, err2 := strconv.Atoi(c.Param("sid"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid V_ID or S_ID"})
		log.Println("Error(handler/operates_on): ", err1, " and ", err2)
		return
	}

	err := repository.DeleteOperatesOn(vID, sID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
