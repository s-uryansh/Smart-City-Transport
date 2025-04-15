package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllIncidents() ([]models.Incident, error) {
	rows, err := db.DB.Query("SELECT INCIDENT_ID, V_ID, DESCRIPTION, REPORT_TIME_DATE FROM incident")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []models.Incident
	for rows.Next() {
		var i models.Incident
		if err := rows.Scan(&i.IncidentID, &i.VID, &i.Description, &i.ReportTimeDate); err != nil {
			return nil, err
		}
		incidents = append(incidents, i)
	}
	return incidents, nil
}

func GetIncidentByID(id int) ([]models.Incident, error) {
	rows, err := db.DB.Query("SELECT INCIDENT_ID, V_ID, DESCRIPTION, REPORT_TIME_DATE FROM incident WHERE V_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []models.Incident
	for rows.Next() {
		var i models.Incident
		if err := rows.Scan(&i.IncidentID, &i.VID, &i.Description, &i.ReportTimeDate); err != nil {
			return nil, err
		}
		incidents = append(incidents, i)
	}
	return incidents, nil
}

func CreateIncident(vid int, i models.Incident) error {
	_, err := db.DB.Exec(
		"INSERT INTO incident (INCIDENT_ID, V_ID, DESCRIPTION, REPORT_TIME_DATE) VALUES (?, ?, ?, ?)",
		i.IncidentID, i.VID, i.Description, i.ReportTimeDate)
	return err
}

func UpdateIncident(id int, i models.Incident) error {
	_, err := db.DB.Exec(
		"UPDATE incident SET V_ID = ?, DESCRIPTION = ?, REPORT_TIME_DATE = ? WHERE INCIDENT_ID = ?",
		i.VID, i.Description, i.ReportTimeDate, id)
	return err
}

func DeleteIncident(id int) error {
	_, err := db.DB.Exec("DELETE FROM incident WHERE INCIDENT_ID = ?", id)
	return err
}
