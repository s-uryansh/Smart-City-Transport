package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
	"log"
	"time"
)

// GetAllVehicles retrieves all vehicles from the database
func GetAllVehicles() ([]models.Vehicle, error) {
	rows, err := db.DB.Query("SELECT VEHICLE_ID, CURRENT_LOCATION, STATUS, LAST_UPDATE FROM vehicle")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {
		var v models.Vehicle
		if err := rows.Scan(&v.VehicleID, &v.CurrentLocation, &v.Status, &v.LastUpdate); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, v)
	}
	return vehicles, nil
}

// GetVehicleByID retrieves a vehicle by its ID
func GetVehicleByID(id int) (models.Vehicle, error) {
	var v models.Vehicle
	err := db.DB.QueryRow("SELECT VEHICLE_ID, CURRENT_LOCATION, STATUS, LAST_UPDATE FROM vehicle WHERE VEHICLE_ID = ?", id).
		Scan(&v.VehicleID, &v.CurrentLocation, &v.Status, &v.LastUpdate)
	return v, err
}

// CreateVehicle inserts a new vehicle into the database
func CreateVehicle(v models.Vehicle) error {
	_, err := db.DB.Exec("INSERT INTO vehicle (VEHICLE_ID, CURRENT_LOCATION, STATUS, LAST_UPDATE) VALUES (?, ?, ?, ?)",
		v.VehicleID, v.CurrentLocation, v.Status, time.Now())
	return err
}

// UpdateVehicle updates an existing vehicle in the database
func UpdateVehicle(id int, v models.Vehicle) error {
	_, err := db.DB.Exec("UPDATE vehicle SET CURRENT_LOCATION = ?, STATUS = ?, LAST_UPDATE = ? WHERE VEHICLE_ID = ?",
		v.CurrentLocation, v.Status, time.Now(), id)
	return err
}
func UpdateVehicleCountRepo(id int) error {
	_, err := db.DB.Exec("UPDATE bookings SET booking_count = ? WHERE vehicle_id = ?", 0, id)
	if err != nil {
		log.Fatal("Error in UpdateVehicleCountRepo: ", err)
	}
	_, err = db.DB.Exec("UPDATE vehicle SET status_change = ? , STATUS = ? WHERE VEHICLE_ID = ?", false, "Available", id)
	if err != nil {
		log.Fatal("Error in UpdateVehicleCountRepo: ", err)
	}
	return nil
}

// DeleteVehicle deletes a vehicle from the database by ID
func DeleteVehicle(id int) error {
	_, err := db.DB.Exec("DELETE FROM vehicle WHERE VEHICLE_ID = ?", id)
	return err
}
