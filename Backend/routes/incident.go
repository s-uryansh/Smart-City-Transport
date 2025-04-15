package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterIncidentRoutes(r *gin.Engine) {
	incident := r.Group("/incident")

	{
		incident.Use(middleware.JWTAuthMiddleware())
		incident.GET("/all", handler.GetAllIncidents)
		incident.GET("/", handler.GetIncidentByID)
		incident.POST("/", handler.CreateIncident)
		incident.PUT("/:id", handler.UpdateIncident)
		incident.DELETE("/:id", handler.DeleteIncident)
	}
}
