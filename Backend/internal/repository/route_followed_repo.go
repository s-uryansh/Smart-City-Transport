package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
	"log"
)

// Get all route_followed entries
func GetAllRouteFollowed() ([]models.RouteFollowed, error) {
	rows, err := db.DB.Query("SELECT VEHICLE_ID, RT_ID FROM route_followed")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.RouteFollowed
	for rows.Next() {
		var rf models.RouteFollowed
		if err := rows.Scan(&rf.VehicleID, &rf.RouteID); err != nil {
			return nil, err
		}
		list = append(list, rf)
	}
	return list, nil
}
func GetRouteByVehicleID(vehicleID int) (int, error) {
	row, err := db.DB.Query("SELECT RT_ID FROM route_followed WHERE VEHICLE_ID = ?", vehicleID)
	if err != nil {
		return 0, err
	}
	defer row.Close()
	var routeID int
	if row.Next() {
		if err := row.Scan(&routeID); err != nil {
			return 0, err
		}
	}
	return routeID, nil
}
func GetRoutesByVehicleID(vehicleID int) ([]int, error) {
	rows, err := db.DB.Query("SELECT RT_ID FROM route_followed WHERE VEHICLE_ID = ?", vehicleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routeIDs []int
	for rows.Next() {
		var routeID int
		if err := rows.Scan(&routeID); err != nil {
			return nil, err
		}
		routeIDs = append(routeIDs, routeID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routeIDs, nil
}

// Insert a new route_followed entry
func CreateRouteFollowed(rf models.RouteFollowed) error {
	log.Println(rf.VehicleID)
	_, err := db.DB.Exec("INSERT INTO route_followed (VEHICLE_ID, RT_ID) VALUES (?, ?)", rf.VehicleID, rf.RouteID)
	return err
}

// Delete a route_followed entry by vehicle ID and route ID
func DeleteRouteFollowed(vehicleID, routeID int) error {
	_, err := db.DB.Exec("DELETE FROM route_followed WHERE VEHICLE_ID = ? AND RT_ID = ?", vehicleID, routeID)
	return err
}
