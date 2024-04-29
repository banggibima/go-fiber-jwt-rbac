package gorm

import (
	"database/sql"

	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	entities := []interface{}{
		&entity.User{},
	}

	if err := Migrate(gormDB, entities...); err != nil {
		return nil, err
	}

	return gormDB, nil
}
