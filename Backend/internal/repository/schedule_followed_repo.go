package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
)

// Get all entries from schedule_followed
func GetAllScheduleFollowed() ([]models.ScheduleFollowed, error) {
	rows, err := db.DB.Query("SELECT R_ID, S_ID FROM schedule_followed")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sfList []models.ScheduleFollowed
	for rows.Next() {
		var sf models.ScheduleFollowed
		if err := rows.Scan(&sf.RID, &sf.SID); err != nil {
			return nil, err
		}
		sfList = append(sfList, sf)
	}
	return sfList, nil
}

func GetScheduleFollowedByID(id int) ([]models.ScheduleFollowed, error) {
	rows, err := db.DB.Query("SELECT R_ID, S_ID FROM schedule_followed WHERE S_ID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followed []models.ScheduleFollowed
	for rows.Next() {
		var r models.ScheduleFollowed
		if err := rows.Scan(&r.RID, &r.SID); err != nil {
			return nil, err
		}
		followed = append(followed, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followed, nil
}

func CreateScheduleFollowed(sf models.ScheduleFollowed) error {
	_, err := db.DB.Exec("INSERT INTO schedule_followed (R_ID, S_ID) VALUES (?, ?)",
		sf.RID, sf.SID)
	return err
}
func DeleteScheduleFollowed(id int) error {
	_, err := db.DB.Exec("DELETE FROM schedule_followed WHERE S_ID = ?", id)
	return err
}
