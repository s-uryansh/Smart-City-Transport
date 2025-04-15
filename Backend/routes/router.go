package routes

import (
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	AllowCredentials: true,
	// }))
	RegisterHumanRoutes(r)
	RegisterAuthRoutes(r)
	RegisterUserRoutes(r)
	RegisterVehicleRoutes(r)
	RegisterIncidentRoutes(r)
	r.Use(middleware.JWTAuthMiddleware())

	// auth := r.Group("/")
	RegisterRouteRoutes(r)
	RegisterScheduleRoutes(r)
	RegisterScheduleFollowedRoutes(r)
	RegisterOperatesOnRoutes(r)
	RegisterAccidentHistoryRoutes(r)
	RegisterMaintenanceRoutes(r)
	RegisterMaintenanceHistoryRoutes(r)
	RegisterPerformsMaintenanceRoutes(r)
	RegisterPaymentRoutes(r)
	RegisterRouteFollowedRoutes(r)

}
