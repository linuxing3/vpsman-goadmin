package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMembersTable(ctx *context.Context) table.Table {

	members := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := members.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Department", "department", db.Varchar)
	info.AddField("Name", "name", db.Varchar).FieldSortable()
	info.AddField("Gender", "gender", db.Varchar)
	info.AddField("Birthday", "birthday", db.Date)
	info.AddField("Etnia", "etnia", db.Varchar)
	info.AddField("Academic_background", "academic_background", db.Varchar)
	info.AddField("Foreign_language", "foreign_language", db.Varchar)
	info.AddField("Political_role", "political_role", db.Varchar)
	info.AddField("Position_and_rank", "position_and_rank", db.Varchar)
	info.AddField("Militant_role", "militant_role", db.Varchar)
	info.AddField("Duty", "duty", db.Varchar)
	info.AddField("From_entity", "from_entity", db.Varchar).FieldSortable().FieldFilterable()
	info.AddField("Arriving_date", "arriving_date", db.Date)
	info.AddField("Rotating_date", "rotating_date", db.Date)
	info.AddField("Sending_entity", "sending_entity", db.Varchar)
	info.AddField("Conyuge_name", "conyuge_name", db.Varchar)
	info.AddField("Conyuge_entity", "conyuge_entity", db.Varchar)
	info.AddField("Conyuge_bonus", "conyuge_bonus", db.Varchar)
	info.AddField("Memo", "memo", db.Varchar)
	info.AddField("Protocol_id", "protocol_id", db.Varchar)
	info.AddField("Is_active", "is_active", db.Varchar)
	info.AddField("Militant", "militant", db.Varchar)
	info.AddField("Appraisals", "appraisals", db.Varchar)
	info.AddField("Designations", "designations", db.Varchar)
	info.AddField("Projects", "projects", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("members").SetTitle("Members").SetDescription("Members")

	formList := members.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Department", "department", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Varchar, form.Text)
	formList.AddField("Birthday", "birthday", db.Date, form.Datetime)
	formList.AddField("Etnia", "etnia", db.Varchar, form.Text)
	formList.AddField("Academic_background", "academic_background", db.Varchar, form.Text)
	formList.AddField("Foreign_language", "foreign_language", db.Varchar, form.Text)
	formList.AddField("Political_role", "political_role", db.Varchar, form.Text)
	formList.AddField("Position_and_rank", "position_and_rank", db.Varchar, form.Text)
	formList.AddField("Militant_role", "militant_role", db.Varchar, form.Text)
	formList.AddField("Duty", "duty", db.Varchar, form.Text)
	formList.AddField("From_entity", "from_entity", db.Varchar, form.Text)
	formList.AddField("Arriving_date", "arriving_date", db.Date, form.Datetime)
	formList.AddField("Rotating_date", "rotating_date", db.Date, form.Datetime)
	formList.AddField("Sending_entity", "sending_entity", db.Varchar, form.Text)
	formList.AddField("Conyuge_name", "conyuge_name", db.Varchar, form.Text)
	formList.AddField("Conyuge_entity", "conyuge_entity", db.Varchar, form.Text)
	formList.AddField("Conyuge_bonus", "conyuge_bonus", db.Varchar, form.Text)
	formList.AddField("Memo", "memo", db.Varchar, form.Text)
	formList.AddField("Protocol_id", "protocol_id", db.Varchar, form.Text)
	formList.AddField("Is_active", "is_active", db.Varchar, form.Text)
	formList.AddField("Militant", "militant", db.Varchar, form.Text)
	formList.AddField("Appraisals", "appraisals", db.Varchar, form.Text)
	formList.AddField("Designations", "designations", db.Varchar, form.Text)
	formList.AddField("Projects", "projects", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("members").SetTitle("Members").SetDescription("Members")

	return members
}
