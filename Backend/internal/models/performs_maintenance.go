package models

// Schema for performs_maintenance

type PerformsMaintenance struct {
	MID     int `json:"m_id" db:"M_ID"`         // Foreign key to Maintenance ID
	StaffID int `json:"staff_id" db:"STAFF_ID"` // Foreign key to Staff ID
}

/*
+----------+---------+------+-----+---------+-------+
| Field    | Type    | Null | Key | Default | Extra |
+----------+---------+------+-----+---------+-------+
| M_ID     | int(11) | NO   | PRI | NULL    |       |
| STAFF_ID | int(11) | NO   | PRI | NULL    |       |
+----------+---------+------+-----+---------+-------+
*/
