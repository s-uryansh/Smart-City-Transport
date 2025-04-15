package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllMaintenanceHistories() ([]models.MaintenanceHistory, error) {
	rows, err := db.DB.Query("SELECT M_ID, V_ID FROM maintenance_history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.MaintenanceHistory
	for rows.Next() {
		var mh models.MaintenanceHistory
		if err := rows.Scan(&mh.MID, &mh.VID); err != nil {
			return nil, err
		}
		list = append(list, mh)
	}
	return list, nil
}

func GetMaintenanceHistoryByMID(vid int) ([]models.MaintenanceHistory, error) {
	rows, err := db.DB.Query("SELECT M_ID, V_ID FROM maintenance_history WHERE V_ID = ?", vid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.MaintenanceHistory
	for rows.Next() {
		var mh models.MaintenanceHistory
		if err := rows.Scan(&mh.MID, &mh.VID); err != nil {
			return nil, err
		}
		list = append(list, mh)
	}
	return list, nil
}

func CreateMaintenanceHistory(mh models.MaintenanceHistory) error {
	_, err := db.DB.Exec("INSERT INTO maintenance_history (M_ID, V_ID) VALUES (?, ?)", mh.MID, mh.VID)
	return err
}

func DeleteMaintenanceHistory(mid, vid int) error {
	_, err := db.DB.Exec("DELETE FROM maintenance_history WHERE M_ID = ? AND V_ID = ?", mid, vid)
	return err
}
