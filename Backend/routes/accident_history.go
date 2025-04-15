package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAccidentHistoryRoutes(rg *gin.Engine) {
	ah := rg.Group("/accident-history")
	{
		ah.Use(middleware.JWTAuthMiddleware())

		ah.GET("/all", handler.GetAllAccidentHistory)
		ah.GET("/", handler.GetAccidentHistoryByID)
		ah.POST("/", handler.CreateAccidentHistory)
		ah.DELETE("/:vid/:iid", handler.DeleteAccidentHistory)
	}
}
