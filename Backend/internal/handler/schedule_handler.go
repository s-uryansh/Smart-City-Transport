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

func GetAllSchedules(c *gin.Context) {
	schedules, err := repository.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schedules"})
		log.Println("Error(handler/schedule): ", err)
		return
	}
	c.JSON(http.StatusOK, schedules)
}

func GetScheduleByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch route"})
		log.Println("Error(handler/schedule): ", err)
		return
	}

	schedules, err := repository.GetSchedulesByVID(h.VID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedules not found"})
		log.Println("Error(handler/schedule :35): ", err)
		return
	}
	log.Println(schedules)
	c.JSON(http.StatusOK, schedules)
}

func NormalizeTimeFields(input []byte) ([]byte, error) {
	var raw map[string]interface{}
	err := json.Unmarshal(input, &raw)
	if err != nil {
		return nil, err
	}

	// Fields to normalize
	timeFields := []string{"departure_time", "arrival_time"}

	for _, field := range timeFields {
		if val, ok := raw[field]; ok {
			if strVal, ok := val.(string); ok && len(strVal) == len("15:04:05") {
				// Prepend a dummy date to make it RFC3339
				raw[field] = "2006-01-02T" + strVal + "Z"
			}
		}
	}

	return json.Marshal(raw)
}

func CreateSchedule(c *gin.Context) {
	// Read the raw request body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		log.Println("Error reading body: ", err)
		return
	}

	// Normalize time-only fields into RFC3339 format
	normalizedJSON, err := NormalizeTimeFields(bodyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to normalize time"})
		log.Println("Error normalizing time: ", err)
		return
	}

	// Now unmarshal into your struct
	var s models.Schedule
	if err := json.Unmarshal(normalizedJSON, &s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error unmarshaling into struct: ", err)
		return
	}
	// log.Println(s.ScheduleID)
	// Save to database
	err = repository.CreateSchedule(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		log.Println("Error(handler/schedule): ", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created"})
}

func UpdateSchedule(c *gin.Context) {
	var s models.Schedule
	if err := c.ShouldBindJSON(&s); err != nil {
		log.Println("JSON Bind Error:", err) // <-- See the exact issue
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := s.ScheduleID
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing schedule ID in body"})
		return
	}

	err := repository.UpdateSchedule(id, s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated"})
}

func DeleteSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		log.Println("Error(handler/schedule): ", err)
		return
	}

	err = repository.DeleteSchedule(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		log.Println("Error(handler/schedule): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted"})
}
