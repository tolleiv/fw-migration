package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tolleiv/fw-migration/config"
	"github.com/tolleiv/fw-migration/config/admin"
	_ "github.com/tolleiv/fw-migration/db/migrations"
)

func main() {

	// Register route
	mux := http.NewServeMux()
	// amount to /admin, so visit `/admin` to view the admin interface
	admin.Admin.MountTo("/admin", mux)

	fmt.Println(fmt.Sprintf("Listening on: %d", config.Config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), mux)
}
