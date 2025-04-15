package handler

import (
	"log"
	"net/http"

	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"SmartCityTransportSystem/pkg/jwt"
	"SmartCityTransportSystem/pkg/utils"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	token, err := jwt.GenerateToken(user.IDNo, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	var h models.Human
	h.IDNo = user.IDNo
	h, err = repository.GetHumanByID(h.IDNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch human details"})
		log.Println("Error (handler/login): ", err)
		return
	}
	c.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)

	c.Set("user_id", user.IDNo)
	c.Set("v_id", h.VID)
	c.Set("username", user.Username)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login successful",
		"user_id":  user.IDNo,
		"username": user.Username,
		"token":    token,
	})
}
