package migrations

import (
	"github.com/tolleiv/fw-migration/app/models"
	"github.com/tolleiv/fw-migration/db"
)

func init() {
	AutoMigrate(&models.Port{})
	AutoMigrate(&models.Proto{})
	AutoMigrate(&models.Application{})
	AutoMigrate(&models.ResourceGroup{})
	AutoMigrate(&models.Resource{})
	AutoMigrate(&models.Action{})
	AutoMigrate(&models.SecurityZone{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
