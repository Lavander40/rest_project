package storage

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), "postgre", "5432", os.Getenv("POSTGRES_DB")))
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}
