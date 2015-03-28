// package database allows for the access to the MySQL database and
// provide easy access the underlying data.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User contains a mapping of the USER table
type User struct {
	ID         int    // user ID
	Name       string // username
	Permission int    // user permission level
}

// Permission contains a mapping of the PERMISSION table
type Permission struct {
	ID          int    // permission ID
	Name        string // permission name
	Description string // permission description
}

// Resource contains a mapping of the RESOURCE table
type Resource struct {
	ID          int    // resource ID
	Name        string // resource name
	Description string // resource description
	User        int    // user ID resource assigned to
	State       int    // state of resource
}

// State contains a mapping of the STATE table
type State struct {
	ID          int    // state ID
	Name        string // state name
	Description string // state description
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
