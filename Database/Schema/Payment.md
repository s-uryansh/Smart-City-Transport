```MySQL
+--------------+---------------+------+-----+---------+-------+
| Field        | Type          | Null | Key | Default | Extra |
+--------------+---------------+------+-----+---------+-------+
| PAYMENT_ID   | int(11)       | NO   | PRI | NULL    |       |
| PASSENGER_ID | int(11)       | YES  | MUL | NULL    |       |
| AMOUNT       | decimal(10,2) | YES  |     | NULL    |       |
| METHOD       | varchar(50)   | YES  |     | NULL    |       |
+--------------+---------------+------+-----+---------+-------+
```