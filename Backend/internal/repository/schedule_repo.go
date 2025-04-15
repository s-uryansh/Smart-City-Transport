package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
	"fmt"
	"log"
	"time"
)

// Get all schedules
func GetAllSchedules() ([]models.Schedule, error) {
	rows, err := db.DB.Query("SELECT SCHEDULE_ID, R_ID, V_ID, DEPARTURE_TIME, ARRIVAL_TIME FROM schedule")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule
	for rows.Next() {
		var s models.Schedule
		if err := rows.Scan(&s.ScheduleID, &s.RID, &s.VID, &s.DepartureTime, &s.ArrivalTime); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

// Get schedule by ID
func GetScheduleByID(id int) (models.Schedule, error) {
	var s models.Schedule
	var depStr, arrStr string

	err := db.DB.QueryRow("SELECT SCHEDULE_ID, R_ID, V_ID, DEPARTURE_TIME, ARRIVAL_TIME FROM schedule WHERE V_ID = ?", id).
		Scan(&s.ScheduleID, &s.RID, &s.VID, &depStr, &arrStr)

	if err != nil {
		return s, err
	}

	// Parse "15:04:05" time format
	s.DepartureTime, err = time.Parse("15:04:05", depStr)
	if err != nil {
		return s, fmt.Errorf("failed to parse departure time: %w", err)
	}

	s.ArrivalTime, err = time.Parse("15:04:05", arrStr)
	if err != nil {
		return s, fmt.Errorf("failed to parse arrival time: %w", err)
	}

	return s, nil
}
func GetSchedulesByVID(id int) ([]models.Schedule, error) {
	rows, err := db.DB.Query(`
		SELECT SCHEDULE_ID, R_ID, V_ID, DEPARTURE_TIME, ARRIVAL_TIME 
		FROM schedule WHERE V_ID = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []models.Schedule

	for rows.Next() {
		var s models.Schedule
		var depStr, arrStr string

		if err := rows.Scan(&s.ScheduleID, &s.RID, &s.VID, &depStr, &arrStr); err != nil {
			return nil, err
		}

		s.DepartureTime, err = time.Parse("15:04:05", depStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse departure time: %w", err)
		}

		s.ArrivalTime, err = time.Parse("15:04:05", arrStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse arrival time: %w", err)
		}

		schedules = append(schedules, s)
	}

	return schedules, nil
}

func GetScheduleByVID(id int) (models.Schedule, error) {
	var s models.Schedule
	var depStr, arrStr string // temporary strings for parsing

	err := db.DB.QueryRow(`
		SELECT SCHEDULE_ID, R_ID, V_ID, DEPARTURE_TIME, ARRIVAL_TIME 
		FROM schedule WHERE V_ID = ?`, id,
	).Scan(&s.ScheduleID, &s.RID, &s.VID, &depStr, &arrStr)
	if err != nil {
		return s, err
	}

	// Parse TIME strings into time.Time (assuming format is "15:04:05")
	s.DepartureTime, err = time.Parse("15:04:05", depStr)
	if err != nil {
		return s, fmt.Errorf("failed to parse departure time: %w", err)
	}

	s.ArrivalTime, err = time.Parse("15:04:05", arrStr)
	if err != nil {
		return s, fmt.Errorf("failed to parse arrival time: %w", err)
	}

	return s, nil
}

// Create new schedule
func CreateSchedule(s models.Schedule) error {
	// log.Println(s.ScheduleID)
	log.Println(s)
	_, err := db.DB.Exec("INSERT INTO schedule (SCHEDULE_ID, R_ID, V_ID, DEPARTURE_TIME, ARRIVAL_TIME) VALUES (?,?, ?, ?, ?)",
		s.ScheduleID, s.RID, s.VID, s.DepartureTime, s.ArrivalTime)
	return err
}

// Update schedule by ID
func UpdateSchedule(id int, s models.Schedule) error {
	_, err := db.DB.Exec(
		"UPDATE schedule SET R_ID = ?, V_ID = ?, DEPARTURE_TIME = ?, ARRIVAL_TIME = ? WHERE SCHEDULE_ID = ?",
		s.RID,
		s.VID,
		s.DepartureTime.Format("15:04:05"), // ← only time part
		s.ArrivalTime.Format("15:04:05"),
		id,
	)

	return err
}

// Delete schedule by ID
func DeleteSchedule(id int) error {
	_, err := db.DB.Exec("DELETE FROM schedule WHERE SCHEDULE_ID = ?", id)
	return err
}
