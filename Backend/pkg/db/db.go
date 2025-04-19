package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	// cfg := config.AppConfig
	user := "root"
	password := "6233"
	host := "35.200.128.58"
	port := "3306"
	name := "smartcity"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening DB(pkg/db/db.go): %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error pinging DB(pkg/db/db.go): %v", err)
	}

	log.Println("Connected to MySQL database")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("MySQL database connection closed")
	}
}
