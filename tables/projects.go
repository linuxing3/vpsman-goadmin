package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProjectsTable(ctx *context.Context) table.Table {

	projects := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := projects.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Type", "type", db.Varchar)
	info.AddField("Status", "status", db.Varchar)
	info.AddField("Is_active", "is_active", db.Varchar)
	info.AddField("Percent_complete", "percent_complete", db.Varchar)
	info.AddField("Expected_start_date", "expected_start_date", db.Date)
	info.AddField("Expected_end_date", "expected_end_date", db.Date)
	info.AddField("Priority", "priority", db.Varchar)
	info.AddField("Department", "department", db.Varchar)
	info.AddField("Tasks", "tasks", db.Varchar)
	info.AddField("Notes", "notes", db.Varchar)
	info.AddField("Actual_start_date", "actual_start_date", db.Date)
	info.AddField("Actual_end_date", "actual_end_date", db.Date)
	info.AddField("Estimated_cost", "estimated_cost", db.Varchar)
	info.AddField("Total_cost", "total_cost", db.Varchar)
	info.AddField("Expense_claim", "expense_claim", db.Varchar)
	info.AddField("Collect_progress", "collect_progress", db.Varchar)
	info.AddField("Frequency", "frequency", db.Varchar)
	info.AddField("From_time", "from_time", db.Datetime)
	info.AddField("To_time", "to_time", db.Datetime)
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)

	info.SetTable("projects").SetTitle("Projects").SetDescription("Projects")

	formList := projects.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Title", "title", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Type", "type", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Status", "status", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Is_active", "is_active", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Percent_complete", "percent_complete", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Expected_start_date", "expected_start_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Expected_end_date", "expected_end_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Priority", "priority", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Department", "department", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Tasks", "tasks", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Notes", "notes", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Actual_start_date", "actual_start_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Actual_end_date", "actual_end_date", db.Date, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Estimated_cost", "estimated_cost", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Total_cost", "total_cost", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Expense_claim", "expense_claim", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Collect_progress", "collect_progress", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Frequency", "frequency", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("From_time", "from_time", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("To_time", "to_time", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("projects").SetTitle("Projects").SetDescription("Projects")

	return projects
}
