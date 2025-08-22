package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetSQliteClient(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
