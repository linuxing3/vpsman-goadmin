package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetFlightsTable(ctx *context.Context) table.Table {

	flights := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := flights.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Departure", "departure", db.Varchar)
	info.AddField("Destination", "destination", db.Varchar)
	info.AddField("Flight_duration", "flight_duration", db.Float)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("flights").SetTitle("Flights").SetDescription("Flights")

	formList := flights.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Departure", "departure", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Destination", "destination", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Flight_duration", "flight_duration", db.Float, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("flights").SetTitle("Flights").SetDescription("Flights")

	return flights
}
