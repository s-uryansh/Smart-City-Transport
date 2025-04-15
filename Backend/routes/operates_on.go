package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOperatesOnRoutes(rg *gin.Engine) {
	operates := rg.Group("/operates_on")
	{
		operates.Use(middleware.JWTAuthMiddleware())
		operates.GET("/all", handler.GetAllOperatesOn)
		operates.POST("/", handler.CreateOperatesOn)
		operates.GET("/", handler.GetOperatesOnByIDs)
		operates.DELETE("/:vid/:sid", handler.DeleteOperatesOn)
	}
}
