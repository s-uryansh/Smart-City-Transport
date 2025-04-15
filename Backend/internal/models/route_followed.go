package models

// Schema for route_followed

type RouteFollowed struct {
	VehicleID int `json:"v_id" db:"VEHICLE_ID"` // Foreign key to Vehicle ID
	RouteID   int `json:"r_id" db:"RT_ID"`      // Foreign key to Route ID
}

/*
+------------+---------+------+-----+---------+-------+
| Field      | Type    | Null | Key | Default | Extra |
+------------+---------+------+-----+---------+-------+
| VEHICLE_ID | int(11) | NO   | PRI | NULL    |       |
| RT_ID      | int(11) | NO   | PRI | NULL    |       |
+------------+---------+------+-----+---------+-------+
*/
