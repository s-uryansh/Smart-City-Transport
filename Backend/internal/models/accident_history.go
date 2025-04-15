package models

//Schema for accident history

type AccidentHistory struct {
	VID int `json:"v_id" db:"V_ID"` // Foreign key (vehicle ID)
	IID int `json:"i_id" db:"I_ID"` // Foreign key (incident ID)
}

/*
+-------+---------+------+-----+---------+-------+
| Field | Type    | Null | Key | Default | Extra |
+-------+---------+------+-----+---------+-------+
| V_ID  | int(11) | NO   | PRI | NULL    |       |
| I_ID  | int(11) | NO   | PRI | NULL    |       |
+-------+---------+------+-----+---------+-------+
*/
