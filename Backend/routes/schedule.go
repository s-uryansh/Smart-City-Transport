package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterScheduleRoutes(rg *gin.Engine) {
	schedule := rg.Group("/schedule")
	{
		schedule.Use(middleware.JWTAuthMiddleware())
		schedule.GET("/all", handler.GetAllSchedules)
		schedule.GET("/", handler.GetScheduleByID)
		schedule.POST("/", handler.CreateSchedule)
		schedule.PUT("/:id", handler.UpdateSchedule)
		schedule.DELETE("/:id", handler.DeleteSchedule)
	}
}
