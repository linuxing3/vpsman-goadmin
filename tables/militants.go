package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMilitantsTable(ctx *context.Context) table.Table {

	militants := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := militants.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Admitted_at", "admitted_at", db.Varchar)
	info.AddField("Formalized_at", "formalized_at", db.Varchar)
	info.AddField("Registered_at", "registered_at", db.Varchar)
	info.AddField("Transfered_at", "transfered_at", db.Date)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("militants").SetTitle("Militants").SetDescription("Militants")

	formList := militants.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Admitted_at", "admitted_at", db.Varchar, form.Text)
	formList.AddField("Formalized_at", "formalized_at", db.Varchar, form.Text)
	formList.AddField("Registered_at", "registered_at", db.Varchar, form.Text)
	formList.AddField("Transfered_at", "transfered_at", db.Date, form.Datetime)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("militants").SetTitle("Militants").SetDescription("Militants")

	return militants
}
