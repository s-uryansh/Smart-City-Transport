package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"SmartCityTransportSystem/pkg/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllScheduleFollowed(c *gin.Context) {
	sfList, err := repository.GetAllScheduleFollowed()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch schedule_followed"})
		log.Println("Error(handler/schedule_followed): ", err)
		return
	}
	c.JSON(http.StatusOK, sfList)
}

// Bad Code Part
func UpdateScheduleFollowedByID(c *gin.Context) {
	// Parse the ID from URL param
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Parse input JSON
	var sf models.ScheduleFollowed
	if err := c.ShouldBindJSON(&sf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error binding JSON: ", err)
		return
	}

	// Perform the update
	_, err = db.DB.Exec(
		"UPDATE schedule_followed SET R_ID = ? WHERE S_ID = ?",
		sf.RID, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule_followed"})
		log.Println("Error updating DB: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ScheduleFollowed updated successfully"})
}

func GetScheduleFollowedByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedule"})
		log.Println("Error(handler/schedule-followed): ", err)
		return
	}
	var schedule models.Schedule
	schedule, err = repository.GetScheduleByVID(h.VID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedule"})
		log.Println("Error ScheduleByVID(handler/schedule-folowed/GetScheduleByVID): ", err)
		return
	}
	log.Println(schedule.VID)
	schedule_id, err := repository.GetScheduleFollowedByID(schedule.ScheduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedule"})
		log.Println("Error(handler/schedule-followed/GetScheduleFollowedByID): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Schedule-Followed ID": schedule_id,
		"Vehicle ID":           schedule.VID})
}

func CreateScheduleFollowed(c *gin.Context) {
	var sf models.ScheduleFollowed
	if err := c.ShouldBindJSON(&sf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/schedule_followed): ", err)
		return
	}
	err := repository.CreateScheduleFollowed(sf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create schedule_followed"})
		log.Println("Error(handler/schedule_followed): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ScheduleFollowed created successfully"})
}
func CreateScheduleFollowedAdmin(c *gin.Context) {
	rid, err := strconv.Atoi(c.Param("R_ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid R_ID"})
		return
	}
	sid, err := strconv.Atoi(c.Param("S_ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid S_ID"})
		return
	}

	sf := models.ScheduleFollowed{
		RID: rid,
		SID: sid,
	}

	log.Println("Creating ScheduleFollowed:", sf)

	err = repository.CreateScheduleFollowed(sf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create schedule_followed"})
		log.Println("Error(handler/schedule_followed): ", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ScheduleFollowed created successfully"})
}

func DeleteScheduleFollowed(C *gin.Context) {
	idStr := C.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		C.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Perform the deletion
	err = repository.DeleteScheduleFollowed(id)
	if err != nil {
		C.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule_followed"})
		log.Println("Error deleting DB: ", err)
		return
	}

	C.JSON(http.StatusOK, gin.H{"message": "ScheduleFollowed deleted successfully"})
}
