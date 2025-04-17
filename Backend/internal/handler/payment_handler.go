package handler

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPayments(c *gin.Context) {
	list, err := repository.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetPaymentByID(c *gin.Context) {
	uid := c.GetInt("user_id")
	log.Println(uid)
	payment, err := repository.GetPaymentByID(uid)
	log.Println(payment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusOK, payment)
}

func CreatePayment(c *gin.Context) {
	uid := c.GetInt("user_id")
	var h models.Human
	h, err := repository.GetHumanByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch operates_on"})
		log.Println("Error(handler/operates_on): ", err)
		return
	}
	hid := h.IDNo
	var p models.Payment
	p.PassengerID = hid
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	if err := repository.CreatePayment(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Payment created"})
}
func CreatePaymentByID(c *gin.Context) {
	var p models.Payment
	var err error
	p.PassengerID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	if err := repository.CreatePayment(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Payment created"})
}
func UpdatePayment(c *gin.Context) {
	userID := c.GetInt("user_id")
	log.Println("User ID: ", userID)
	var p models.Payment
	p.PassengerID = userID
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	if err := repository.UpdatePayment(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment updated"})
}

func DeletePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("payment_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	if err := repository.DeletePayment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		log.Println("Error(handler/payments): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted"})
}
