package postgres

import (
	"database/sql"
)

func Connect(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
