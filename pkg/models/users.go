package models

import (
	"github.com/dr4g0n369/libraryManagement/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(user *types.Login) error { // Insert a new user
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}
	result, err := db.Exec(`INSERT INTO users (username, password, role) VALUES (?, sha1(?), ?)`, user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}
	defer db.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id

	return nil
}

func LoginUser(user *types.Login) error { // Query a single user
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "SELECT id, role FROM users WHERE username = ? and password = sha1(?)"
	if err := db.QueryRow(query, user.Username, user.Password).Scan(&user.Id, &user.Role); err != nil {
		return err
	}

	return nil
}

func GetUserDetails(user *types.Login) error { // Query a single user
	db, err := ConnectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "SELECT username, role FROM users WHERE id = ?"
	if err := db.QueryRow(query, user.Id).Scan(&user.Username, &user.Role); err != nil {
		return err
	}

	return nil
}
