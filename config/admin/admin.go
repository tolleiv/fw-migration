package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/tolleiv/fw-migration/app/models"
	"github.com/tolleiv/fw-migration/db"
)

var Admin *admin.Admin

func init() {
	Admin = admin.New(&qor.Config{DB: db.DB})

	securityzone := Admin.AddResource(&models.SecurityZone{}, &admin.Config{Menu: []string{"Inputs"}, Priority: 99})

	port := Admin.AddResource(&models.Port{}, &admin.Config{Menu: []string{"Inputs"}, Priority: 4})
	port.SearchAttrs("Name", "Num")
	port.EditAttrs(
		&admin.Section{
			Title: "Details",
			Rows: [][]string{
				{"Name"},
				{"Num"},
			}},
	)
	port.NewAttrs(port.EditAttrs())

	application := Admin.AddResource(&models.Application{}, &admin.Config{Menu: []string{"Inputs"}, Priority: 2})
	application.SearchAttrs("Name", "Source.Name", "Destination.Name", "Proto")
	application.EditAttrs(
		&admin.Section{
			Title: "Details",
			Rows: [][]string{
				{"Name"},
				{"Source", "Destination", "Proto"},
			}},
	)
	application.NewAttrs(application.EditAttrs())
	application.Meta(&admin.Meta{Name: "Source", Config: &admin.SelectOneConfig{AllowBlank: true}})
	application.Meta(&admin.Meta{Name: "Destination", Config: &admin.SelectOneConfig{AllowBlank: true}})
	application.Meta(&admin.Meta{Name: "Proto", Config: &admin.SelectOneConfig{AllowBlank: true}})
	application.Filter(&admin.Filter{
		Name:   "Source",
		Config: &admin.SelectOneConfig{RemoteDataResource: port},
	})
	application.Filter(&admin.Filter{
		Name:   "Destination",
		Config: &admin.SelectOneConfig{RemoteDataResource: port},
	})

	resourcegroup := Admin.AddResource(&models.ResourceGroup{}, &admin.Config{Menu: []string{"Inputs"}, Priority: 3})
	resourcegroup.SearchAttrs("Name")
	resourcegroup.Filter(&admin.Filter{
		Name:   "SecurityZone",
		Config: &admin.SelectOneConfig{RemoteDataResource: securityzone},
	})

	Admin.AddResource(&models.Resource{}, &admin.Config{Menu: []string{"Inputs"}})

	action := Admin.AddResource(&models.Action{}, &admin.Config{Menu: []string{"Inputs"}, Priority: 1})
	action.SearchAttrs("Application", "Source.Name", "Destination.Name")
	action.Meta(&admin.Meta{Name: "Application", Config: &admin.SelectOneConfig{AllowBlank: true}})
	action.Meta(&admin.Meta{Name: "Source", Config: &admin.SelectOneConfig{AllowBlank: true}})
	action.Meta(&admin.Meta{Name: "Destination", Config: &admin.SelectOneConfig{AllowBlank: true}})
	action.Meta(&admin.Meta{Name: "Action", Config: &admin.SelectOneConfig{Collection: []string{"ACCEPT", "DENY", "REJECT"}}})
	action.EditAttrs(
		&admin.Section{
			Title: "Details",
			Rows: [][]string{
				{"Application"},
				{"Source", "Action", "Destination"},
			}},
	)
	action.NewAttrs(action.EditAttrs())
	action.Filter(&admin.Filter{
		Name:   "Application",
		Config: &admin.SelectOneConfig{RemoteDataResource: application},
	})
	action.Filter(&admin.Filter{
		Name:   "Source",
		Config: &admin.SelectOneConfig{RemoteDataResource: resourcegroup},
	})
	action.Filter(&admin.Filter{
		Name:   "Destination",
		Config: &admin.SelectOneConfig{RemoteDataResource: resourcegroup},
	})

	proto := Admin.AddResource(&models.Proto{}, &admin.Config{Menu: []string{"Inputs"}})
	proto.EditAttrs(
		&admin.Section{
			Title: "Details",
			Rows: [][]string{
				{"Name"},
			}},
	)
	proto.NewAttrs(proto.EditAttrs())
	proto.SearchAttrs("Name")

	Admin.AddSearchResource(action, application, resourcegroup)
}
