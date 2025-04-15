package models

// Schema for payment

type Payment struct {
	PaymentID   int     `json:"payment_id" db:"PAYMENT_ID"`     // Primary key
	PassengerID int     `json:"passenger_id" db:"PASSENGER_ID"` // Foreign key to passenger
	Amount      float64 `json:"amount" db:"AMOUNT"`             // Payment amount
	Method      string  `json:"method" db:"METHOD"`             // Payment method
}

/*
+--------------+---------------+------+-----+---------+-------+
| Field        | Type          | Null | Key | Default | Extra |
+--------------+---------------+------+-----+---------+-------+
| PAYMENT_ID   | int(11)       | NO   | PRI | NULL    |       |
| PASSENGER_ID | int(11)       | YES  | MUL | NULL    |       |
| AMOUNT       | decimal(10,2) | YES  |     | NULL    |       |
| METHOD       | varchar(50)   | YES  |     | NULL    |       |
+--------------+---------------+------+-----+---------+-------+
*/
