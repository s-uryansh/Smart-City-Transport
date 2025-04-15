package models

import "time"

// Schema for incident

type Incident struct {
	IncidentID     int       `json:"incident_id" db:"INCIDENT_ID"`           // Primary key
	VID            int       `json:"v_id" db:"V_ID"`                         // Foreign key (Vehicle ID)
	Description    string    `json:"description" db:"DESCRIPTION"`           // Text description
	ReportTimeDate time.Time `json:"report_time_date" db:"REPORT_TIME_DATE"` // Date and time of the report
}

/*
+------------------+----------+------+-----+---------+-------+
| Field            | Type     | Null | Key | Default | Extra |
+------------------+----------+------+-----+---------+-------+
| INCIDENT_ID      | int(11)  | NO   | PRI | NULL    |       |
| V_ID             | int(11)  | YES  | MUL | NULL    |       |
| DESCRIPTION      | text     | YES  |     | NULL    |       |
| REPORT_TIME_DATE | datetime | YES  |     | NULL    |       |
+------------------+----------+------+-----+---------+-------+
*/
