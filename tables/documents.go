package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetDocumentsTable(ctx *context.Context) table.Table {

	documents := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := documents.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Year", "year", db.Varchar)
	info.AddField("Date", "date", db.Date)
	info.AddField("Classi_level", "classi_level", db.Varchar)
	info.AddField("Category", "category", db.Varchar)
	info.AddField("In_or_out", "in_or_out", db.Varchar)
	info.AddField("Sending_code", "sending_code", db.Varchar)
	info.AddField("Ordered_number", "ordered_number", db.Varchar)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Content", "content", db.Varchar)
	info.AddField("To_entity", "to_entity", db.Varchar)
	info.AddField("Copy_entity", "copy_entity", db.Varchar)
	info.AddField("Attachment", "attachment", db.Varchar)
	info.AddField("Keyword", "keyword", db.Varchar)
	info.AddField("Work_entity", "work_entity", db.Varchar)
	info.AddField("Author", "author", db.Varchar)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("documents").SetTitle("Documents").SetDescription("Documents")

	formList := documents.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Year", "year", db.Varchar, form.Text)
	formList.AddField("Date", "date", db.Date, form.Datetime)
	formList.AddField("Classi_level", "classi_level", db.Varchar, form.Text)
	formList.AddField("Category", "category", db.Varchar, form.Text)
	formList.AddField("In_or_out", "in_or_out", db.Varchar, form.Text)
	formList.AddField("Sending_code", "sending_code", db.Varchar, form.Text)
	formList.AddField("Ordered_number", "ordered_number", db.Varchar, form.Text)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.Text)
	formList.AddField("To_entity", "to_entity", db.Varchar, form.Text)
	formList.AddField("Copy_entity", "copy_entity", db.Varchar, form.Text)
	formList.AddField("Attachment", "attachment", db.Varchar, form.Text)
	formList.AddField("Keyword", "keyword", db.Varchar, form.Text)
	formList.AddField("Work_entity", "work_entity", db.Varchar, form.Text)
	formList.AddField("Author", "author", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("documents").SetTitle("Documents").SetDescription("Documents")

	return documents
}
