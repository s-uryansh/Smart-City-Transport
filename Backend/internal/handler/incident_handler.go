package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllIncidents(c *gin.Context) {
	data, err := repository.GetAllIncidents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incidents"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetIncidentByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	vid := h.VID
	data, err := repository.GetIncidentByID(vid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateIncident(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	vid := h.VID
	var i models.Incident
	i.VID = vid
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	if i.ReportTimeDate.IsZero() {
		i.ReportTimeDate = time.Now()
	}
	if err := repository.CreateIncident(vid, i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create incident"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Incident created"})
}

func UpdateIncident(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch human"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident ID"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	var i models.Incident
	i.VID = h.VID
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	if i.ReportTimeDate.IsZero() {
		i.ReportTimeDate = time.Now()
	}
	if err := repository.UpdateIncident(id, i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update incident"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Incident updated"})
}

func DeleteIncident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident ID"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	if err := repository.DeleteIncident(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete incident"})
		log.Println("Error(handler/incident): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Incident deleted"})
}
