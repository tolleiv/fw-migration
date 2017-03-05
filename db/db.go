package db

import (
	"os"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "demo.db")
	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}
	} else {
		panic(err)
	}
}
