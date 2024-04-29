package gorm

import (
	"gorm.io/gorm"
)

func Migrate(gormDB *gorm.DB, entities ...interface{}) error {
	for _, entity := range entities {
		if err := CreateTable(gormDB, entity); err != nil {
			return err
		}
	}

	return nil
}

func CreateTable(gormDB *gorm.DB, entity interface{}) error {
	exists := gormDB.Migrator().HasTable(entity)
	if !exists {
		if err := gormDB.Migrator().CreateTable(entity); err != nil {
			return err
		}
	}

	return nil
}

func AddColumn(gormDB *gorm.DB, entity interface{}, column string) error {
	exists := gormDB.Migrator().HasColumn(entity, column)
	if !exists {
		if err := gormDB.Migrator().AddColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}

func AlterColumn(gormDB *gorm.DB, entity interface{}, column string) error {
	exists := gormDB.Migrator().HasColumn(entity, column)
	if exists {
		if err := gormDB.Migrator().AlterColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}
