package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllPayments() ([]models.Payment, error) {
	rows, err := db.DB.Query("SELECT PAYMENT_ID, PASSENGER_ID, AMOUNT, METHOD FROM payment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var p models.Payment
		if err := rows.Scan(&p.PaymentID, &p.PassengerID, &p.Amount, &p.Method); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}

func GetPaymentByID(id int) ([]models.Payment, error) {
	rows, err := db.DB.Query("SELECT PAYMENT_ID, PASSENGER_ID, AMOUNT, METHOD FROM payment WHERE PASSENGER_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []models.Payment
	for rows.Next() {
		var p models.Payment
		if err := rows.Scan(&p.PaymentID, &p.PassengerID, &p.Amount, &p.Method); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	return payments, nil
}

func CreatePayment(p models.Payment) error {
	_, err := db.DB.Exec("INSERT INTO payment (PAYMENT_ID, PASSENGER_ID, AMOUNT, METHOD) VALUES (?, ?, ?, ?)",
		p.PaymentID, p.PassengerID, p.Amount, p.Method)
	return err
}

func UpdatePayment(p models.Payment) error {
	_, err := db.DB.Exec("UPDATE payment SET AMOUNT = ?, METHOD = ? WHERE PAYMENT_ID = ?",
		p.Amount, p.Method, p.PaymentID)
	return err
}

func DeletePayment(id int) error {
	_, err := db.DB.Exec("DELETE FROM payment WHERE PAYMENT_ID = ?", id)
	return err
}
