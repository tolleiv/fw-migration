package api

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/tolleiv/fw-migration/app/models"
	"github.com/tolleiv/fw-migration/db"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})
	API.AddResource(&models.Action{})
}
