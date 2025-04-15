package models

import "time"

// Schema for schedule

type Schedule struct {
	ScheduleID    int       `json:"schedule_id" db:"SCHEDULE_ID"`       // Primary key
	RID           int       `json:"r_id" db:"R_ID"`                     // Foreign key (route ID)
	VID           int       `json:"v_id" db:"V_ID"`                     // Foreign key (vehicle ID)
	DepartureTime time.Time `json:"departure_time" db:"DEPARTURE_TIME"` // Departure time
	ArrivalTime   time.Time `json:"arrival_time" db:"ARRIVAL_TIME"`     // Arrival time
}

/*
+----------------+---------+------+-----+---------+-------+
| Field          | Type    | Null | Key | Default | Extra |
+----------------+---------+------+-----+---------+-------+
| SCHEDULE_ID    | int(11) | NO   | PRI | NULL    |       |
| R_ID           | int(11) | YES  | MUL | NULL    |       |
| V_ID           | int(11) | YES  | MUL | NULL    |       |
| DEPARTURE_TIME | time    | YES  |     | NULL    |       |
| ARRIVAL_TIME   | time    | YES  |     | NULL    |       |
+----------------+---------+------+-----+---------+-------+
*/
