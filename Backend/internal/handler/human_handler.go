package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHumans(c *gin.Context) {
	humans, err := repository.GetAllHumans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch humans"})
		log.Println("Error(handler/human): ", err)
		return
	}
	c.JSON(http.StatusOK, humans)
}

func GetHuman(c *gin.Context) {
	userID := c.GetInt("user_id")
	// id, _ := strconv.Atoi(c.Param("id")) //ID passed by user
	// log.Println("User got his details")
	human, err := repository.GetHumanByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Human not found"})
		log.Println("Error(handler/human): ", err)
		return
	}

	if uint(human.IDNo) != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, human)
}

func CreateHuman(c *gin.Context) {
	var human models.Human
	human.IDNo = c.GetInt("user_id")
	// id, _ := strconv.Atoi(c.Param("IDNo"))
	// human.IDNo = id
	if err := c.ShouldBindJSON(&human); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/human): ", err)
		return
	}
	log.Println(human.VID)
	if err := repository.CreateHuman(human); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create human"})
		log.Println("Error(handler/human): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Human created"})
}

func UpdateHuman(c *gin.Context) {
	userID := c.GetInt("user_id")
	// id, _ := strconv.Atoi(c.Param("id"))

	var human models.Human
	if err := c.ShouldBindJSON(&human); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/human): ", err)
		return
	}

	existingHuman, err := repository.GetHumanByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Human not found"})
		return
	}

	if uint(existingHuman.IDNo) != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if human.VID == 0 {
		human.VID = existingHuman.VID
	}

	if err := repository.UpdateHuman(userID, human); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update human"})
		log.Println("Error(handler/human): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Human updated"})
}

func DeleteHuman(c *gin.Context) {
	userID := c.GetInt("user_id")
	// id, _ := strconv.Atoi(c.Param("id"))

	human, err := repository.GetHumanByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Human not found"})
		return
	}

	if uint(human.IDNo) != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := repository.DeleteHuman(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete human"})
		log.Println("Error(handler/human): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Human deleted"})
}
