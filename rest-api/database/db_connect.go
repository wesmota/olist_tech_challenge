package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("database is connected")
	}

	return db
}
