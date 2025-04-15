package models

import "time"

// Schema for vehicle

type Vehicle struct {
	VehicleID       int       `json:"vehicle_id" db:"VEHICLE_ID"`             // Primary key
	CurrentLocation string    `json:"current_location" db:"CURRENT_LOCATION"` // Current location of vehicle
	Status          string    `json:"status" db:"STATUS"`                     // Status (e.g., Active, Maintenance, etc.)
	LastUpdate      time.Time `json:"-" db:"LAST_UPDATE"`                     // Last updated timestamp

}

/*
+------------------+--------------+------+-----+---------+-------+
| Field            | Type         | Null | Key | Default | Extra |
+------------------+--------------+------+-----+---------+-------+
| VEHICLE_ID       | int(11)      | NO   | PRI | NULL    |       |
| CURRENT_LOCATION | varchar(100) | YES  |     | NULL    |       |
| STATUS           | varchar(50)  | YES  |     | NULL    |       |
| LAST_UPDATE      | datetime     | YES  |     | NULL    |       |
+------------------+--------------+------+-----+---------+-------+
*/
