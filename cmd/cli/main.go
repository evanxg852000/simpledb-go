package main

import (
	"fmt"
	"os"
	"path"

	"github.com/c-bata/go-prompt"
	"github.com/evanxg852000/simpledb/internal/plan"
	"github.com/evanxg852000/simpledb/internal/server"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	fmt.Println("Welcome SimpleDB CLI!")
	workspaceDir := "./data"
	err := os.MkdirAll(workspaceDir, 0755)
	if err != nil {
		fmt.Println("error initializing database: ", err)
	}

	dbDir := path.Join(workspaceDir, "students_db")
	db := server.NewSimpleDB(dbDir, 4096, 10)

	if db.FileManager().IsNew() {
		prepareUsersTable(db)
	}

	//for debugin
	// tx := db.NewTx()
	// c, err := db.Planner().ExecuteQuery("select * from users where id = 3", tx)
	// _, rows := fetchResult(c)
	// _ = rows

	for {
		sqlInput := prompt.Input(">> ", completer,
			prompt.OptionTitle("sql-prompt"),
			prompt.OptionHistory([]string{"select id, name from users"}),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
		)

		if sqlInput == ".exit" {
			os.Exit(0)
		}

		tx := db.NewTx()
		stmtResult, err := db.Planner().ExecuteQuery(sqlInput, tx)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			continue
		}

		tx.Commit()

		if affectedRows, ok := stmtResult.(int64); ok {
			fmt.Println("Affected rows:", affectedRows)
			continue
		}

		headers, rows := fetchResult(stmtResult)
		renderDataTable(headers, rows)
		fmt.Println("")
	}

}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func renderDataTable(headers table.Row, rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(headers)
	t.AppendRows(rows)
	t.Render()
}

func prepareUsersTable(db *server.SimpleDB) {
	t1 := db.NewTx()
	_, err := db.Planner().ExecuteQuery("create table users(id int, name varchar(10))", t1)
	if err != nil {
		t1.Rollback()
		panic(err)
	}
	t1.Commit()

	names := []string{"Evan", "John", "Jane", "Rodriguez", "Samuel", "Mauris"}
	t2 := db.NewTx()
	for id, name := range names {
		sqlStmt := fmt.Sprintf("insert into users(id, name) values(%d, '%s')", id+1, name)
		_, err = db.Planner().ExecuteQuery(sqlStmt, t2)
		if err != nil {
			t2.Rollback()
			panic(err)
		}
	}

	t2.Commit()
}

func fetchResult(result any) (table.Row, []table.Row) {
	plan := result.(plan.Plan)
	schema := plan.Schema()
	scan := plan.Open()
	rows := make([]table.Row, 0)
	for scan.Next() {
		row := make([]string, 0)
		for _, fieldName := range schema.Fields() {
			data := scan.GetValue(fieldName)
			row = append(row, data.String())
		}
		rows = append(rows, makeRow(row))
	}
	return makeRow(schema.Fields()), rows
}

func makeRow(data []string) []any {
	rows := make([]any, 0, len(data))
	for _, item := range data {
		rows = append(rows, item)
	}
	return rows
}
