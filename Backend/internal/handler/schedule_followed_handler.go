package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"

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
