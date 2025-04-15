package models

// Schema for route

type Route struct {
	RID         int     `json:"r_id" db:"R_ID"`                 // Primary key
	JourneyTime string  `json:"journey_time" db:"JOURNEY_TIME"` // Time of journey
	StartPoint  string  `json:"start_point" db:"START_POINT"`   // Starting location
	EndPoint    string  `json:"end_point" db:"END_POINT"`       // Ending location
	Distance    float64 `json:"distance" db:"DISTANCE"`         // Distance of the route
}

/*
+--------------+---------------+------+-----+---------+-------+
| Field        | Type          | Null | Key | Default | Extra |
+--------------+---------------+------+-----+---------+-------+
| R_ID         | int(11)       | NO   | PRI | NULL    |       |
| JOURNEY_TIME | time          | YES  |     | NULL    |       |
| START_POINT  | varchar(100)  | YES  |     | NULL    |       |
| END_POINT    | varchar(100)  | YES  |     | NULL    |       |
| DISTANCE     | decimal(10,2) | YES  |     | NULL    |       |
+--------------+---------------+------+-----+---------+-------+
*/
