package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterVehicleRoutes(r *gin.Engine) {
	vehicle := r.Group("/vehicles")
	{
		vehicle.POST("/", handler.CreateVehicle) // Create a new vehicle
		vehicle.Use(middleware.JWTAuthMiddleware())
		vehicle.GET("/all", handler.GetAllVehicles) // Get all vehicles
		vehicle.GET("/", handler.GetVehicleByID)    // Get vehicle by ID
		vehicle.PUT("/", handler.UpdateVehicle)     // Update vehicle by ID
		vehicle.DELETE("/", handler.DeleteVehicle)  // Delete vehicle by ID
	}
}
