package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterScheduleFollowedRoutes(rg *gin.Engine) {
	sf := rg.Group("/schedule-followed")
	{
		sf.Use(middleware.JWTAuthMiddleware())
		sf.GET("/all", handler.GetAllScheduleFollowed)
		sf.GET("/", handler.GetScheduleFollowedByID)
		sf.POST("/", handler.CreateScheduleFollowed)
	}
}
