package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Connect(connectionString string) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type Sqlite3 struct {
	db *sql.DB
}

func (s *Sqlite3) Connect(file string) error {
	db, err := sql.Open("sqlite3", file)	
	if err != nil {
		return fmt.Errorf("failed to open sqlite db: %w", err)
	}
	s.db = db
	return nil
}

func (s *Sqlite3) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("Failed to close sqllite db: %w", err)
	}
	return nil
}

func (s *Sqlite3) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func (s *Sqlite3) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *Sqlite3) EnsureSchema() error {
	query := `
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			shape TEXT NOT NULL,
			flour TEXT NOT NULL,
			topping TEXT NOT NULL,
			scoring TEXT NOT NULL,
			order_date TIMESTAMP NOT NULL,
			status TEXT NOT NULL
		);`
	_, err := s.db.Exec(query)
	return err
}

