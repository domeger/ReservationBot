// package database allows for the access to the MySQL database
package database

/*
Database Design
---------------

PK - primary key
FK - foreign key
AI - auto increment

Table Schema
------------

USERS
	ID INT PK AI
	NAME VARCHAR(60)
	PERMISSION_ID INT FK -> PERMISSION.ID

PERMISSIONS
	ID INT PK AI
	NAME VARCHAR(24)
	DESCRIPTION VARCHAR(255)

RESOURCES
	ID INT PK AI
	NAME VARCHAR(60)
	DESCRIPTION VARCHAR(255)
	STATE_ID INT FK -> STATE.ID
	USER_ID INT FK -> USERS.ID

STATES
	ID INT PK AI
	NAME VARCHAR(60)
	DESCRIPTION VARCHAR(255)
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int
	Name       string
	Permission int
}

type Permission struct {
	ID          int
	Name        string
	Description string
}

type Resource struct {
	ID          int
	Name        string
	Description string
	User        int
	State       int
}

type State struct {
	ID          int
	Name        string
	Description string
}

// Connect will provide the caller with a db connection
func Connect(conf map[string]string) (sql.Conn, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", conf["username"], conf["password"], conf["host"], conf["database"]))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetResources
func GetResources(conn driver.Conn) ([]resources, error) {
	return
}

// GetUsers
func GetUsers(conn driver.Conn) ([]users, error) {
	return
}

// GetStatus
func GetStatus(conn driver.Conn, res string) error {
	return nil
}
