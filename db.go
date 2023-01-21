package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	FirSyncDBPath = "./firsync.db"
)

func createTables(db *sqlx.DB) error {
	// Create the users table
	usersTable := `
        CREATE TABLE users (
            id INTEGER PRIMARY KEY,
            public_key TEXT NOT NULL UNIQUE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `

	// Create the projects table
	projectsTable := `
        CREATE TABLE projects (
            id INTEGER PRIMARY KEY,
            user_id INTEGER NOT NULL,
            name TEXT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
        );
    `

	// Create the commits table
	commitsTable := `
        CREATE TABLE commits (
            id INTEGER PRIMARY KEY,
            project_id INTEGER NOT NULL,
            message TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
        );
    `

	// Execute the table creation statements
	if _, err := db.Exec(usersTable); err != nil {
		return err
	}
	if _, err := db.Exec(projectsTable); err != nil {
		return err
	}
	if _, err := db.Exec(commitsTable); err != nil {
		return err
	}

	return nil
}

func setupDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", FirSyncDBPath)
	if err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func destroyDB() error {
	err := os.Remove(FirSyncDBPath)
	if err != nil {
		return fmt.Errorf("error deleting the DB file: %v", err)
	}
	return nil
}
