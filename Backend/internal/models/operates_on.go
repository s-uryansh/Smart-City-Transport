package models

// Schema for operates_on

type OperatesOn struct {
	VID int `json:"v_id" db:"V_ID"` // Foreign key to Vehicle ID
	SID int `json:"s_id" db:"S_ID"` // Foreign key to Schedule ID
}

/*
+-------+---------+------+-----+---------+-------+
| Field | Type    | Null | Key | Default | Extra |
+-------+---------+------+-----+---------+-------+
| V_ID  | int(11) | NO   | PRI | NULL    |       |
| S_ID  | int(11) | NO   | PRI | NULL    |       |
+-------+---------+------+-----+---------+-------+
*/
