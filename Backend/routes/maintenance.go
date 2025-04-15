package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMaintenanceRoutes(rg *gin.Engine) {
	m := rg.Group("/maintenance")
	{
		m.Use(middleware.JWTAuthMiddleware())
		m.GET("/all", handler.GetAllMaintenance)
		m.GET("/", handler.GetMaintenanceByID)
		m.POST("/", handler.CreateMaintenance)
		m.PUT("/:id", handler.UpdateMaintenance)
		m.DELETE("/:id", handler.DeleteMaintenance)
	}
}
