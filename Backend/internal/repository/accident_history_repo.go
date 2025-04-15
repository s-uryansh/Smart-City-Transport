package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

func GetAllAccidentHistory() ([]models.AccidentHistory, error) {
	rows, err := db.DB.Query("SELECT V_ID, I_ID FROM accident_history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []models.AccidentHistory
	for rows.Next() {
		var ah models.AccidentHistory
		if err := rows.Scan(&ah.VID, &ah.IID); err != nil {
			return nil, err
		}
		histories = append(histories, ah)
	}
	return histories, nil
}
func GetAccidentHistoryByID(id int) ([]models.AccidentHistory, error) {
	rows, err := db.DB.Query("SELECT V_ID, I_ID FROM accident_history WHERE V_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accidents []models.AccidentHistory
	for rows.Next() {
		var i models.AccidentHistory
		if err := rows.Scan(&i.IID, &i.VID); err != nil {
			return nil, err
		}
		accidents = append(accidents, i)
	}
	return accidents, nil
}
func CreateAccidentHistory(ah models.AccidentHistory) error {
	_, err := db.DB.Exec("INSERT INTO accident_history (V_ID, I_ID) VALUES (?, ?)", ah.VID, ah.IID)
	return err
}

func DeleteAccidentHistory(vid, iid int) error {
	_, err := db.DB.Exec("DELETE FROM accident_history WHERE V_ID = ? AND I_ID = ?", vid, iid)
	return err
}
