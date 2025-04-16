package models

type User struct {
	IDNo     int    `json:"id_no" db:"ID_NO"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"pmassword"`
	Role     string `json:"roles" db:"roles"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
+----------+--------------+------+-----+---------+-------+
| Field    | Type         | Null | Key | Default | Extra |
+----------+--------------+------+-----+---------+-------+
| ID_NO    | int(11)      | NO   | PRI | NULL    |       |
| username | varchar(50)  | NO   |     | NULL    |       |
| password | varchar(255) | NO   |     | NULL    |       |
| roles    | varchar(40)  | YES  |     | user    |       |
+----------+--------------+------+-----+---------+-------+
*/
