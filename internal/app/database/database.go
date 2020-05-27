package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Open() error {
	data, err := sql.Open("postgres", db.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := data.Ping(); err != nil {
		return err
	}

	db.db = data

	return nil
}

func (db *Database) Close() {
	db.db.Close()
}
