package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"os"
)

func main() {
	//os.Remove("test-db.sqlite")

	_, err := os.Stat("test-db.sqlite")
	if os.IsNotExist(err) {
		createSqliteDb()
	} else {
		sqliteQuery()
	}
}

func createSqliteDb() {
	// Create SQLite file
	file, err := os.Create("test-db.sqlite")
	if err != nil {
		panic(err)
	}
	// Close the file
	file.Close()

	// Open SQLite file for Query
	sqliteDBFile, err := sql.Open("sqlite3", "test-db.sqlite")
	// Close sqlite file when all operation is done.
	if err != nil {
		panic(err)
	}
	defer sqliteDBFile.Close()

	// Create Table
	prepareStatement, _ := sqliteDBFile.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY , name TEXT, department TEXT)")
	prepareStatement.Exec()

	// Insert sample data
	prepareStatement, _ = sqliteDBFile.Prepare("INSERT INTO people (name, department) VALUES (?,?)")
	prepareStatement.Exec("User1", "Dept 1")
	prepareStatement.Exec("User2", "Dept 2")
	prepareStatement.Exec("User3", "Dept 3")
	prepareStatement.Exec("User4", "Dept 4")
	sqliteQuery()
}

func sqliteQuery() {
	// Open SQLite file for Query
	sqliteDBFile, err := sql.Open("sqlite3", "test-db.sqlite")
	// Close sqlite file when all operation is done.
	if err != nil {
		panic(err)
	}
	defer sqliteDBFile.Close()
	rows, err := sqliteDBFile.Query("SELECT * FROM people")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, department string
		rows.Scan(&id, &name, &department)
		fmt.Println(id, name, department)
	}
}
