package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		description TEXT,
		location TEXT,
		dateTime TEXT,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createMetricsTable := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		endpoint TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createMetricsTable)
	if err != nil {
		panic("Could not create metrics table.")
	}
}
