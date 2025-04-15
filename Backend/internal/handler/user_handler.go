package handler

import (
	"log"
	"net/http"

	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.GetInt("user_id")
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	err := repository.CreateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
	var u models.User
	u.IDNo = c.GetInt("user_id")
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	err := repository.UpdateUser(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	id := c.GetInt("user_id")

	err := repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Error (handler/user_handler.go)")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
