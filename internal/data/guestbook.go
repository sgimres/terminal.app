package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"charm.land/log/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type GuestbookEntry struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

var db *sql.DB

func InitGuestbookDB() error {
	var err error

	url, ok := os.LookupEnv("TURSO_DATABASE_URL")

	if !ok {
		log.Error("TURSO_DATABASE_URL not set")
	}

	token, ok := os.LookupEnv("TURSO_AUTH_TOKEN")

	if !ok {
		log.Error("TURSO_AUTH_TOKEN not set")
	}

	url = fmt.Sprintf("%s?authToken=%s", url, token)

	db, err = sql.Open("libsql", url)
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
	return err
}

func AddGuestbookEntry(name, description string, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO guestbook (name, description) VALUES (?, ?)",
		name, description,
	)
	return err
}

func CloseGuestbookDB() {
	if db != nil {
		_ = db.Close()
	}
}

func GetDB() *sql.DB {
	return db
}
