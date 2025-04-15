package models

// Schema for maintenance history

type MaintenanceHistory struct {
	MID int `json:"m_id" db:"M_ID"` // Foreign key to Maintenance ID
	VID int `json:"v_id" db:"V_ID"` // Foreign key to Vehicle ID
}

/*
+-------+---------+------+-----+---------+-------+
| Field | Type    | Null | Key | Default | Extra |
+-------+---------+------+-----+---------+-------+
| M_ID  | int(11) | NO   | PRI | NULL    |       |
| V_ID  | int(11) | NO   | PRI | NULL    |       |
+-------+---------+------+-----+---------+-------+
*/
