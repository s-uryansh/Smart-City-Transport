package routes

import (
	"SmartCityTransportSystem/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.LoginUser)
		auth.POST("/logout", handler.LogoutUser)
	}
}
