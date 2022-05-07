package repository

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	articleTable   = "articles"
	translateTable = "translate"
)

func NewSQLiteDB(dbFilePath string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(50)
	conn.SetConnMaxIdleTime(10 * time.Second)
	conn.SetConnMaxLifetime(10 * time.Second)

	sqlStmt := `	
	CREATE TABLE IF NOT EXISTS ` + articleTable + ` (
		id INTEGER PRIMARY KEY,
		channel_id VARCHAR(255) NOT NULL,
		signature VARCHAR(255) NOT NULL,
		html_message TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS ` + translateTable + ` (
		id INTEGER PRIMARY KEY,
		ru VARCHAR(255) NOT NULL,
		en VARCHAR(255) NOT NULL
	);
	`

	_, err = conn.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
