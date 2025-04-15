package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPerformsMaintenanceRoutes(rg *gin.Engine) {
	pm := rg.Group("/performs-maintenance")
	{
		pm.Use(middleware.JWTAuthMiddleware())
		pm.GET("/all", handler.GetAllPerformsMaintenance)
		pm.GET("/", handler.GetPerformsMaintenanceByMID)
		pm.POST("/", handler.CreatePerformsMaintenance)
		pm.DELETE("/:m_id/:staff_id", handler.DeletePerformsMaintenance)
	}
}
