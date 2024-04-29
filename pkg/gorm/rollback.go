package gorm

import (
	"gorm.io/gorm"
)

func Rollback(gormDB *gorm.DB, entities ...interface{}) error {
	for _, entity := range entities {
		if err := DropTable(gormDB, entity); err != nil {
			return err
		}
	}

	return nil
}

func DropTable(gormDB *gorm.DB, entity interface{}) error {
	exist := gormDB.Migrator().HasTable(entity)
	if exist {
		if err := gormDB.Migrator().DropTable(entity); err != nil {
			return err
		}
	}

	return nil
}

func DropColumn(gormDB *gorm.DB, entity interface{}, column string) error {
	exist := gormDB.Migrator().HasColumn(entity, column)
	if exist {
		if err := gormDB.Migrator().DropColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}
