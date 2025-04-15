package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterHumanRoutes(r *gin.Engine) {
	h := r.Group("/humans")
	{
		h.POST("/", handler.CreateHuman)
		h.Use(middleware.JWTAuthMiddleware())
		h.GET("/all", handler.GetHumans)
		h.GET("/", handler.GetHuman)
		h.PUT("/", handler.UpdateHuman)
		h.DELETE("/", handler.DeleteHuman)
	}
}
