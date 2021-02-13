package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetVehiclesTable(ctx *context.Context) table.Table {

	vehicles := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := vehicles.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Departure", "departure", db.Varchar)
	info.AddField("License_plate", "license_plate", db.Varchar)
	info.AddField("Mark", "mark", db.Varchar)
	info.AddField("Model", "model", db.Varchar)
	info.AddField("Details", "details", db.Varchar)
	info.AddField("Last_odometer_value", "last_odometer_value", db.Varchar)
	info.AddField("Location", "location", db.Varchar)
	info.AddField("Chassis_no", "chassis_no", db.Varchar)
	info.AddField("Value", "value", db.Varchar)
	info.AddField("Insurance_details", "insurance_details", db.Varchar)
	info.AddField("Insurance_company", "insurance_company", db.Varchar)
	info.AddField("Policy_no", "policy_no", db.Varchar)
	info.AddField("Start_date", "start_date", db.Date)
	info.AddField("End_date", "end_date", db.Date)
	info.AddField("Additional_details", "additional_details", db.Varchar)
	info.AddField("Fueltype", "fueltype", db.Varchar)
	info.AddField("Fuel_uom", "fuel_uom", db.Varchar)
	info.AddField("Color", "color", db.Varchar)
	info.AddField("Wheels", "wheels", db.Varchar)
	info.AddField("Door", "door", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("vehicles").SetTitle("Vehicles").SetDescription("Vehicles")

	formList := vehicles.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Departure", "departure", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("License_plate", "license_plate", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Mark", "mark", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Model", "model", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Details", "details", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Last_odometer_value", "last_odometer_value", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Location", "location", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Chassis_no", "chassis_no", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Value", "value", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Insurance_details", "insurance_details", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Insurance_company", "insurance_company", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Policy_no", "policy_no", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Start_date", "start_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("End_date", "end_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Additional_details", "additional_details", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Fueltype", "fueltype", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Fuel_uom", "fuel_uom", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Color", "color", db.Varchar, form.Color).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Wheels", "wheels", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Door", "door", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("vehicles").SetTitle("Vehicles").SetDescription("Vehicles")

	return vehicles
}
