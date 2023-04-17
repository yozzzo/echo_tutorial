package models

import (
	"echo-tutorial/config"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open(config.Config.DBName))

	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Todo{})
}
