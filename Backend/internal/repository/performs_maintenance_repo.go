package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllPerformsMaintenance() ([]models.PerformsMaintenance, error) {
	rows, err := db.DB.Query("SELECT M_ID, STAFF_ID FROM performs_maintenance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PerformsMaintenance
	for rows.Next() {
		var pm models.PerformsMaintenance
		if err := rows.Scan(&pm.MID, &pm.StaffID); err != nil {
			return nil, err
		}
		list = append(list, pm)
	}
	return list, nil
}

func GetPerformsMaintenanceByMID(mid int) ([]models.PerformsMaintenance, error) {
	rows, err := db.DB.Query("SELECT M_ID, STAFF_ID FROM performs_maintenance WHERE M_ID = ?", mid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PerformsMaintenance
	for rows.Next() {
		var pm models.PerformsMaintenance
		if err := rows.Scan(&pm.MID, &pm.StaffID); err != nil {
			return nil, err
		}
		list = append(list, pm)
	}
	return list, nil
}

func CreatePerformsMaintenance(pm models.PerformsMaintenance) error {
	_, err := db.DB.Exec("INSERT INTO performs_maintenance (M_ID, STAFF_ID) VALUES (?, ?)", pm.MID, pm.StaffID)
	return err
}

func DeletePerformsMaintenance(mid, staffID int) error {
	_, err := db.DB.Exec("DELETE FROM performs_maintenance WHERE M_ID = ? AND STAFF_ID = ?", mid, staffID)
	return err
}
