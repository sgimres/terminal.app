package data

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type GuestbookEntry struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

var db *sql.DB

func InitGuestbookDB() error {
	dbPath := "./db.sqlite3"
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS guestbook (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func AddGuestbookEntry(name, description string) error {
	if db == nil {
		return errors.New("database not initialized")
	}
	_, err := db.Exec(
		"INSERT INTO guestbook (name, description) VALUES (?, ?)",
		name, description,
	)
	return err
}

func CloseGuestbookDB() {
	if db != nil {
		db.Close()
	}
}
