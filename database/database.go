// package database allows for the access to the MySQL database and
// provide easy access the underlying data.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connect will provide the caller with a db connection
func Connect(user, pass, host, database string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, database))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Resources returns a slice of strings containing resources
func Resources(conn *sql.DB) ([]string, error) {
	rows, err := conn.Query("SELECT NAME FROM RESOURCES")
	if err != nil {
		return nil, err
	}
	var resources []string
	for rows.Next() {
		var resource string
		err = rows.Scan(&resource)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

// Users returns a string slice containing users
func Users(conn *sql.DB) ([]string, error) {
	rows, err := conn.Query("SELECT NAME FROM USERS")
	if err != nil {
		return nil, err
	}
	var users []string
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return nil, err
		}
		users = append(users, username)
	}
	return users, nil
}

// ResourceStatus returns the status for the given resource
func ResourceStatus(conn *sql.DB, resource string) (string, error) {
	rows, err := conn.Query(fmt.Sprintf(`SELECT STATES.NAME FROM STATES JOIN RESOURCES ON RESOURCES.STATE = STATES.ID WHERE RESOURCES.NAME = '%s'`, resource))
	if err != nil {
		return "", err
	}
	var status string
	for rows.Next() {
		err = rows.Scan(&status)
		if err != nil {
			return "", err
		}
	}
	return status, nil
}
