package gorm

import (
	"gorm.io/gorm"
)

func Rollback(tx *gorm.DB, entities ...interface{}) error {
	for _, entity := range entities {
		if err := DropTable(tx, entity); err != nil {
			return err
		}
	}

	return nil
}

func DropTable(tx *gorm.DB, entity interface{}) error {
	exist := tx.Migrator().HasTable(entity)
	if exist {
		if err := tx.Migrator().DropTable(entity); err != nil {
			return err
		}
	}

	return nil
}

func DropColumn(tx *gorm.DB, entity interface{}, column string) error {
	exist := tx.Migrator().HasColumn(entity, column)
	if exist {
		if err := tx.Migrator().DropColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}
