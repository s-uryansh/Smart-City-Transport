package routes

import (
	"SmartCityTransportSystem/internal/handler"
	"SmartCityTransportSystem/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(rg *gin.Engine) {
	pay := rg.Group("/payments")
	{
		pay.Use(middleware.JWTAuthMiddleware())
		pay.GET("/all", handler.GetAllPayments)
		pay.GET("/", handler.GetPaymentByID)
		pay.POST("/", handler.CreatePayment)
		pay.PUT("/", handler.UpdatePayment)
		pay.DELETE("/:payment_id", handler.DeletePayment)
	}
}
