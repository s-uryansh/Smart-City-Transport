package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

// Get all operates_on records
func GetAllOperatesOn() ([]models.OperatesOn, error) {
	rows, err := db.DB.Query("SELECT V_ID, S_ID FROM operates_on")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.OperatesOn
	for rows.Next() {
		var op models.OperatesOn
		if err := rows.Scan(&op.VID, &op.SID); err != nil {
			return nil, err
		}
		list = append(list, op)
	}
	return list, nil
}

// Create a new operates_on entry
func CreateOperatesOn(op models.OperatesOn) error {
	_, err := db.DB.Exec("INSERT INTO operates_on (V_ID, S_ID) VALUES (?, ?)",
		op.VID, op.SID)
	return err
}

// Get a specific operates_on entry by V_ID and S_ID
func GetOperatesOnByIDs(vID int) ([]models.OperatesOn, error) {
	rows, err := db.DB.Query("SELECT V_ID, S_ID FROM operates_on WHERE V_ID = ?", vID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.OperatesOn
	for rows.Next() {
		var op models.OperatesOn
		if err := rows.Scan(&op.VID, &op.SID); err != nil {
			return nil, err
		}
		list = append(list, op)
	}
	return list, err

}

// Delete a specific operates_on entry
func DeleteOperatesOn(vID, sID int) error {
	_, err := db.DB.Exec("DELETE FROM operates_on WHERE V_ID = ? AND S_ID = ?", vID, sID)
	return err
}
