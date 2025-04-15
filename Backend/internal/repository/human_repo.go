package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
	"log"
)

// Get all humans
func GetAllHumans() ([]models.Human, error) {
	rows, err := db.DB.Query("SELECT ID_NO, FNAME, LNAME, DOB, AGE, V_ID FROM human")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var humans []models.Human
	for rows.Next() {
		var h models.Human
		if err := rows.Scan(&h.IDNo, &h.FName, &h.LName, &h.DOB, &h.Age, &h.VID); err != nil {
			return nil, err
		}
		humans = append(humans, h)
	}
	return humans, nil
}

// Get one human by ID
func GetHumanByID(id int) (models.Human, error) {
	var h models.Human
	err := db.DB.QueryRow("SELECT ID_NO, FNAME, LNAME, DOB, AGE, V_ID FROM human WHERE ID_NO = ?", id).Scan(
		&h.IDNo, &h.FName, &h.LName, &h.DOB, &h.Age, &h.VID,
	)
	// if err == nil {
	// 	log.Println("Human retrieved successfully")
	// }
	return h, err
}

// Create a new human
func CreateHuman(h models.Human) error {
	_, err := db.DB.Exec("INSERT INTO human (ID_NO, FNAME, LNAME, DOB, AGE, V_ID) VALUES (?, ?, ?, ?, ?, ?)",
		h.IDNo, h.FName, h.LName, h.DOB, h.Age, h.VID,
	)
	return err
}

// Update a human
func UpdateHuman(id int, h models.Human) error {
	// log.Println(id)
	_, err := db.DB.Exec("UPDATE human SET FNAME = ?, LNAME = ?, DOB = ?, AGE = ?, V_ID = ? WHERE ID_NO = ?",
		h.FName, h.LName, h.DOB, h.Age, h.VID, id,
	)
	if err != nil {
		log.Println(h.VID)
	}
	return err
}

// Delete a human
func DeleteHuman(id int) error {
	_, err := db.DB.Exec("DELETE FROM human WHERE ID_NO = ?", id)
	return err
}
