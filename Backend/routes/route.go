package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouteRoutes(rg *gin.Engine) {
	routes := rg.Group("/routes")
	{
		routes.Use(middleware.JWTAuthMiddleware())
		routes.GET("/all", handler.GetAllRoutes)
		routes.GET("/:id", handler.GetRouteByID)
		routes.POST("/:id", handler.CreateRoute)
		routes.PUT("/:id", handler.UpdateRoute)
		routes.DELETE("/:id", handler.DeleteRoute)
	}
}
