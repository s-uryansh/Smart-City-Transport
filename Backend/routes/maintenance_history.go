package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterMaintenanceHistoryRoutes(rg *gin.Engine) {
	mh := rg.Group("/maintenance-history")
	{
		mh.Use(middleware.JWTAuthMiddleware())
		mh.GET("/all", handler.GetAllMaintenanceHistories)
		mh.GET("/", handler.GetMaintenanceHistoryByMID)
		mh.POST("/", handler.CreateMaintenanceHistory)
		mh.DELETE("/:m_id/:v_id", handler.DeleteMaintenanceHistory)
	}
}
