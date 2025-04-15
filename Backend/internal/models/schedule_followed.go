package models

// Schema for schedule_followed

type ScheduleFollowed struct {
	RID int `json:"r_id" db:"R_ID"` // Foreign key (route ID)
	SID int `json:"s_id" db:"S_ID"` // Foreign key (staff ID)
}

/*
+-------+---------+------+-----+---------+-------+
| Field | Type    | Null | Key | Default | Extra |
+-------+---------+------+-----+---------+-------+
| R_ID  | int(11) | NO   | PRI | NULL    |       |
| S_ID  | int(11) | NO   | PRI | NULL    |       |
+-------+---------+------+-----+---------+-------+
*/
