package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouteFollowedRoutes(rg *gin.Engine) {
	route := rg.Group("/route-followed")
	{
		route.Use(middleware.JWTAuthMiddleware())
		route.GET("/all", handler.GetAllRouteFollowed)
		route.GET("/", handler.GetRouteByVehicleID)
		route.POST("/", handler.CreateRouteFollowed)
		route.DELETE("/:route_id", handler.DeleteRouteFollowed)
	}
}
