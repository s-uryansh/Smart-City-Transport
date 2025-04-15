package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

// For JSON unmarshaling
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // remove quotes
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

// For SQL insertion
func (cd CustomDate) Value() (driver.Value, error) {
	return cd.Time.Format("2006-01-02"), nil
}
func (cd *CustomDate) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan DOB: %v", value)
	}
	cd.Time = t
	return nil
}

//Schema for human

type Human struct {
	IDNo  int        `json:"id_no" db:"ID_NO"` // Primary key
	LName string     `json:"lname" db:"LNAME"` // Last name
	FName string     `json:"fname" db:"FNAME"` // First name
	DOB   CustomDate `json:"dob" db:"DOB"`     // Date of birth
	Age   int        `json:"age" db:"AGE"`     // Age
	VID   int        `json:"v_id" db:"V_ID"`   // Foreign key (vehicle ID)
}

/*
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| ID_NO | int(11)     | NO   | PRI | NULL    |       |
| FNAME | varchar(50) | YES  |     | NULL    |       |
| LNAME | varchar(50) | YES  |     | NULL    |       |
| DOB   | date        | YES  |     | NULL    |       |
| AGE   | int(11)     | YES  |     | NULL    |       |
| V_ID  | int(11)     | YES  | MUL | NULL    |       |
+-------+-------------+------+-----+---------+-------+
*/
