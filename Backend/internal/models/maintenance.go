package models

// Schema for maintenance

type Maintenance struct {
	MaintenanceID int    `json:"maintenance_id" db:"MAINTENANCE_ID"` // Primary key
	VID           int    `json:"v_id" db:"V_ID"`                     // Foreign key (Vehicle ID)
	IssueReported string `json:"issue_reported" db:"ISSUE_REPORTED"` // Description of issue
	RepairStatus  string `json:"repair_status" db:"REPAIR_STATUS"`   // Status of repair
}

/*
+----------------+-------------+------+-----+---------+-------+
| Field          | Type        | Null | Key | Default | Extra |
+----------------+-------------+------+-----+---------+-------+
| MAINTENANCE_ID | int(11)     | NO   | PRI | NULL    |       |
| V_ID           | int(11)     | YES  | MUL | NULL    |       |
| ISSUE_REPORTED | text        | YES  |     | NULL    |       |
| REPAIR_STATUS  | varchar(50) | YES  |     | NULL    |       |
+----------------+-------------+------+-----+---------+-------+
*/
