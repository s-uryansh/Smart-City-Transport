package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutUser(c *gin.Context) {
	// Clear cookie by setting max age to -1
	c.SetCookie("Authorization", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
