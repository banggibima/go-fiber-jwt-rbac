package postgres

import (
	"database/sql"

	"github.com/banggibima/go-fiber-jwt-rbac/config"
	_ "github.com/lib/pq"
)

func New(config *config.Config) (*sql.DB, error) {
	driver := config.Postgres.Driver
	connection := config.Postgres.Connection

	db, err := sql.Open(driver, connection)
	if err != nil {
		return nil, err
	}

	if err := Connect(db); err != nil {
		return nil, err
	}

	return db, nil
}
