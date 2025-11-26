package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connection established")
	return createTables()
}

// createTables creates all necessary tables
func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS generals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			branch TEXT NOT NULL,
			rank TEXT NOT NULL,
			birth_date TEXT,
			death_date TEXT,
			biography TEXT,
			photo_url TEXT,
			notable_battles TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS terms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			term TEXT NOT NULL,
			definition TEXT NOT NULL,
			category TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS battles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			date TEXT,
			location TEXT,
			description TEXT,
			participants TEXT,
			outcome TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			general_id INTEGER NOT NULL,
			quote_text TEXT NOT NULL,
			context TEXT,
			date TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (general_id) REFERENCES generals(id)
		)`,
	}

	for _, query := range queries {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	log.Println("Database tables created successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
