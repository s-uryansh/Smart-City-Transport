package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllMaintenance() ([]models.Maintenance, error) {
	rows, err := db.DB.Query("SELECT MAINTENANCE_ID, V_ID, ISSUE_REPORTED, REPAIR_STATUS FROM maintenance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(&m.MaintenanceID, &m.VID, &m.IssueReported, &m.RepairStatus); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

func GetMaintenanceByID(id int) ([]models.Maintenance, error) {
	rows, err := db.DB.Query("SELECT MAINTENANCE_ID, V_ID, ISSUE_REPORTED, REPAIR_STATUS FROM maintenance WHERE V_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(&m.MaintenanceID, &m.VID, &m.IssueReported, &m.RepairStatus); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, err
}
func GetMaintenanceByMID(id int) ([]models.Maintenance, error) {
	rows, err := db.DB.Query("SELECT MAINTENANCE_ID, V_ID, ISSUE_REPORTED, REPAIR_STATUS FROM maintenance WHERE MAINTENANCE_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Maintenance
	for rows.Next() {
		var m models.Maintenance
		if err := rows.Scan(&m.MaintenanceID, &m.VID, &m.IssueReported, &m.RepairStatus); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, err
}
func CreateMaintenance(m models.Maintenance) error {
	_, err := db.DB.Exec(
		"INSERT INTO maintenance (MAINTENANCE_ID, V_ID, ISSUE_REPORTED, REPAIR_STATUS) VALUES (?, ?, ?, ?)",
		m.MaintenanceID, m.VID, m.IssueReported, m.RepairStatus)
	return err
}

func UpdateMaintenance(id int, m models.Maintenance) error {
	_, err := db.DB.Exec(
		"UPDATE maintenance SET V_ID = ?, ISSUE_REPORTED = ?, REPAIR_STATUS = ? WHERE MAINTENANCE_ID = ?",
		m.VID, m.IssueReported, m.RepairStatus, id)
	return err
}

func DeleteMaintenance(id int) error {
	_, err := db.DB.Exec("DELETE FROM maintenance WHERE MAINTENANCE_ID = ?", id)
	return err
}
