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
		vehicle.POST("/:id", handler.UpdateVehicleCount)
		vehicle.Use(middleware.JWTAuthMiddleware())
		vehicle.GET("/all", handler.GetAllVehicles) // Get all vehicles
		vehicle.GET("/", handler.GetVehicleByID)    // Get vehicle by ID
		vehicle.PUT("/", handler.UpdateVehicle)
		vehicle.PUT("/:id", handler.UpdateVehicleByID) // Update vehicle by ID
		vehicle.DELETE("/", handler.DeleteVehicle)     // Delete vehicle by ID
		vehicle.DELETE("/:id", handler.DeleteVehicle)  // Delete vehicle by ID
	}
}
