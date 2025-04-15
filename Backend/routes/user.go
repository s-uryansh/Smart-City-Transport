package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/", handler.CreateUser)
		users.Use(middleware.JWTAuthMiddleware())

		users.GET("/all", handler.GetAllUsers)
		users.GET("/", handler.GetUserByID)
		users.PUT("/", handler.UpdateUser)
		users.DELETE("/", handler.DeleteUser)
	}
}
