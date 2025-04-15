package repository

import (
	"SmartCityTransportSystem/internal/models"
	"SmartCityTransportSystem/pkg/db"
	"SmartCityTransportSystem/pkg/utils"
	"fmt"
)

// check for duplicate username
func IsUsernameTaken(username string) (bool, error) {
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
	return exists, err
}
func GetAllUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT ID_NO, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.IDNo, &u.Username, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserByID(id int) (*models.User, error) {
	row := db.DB.QueryRow("SELECT ID_NO, username, password FROM users WHERE ID_NO = ?", id)
	var u models.User
	err := row.Scan(&u.IDNo, &u.Username, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	row := db.DB.QueryRow("SELECT ID_NO, username, password FROM users WHERE username = ?", username)
	var user models.User
	err := row.Scan(&user.IDNo, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(u models.User) error {
	//no duplicate username
	exists, err := IsUsernameTaken(u.Username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username already taken")
	}

	//password encryption
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	//add to db
	_, err = db.DB.Exec("INSERT INTO users (ID_NO, username, password) VALUES (?, ?, ?)",
		u.IDNo, u.Username, hashedPassword)
	return err
}

func UpdateUser(u models.User) error {
	_, err := db.DB.Exec("UPDATE users SET username = ?, password = ? WHERE ID_NO = ?", u.Username, u.Password, u.IDNo)
	return err
}

func DeleteUser(id int) error {
	_, err := db.DB.Exec("DELETE FROM users WHERE ID_NO = ?", id)
	return err
}
