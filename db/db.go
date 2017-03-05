package db

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/tolleiv/fw-migration/config"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("sqlite3", config.Config.DB.Name)
	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}
	} else {
		panic(err)
	}
}
