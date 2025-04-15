```MySQL
+----------------+-------------+------+-----+---------+-------+
| Field          | Type        | Null | Key | Default | Extra |
+----------------+-------------+------+-----+---------+-------+
| MAINTENANCE_ID | int(11)     | NO   | PRI | NULL    |       |
| V_ID           | int(11)     | YES  | MUL | NULL    |       |
| ISSUE_REPORTED | text        | YES  |     | NULL    |       |
| REPAIR_STATUS  | varchar(50) | YES  |     | NULL    |       |
+----------------+-------------+------+-----+---------+-------+
```