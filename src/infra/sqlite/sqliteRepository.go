package sqlite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type connectorManager interface {
	getConnection() (*sql.Conn, error)
	closeConnection(conn *sql.Conn)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (instance DatabaseConnectionManager) getConnection() (*sql.Conn, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	userTable := `
		CREATE TABLE IF NOT EXISTS user (
		    			id TEXT PRIMARY KEY,
		    			name TEXT,
		    			email TEXT UNIQUE
		    		);
	`

	_, err = db.Exec(userTable)

	if err != nil {
		return nil, err
	}

	conn, _ := db.Conn(context.Background())

	return conn, nil
}

func (instance DatabaseConnectionManager) closeConnection(conn *sql.Conn) {
	err := conn.Close()

	if err != nil {
		panic(err)
	}
}
