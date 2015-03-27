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
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct{}
type Permissions struct{}
type Resource struct{}
type States struct{}

// Connect will provide the caller with a db connection
func Connect(conf *config) (sql.Conn, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", conf.db.username, conf.db.password, conf.db.host, conf.db.database))
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
