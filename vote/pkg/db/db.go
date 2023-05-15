package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/ibmdb/go_ibm_db" // Import the DB2 driver
)

// read the DB connection details from the env vars
var db2ConnStr = fmt.Sprintf(
	"HOSTNAME=%s;PORT=%s;DATABASE=%s;UID=%s;PWD=%s",
	os.Getenv("DB_HOSTNAME"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_DATABASE"),
	os.Getenv("DB_USERNAME"),
	os.Getenv("DB_PASSWORD"),
)

var db2Conn *sql.DB

// Initializes a new DB2 connection if db2Conn is nil i.e., there is no active connection.
// If there is an active connection, it will return the existing connection.
func GetDB2Connection() (*sql.DB, error) {
	if db2Conn == nil {
		var err error
		db2Conn, err = sql.Open("go_ibm_db", db2ConnStr)
		if err != nil {
			return nil, err
		}
	}
	return db2Conn, nil
}

// CloseDB2Connection closes the DB2 connection if it is active.
func CloseDB2Connection() error {
	if db2Conn != nil {
		return db2Conn.Close()
	}
	return nil
}
