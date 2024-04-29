package gorm

import (
	"gorm.io/gorm"
)

func Migrate(tx *gorm.DB, entities ...interface{}) error {
	for _, entity := range entities {
		if err := CreateTable(tx, entity); err != nil {
			return err
		}
	}

	return nil
}

func CreateTable(tx *gorm.DB, entity interface{}) error {
	exists := tx.Migrator().HasTable(entity)
	if !exists {
		if err := tx.Migrator().CreateTable(entity); err != nil {
			return err
		}
	}

	return nil
}

func AddColumn(tx *gorm.DB, entity interface{}, column string) error {
	exists := tx.Migrator().HasColumn(entity, column)
	if !exists {
		if err := tx.Migrator().AddColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}

func AlterColumn(tx *gorm.DB, entity interface{}, column string) error {
	exists := tx.Migrator().HasColumn(entity, column)
	if exists {
		if err := tx.Migrator().AlterColumn(entity, column); err != nil {
			return err
		}
	}

	return nil
}
