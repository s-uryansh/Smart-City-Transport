package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

// Get all routes
func GetAllRoutes() ([]models.Route, error) {
	rows, err := db.DB.Query("SELECT R_ID, JOURNEY_TIME, START_POINT, END_POINT, DISTANCE FROM route")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var r models.Route
		if err := rows.Scan(&r.RID, &r.JourneyTime, &r.StartPoint, &r.EndPoint, &r.Distance); err != nil {
			return nil, err
		}
		routes = append(routes, r)
	}
	return routes, nil
}

// Get one route by ID
func GetRouteByID(id int) (models.Route, error) {
	var r models.Route
	err := db.DB.QueryRow("SELECT R_ID, JOURNEY_TIME, START_POINT, END_POINT, DISTANCE FROM route WHERE R_ID = ?", id).
		Scan(&r.RID, &r.JourneyTime, &r.StartPoint, &r.EndPoint, &r.Distance)
	if err != nil {
		return r, err
	}
	return r, nil
}

// Insert route (assumes R_ID is auto-incremented by DB)
func CreateRoute(r models.Route) error {
	_, err := db.DB.Exec(
		"INSERT INTO route (R_ID ,JOURNEY_TIME, START_POINT, END_POINT, DISTANCE) VALUES (?,?, ?, ?, ?)",
		r.RID, r.JourneyTime, r.StartPoint, r.EndPoint, r.Distance,
	)
	return err
}

// Update route
func UpdateRoute(id int, r models.Route) error {
	_, err := db.DB.Exec(
		"UPDATE route SET JOURNEY_TIME = ?, START_POINT = ?, END_POINT = ?, DISTANCE = ? WHERE R_ID = ?",
		r.JourneyTime, r.StartPoint, r.EndPoint, r.Distance, id,
	)
	return err
}

// Delete route
func DeleteRoute(id int) error {
	_, err := db.DB.Exec("DELETE FROM route WHERE R_ID = ?", id)
	return err
}
