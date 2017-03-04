package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/qor"
	"github.com/qor/admin"
)

type Port struct {
	gorm.Model
	Num string
}

type Proto struct {
	gorm.Model
	Name string
}

type Application struct {
	gorm.Model
	Name string
	Source Port
	Destination Port
	Proto Proto
}
func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&Port{}, &Proto{}, &Application{})

	// Initalize
	Admin := admin.New(&qor.Config{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&Port{})
	Admin.AddResource(&Proto{})
	Admin.AddResource(&Application{})

	// Register route
	mux := http.NewServeMux()
	// amount to /admin, so visit `/admin` to view the admin interface
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)
}