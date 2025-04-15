```MySQL
+----------------+---------+------+-----+---------+-------+
| Field          | Type    | Null | Key | Default | Extra |
+----------------+---------+------+-----+---------+-------+
| SCHEDULE_ID    | int(11) | NO   | PRI | NULL    |       |
| R_ID           | int(11) | YES  | MUL | NULL    |       |
| V_ID           | int(11) | YES  | MUL | NULL    |       |
| DEPARTURE_TIME | time    | YES  |     | NULL    |       |
| ARRIVAL_TIME   | time    | YES  |     | NULL    |       |
+----------------+---------+------+-----+---------+-------+
```